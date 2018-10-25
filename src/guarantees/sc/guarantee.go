package sc

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"guarantees/com"
	"guarantees/data"
	"guarantees/data/guarantee/additional"
	"guarantees/data/guarantee/additional/bankpars"
	"guarantees/data/guarantee/primary/guarantee"
	"guarantees/data/guarantee/primary/statement"
	"regexp"
)

//todo change this trash
func (s *SmartContract) createGuaranteeByStatementId(APIstub shim.ChaincodeStubInterface, args []string, simulate string) peer.Response {
	element := com.FPath.Path.PushBack("s.createGuaranteeByStatementId")
	defer com.FPath.Path.Remove(element)

	if len(args) != 1 {
		return com.IncorrectNumberOfArgsError(args, 1)
	}

	statementId := args[0]

	/////////////////////////////////////////////////////////////////
	statementObject := statement.Statement{Id: statementId}
	response := data.QueryById(&statementObject, APIstub)
	if response.Status >= com.ERRORTHRESHOLD {
		return response
	}
	/////////////////////////////////////////////////////////////////

	/////////////////////////////////////////////////////////////////
	bankPars := bankpars.BankPars{}

	response = data.SetIdByRelationId(&bankPars, APIstub, statementObject.StatementSigned.Guarantor.Id)
	if response.Status >= com.ERRORTHRESHOLD {
		return response
	}

	response = data.QueryById(&bankPars, APIstub)
	if response.Status >= com.ERRORTHRESHOLD {
		return response
	}
	/////////////////////////////////////////////////////////////////

	guaranteeObject := guarantee.Guarantee{}
	guaranteeObject.GuaranteeSigned.StatementFields = statementObject.StatementSigned
	guaranteeObject.GuaranteeSigned.BankPars = bankPars

	newPars := []additional.Par{}
	for _, par := range guaranteeObject.GuaranteeSigned.StatementFields.GType.Pars {
		if par.Name != "obligations" && par.Name != "additionalConditions" {
			if par.DisplayType != "checkbox" || par.Value != "false" {
				if par.DisplayType == "multiCheckbox" {
					isAnyOptionChosenTrue := false
					newOptions := []additional.Option{}
					for _, option := range par.Options {
						if option.Chosen != "false" {
							isAnyOptionChosenTrue = true
							option.VisibleConditions = []additional.VisibleCondition{}
							newOptions = append(newOptions, option)
						}
					}
					par.Options = newOptions
					if isAnyOptionChosenTrue {
						matched, err := regexp.MatchString(par.RegularExpression, par.Value)
						if err != nil {
							return com.RegExpMatchError(err, par.RegularExpression, par.Value)
						}
						if !matched {
							return com.RegExpValidationError(par.RegularExpression, par.Value)
						}
						par.ValidationConditions = []additional.ValidationCondition{}
						newPars = append(newPars, par)
					}

				} else {
					if par.Value != "" {
						if par.Caption == "Предполагаемая дата выдачи гарантии" {
							par.Caption = "Дата выдачи гарантии"
						}
						matched, err := regexp.MatchString(par.RegularExpression, par.Value)
						if err != nil {
							return com.RegExpMatchError(err, par.RegularExpression, par.Value)
						}
						if !matched {
							return com.RegExpValidationError(par.RegularExpression, par.Value)
						}
						par.ValidationConditions = []additional.ValidationCondition{}
						newPars = append(newPars, par)
					}
				}
			}
		}
	}
	guaranteeObject.GuaranteeSigned.StatementFields.GType.Pars = newPars

	guaranteeObject.Status = "validationErr"

	data.Put(&guaranteeObject, APIstub, simulate)

	guaranteeOutObject := guaranteeObject.ToOut()

	return com.SuccessPayloadResponse(&guaranteeOutObject)
}

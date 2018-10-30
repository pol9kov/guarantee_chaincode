package bankpars

const (
	// Name of entity (for logs)
	ENTITY_NAME = "BankPars"

	// Xml tag name
	XML_TAG = "bank_parses"

	// Object type names (for storage)
	KEY = "BANKPARS" // default
)

func (bankPars BankPars) CreateValidation() bool {
	return true
}

func (bankPars BankPars) ChangeValidation(newBankParsInterface interface{}) bool {
	_ = newBankParsInterface.(*BankPars)
	return false
}

func (bankPars BankPars) GetKeyObjectType() string {
	return KEY
}

func (bankPars BankPars) GetIndexes() [][]string {
	return [][]string{}
}

func (bankPars BankPars) GetEntityName() string {
	return ENTITY_NAME
}
func (bankPars BankPars) GetTagName() string {
	return XML_TAG
}
func (bankPars *BankPars) SetRelationTxId(relationTxId string) {
	bankPars.RelationTxId = relationTxId
}
func (bankPars *BankPars) GetRelationTxId() string {
	return bankPars.RelationTxId
}
func (bankPars *BankPars) SetId(id string) {
	bankPars.Id = id
}
func (bankPars BankPars) GetId() string {
	return bankPars.Id
}
func (bankPars *BankPars) SetKey(key string) {
	bankPars.Key = key
}
func (bankPars BankPars) GetKey() string {
	return bankPars.Key
}
func (bankPars *BankPars) SetMSPId(MSPId string) {
	bankPars.MSPId = MSPId
}
func (bankPars BankPars) GetMSPId() string {
	return bankPars.MSPId
}

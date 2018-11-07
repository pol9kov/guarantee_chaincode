package guarantees

import (
	"testing"
)

func Test_LifeCircle(t *testing.T) {
	stub := InitGuarantor(t)

	CreateGType(t, stub)
	CreateBankPars(t, stub)
	CreateGuarantee(t, stub)
}

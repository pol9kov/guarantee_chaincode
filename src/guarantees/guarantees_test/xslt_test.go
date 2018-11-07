package guarantees

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func Test_XSLT(t *testing.T) {
	type document map[string]interface{}

	xslFile := "testdata/transform/makeGuarantee.xsl"
	xmlFile := "testdata/transform/generateGuarantee.xml"
	outFile := "testdata/transform/outputGuarantee.xml"

	cmd := exec.Cmd{
		Args: []string{"xsltproc", "-o", outFile, xslFile, xmlFile},
		Env:  os.Environ(),
		Path: "/usr/bin/xsltproc",
	}

	_, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

}

package clamavrestsdk

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed testfiles/eicar_test_virus_com.zip
var testdocument []byte

//go:embed testfiles/helloworld.txt
var testdocument2 []byte

func TestFileUpload(t *testing.T) {
	i := NewclamAVRestInstance("127.0.0.1", 33779, false)
	isInfected, err := i.ScanFile(testdocument)
	if err != nil {
		panic(err)
	}

	fmt.Println("result", isInfected)

	fmt.Println("-----------------------------------")

	i2 := NewclamAVRestInstance("127.0.0.1", 33779, false)
	isInfected2, err2 := i2.ScanFile(testdocument2)
	if err2 != nil {
		panic(err2)
	}

	fmt.Println("result", isInfected2)
}

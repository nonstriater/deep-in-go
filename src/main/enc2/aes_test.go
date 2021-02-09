package enc2

import (
	"fmt"
	"testing"
)

func Test_enc(t *testing.T) {
	token := "8973aa0f19b98a2f28f70a37f80229d1"
	str := "fd0627f6b2f9499e74a425ffdb67c774"
	dec, err := AesCBCDecrypte(str, token);
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(dec)

	if dec != "800|1467792422" {
		t.Errorf("not equal. %s", dec)
	}
}


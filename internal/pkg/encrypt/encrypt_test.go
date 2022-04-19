package encrypt

import (
	"fmt"
	"testing"
)

func TestCipher(t *testing.T) {
	pt := "hello"
	ct := Encrypt([]byte(pt))
	fmt.Println(ct)
	ptb, _ := Decrypt(ct)
	fmt.Println(string(ptb))
}

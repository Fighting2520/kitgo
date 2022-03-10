package encryt

import (
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	md5 := Md5("123456", "")
	md52 := Md5("123456", "123456")
	fmt.Println(md5)
	fmt.Println(md52)
}

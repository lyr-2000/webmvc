package sqlutil

import (
	"fmt"
	"log"
	"testing"
)

func Test_print(t *testing.T) {
	var A struct {
		Value string
		B     string `col:"b"`
	}
	fmt.Println(AsUnderLine("userId"))
	A.Value = "aaa"
	A.B = "aaa"
	mp := AsColMap(&A)
	log.Println(mp)
}


func Test_camel(t *testing.T) {
	log.Println(AsCamelName("studentUserName"))
	log.Println(AsCamelName("student_user_Name"))
	log.Println(AsBigCamelName("student_user_Name"))
	log.Println(AsBigCamelName("studentUserName"))
}
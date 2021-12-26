package parseutil

import (
	"fmt"
	"log"
	"testing"
)

func Test_print(t *testing.T) {
	mp := make(map[string]interface{}, 0)
	mp["a"] = 1
	a := Split("a;b;c;|d=e=f")
	log.Printf("%#v", a)
}

func Test_paraseParam(t *testing.T) {
	mp := make(map[string]interface{}, 0)
	//mp["b"] = "11.2"
	mp["e"] = 24.8
	mp["c"] = "-1.888"
	var Out struct {
		Name  int `alias:"a,b,c,e,d"`
		Nameb string    `alias:"e,d,b"`
	}
	CopyFromMap(&Out, mp)
	log.Printf("%+v\n", Out)
}

func Test_copy_out(t *testing.T) {
	var Out struct {
		Name  int `alias:"userId,user_id,a,b,c,e,d"`
		Nameb string    `alias:"userName,user_name,e,d,b"`
	}
	Out.Name = 666
	Out.Nameb = "userId"
	fmt.Println(CopyOutMap(Out))
}

//func Test_parse(t *testing.T) {
//	var a interface{} = float32(12)
//	var b = a.(int)
//
//	fmt.Println(" ok = ", b)
//
//}

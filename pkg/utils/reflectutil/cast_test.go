package reflectutil

import (
	"log"
	"testing"
	"time"
)

func Test_set_addr(t *testing.T) {
	var i = -1
	SetPrimitive(&i,"32")

	var cur = time.Now()
	log.Println("result = ",i)
	SetPrimitive(&cur,time.Now().Unix())
	log.Println(cur)
}


func Test_time(t*testing.T) {
	var cur time.Time
	SetPrimitive(&cur, "2017-02-27 17:30:20")
	log.Println("res = ",cur)
}
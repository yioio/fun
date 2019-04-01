package main

import (
	"github.com/yioio/fun/funArray"
	"log"
)

func main() {
	ret1 := funArray.In_array("cat", []interface{}{"cat", "dog"})
	ret2 := funArray.In_array_string("cat", []string{"cat", "dog"})
	log.Println(ret1)
	log.Println(ret2) 
}
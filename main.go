package main

import (
	"github.com/yioio/fun/funArray"
	"log"
)

func main() {
	ret := funArray.In_array("cat", []interface{}{"cat", "dog"})
	log.Println(ret)
}
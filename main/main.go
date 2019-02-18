package main

import(
	"reflect"
	"log"
)

func main() {
	ret := reflect.TypeOf(map[string]string{"123Key":"222"})
	log.Println(reflect.TypeOf(ret))
}
package main

import (
	"fmt"
)

type objectArray struct {
	Value1 string
	Value2 string
	Value3 string
	Value4 string
}

func main() {

	objArray := []objectArray{}

	for 1 > 0 {
		objArray = append(objArray, objectArray{
			Value1: "asdfsdfafasdfasfsfsafasdfasdfsadfadsfasdfasdfadsfasdfadsfadsfasdfasdfsadfsadfasdfasdfsadfasasdfsdfafasdfasfsfsafasdfasdfsadfadsfasdfasdfadsfasdfadsfadsfasdfasdfsadfsadfasdfasdfsadfas",
			Value2: "asdfsdfafasdfasfsfsafasdfasdfsadfadsfasdfasdfadsfasdfadsfadsfasdfasdfsadfsadfasdfasdfsadfasasdfsdfafasdfasfsfsafasdfasdfsadfadsfasdfasdfadsfasdfadsfadsfasdfasdfsadfsadfasdfasdfsadfas",
			Value3: "asdfsdfafasdfasfsfsafasdfasdfsadfadsfasdfasdfadsfasdfadsfadsfasdfasdfsadfsadfasdfasdfsadfasasdfsdfafasdfasfsfsafasdfasdfsadfadsfasdfasdfadsfasdfadsfadsfasdfasdfsadfsadfasdfasdfsadfas",
			Value4: "asdfsdfafasdfasfsfsafasdfasdfsadfadsfasdfasdfadsfasdfadsfadsfasdfasdfsadfsadfasdfasdfsadfasasdfsdfafasdfasfsfsafasdfasdfsadfadsfasdfasdfadsfasdfadsfadsfasdfasdfsadfsadfasdfasdfsadfas",
		})
	}
	fmt.Println(objArray)
}

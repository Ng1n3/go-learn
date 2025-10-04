package main

import (
	"fmt"
	"maps"
)

func main() {

	// var mapVariabel map[keyType]valueType
	mayMap := make(map[string]int)

	fmt.Println(mayMap)

	fmt.Println(mayMap)

	mayMap["key1"] = 9
	mayMap["code"] = 18

	fmt.Println(mayMap)
	fmt.Println(mayMap["key"])

	delete(mayMap, "key1")
	fmt.Println(mayMap)

	mayMap["key1"] = 11
	mayMap["key2"] = 14
	mayMap["key3"] = 25
	fmt.Println(mayMap)

	clear(mayMap)
	fmt.Println(mayMap)

	mayMap["key1"] = 11
	mayMap["key2"] = 14
	mayMap["key3"] = 25

	value, unkownvalue := mayMap["key1"]
	fmt.Println("Value", value)
	fmt.Println("unkownvalue:", unkownvalue)

	mayMap2 := map[string]int{"a": 1, "b": 2}

	mayMap3 := map[string]int{"a": 1, "b": 2}

	fmt.Println(mayMap2)

	if maps.Equal(mayMap, mayMap3) {
		fmt.Println("mayMap and mayMap2 are equal")
	}

	for k, v := range mayMap2 {
		fmt.Println(k, v)
	}

	var mayMap4 map[string]string

	if mayMap4 == nil {
		fmt.Println("The map is initalized to nil value")
	} else {
		fmt.Println("The map is not initalized to nil value")
	}

	val, ok := mayMap4["key"]
	if !ok {
		fmt.Println("Key does not exist in mayMap4", val)
	}

	mayMap4 = make(map[string]string)
	mayMap4["key"] = "tom"
	fmt.Println(mayMap4)

	fmt.Println("myMap length is", len(mayMap))

	myMap5 := make(map[string]map[string]string)
	fmt.Println(myMap5)
}

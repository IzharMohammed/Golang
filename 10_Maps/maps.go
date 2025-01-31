package main

import (
	"fmt"
	"maps"
)

func main() {
	// creating maps
	// m := make(map[string]string)
	// m["name"] = "izhar"
	// m["role"] = "Full stack web dev"
	// fmt.Println(m["name"], m["role"])
	// fmt.Println(m["phone"])

	// IMP:- In go maps if key does not exist then it returns zero value
	// Creation of maps with make(we use it if we dont know values in advance)
	// m1 := make(map[string]int)
	// m1["izhar"] = 1
	// m1["random"] = 10
	// fmt.Println(m1["random"], m1["izhar"], m1["who"])
	// fmt.Println(m1)
	// delete(m1, "random")
	// fmt.Println(m1)
	// clear(m1)
	// fmt.Println(m1)

	// creation of maps without make, use it if we know values
	m := map[string]int{"price": 10, "phones": 3}
	// fmt.Println(m)

	// value, bool values for getting map values  
	v, ok := m["phones"]
	
	if ok {
		fmt.Println("ok",v)
	}else{
		fmt.Println("not ok")
	}

	m2 := map[string]int{"price": 10, "phones": 3}
	m3 := map[string]int{"price": 10, "phones": 3}
	fmt.Println(maps.Equal(m2, m3))

	

}

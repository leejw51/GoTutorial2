package pear

import "fmt"

func testMap() {
	fmt.Println("pear")
	m := make(map[string]int)
	m["red"] = 100
	m["blue"] = 200
	if found, ok := m["red"]; ok {
		fmt.Printf("%v\n", found)
	}
}

func testList() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%v\n", i)
	}
}
func TestPear() {

}

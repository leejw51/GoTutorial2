package pear
import "fmt"


func testMap() {
	fmt.Println("pear")
	m := make(map[string]int)
	m["red"]=100
	m["blue"]=200
	if found, ok := m["red"]; ok {
		fmt.Printf("%v\n", found)
	}
}
func TestPear(){
	
}
package apple

import (
	"fmt"
	"time"
)

func test0(from string) {
	for i := 0; i < 5; i++ {
		fmt.Println(from, ":", i)
	}
}
func TestThread() {
	fmt.Println("test thread")
	test0("direct")
	go test0("go routine")
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
}

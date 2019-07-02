package main

import (
	"fmt"
	"test/sorting"
)

//	"time"

//	"reflect"
func testSort() {
	in := []int{6, 10, 1, 4, 7, 2, 8, 9, 3, 5}
	fmt.Println(in)
	sorting.QSorting().Sorting(in)
	fmt.Println(in)
}
func sum(c chan int) {
	for i := 0; i < 100; i++ {
		//time.Sleep(time.Second * 2)
		c <- i
	}
	fmt.Println("sum")
	//close(c)
}

func main() {
	fmt.Println("Good!")
}

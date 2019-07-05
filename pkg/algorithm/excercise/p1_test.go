package excercise

import (
	"fmt"
	"testing"
)

func TestMinSubarrayLen(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	k := int(7)
	len := MinSubarrayLen(a, k)
	fmt.Println(len)
}

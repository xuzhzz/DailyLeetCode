package __add_two_number

import (
	"testing"
)

func TestAddTwoNum(t *testing.T)  {
	l1 := build(1)
	l1.printl()
	l2 := build(9)
	l2.printl()
	//res := build(7, 0, 8)
	AddTwoNumbers(l1, l2).printl()
}

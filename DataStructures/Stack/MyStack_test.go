package Stack

import (
	"testing"
)

/* Declare an object of MyStack */
var stack MyStack

/* Method to get an empty stack object of MyStack */
func InitStack() *MyStack {

	if stack.items == nil {
		stack := MyStack{}
		stack.CreateStack()
	}

	return &stack
}

/* Test Pushing into stack */
func TestPush(t *testing.T) {
	stack := InitStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	if size := len(stack.items); size != 3 {
		t.Errorf("Wrong size count of the stack. Should be 3, but fot %d\n", size)
	}
}

/* Func test pop method */
func TestPop(t *testing.T) {

	if len(stack.items) == 0 {
		t.Errorf("Incorrect size of stack. Expected 3 but got 0")
	}

	if stack.IsEmpty() {
		t.Errorf("Error in IsEmpty() method. Expected false, got true ")
	}

	val_1 := stack.Pop()

	if *val_1 != 3 {
		t.Errorf("Incorrect popped value. Expected 3, got %d", val_1)
	}

	val_2 := stack.Pop()
	if *val_2 != 2 {
		t.Errorf("Incorrect popped value. Expected 2, got %d", val_1)
	}

	val_3 := stack.Pop()
	if *val_3 != 1 {
		t.Errorf("Incorrect popped value. Expected 1, got %d", val_1)
	}

	if !stack.IsEmpty() {
		t.Errorf("Error in 2nd IsEmpty(). Expected true but got false.")
	}

}

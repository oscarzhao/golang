/*************************************************************************
 *  @author:  Programming In Go, chapter 1.5
 *  @date created:  2015-04-10
 *  @purpose:    stack
 ************************************************************************/
package stack

import "errors"

type Stack []interface{}

func (stack Stack) Len() int {
	return len(stack)
}

func (stack Stack) Cap() int {
	return cap(stack)
}

func (stack *Stack) Push(x interface{}) {
	*stack = append(*stack, x)
}

func (stack Stack) Top() (interface{}, error) {
	if len(stack) == 0 {
		return nil, errors.New("cannot Top an empty stack")
	}
	return stack[len(stack)-1], nil
}

func (stack *Stack) Pop() (interface{}, error) {
	// create temporary theStack as a ref to *stack
	// for convenience, not necessary
	// in the text book page 17
	if len(*stack) == 0 {
		return nil, errors.New("cannot Pop an empty stack")
	}
	ret := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return ret, nil
}

package main

import (
	lls "github.com/emirpasic/gods/stacks/linkedliststack"
)

type CompositeIterator struct {
	pos   int
	stack *lls.Stack
}

func NewCompositeIterator(iterator IIterator) *CompositeIterator {
	c := &CompositeIterator{}
	c.stack = lls.New()
	c.stack.Push(iterator)
	return c
}

func (c *CompositeIterator) Next() interface{} {
	if c.HasNext() {
		iteratorS, _ := c.stack.Peek()
		iterator := iteratorS.(IIterator)
		component, ok := iterator.Next().(IMenuComponent)
		if ok {
			c.stack.Push(component.CreateIterator())
		}
		return component
	}
	return nil
}

func (c *CompositeIterator) HasNext() bool {
	if c.stack.Empty() {
		return false
	} else {
		iteratorS, _ := c.stack.Peek()
		iterator := iteratorS.(IIterator)
		if !iterator.HasNext() {
			c.stack.Pop()
			return c.HasNext()
		} else {
			return true
		}
	}
}

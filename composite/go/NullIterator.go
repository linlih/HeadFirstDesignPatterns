package main

type NullIterator struct {
}

func NewNullIterator() *NullIterator {
	return &NullIterator{}
}

func (n *NullIterator) Next() interface{} {
	return nil
}

func (n *NullIterator) HasNext() bool {
	return false
}

package main

type IIterator interface {
	HasNext() bool
	Next() interface{}
}

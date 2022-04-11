package main

type IMenu interface {
	CreateIterator() IIterator
}
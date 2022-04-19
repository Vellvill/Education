/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/

package main

type Human struct {
	age  int
	name string
	Action
}
type Action struct {
	run  bool
	fire bool
}

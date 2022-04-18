package main

type Human struct {
	age  int
	name string
	Action
}
type Action struct {
	run   bool
	fire  bool
	sleep bool
}

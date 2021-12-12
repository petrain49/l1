package main

import (
	"fmt"
)

type Human struct {
	head   string
	body   string
	foot   string
}

func (h *Human) printHead() {
	fmt.Println(h.head)
}

func (h *Human) printBody() {
	fmt.Println(h.body)
}

func (h *Human) printFoot() {
	fmt.Println(h.foot)
}

type Action struct {
	Human
	velocity int
	direction int
}

func (a *Action) printVelocity() {
	fmt.Println(a.velocity)
}

func (a *Action) printDirection() {
	fmt.Println(a.direction)
}

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/
func one() {
	h := Human{"Big", "Normal", "Short"}
	a := Action{h, 1, 2}

	h.printHead()
	h.printBody()
	h.printFoot()

	a.printVelocity()
	a.printDirection()

	a.printHead()
	a.printBody()
	a.printFoot()
}
// Copyright 2024 JOK Inc. All Rights Reserved.
// Author: easytojoin@163.com (jok)

package facade

import "fmt"

type SystemA struct {}

func (a *SystemA) OperationA() string  {
	return "A operation"
}

type SystemB struct {}

func (b *SystemB) OperationB() string {
	return "B operation"
}

type System struct {
	a SystemA
	b SystemB
}

func (s *System) Operation() {
	fmt.Println(s.a.OperationA())
	fmt.Println(s.b.OperationB())
}


func Facade() {
	s := System {a : SystemA{}, b : SystemB{}}
	s.Operation()
}
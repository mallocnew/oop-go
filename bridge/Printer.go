// Copyright 2024 JOK Inc. All Rights Reserved.
// Author: easytojoin@163.com (jok)

package bridge_pattern

import "fmt"

type Printer interface {
	PrintPdf(name string)
}

type Epson struct{}

func (e *Epson) PrintPdf(name string) {
	fmt.Printf("Epson PrintPdf: %s\n", name)
}

type Hp struct{}

func (h *Hp) PrintPdf(name string) {
	fmt.Printf("Hp PrintPdf: %s\n", name)
}

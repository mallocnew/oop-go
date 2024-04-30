// Copyright 2024 JOK Inc. All Rights Reserved.
// Author: easytojoin@163.com (jok)

package bridge_pattern

import "fmt"

type Computer interface {
	Print()
	SetPrinter(p Printer)
}

type Windows struct {
	printer Printer
}

func (w *Windows) Print() {
	fmt.Println("Windows Print")
	w.printer.PrintPdf("windows file")
}

func (w *Windows) SetPrinter(p Printer) {
	w.printer = p
}

type Mac struct {
	printer Printer
}

func (m *Mac) Print() {
	fmt.Println("Mac Print")
	m.printer.PrintPdf("mac file")
}

func (m *Mac) SetPrinter(p Printer) {
	m.printer = p
}

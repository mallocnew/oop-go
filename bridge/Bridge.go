// Copyright 2024 JOK Inc. All Rights Reserved.
// Author: easytojoin@163.com (jok)

package bridge_pattern

func PrintFile() {
	hpPrinter := &Hp{}
	epsonPrinter := &Epson{}

	mac := &Mac{}
	mac.SetPrinter(hpPrinter)
	mac.Print()
	mac.SetPrinter(epsonPrinter)
	mac.Print()

	win := &Windows{}
	win.SetPrinter(hpPrinter)
	win.Print()
	win.SetPrinter(epsonPrinter)
	win.Print()
}

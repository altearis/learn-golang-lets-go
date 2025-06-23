package main

import (
	"fmt"
)

func main() {

	var name string
	name = "Indra"

	fmt.Println(name)

	umur := 30
	fmt.Println(umur)

	fmt.Sprintf("Nama saya %s, umur saya %d tahun\n", name, umur)
	for i := 1; i <= 5; i++ {
		//TIP <p>To start your debugging session, right-click your code in the editor and select the Debug option.</p> <p>We have set one <icon src="AllIcons.Debugger.Db_set_breakpoint"/> breakpoint
		// for you, but you can always add more by pressing <shortcut actionId="ToggleLineBreakpoint"/>.</p>
		fmt.Println("i =", 100/i)
	}
}

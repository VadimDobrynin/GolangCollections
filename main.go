package main

import (
	"Collections/myCollections"
	"fmt"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	for i := 1; i <= 5; i++ {
		//TIP <p>To start your debugging session, right-click your code in the editor and select the Debug option.</p> <p>We have set one <icon src="AllIcons.Debugger.Db_set_breakpoint"/> breakpoint
		// for you, but you can always add more by pressing <shortcut actionId="ToggleLineBreakpoint"/>.</p>
		fmt.Println("i =", 100/i)
	}
	myCollections.SaYHallo()

	ml := myCollections.NewSinglyLinkedList(777)
	ml.Prepend(666)
	ml.Prepend(555)
	ml.Prepend(444)
	ml.Prepend(333)
	ml.Prepend(222)
	ml.Prepend(111)
	ml.Append(1)
	ml.Append(2)
	ml.Append(3)
	ml.Append(4)
	ml.Append(5)
	ml.Print()
	fmt.Println("------------------------------")
	ml.Print()
	fmt.Println("------------------------------")
	ml.RemoveByValue(3)
	ml.RemoveByValue(1)
	ml.RemoveByValue(5)
	ml.RemoveByValue(5)
	ml.RemoveByValue(111)
	fmt.Println("после удаления значения 3")
	ml.Print()
	ml.RemoveByIndex(0)
	ml.RemoveByIndex(6)
	fmt.Println("послеs удаления по индексу 0")
	ml.Print()

	fmt.Println("-----------------------------------")
	fmt.Println(ml.Find(333))
}

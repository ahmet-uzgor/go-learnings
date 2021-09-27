package main

import (
	"fmt"
)

type names []string

func main() {
	nameList := names{
		"ahmet",
		"mehmet",
		"cemil",
	}

	// nameList.printNames()
	nameList = newNameList()
	// nameList.printNames()
	nameList = nameList.appendNewName("hello ahmet")
	nameList.printNames()
}

func (n names) printNames() {
	for i, name := range n {
		fmt.Println(i, name)
	}
}

func newNameList() names {
	nameList := names{
		"name1",
		"name2",
		"name3",
	}

	return nameList
}

func (n names) appendNewName(name string) names {
	nameList := append(n, name)
	return nameList
}

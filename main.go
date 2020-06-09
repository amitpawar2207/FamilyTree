package main

import (
	"fmt"
	t "geektrust/familytree"
	f "geektrust/fileoperations"
	"os"
)

func main() {

	//Generating king shan's predefined family tree
	familyMembers, err := t.GenerateFamilyTree(t.FamilyMemberFilePath)
	if err != nil {
		fmt.Println(err)
	}

	if os.Args[1] == "" {
		fmt.Println("No filepath entered")
		os.Exit(0)
	}

	inputs, err1 := f.ReadInputFile(os.Args[1])
	if err1 != nil {
		fmt.Println(err1)
	} else {
		for _, input := range inputs {
			fmt.Println(t.ProcessInputs(input, familyMembers))
		}
	}
}

package one_level_of_indentation

import "fmt"

import (
	"github.com/rafos/object-calisthenics-in-go/one_level_of_indentation/after"
	"github.com/rafos/object-calisthenics-in-go/one_level_of_indentation/before"
)

func OneLevelOfIndentationExample() {
	fmt.Println("*** OneLevelOfIndentationExample ***")

	var data = [][]string{
		{" ", " ", " ", "x", " ", "o", " ", " ", "o", "x"},
		{" ", "o", "x", " ", " ", " ", " ", "x", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " ", " ", "x"},
		{"x", " ", " ", "x", "o", "x", " ", " ", " ", " "},
		{" ", "x", " ", "o", " ", " ", " ", "o", "x", " "},
		{" ", " ", " ", " ", " ", " ", "x", " ", " ", " "},
		{" ", " ", "o", " ", "x", " ", "o", " ", " ", " "},
		{"o", " ", " ", "x", " ", " ", " ", " ", "x", " "},
		{" ", "x", " ", " ", "o", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", "o", " ", " ", " ", "o"},
	}

	beforeBoard := before.NewBoard(data)
	fmt.Println("The Board (before)")
	fmt.Println(beforeBoard.Board())

	afterBoard := after.NewBoard(data)
	fmt.Println("The Board (after)")
	fmt.Println(afterBoard.Board())
}

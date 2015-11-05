package one_dot_per_line

import (
    "github.com/rafos/object-calisthenics-in-go/one_dot_per_line/before"
    "fmt"
    "github.com/rafos/object-calisthenics-in-go/one_dot_per_line/after"
)

func OneDotPerLineExample() {
    fmt.Println("*** OneDotPerLineExample ***")
    beforeBoard := before.NewBoard()
    fmt.Println("Before:", beforeBoard.BoardRepresentation())

    fmt.Println()

    afterBoard := after.NewBoard()
    fmt.Println("After:", afterBoard.BoardRepresentation())
}
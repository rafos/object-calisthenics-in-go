package first_class_collections

import (
	"fmt"

	"github.com/rafos/object-calisthenics-in-go/first_class_collections/after"
	"github.com/rafos/object-calisthenics-in-go/first_class_collections/before"
)

func FirstClassCollectionsExample() {
	fmt.Println("*** FirstClassCollectionsExample ***")

	beforePerson := before.NewPerson("Foo")
	beforePerson.AddFriend("Bar")
	beforePerson.AddFriend("Baz")
	fmt.Println("Before:", beforePerson)
	beforePerson.RemoveFriend("Bar")
	fmt.Println("Before:", beforePerson)
	fmt.Println("Before:", beforePerson.GetFriends())

	fmt.Println()

	afterPerson := after.NewPerson("Foo")
	afterPerson.AddFriend("Bar")
	afterPerson.AddFriend("Baz")
	fmt.Println("After:", afterPerson)
	afterPerson.RemoveFriend("Bar")
	fmt.Println("After:", afterPerson)
	fmt.Println("After:", afterPerson.GetFriends())
}

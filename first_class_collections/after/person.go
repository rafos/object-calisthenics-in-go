package after

import "fmt"

type person struct {
	name    string
	friends *friends
}

func NewPerson(name string) *person {
	return &person{
		name:    name,
		friends: NewFriends(),
	}
}

func (p *person) AddFriend(name string) {
	p.friends.Add(name)
}

func (p *person) RemoveFriend(name string) {
	p.friends.Remove(name)
}

func (p *person) GetFriends() *friends {
	return p.friends
}

func (p *person) String() string {
	return fmt.Sprintf("%s [%v]", p.name, p.friends)
}

package before

import "fmt"

type person struct {
	name    string
	friends []string
}

type friend struct {
	name string
}

func NewPerson(name string) *person {
	return &person{
		name:    name,
		friends: []string{},
	}
}

func (p *person) AddFriend(name string) {
	p.friends = append(p.friends, name)
}

func (p *person) RemoveFriend(name string) {
	new := []string{}
	for _, friend := range p.friends {
		if friend != name {
			new = append(new, friend)
		}
	}
	p.friends = new
}

func (p *person) GetFriends() []string {
	return p.friends
}

func (p *person) String() string {
	return fmt.Sprintf("%s %v", p.name, p.friends)
}

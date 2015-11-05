package after

import (
	"strings"
)

type friends struct {
	data []string
}

func NewFriends() *friends {
	return &friends{
		data: []string{},
	}
}

func (f *friends) Add(name string) {
	f.data = append(f.data, name)
}

func (f *friends) Remove(name string) {
	new := []string{}
	for _, friend := range f.data {
		if friend != name {
			new = append(new, friend)
		}
	}
	f.data = new
}

func (f *friends) String() string {
	return strings.Join(f.data, " ")
}

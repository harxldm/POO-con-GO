package main

import "fmt"

type user struct {
	name      string
	age       int
	user_name string
	password  string
	comp      map[uint]string
}
func Newuser(name string, age int, user_name string) *user {
	if age == 0 {
		age = 18
	}
	return &user {
		name: name,
		age: age,
		user_name: user_name,
	}
}

func (User *user) changename(name2 string) {
	User.name = name2
	//fmt.Println(User.name)
}
func (User user) printclases() {
	text := "Las competencias son: "
	for _, clases := range User.comp {
		text += clases + ", "
	}
	fmt.Println(text[:len(text)-2])
}
func main() {
	Go := user {
		name:      "Harold",
		age:       19,
		user_name: "harxldm",
		password:  "1234",
		comp: map[uint]string{
			1: "creativo",
			2: "Honesto",
		},
	}

	Go.changename("Paula")
	fmt.Println(Go.name)
}

package main

import "fmt"

type Cow struct {
	name string
}
type Bird struct {
	name string
}
type Snake struct {
	name string
}
type Animal interface {
	Move()
	Eat()
	Speak()
}

func (c Cow) Eat() {
	fmt.Println("Grass")
}
func (c Cow) Move() {
	fmt.Println("Walk")
}
func (c Cow) Speak() {
	fmt.Println("Moo")
}
func (b Bird) Eat() {
	fmt.Println("Worms")
}
func (b Bird) Move() {
	fmt.Println("Fly")
}
func (b Bird) Speak() {
	fmt.Println("Peep")
}
func (s Snake) Eat() {
	fmt.Println("Mice")
}
func (s Snake) Move() {
	fmt.Println("Slither")
}
func (s Snake) Speak() {
	fmt.Println("Hsss")
}
func CreateAnimal(aname, info string, s map[string]Animal) {
	var na Animal
	switch info {
	case "cow":
		na = Cow{name: aname}

	case "bird":
		na = Bird{name: aname}
	case "snake":
		na = Snake{name: aname}
	}
	s[aname] = na
	fmt.Println("Created it")
}
func QueryAnimal(name, job string, s map[string]Animal) {
	switch job {
	case "eat":
		s[name].Eat()
	case "move":
		s[name].Move()
	case "speak":
		s[name].Speak()
	}
}

func main() {
	var command, name, info string
	animals := map[string]Animal{}

	for {
		fmt.Print(">")
		fmt.Scanln(&command, &name, &info)

		switch command {
		case "newanimal":
			CreateAnimal(name, info, animals)
		case "query":
			QueryAnimal(name, info, animals)
		}
	}

}

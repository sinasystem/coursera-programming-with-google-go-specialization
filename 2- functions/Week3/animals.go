package main

import (
	"fmt"
)

type animal struct {
	food, locomotion, noise string
}

func (an animal) Eat() {
	fmt.Println(an.food)
}
func (an animal) Move() {
	fmt.Println(an.locomotion)
}
func (an animal) Speak() {
	fmt.Println(an.noise)
}

func main() {

	cow := animal{
		food:       "grass",
		locomotion: "walk",
		noise:      "moo",
	}
	bird := animal{
		food:       "worms",
		locomotion: "fly",
		noise:      "peep",
	}
	snake := animal{
		food:       "mice",
		locomotion: "slither",
		noise:      "hsss",
	}
	animalMap := map[string]animal{
		"cow":   cow,
		"bird":  bird,
		"snake": snake,
	}
	var name, task string
	for {
		fmt.Print("> ")
		fmt.Scan(&name)
		fmt.Scan(&task)

		switch task {
		case "eat":
			animalMap[name].Eat()
		case "move":
			animalMap[name].Move()
		case "speak":
			animalMap[name].Speak()
		}
	}
}

package main

import "fmt"

// Animal interface defines a method for producing a sound.
type Animal interface {
    MakeSound() string
}

// Dog type implements the Animal interface.
type Dog struct{}

func (d Dog) MakeSound() string {
    return "Woof!"
}

func PrintAnimalSound(animal Animal) {
    fmt.Println("Animal sound:", animal.MakeSound())
}

func main() {
    dog := Dog{}

    PrintAnimalSound(dog) // Output: Animal sound: Woof!
}

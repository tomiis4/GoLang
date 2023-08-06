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

// Cat type implements the Animal interface.
type Cat struct{}

func (c Cat) MakeSound() string {
    return "Meow!"
}

func PrintAnimalSound(animal Animal) {
    fmt.Println("Animal sound:", animal.MakeSound())
}

func main() {
    dog := Dog{}
    cat := Cat{}

    PrintAnimalSound(dog) // Output: Animal sound: Woof!
    PrintAnimalSound(cat) // Output: Animal sound: Meow!
}

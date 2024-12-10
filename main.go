package main

import "fmt"

// Структура Human
type Human struct {
	Name string
	Age  int
}

// Метод структуры Human
func (h *Human) Speak() {
	fmt.Printf("Hi, my name is %s and I am %d years old.\n", h.Name, h.Age)
}

// Структура Action
type Action struct {
	Human // Встраивание структуры Human
	Job   string
}

// Метод структуры Action
func (a *Action) IntroduceAndWork() {
	fmt.Printf("Hi, my name is %s, I am %d years old, and I work as a %s.\n", a.Name, a.Age, a.Job)
}

func main() {
	// Создаем экземпляр структуры Action
	action := Action{
		Human: Human{Name: "John", Age: 30},
		Job:   "Programmer",
	}

	// Вызываем метод из структуры Human
	action.Speak()

	// Вызываем метод из структуры Action
	action.IntroduceAndWork()
}

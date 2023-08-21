package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (p *Human) SetName(name string) {
	p.Name = name
}
func (p *Human) SetAge(age int) {
	p.Age = age
}
func (p Human) Represent() {
	fmt.Println("My name is ", p.Name, "I am ", p.Age)
}

type Action struct {
	Human
	Name string
}

func (a Action) GetName() string {
	return a.Name
}

// переопределение
func (a Action) Represent() {
	fmt.Printf("Action %s of %s\n", a.Name, a.Human.Name)
}

func main() {

	human := Human{}
	human.SetAge(22)
	human.SetName("Ivan")
	human.Represent()
	Action := Action{
		Name:  "living",
		Human: human,
	}
	Action.GetName()
	//Наследуемый метод переопределен, поэтому вызовется для Action
	Action.Represent()
	// Наследуемый метод не переопределен, следовательно, метод изменит Name структуры Human, а не Action
	Action.SetName("Oleg")
	//Также вызовется для Action
	Action.Represent()

}

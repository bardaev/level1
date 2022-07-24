package tasks

func Task1() {
	var h Human = Human{Name: "Alex", Action: Action{}}

	PrintResult(1, h.DoSmth())
}

type Human struct {
	Name string
	Action
}

type Action struct {
}

func (a *Action) DoSmth() string {
	return "Action"
}

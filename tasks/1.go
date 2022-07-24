package tasks

// Добавив структуру Action в структуру Human, произошло встраивание метода
// Метод DoSmth доступен как h.DoSmth(), так и h.Action.DoSmth()
func Task1() {
	var h Human = Human{Name: "Alex", Action: Action{}}

	PrintResult(1, "Встраивание метода", h.DoSmth())
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

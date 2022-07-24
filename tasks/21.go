package tasks

import "fmt"

func Task21() {
	var user *user = &user{}
	var linux *linux = &linux{}
	var windows *windows = &windows{}

	PrintResult(21.1, "Вызов программы windows", nil)
	user.RunProgramOnComputer(windows)

	var linuxAdapterWine *linuxAdapterWine = &linuxAdapterWine{
		linuxComp: linux,
	}

	PrintResult(21.2, "Вызов программы windows на компьютере linux через адаптер wine", nil)
	user.RunProgramOnComputer(linuxAdapterWine)
}

type program interface {
	RunProgram()
}

type windows struct{}

func (w *windows) RunProgram() {
	fmt.Println("Run windows program")
}

type linux struct{}

func (l *linux) RunLinuxProgram() {
	fmt.Println("Run linux program")
}

type linuxAdapterWine struct {
	linuxComp *linux
}

func (lin *linuxAdapterWine) RunProgram() {
	lin.linuxComp.RunLinuxProgram()
}

type user struct{}

func (u *user) RunProgramOnComputer(p program) {
	p.RunProgram()
}

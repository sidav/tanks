package main

type tankAi struct {
	chanceToRotateOneFrom int
}

func initSimpleTankAi() *tankAi {
	return &tankAi{chanceToRotateOneFrom: 10}
}

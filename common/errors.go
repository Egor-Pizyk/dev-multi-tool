package common

type ErrorState int

const (
	SomethingWrong ErrorState = iota
)

var ErrorsEnum = map[ErrorState]string{
	SomethingWrong: "Something went wrong...",
}

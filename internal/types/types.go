package types

type Logger interface {
	Warn(...any)
	Error(...any)
	Errorf(string, ...any)
}

package exception

type Exception struct {
	Code string
	Human string
	Values map[string]string
}

func (e Exception) human() string {
	return e.Human
}

func (e Exception) Error() string {
	return e.human()
}

package support

type Exception struct {
	Message string
	Field   string
}

func (e Exception) Error() string {
	return "todo test join message with field"
}

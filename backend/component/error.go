package component

type NotFoundError struct{}

func NewNotFoundError() *NotFoundError {
	return &NotFoundError{}
}

func (e *NotFoundError) Error() string {
	return "not found"
}

func IsNotFoundError(err error) bool {
	_, ok := err.(*NotFoundError)
	return ok
}

package models

type notFound struct {
	Message string
}

func NotFoundError(message string) error {
	return &notFound{Message: message}
}

func (err *notFound) Error() string {
	return err.Message
}

type validation struct {
	Message string
}

func ValidationError(message string) error {
	return &validation{Message: message}
}

func (err *validation) Error() string {
	return err.Message
}

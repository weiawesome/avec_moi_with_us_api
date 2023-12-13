package utils

type Error struct {
	Error string `json:"error"`
}

type ExistError struct {
	Message string
}

func (err ExistError) Error() string {
	return err.Message
}

type AuthError struct {
	Message string
}

func (err AuthError) Error() string {
	return err.Message
}

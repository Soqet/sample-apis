package db

type DbError struct {
	message string
}

func (err *DbError) Error() string {
	return err.message
}

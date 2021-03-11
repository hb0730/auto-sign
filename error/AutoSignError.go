package error

import "fmt"

type AutoSignError struct {
	Module  string
	Message string
}

var AsError error = &AutoSignError{}

func (e *AutoSignError) Error() string {
	return fmt.Sprintf("Parameter verification failed: %v %v \n", e.Module, e.Message)
}

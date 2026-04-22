package service

import (
	"fmt"
	"errors"
)

// ValidationError 表示业务校验失败；指针接收者便于 errors.As。
type ValidationError struct {
	Field   string
	Message string
}

var dataError = errors.New("message has issue")

func checkField() error{
	return fmt.Errorf("The request: %w", dataError)
}

func (e *ValidationError) Error() string {
	

	// First try by me
	// if e.Field == "query"{
	// 	fmt.Println("error is", errors.Is(errors.New("message has issue"), dataError))
	// }

	// 2rd: 有解释
	err := checkField();
	fmt.Println("direct check: ", err == dataError)
	fmt.Println("use is: ", errors.Is(err, dataError))

	return fmt.Sprintf("validation failed on %s: %s", e.Field, e.Message)
}


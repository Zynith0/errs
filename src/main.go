package error

import "fmt"

func Error(err error, msg string) {
	if err != nil {
		fmt.Println(msg, err)
	}
}

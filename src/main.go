package errs

import "fmt"

func Error(msg string, err error) {
	if err != nil {
		fmt.Println(msg, err)
	}
}

func SkillIssue(err error) {
	if err != nil {
		fmt.Println("skill issue", err)
	}
}

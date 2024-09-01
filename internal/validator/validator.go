package validator

import (
	"fmt"
	"regexp"
)

type Validator struct {
	Regex string
}

func (v *Validator) CheckInput(message string) (bool, error) {
	re, err := regexp.Compile(v.Regex)
	if err != nil {
		return false, fmt.Errorf("invalid regex pattern: %v", err)
	}
	match := re.MatchString(message)
	return match, nil
}

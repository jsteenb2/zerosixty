package company

import (
	"fmt"
	"strconv"
)

type PhoneNumber string

func (p PhoneNumber) HumanReadable() string {
	return fmt.Sprintf("(%s) %s-%s", p[0:3], p[3:6], p[6:])
}

func (p PhoneNumber) AreaCode() int {
	ac, _ := strconv.Atoi(string(p[0:3]))
	return ac
}

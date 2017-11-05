package company

import (
	"fmt"
)

type Address struct {
	ZipCode int
	Street  string
	Number  int
	City    string
	State   string
}

func (a Address) MailFormat() string {
	return fmt.Sprintf("\n%d %s\n%s, %s %d", a.Number, a.Street, a.City, a.State, a.ZipCode)
}

func (Address) WhoAmI() string {
	return "address"
}

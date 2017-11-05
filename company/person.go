package company

import (
	"fmt"
)

type Person string

func (p *Person) Greet() string {
	//return fmt.Sprintf("Hello, I'm %s", p) // prints out the address of person
	return fmt.Sprintf("Hello, I'm %s", *p) // * dereferences the pointer to actual value
}

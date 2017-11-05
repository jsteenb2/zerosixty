package company_test

import (
	"testing"

	c "github.com/jsteenb2/zerosixty/company" // can alias package names, not recommended for production
)

/*
	Interfaces are where go really shines
	Go is a structurally type language, and all interfaces are satisfied implicitly
	bye bye implements keyword

	You can achieve polymorphic **behavior** using interfaces

	Interfaces are a value-less type
	Interfaces are a 2 word data structure
	1st word: is pointer to interface table (iTable)
		iTable implements the following:
			describes type of value being stored (in 2nd word)
			pointer to code associated with concrete methods in interface
	2nd word: is pointer to a shared or copy of the type implementing the interface

*/

type WhoAmIer interface {
	WhoAmI() string
}

func TestInterfaceSatisfiedImplicitly(t *testing.T) {
	whoamiSlice := []WhoAmIer{
		c.Contact{},
		c.Address{},
		c.PhoneNumber(5558675309),
		c.Company{},
		c.Person("ex"),
	}

	for _, v := range whoamiSlice {
		t.Log(v.WhoAmI())
	}

	// the line below does not work
	// the compiler knows nothing of the implementation of Address struct
	// based on the WhoAmIer interface
	// t.Log(whoamiSlice[1].MailFormat())

}

func isCompany(who WhoAmIer) bool {
	return who.WhoAmI() == "company"
}

func TestInterfaceAsInputType(t *testing.T) {
	whoamiSlice := []WhoAmIer{
		c.Contact{},
		c.Address{},
		c.PhoneNumber(5558675309),
		c.Company{},
		c.Person("ex"),
	}

	for idx, v := range whoamiSlice {
		if isCompany(v) {
			t.Logf("company found at idx: %d", idx)
		}
	}
}

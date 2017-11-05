package company_test

import (
	"testing"

	"github.com/jsteenb2/zerosixty/company"
)

// pointers (pass-by-reference) are likely going to cause your value to escape to the heap
// pass-by-value almost always does not escape and lives on the stack (no GC)

func TestStructPointers(t *testing.T) {
	threeve := newTestCompany()
	tad := &threeve.PointOfContact[0] // pre-pending with & returns pointer to value
	t.Log(tad.Greet())

	address := company.Address{63102, "Third Street", 900, "Shreevesport", "CA"}
	deliveryCosts := threeve.Delivery(address) // automatically converts value semantic to pointer semantics

	t.Logf("delivery cost: $%0.2f", deliveryCosts)
}

func newTestCompany() company.Company {
	address := company.Address{12345, "3rd Street", 3, "Shreevesport", "DE"}

	contact := company.Contact{
		PointOfContact: []company.Person{"Tad", "Pole"},
		PhoneNumber:    company.PhoneNumber(5558675309),
		Address:        address,
	}

	return company.Company{
		Name:    "Threeve",
		Contact: contact,
	}
}

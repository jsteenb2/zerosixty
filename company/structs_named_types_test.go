package company_test

import (
	"testing"

	"github.com/jsteenb2/zerosixty/company"
)

func TestNamedTypeMethods(t *testing.T) {
	// named types can have methods

	local := company.PhoneNumber("3005551234")

	t.Logf("HR: %s", local.HumanReadable())
	t.Logf("area code: %d", local.AreaCode())
}

func TestStructTypeMethods(t *testing.T) {
	// struct types have methods as well

	a := company.Address{63102, "Spruce Street", 900, "St. Louis", "MO"}

	t.Log(a.MailFormat())
}

func TestStructFieldPromotion(t *testing.T) {
	// a struct with an embedded struct(s) can access
	// the embedded struct fields directly
	// field promotion happens unless there is name collision

	address := company.Address{
		ZipCode: 63102,
		Street:  "Spruce Street",
		Number:  900,
	}

	contact := company.Contact{
		PointOfContact: []company.Person{"Pops", "RedGoatTea", "Trigger"},
		PhoneNumber:    "3146782200",
		Address:        address,
	}

	async := company.Company{"Asynchrony", contact}

	t.Log(async.Street) // three levels deep
	t.Log(async.PhoneNumber)
}

func TestStructEmbeddedNameCollision(t *testing.T) {
	address := company.Address{63102, "Spruce Street", 900, "St. Louis", "MO"}

	contact := company.Contact{
		PointOfContact: []company.Person{"RedGoatTea", "Trigger"},
		PhoneNumber:    "3146782200",
		Address:        address,
	}

	async := company.Company{"Asynchrony", contact}

	t.Log(async.WhoAmI())
	t.Log(contact.WhoAmI())
	t.Log(address.WhoAmI())
}

func TestStructEmbeddedMethodPromotion(t *testing.T) {
	// embedded type methods are also promoted same as fields

	address := company.Address{63102, "Spruce Street", 900, "St. Louis", "MO"}

	contact := company.Contact{
		PointOfContact: []company.Person{"Pops", "RedGoatTea", "Trigger"},
		PhoneNumber:    "3146782200",
		Address:        address,
	}

	async := company.Company{"Asynchrony", contact}

	t.Log(async.MailFormat())
	t.Log(contact.MailFormat())
	t.Log(address.MailFormat())
}

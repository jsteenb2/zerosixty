package company

type Contact struct {
	PointOfContact []Person
	PhoneNumber    PhoneNumber
	Address
}

func (Contact) WhoAmI() string {
	return "contact"
}

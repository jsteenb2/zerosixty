package company

type Company struct {
	Name string
	Contact
}

func (c *Company) Delivery(address Address) float64 {
	cost := 0.0
	if c.State != address.State {
		cost += 5.00
	}
	if c.City != address.City {
		cost += 2.00
	}
	if c.ZipCode != address.ZipCode {
		cost += 0.5
	}
	return cost
}

func (c Company) Founder() Person {
	return c.PointOfContact[0]
}

func (Company) WhoAmI() string {
	return "company"
}

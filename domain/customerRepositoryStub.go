package domain


// adapter
type CustomerRepositoryStub struct {
	customers []Customer
}

// FindAll customers
func (c *CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return c.customers, nil
}

// creates an instance of customerrepository [adapter]
func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1", Name: "jose", City: "Cidade tal", ZipCode: "Tal", BirthdayDate: "1970-01-01", Status: "1"},
		{Id: "2", Name: "Maria", City: "Cidade tal", ZipCode: "Ta1l", BirthdayDate: "1970-02-03", Status: "1"},
	}
	return CustomerRepositoryStub{
		customers: customers,
	}
}
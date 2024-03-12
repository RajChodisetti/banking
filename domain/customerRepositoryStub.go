package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRespositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Ashish", "Hyderabad", "50001", "2000-01-01", "1"},
		{"1002", "Nehra", "Delhi", "200001", "1999-06-06", "1"},
	}
	return CustomerRepositoryStub{customers}

}

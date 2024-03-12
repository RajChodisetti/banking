package app

import (
	"bankingweb/service"
	"encoding/json"
	"encoding/xml"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"Zipcode" xml:"zip"`
}

//	func greet(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprint(w, "Hello World")
//	}
type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// customers := []Customer{
	// 	{"Ashish", "News Delhi", "50001"},
	// 	{"Rob", "Old Delhi", "50002"},
	//}
	customers, _ := ch.service.GetAllCustomer()

	if r.Header.Get("Content-TYpe") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

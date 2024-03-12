package app

import (
	"bankingweb/domain"
	"bankingweb/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	//mux := http.NewServeMux()
	router := mux.NewRouter()

	//wiring
	//ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRespositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRespositoryDb())}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	// router.HandleFunc("/greet", greet)

	// //request matcher regex followed after :
	// router.HandleFunc("/customers/{customer_id:[0-9]+}", getAllCustomer)

	// //method matcher .methods
	// router.HandleFunc("/customers", createCustomers).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

// func getAllCustomer(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	fmt.Fprint(w, vars["customer_id"])

// }
// func createCustomers(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Post request recieved")
// }

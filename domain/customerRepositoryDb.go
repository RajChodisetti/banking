package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {
	findallcustomers := "select id, name, city, zipcode, dateofbirth, status from customer"
	rows, err := d.client.Query(findallcustomers)
	if err != nil {
		log.Println("Error encountered while querying: ", err)
		return nil, err
	}
	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
		if err != nil {
			log.Println("error while mapping it to struct")
		}
		customers = append(customers, c)

	}
	return customers, nil
}

func NewCustomerRespositoryDb() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "root:new_password@tcp(localhost:3306)/db")
	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client}
}

package crud

import (
	"errors"
	"time"
)

//IndexPage represents the content of the index page, available on "/"
//The index page shows a list of all customers stored on db
type IndexPage struct {
	AllCustomers []Customer
}

var ErrorNoRecord = errors.New("MOdels: no matching record Found")

type CustomerPage struct {
	TargetCustomer Customer
}

//Customer represents a customer object
type Customer struct {
	ID      int
	FName   string
	LName   string
	DOB     time.Time
	Gender  string
	Email   string
	Address string
}

func (c Customer) DOBstr() string {
	return c.DOB.Format("2006-01-02")
}

//ErrorPage represents shows an error message, available on "/customer.html"
type ErrorPage struct {
	ErrorMsg string
}

/*package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record found")

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
} */

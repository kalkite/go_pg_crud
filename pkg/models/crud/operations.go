package crud

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
)

type CrudMOdel struct {
	DB *sql.DB
}

func (m *CrudMOdel) GetCustomer(customerID int) (Customer, error) {
	//Fetch
	res := Customer{}

	var id int
	var fname string
	var lname string
	var dob pq.NullTime
	var gender string
	var email string
	var address string

	err := m.DB.QueryRow(`SELECT id, fname, lname, dob, gender, email, address FROM customers where id = $1`, customerID).Scan(&id, &fname, &lname, &dob, &gender, &email, &address)
	if err == nil {
		res = Customer{ID: id, FName: fname, LName: lname, DOB: dob.Time, Gender: gender, Email: email, Address: address}
	}

	return res, err
}
func (m *CrudMOdel) AllCustomers() ([]Customer, error) {

	customers := []Customer{}

	rows, err := m.DB.Query(`SELECT id, fname, lname, dob, gender, email, address FROM customers order by id`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var fname string
		var lname string
		var dob pq.NullTime
		var gender string
		var email string
		var address string

		err = rows.Scan(&id, &fname, &lname, &dob, &gender, &email, &address)
		if err != nil {
			return customers, err
		}

		currentCustomer := Customer{ID: id, FName: fname, LName: lname, Gender: gender, Email: email, Address: address}

		if dob.Valid {
			currentCustomer.DOB = dob.Time
		}

		customers = append(customers, currentCustomer)
	}

	return customers, err
}

func (m *CrudMOdel) SearchCustomers(name string) ([]Customer, error) {
	//Search
	customers := []Customer{}

	rows, err := m.DB.Query(`SELECT id, fname, lname, dob, gender, email, address FROM customers WHERE fname=$1 OR lname=$1`, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var fname string
		var lname string
		var dob pq.NullTime
		var gender string
		var email string
		var address string

		err = rows.Scan(&id, &fname, &lname, &dob, &gender, &email, &address)
		if err != nil {
			return customers, err
		}

		currentCustomer := Customer{ID: id, FName: fname, LName: lname, Gender: gender, Email: email, Address: address}
		if dob.Valid {
			currentCustomer.DOB = dob.Time
		}

		customers = append(customers, currentCustomer)
	}

	return customers, err
}

func (m *CrudMOdel) InsertCustomer(fname, lname string, dob time.Time, gender string, email string, address string) (int, error) {
	//Create
	var customerID int
	err := m.DB.QueryRow(`INSERT INTO customers(fname, lname, dob, gender, email, address) VALUES($1, $2, $3, $4, $5, $6) RETURNING id`, fname, lname, dob, gender, email, address).Scan(&customerID)

	if err != nil {
		return 0, err
	}

	// fmt.Printf("Last inserted ID: %v\n", customerID)
	return customerID, err
	// return -1, err
}

func (m *CrudMOdel) UpdateCustomer(id int, fname, lname string, dob time.Time, gender string, email string, address string) (int, error) {
	//Update
	res, err := m.DB.Exec(`UPDATE customers set fname=$1, lname=$2, dob=$3, gender=$4, email=$5, address=$6 where id=$7 RETURNING id`, fname, lname, dob, gender, email, address, id)
	if err != nil {
		return 0, err
	}

	rowsUpdated, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowsUpdated), err
	// fmt.Print(rowsUpdated)
	// return -1, err
}

func (m *CrudMOdel) RemoveCustomer(customerID int) (int, error) {
	//Delete
	res, err := m.DB.Exec(`delete from customers where id = $1`, customerID)
	if err != nil {
		return 0, err
	}

	rowsDeleted, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowsDeleted), nil
	// print(rowsDeleted)
	// return -1, err
}

/*type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	stmt := `INSERT INTO snippets(title, content, created, expires)
			VALUES (?, ?, ?, ?, ?);`

	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}
	id, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(id), err
}*/

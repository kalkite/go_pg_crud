package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"

	"wallester.eu/snippetbox/pkg/models/crud"
)

func (app *application) handleSaveCustomer(w http.ResponseWriter, r *http.Request) {

	// if r.URL.Path != "/" {
	// 	app.notFound(w)
	// 	return
	// }

	var id = 0
	var err error
	r.ParseForm()
	params := r.PostForm
	idStr := params.Get("id")

	if len(idStr) > 0 {
		id, err = strconv.Atoi(idStr)
		if err != nil {
			app.renderErrorPage(w, err)
			return
		}
	}

	fname := strings.Title(strings.ToLower(params.Get("fname")))
	lname := strings.Title(strings.ToLower(params.Get("lname")))
	gender := params.Get("gender")
	email := params.Get("email")
	address := params.Get("address")

	dobStr := params.Get("dob")
	var dob time.Time

	if len(dobStr) > 0 {
		dob, err = time.Parse("2006-01-02", dobStr)
		if err != nil {
			app.renderErrorPage(w, err)
		}
	}

	if id == 0 {
		_, err = app.operations.InsertCustomer(fname, lname, dob, gender, email, address)
	} else {
		_, err = app.operations.UpdateCustomer(id, fname, lname, dob, gender, email, address)
	}

	if err != nil {
		app.renderErrorPage(w, err)
	}

	http.Redirect(w, r, "/", 302)

	/*title := "O snail"
	content := "hello"
	expires := "7"
	id, err := fmt.Fprint(1)

	//app.operations.get

	//id, err = app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}
	// Redirect the user to the relevant page for the snippet.
	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther) */

}

func (app *application) handleListCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := app.operations.AllCustomers()
	if err != nil {
		app.renderErrorPage(w, err)
		return
	}

	buf, err := ioutil.ReadFile("./ui/html/index.html")
	if err != nil {
		app.renderErrorPage(w, err)
		return
	}

	var page = crud.IndexPage{AllCustomers: customers}
	indexPage := string(buf)
	t := template.Must(template.New("indexPage").Parse(indexPage))
	t.Execute(w, page)
}

func (app *application) handleSearchCustomer(w http.ResponseWriter, r *http.Request) {
	var customers []crud.Customer
	var err error
	r.ParseForm()
	params := r.PostForm
	name := params.Get("searchname")
	name = strings.Title(strings.ToLower(name))

	if len(name) == 0 {
		customers, err = app.operations.AllCustomers()
	} else {
		customers, err = app.operations.SearchCustomers(name)
	}

	if err != nil {
		app.renderErrorPage(w, err)
		return
	}

	buf, err := ioutil.ReadFile("./ui/html/index.html")
	if err != nil {
		app.renderErrorPage(w, err)
		return
	}

	var page = crud.IndexPage{AllCustomers: customers}
	indexPage := string(buf)
	t := template.Must(template.New("indexPage").Parse(indexPage))
	err = t.Execute(w, page)
	if err != nil {
		app.serverError(w, err)
	}

}

func (app *application) handleViewCustomer(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	idStr := params.Get("id")

	var currentCustomer = crud.Customer{}
	currentCustomer.DOB = time.Now()

	if len(idStr) > 0 {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			app.renderErrorPage(w, err)
			return
		}

		currentCustomer, err = app.operations.GetCustomer(id)
		if err != nil {
			app.renderErrorPage(w, err)
			return
		}
	}

	buf, err := ioutil.ReadFile("./ui/html/customer.html")
	if err != nil {
		app.renderErrorPage(w, err)
		return
	}

	var page = crud.CustomerPage{TargetCustomer: currentCustomer}
	customerPage := string(buf)
	t := template.Must(template.New("customerPage").Parse(customerPage))
	t.Execute(w, page)

}

func (app *application) handleDeleteCustomer(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	idStr := params.Get("id")

	if len(idStr) > 0 {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			app.renderErrorPage(w, err)
			return
		}

		n, err := app.operations.RemoveCustomer(id)
		if err != nil {
			app.renderErrorPage(w, err)
			return
		}

		fmt.Printf("Rows removed: %v\n", n)
	}
	http.Redirect(w, r, "/", 302)
}

func (app *application) renderErrorPage(w http.ResponseWriter, errorMsg error) {
	buf, err := ioutil.ReadFile("./ui/html/error.html")

	if err != nil {
		log.Printf("%v\n", err)
		fmt.Fprintf(w, "%v\n", err)
		return
	}

	var page = crud.ErrorPage{ErrorMsg: errorMsg.Error()}
	errorPage := string(buf)
	t := template.Must(template.New("errorPage").Parse(errorPage))
	t.Execute(w, page)

}

package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	
	
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.handleListCustomers)
	mux.HandleFunc("/customer.html", app.handleViewCustomer)
	mux.HandleFunc("/save", app.handleSaveCustomer)
	mux.HandleFunc("/delete", app.handleDeleteCustomer)
	mux.HandleFunc("/search", app.handleSearchCustomer)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}

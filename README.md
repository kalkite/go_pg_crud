# go_pg_crud

This is a sample web application for CRUD operations by using go and postgreSQL. 

# Project Name

Brief description of your project.

## Table of Contents

- [Project Structure](#project-structure)
- [Description](#description)
- [Installation](#installation)
- [Usage](#usage)
- [Screenshots](#screenshots)
- [License](#license)

## Project Structure



![2.png](https://github.com/chittibc/go_pg_crud/blob/master/2.png)


![1.png](https://github.com/chittibc/go_pg_crud/blob/master/1.png)

# How to Use

* Install PostgreSQL on your system. You can skip this step if already installed.
* Clone this repository.
* Execute the `customers.sql` file into your PostgreSQL client. This will import sample database and tables that will be used for this example.
* Modify `main.go` file, You must configure the PostgreSQL database connection. Change your postgreSQL username and passoword in db connectin. 
* go to your project and Run `$go run cmd/web/*` or if you want to add command line interface port you give , `$go run cmd/web/* :9999` 
* Open your web browser, and navigate to `http://localhost:4000`
* if you get any error in importing pacakges, `go mod init project_name`, 

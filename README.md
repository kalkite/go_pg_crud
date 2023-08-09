# Go Web Application backend.  

## Description

This project is a Go web application designed for creating and managing code snippets. It offers features like user registration (signup), authentication and authorization (login), and a dynamic home page for viewing and organizing snippets. The application leverages Go's strong concurrency support.

## Features

- User signup and login functionality.
- Database-driven responses for storing and retrieving user data and snippets using Mysql. 
- Dynamic HTML templates for rendering dynamic content.
- Middleware for security headers. 
- Advanced routing to handle different URL paths and HTTP methods.
- Stateful HTTP interaction to maintain user sessions and data.
- Enhanced security measures:
  - Running HTTPS servers for secure communication.
  - Connection timeouts to prevent idle connections.
  - Configuring HTTPS settings for improved security.
  - User authentication and context-based authorization.
- Use of request context for managing authentication and authorization.


![signup.png](https://github.com/kalkite/go_web_application/blob/master/signup.png)

![login.png](https://github.com/kalkite/go_web_application/blob/master/login.png)

![create_snippet.png](https://github.com/kalkite/go_web_application/blob/master/create_snippet.png)

# How to Use

* Install PostgreSQL on your system. You can skip this step if already installed.
* Clone this repository.
* Execute the `customers.sql` file into your PostgreSQL client. This will import sample database and tables that will be used for this example.
* Modify `main.go` file, You must configure the PostgreSQL database connection. Change your postgreSQL username and passoword in db connectin. 
* go to your project and Run `$go run cmd/web/*` or if you want to add command line interface port you give , `$go run cmd/web/* :9999` 
* Open your web browser, and navigate to `http://localhost:4000`
* if you get any error in importing pacakges, `go mod init project_name`, 

#  Events Management REST API
 An event management pilot project that allows basic creation of events, users and registering using JWT tokens and written in Go.

## 📁 Folder Structure

```
├── api-test            # REST requests for databases
├── db			# Database script to create db using SQLite3
├── middlewares		# User authentication scripts
├── models		# Scripts to handle queries
├── routes   		# Routing scripts for each type of operations
├── screenshots		# Screenshots for future use
├── utils		# Utility scripts for hashing/jwt
├── .gitignore		# 
├── LICENSE 		#
├── README 		# 
├── api.db 		# SQLite db file for testing
├── go.mod		# Project dependencies
├── go.sum 		# 
├── main.go 		# Main file for running 

```
## ⭐ Screenshots
![console screenshot for Gin](https://github.com/isaacchunn/go-rest-api/blob/main/screenshots/console.png)

## 📌 Setting Up

> [!NOTE]
The version of Go used is 1.22.2.
Requirements:
1. Gin Framework
2. JWT
3. SQLite3
4. crypto/bcrypt

1. Open terminal in root directory.
2. Install dependencies using go commands
	```go
	go get
	```
	```go
	go mod tidy
	```
3. Run the main file to start the server.
	```go
	//Run main.go file 
	go run main.go
	//Or use 
	go run .
	```
4. Building into exe
	```go
	//Builds the files
	go build
	```
> [!WARNING]
> The frontend application is still in development and its very difficult to run commands on this API without a clear UI, even with Postman. 




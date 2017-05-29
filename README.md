# JSON Service for Heroku/PostgreSQL (Go)
Just a project more for myself to refer back to when working on future GoLang projects. Generates .JSON services from Heroku/PostgreSQL db queries/data. Service pulls data from **basic_app** production ( https://basic-app.herokuapp.com/ )

# Setup
***To Run:***

*Set Environment Variables via Commands or in Bash File*

export BASIC_APP_DATABASE_URL="postgres://restofurl"

export PORT="8080"

  ```
  git clone https://github.com/brianhodges/golang-jsonservice
  cd golang-jsonservice
  go run main.go
  ```
*Then simply navigate in your browser to:* 
 
 **DEVELOPMENT**
 
 All Users -> 
 
    http://localhost:8080/users.json
 
 Individual User -> 
 
    http://localhost:8080/users/1.json
 
 **PRODUCTION (HEROKU)**
 
 All Users -> 
 
    https://golang-jsonservice.herokuapp.com/users.json
 
 Individual User -> 
 
    https://golang-jsonservice.herokuapp.com/users/1.json

# DISCLAIMER
Your PostgreSQL Database must have a USERS Table with matching fields. For my particular schema check: https://github.com/brianhodges/basic_app

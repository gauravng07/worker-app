#Hello-Fresh

# Local setup
* Install version `go1.14 `
* Install docker
* Run `local-setup` to setup hooks and other tools

# Commands for local development
* Run app :`make run`
* All tests : `make test`
* Or Just: `go run main.go`

# Folder Structure
- **internal:** contains application code

#Docker
* docker build --tag hello-fresh:1.0 .
* docker run --publish 9000:9000 --detach --name bb hello-fresh:1.0

#Functionality NFR
1. `Upload file`
    * Once user run the application
    * Open localhost:9000
    * Upload file and submit.
    * Response will be on browser in JSON format
2. `Search By Recipe`
    * Give recipe name
    * Search
3. `Post Code with delivery time`
    * Give postcode
    * Delivery Start Time
    * Delivery End Time 

#Supporting Notes
* Used to mux and html forms to provide client interface
* Leveraging go channel as queue and worker
# go-api-server

A simple api server written in go

- Sending a POST request to http://localhost:8080/people of an object 'Person' will add the object to an internal map.
- Sending a GET request to http://localhost:8080/{name} of an object 'Person' will retrieve that object
- Sending a GET request to http://localhost:8080/people will return all 'Person' objects stored in the internal map and
  dump them in a file 'data.json'

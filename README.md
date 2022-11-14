# go-mongodb-urlTracker
Endpoints are:
```
/ping
```
## /ping  
Send a GET request.  
Returns a string:  
```
pong
```

## Errors
All the endpoints return an error in json format if something goes wrong. Ex)
```
{
   "Message": <message>,
   "Status": <status>,
   "Error": <error>
}
```
### Used:
lang: **go**  
mux: **github.com/gin-gonic/gin**  
mongodb driver: **go.mongodb.org/mongo-driver**  
### How to run app:
First, get libs and source code.
```
go get github.com/gin-gonic/gin
```
```
go get go.mongodb.org/mongo-driver
```
```
go run main.go
```
It runs on port 8080 by default.

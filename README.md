# HTTP-KeyStore-App

## 🚩 Table of Contents

 - [Description](#description)
 - [Features](#features)
 - [Repository Structure](#repository-structure)
 - [How to Run and Test?](#how-to-run-and-test)
 - [Sample Execution Result](#sample-execution-result)

 ## Description
 This application provides a simple in-memory cache with an HTTP interface.

 ## Features

 ## Repository Structure

 ### Directory tree
    . 
    ├── cmd
    |      └──server
    |             └──server.go             # main function of application i.e; http-keystore-app starts here
    |                         
    ├── internal                                  # handlers for key-value storing and retrieving
    |      └──handlers
    |              └──handlers.go
    |              └──store_handler.go
    |              └──load_handler.go
    |              └──store_handler_test.go         # unit tests for handlers     
    |              └──load_handler_test.go                  
    |                               
    |      
    |      
    |
    ├── docs                                     # contains screenshots of run results for user reference 
    |      └──images                        
    |                  
    ├── vendor                                   # contains application dependencies
    └── README.md


[`cmd`](https://github.com/saikiranambati942/http-keystore-app/tree/master/cmd "API documentation") package:
------------------------------------------------------------------------------------------------------------------

 `cmd` package is the initial point of the application where `server` is the placeholder for server.go(starting point of application)

 [`internal`](https://github.com/saikiranambati942/http-keystore-app/tree/master/internal "API documentation") package:
------------------------------------------------------------------------------------------------------------------------

 `internal` package contains the private code internal to our application, has below `handlers`:

StoreHandler handler handles the POST requests of  storing a value with respect to a key (`store_handler.go`)
LoadHandler handler handles the GET requests of  retrieving a value with respect to a key (`load_handler.go`)
This package also contains the unit test cases covered for all the handlers.

[`vendor`](https://github.com/saikiranambati942/http-keystore-app/tree/master/vendor "API documentation") package:
------------------------------------------------------------------------------------------------------------------------

`vendor` folder contains application dependencies, which includes all the packages needed to support builds and tests of application


## How to Run and Test?
After cloning the repository (https://github.com/saikiranambati942/http-keystore-app.git), run the below command from the root directory to start the http server on localhost:8080

```
go run cmd/server/server.go
```

To store a value with respect to a key, trigger the POST "/{key}" endpoint with the below request format:
```
{
  "value": string,
}
```
For example:
```
{
  "value": "v1",
}
```
To retrieve a value with respect to a key, trigger the GET "/{key}" endpoint with the key specifying in url.

For example:
```
http://localhost:8080/k1
```

## Sample Execution Result
#### Step1: Start server using the below command from the root directory

```
go run cmd/server/server.go
```
#### Step2: Store a value with respect to a key

![](https://github.com/saikiranambati942/http-keystore-app/blob/main/docs/images/store_key_value.png)

#### Step3: Retrieve a value with respect to a key

![](https://github.com/saikiranambati942/http-keystore-app/blob/main/docs/images/load_key_present.png)

#### Step3: Try to retrieve a value of a key which is not present

![](https://github.com/saikiranambati942/http-keystore-app/blob/main/docs/images/load_key_notpresent.png)

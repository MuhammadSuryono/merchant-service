# Okta Backend Framework

## Installation

To install

1. The first need [Go](https://golang.org/) installed (**version 1.11+ is required**), then you can use the below Go command to clone this framework

```sh
$ git clone https://bitbucket.emhusnan.id/scm/~yono/framework-backend.git
```

## Quick start

```sh
# assume the following setting in .env file
$ cat .env.example
```


```sh
# Copy all the settings in the .env.example file, then create a new file and paste it
$ touch .env
```
 
```sh
# assume the following codes in main.go file
$ cat main.go
```

```go
package main

import (
	"os"

	_http "github.com/MuhammadSuryono1997/framework-okta/base/http"
)

func main() {
	server := _http.CreateHttpServer()

	port := os.Getenv("PORT")
	if port == "" {
		port = "5001"
	}

	server.Run(":" + port) // Listen based by .env
}
```

```
# run example.go and visit 0.0.0.0:8080/ping (for windows "localhost:8080/ping") on browser
$ go run main.go
```

## Update Module

```sh
# Update all module
$ go get -u all

# Update some module not spesific
$ go get -u github.com/MuhammadSuryono1997/framework-okta

# Update some module spesific
$ go get -u github.com/MuhammadSuryono1997/framework-okta@v0.0.25

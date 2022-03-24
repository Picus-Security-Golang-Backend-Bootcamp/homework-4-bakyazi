package main

import (
	"fmt"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-bakyazi/internal/api"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-bakyazi/internal/library"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

const (
	EnvResponseTimeoutKey string = "PATIKE_RESP_TIMEOUT_IN_MSEC"
	DefaultRespTimeout           = time.Millisecond * 5000
	SuccessPrefix                = "\nRESULT:\n\t"
	ErrorPrefix                  = "\nERROR:\n\t"
)

// init loads books slice as a library.BookList
func init() {
	err := loadEnv()
	if err != nil {
		log.Printf("cannot load env,err: %v\n", err)
	}
	library.Init()
}

// loadEnv loads env variables from .env or .env.local
func loadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	// if LOCAL mode enable env variables are overwritten by .env.local
	if os.Getenv("PATIKA_ENV_PROFILE") == "LOCAL" {
		err := godotenv.Overload(".env.local")
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	api.Init()
}

//printUsage prints usage
func printUsage() {
	usage := `
### USAGE ###

Commands: 
- search => to search and list books with specified arguments, arguments are searched in books' name, author, isdn, stockCode and ID attributes
	e.g:
		$ ./bin/library search moby dick
- list => to show list of all books
	e.g:
		$ ./bin/library list
- delete => to delete the book by specified ID
	e.g:
		$ ./bin/library delete 5
- buy => to buy the book specified by the ID in the specified amount, first argument is ID of the book and second argument is the amount desired to be bought
	e.g:
		$ ./bin/library buy 5 10
- clear => to operate hard delete for all tables. only recommended for reset DB with sample inputs
	e.g:
		$ ./bin/library clear


### ENVIRONMENT VARIABLES ###

- PATIKA_ENV_PROFILE => Run mode (PROD/LOCAL)
- PATIKA_DB_HOST => Host/IP Address for DB connection (localhost, 127.0.0.1)
- PATIKA_DB_PORT => Port of DB connection
- PATIKA_DB_USER => DB Username  
- PATIKA_DB_PASSWORD => DB Password
- PATIKA_DB_NAME => DB Name
- PATIKE_RESP_TIMEOUT_IN_MSEC => Response Timeout in milliseconds
`
	fmt.Println(usage)
}

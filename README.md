# homework-4-bakyazi

Author: Berkay Akyazi

## build

```shell
$ go mod download
$ go build -o bin/library cmd/hw4/main.go
```

## usage

## environment variables

There should be two files (`.env` or `.env.local`) at root directory of the project

`.env` file must exist and `PATIKA_ENV_PROFILE` must be set. If this parameter is `PROD` then project uses values from `.env` otherwise it overwrites values with `.env.local`


```
### ENVIRONMENT VARIABLES ###

- PATIKA_ENV_PROFILE => Run mode (PROD/LOCAL)
- PATIKA_DB_DRIVER => (postgresql/mysql)
- PATIKA_DB_HOST => Host/IP Address for DB connection (localhost, 127.0.0.1)
- PATIKA_DB_PORT => Port of DB connection
- PATIKA_DB_USER => DB Username  
- PATIKA_DB_PASSWORD => DB Password
- PATIKA_DB_NAME => DB Name
```



## Authorization

GET methods don't require Authorization. Beside that, other methods of every endpoint requires Authorization header. Authorization header value must start with `Bearer `. 

e.g `Bearer tokenfsdlkfjskldfkls`

## endpoints

You can test endpoint by importing my POSTMAN collection export [file](postman_collection.json)

### /authors

###### GET
To receive all authors in DB except deleted ones.

http://localhost:8090/authors

###### POST

To create new author. Accepts application/json

```json
{
  "name": "Berkay Akyazi"
}
```

and it returns the author with id

```json
{
  "id": 148,
  "name": "Berkay Akyazi"
}
```
cURL example:
```shell
curl --location --request POST 'http://localhost:8090/authors' \
--header 'Authorization: Bearer berkayakyazi123123' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Berkay Akyazi"
}'
```

### /authors/{id}

###### GET
To retrieve the author with given id if it exists

http://localhost:8090/authors/1
###### PUT
To update existing author with given id

cURL example

```shell
curl --location --request PUT 'http://localhost:8090/authors/148' \
--header 'Authorization: Bearer 434249873294' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Berkay Akyazi Jr."
}'
```

###### DELETE

To delete existing author with given id. It also deletes all books of that author.

cURL example
```shell
curl --location --request DELETE 'http://localhost:8090/authors/148' \
--header 'Authorization: Bearer fjksdhfkjsdfs'
```

### /books
###### GET
To receive all books in DB except deleted ones.
http://localhost:8090/books

###### POST
To create new book. Accepts application/json
```json
{
    "name": "Test Book",
    "author_id": 148,
    "stock_code": "6781",
    "isbn": "396-85-54496-53-9",
    "page_count": 726,
    "price": 55,
    "stock_amount": 45
}
```

and it returns the book with id
```json
{
    "id": 202,
    "name": "Test Book",
    "author_id": 148,
    "stock_code": "6781",
    "isbn": "396-85-54496-53-9",
    "page_count": 726,
    "price": 55,
    "stock_amount": 45,
    "deleted": null
}
```

cURL example:
```shell
curl --location --request POST 'http://localhost:8090/books' \
--header 'Authorization: Bearer fsdjkfhsdf' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Test Book",
    "author_id": 148,
    "stock_code": "6781",
    "isbn": "396-85-54496-53-9",
    "page_count": 726,
    "price": 55,
    "stock_amount": 45
}'
```

### /books/{id}
###### GET
To retrieve the book with given id
http://localhost:8090/books/24
###### PUT
To update existing book with given id

cURL example:
```shell
curl --location --request PUT 'http://localhost:8090/books/24' \
--header 'Authorization: Bearer fsdfdsfsdfdsfdsfsd' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Test2 Book",
    "author_id": 148,
    "stock_code": "5929",
    "isbn": "638-69-17646-66-6",
    "page_count": 411,
    "price": 76,
    "stock_amount": 54,
    "deleted": null
}'
```

###### DELETE
To delete existing book with given id

cURL example:
```shell
curl --location --request DELETE 'http://localhost:8090/books/24' \
--header 'Authorization: Bearer fdsfsdfsdfdsfsdfs'
```

### /books/search
###### GET
To filter book list with given parameter in query parameter named `query`

http://localhost:8090/books/search?query=james

### /books/buy
###### POST

To buy a book by specified amount. It decrease "stock_amount" value of the book.

example json body
```json
{
    "book_id": 200,
    "quantity": 10
}
```

cURL example:
```shell
curl --location --request POST 'http://localhost:8090/books/buy' \
--header 'Authorization: Bearer fdskjfhsdjkf' \
--header 'Content-Type: application/json' \
--data-raw '{
    "book_id": 200,
    "quantity": 10
}'
```
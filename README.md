# book-service-rest-api

Simple Rest API for managing book's CRUD.
App retrieves book info from [OpenLibrary](https://openlibrary.org/dev/docs/api/books) api by ISBN

Authentication
All requests, excluding /signup, should have these headers to be authorizes:

* Key: user key
* Sign: MD5 sign of the following string "{Method}+{URL}+{Body}+{UserSecret}"

Example:
User:
* Key: "MyUserKey"
* Secret: "MyUserSecret"

REQUEST:
* Method: "POST"
* URL: "http://mydomain.com/books"
* Body: "{"isbn":"9781118464465"}"

AUTH HEADERS:
String to sign should be as follows "POSThttp://mydomain.com/books{"isbn":"9781118464465"}MyUserSecret"

* Key: "MyUserKey"
* Sign: "a616317d753f2b4520d0717395b41a21"

# POST Create new user
### Request
```
 curl --location --request POST '/signup' \
--header 'Key: {Key}' \
--header 'Sign: {Sign}' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name":     "Jackson",
    "key":      "MyKey",
    "secret":   "MySecret"
}'
```
### Response
```
{
  "data": {
    "id": 32,
    "name": "Jackson",
    "key": "MyKey",
    "secret": "MySecret"
  },
  "isOk": true,
  "message": "ok"
}
```
## GETGet user info
### Request
```
curl --location --request GET '/myself' \
--header 'Key: {key}' \
--header 'Sign: {sign}'
```

### Response
```
{
  "data": {
    "id": 32,
    "name": "Jackson",
    "key": "MyKey",
    "secret": "MySecret"
  },
  "isOk": true,
  "message": "ok"
}
```
## POST Create a book
### Request
```
curl --location --request POST '/books' \
--header 'Key: {key}' \
--header 'Sign: {sign}' \
--data-raw '{
    "isbn":"9781118464465"
}'
```
### Response
```
{
  "data": {
    "book": {
      "id": 21,
      "isbn": "9781118464465",
      "title": "Raspberry Pi User Guide",
      "author": "Eben Upton",
      "published": 2012,
      "pages": 221
    },
    "status": 2
  },
  "isOk": true,
  "message": "ok"
}
```

## GET Get all books
### Request
```
curl --location --request GET '/books' \
--header 'Key: {key}' \
--header 'Sign: {sign}'
```
### Response
```
{
  "data": [
    {
      "book": {
        "id": 21,
        "isbn": "9781118464465",
        "title": "Raspberry Pi User Guide",
        "author": "Eben Upton",
        "published": 2012,
        "pages": 221
      },
      "status": 0
    }
  ],
  "isOk": true,
  "message": "ok"
}
```
## PATCH Edit a book
### Request
```
curl --location --request PATCH '/books/:id' \
--header 'Key: {key}' \
--header 'Sign: {sign}' \
--data-raw '{
    "book": {
        "isbn": "9781118464465",
        "title": "Raspberry Pi User Guide",
        "author": "Eben Upton",
        "published": 2012,
        "pages": 221
    },
    "status": 1
}'
```
### Response
```
{
  "data": {
    "book": {
      "id": 21,
      "isbn": "9781118464465",
      "title": "Raspberry Pi User Guide",
      "author": "Eben Upton",
      "published": 2012,
      "pages": 221
    },
    "status": 1
  },
  "isOk": true,
  "message": "ok"
}
```

## DELETE Delete a book
### Request
```
curl --location --request GET '/books' \
--header 'Key: {key}' \
--header 'Sign: {sign}'
```
### Response
```
{
  "data": "Successfully deleted",
  "isOk": true,
  "message": "ok"
}
```

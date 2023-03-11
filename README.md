# Bookstore-API
API for the bookstore with user/admin functionality.

# How to use?

Each response has a *status*, *error* and *data* fields. If the code is 200, then everything is fine and the request was completed successfully, otherwise an error code will be returned. The *error* field contains a description of the error. Request and response both are sent in JSON format.

**Prefix for all requests - /api**

## Without authorization

**GET** /books - get the list of books avaliable at the store

##### Output
```
{
    "status": 200,
    "error": "",
    "data": {
        "books": [
            {
                "id": "27dbfc56ea8f9a107376faf2880333c9",
                "title": "cool book",
                "series": "cool series",
                "price": 500,
                "picture": "http://somedomain/pictures/books/27dbfc56ea8f9a107376faf2880333c9_.png",
                "publisher": "cool pub",
                "language": "eng",
                "description": "cool desc",
                "authors_id": [
                    "e1ddbfba9a5b6b3602eb284a33fa0d83"
                ]
            }
        ]
    }
}
```

**GET** /authors - get the list of authors of the books avaliable at the store

##### Output
```
{
    "status": 200,
    "error": "",
    "data": {
        "authors": [
            {
                "id": "2c79b17fb3d2f2039c918b4ef7d85776",
                "name": "aaaa",
                "surname": "bbb",
                "bio": "bio"
            },
            {
                "id": "e1ddbfba9a5b6b3602eb284a33fa0d83",
                "name": "a",
                "surname": "b",
                "bio": "biobio"
            }
        ]
    }
}
```

## Registration and authorization

**POST** /register
##### Input
```
{
    "name": "name",
    "surname": "suraname",
    "email": "mail@gmail.com",
    "phone_number": "+1234",
    "password": "password"
}
```
##### Output
```
{
    "status": 200,
    "error": "",
    "data": {
        "token": "abcde"
    }
}
```

**POST** /login
##### Input
```
{
    "email": "mail@gmail.com",
    "password": "1234"
}
```
##### Output
```
{
    "status": 200,
    "error": "",
    "data": {
        "token": "abcde"
    }
}
```

#### The token is used for all the below User and Admin requests. You need to add it in the header like this (without {}):

```
Authorization: Bearer {token from the register/login response}
```

## User

**POST** /address/create - create the address for orders delivery
##### Input
```
{
    "country": "country",
    "city": "city",
    "street": "street",
    "house_number": "house num"
}
```
##### Output
```
{
    "status": 200,
    "error": "",
    "data": {
        "country": "country",
        "city": "city",
        "street": "street",
        "house_number": "house num",
        "apartment_number": "",
        "floor": 0
    }
}
```

**POST** /address/update
##### Input
```
{
    "apartment_number": "apt num",
    "floor": 10
}
```
##### Output
```
{
    "status": 200,
    "error": "",
    "data": {
        "country": "country",
        "city": "city",
        "street": "street",
        "house_number": "house num",
        "apartment_number": "apt num",
        "floor": 10
    }
}
```

**GET** /address - get current user delivery address
##### Output
```
{
    "status": 200,
    "error": "",
    "data": {
        "country": "country",
        "city": "city",
        "street": "street",
        "house_number": "house num",
        "apartment_number": "apt num",
        "floor": 10
    }
}
```

**DELETE** /address
##### Output
```
{
    "status": 200,
    "error": "",
    "data": ""
}
```

**POST** /user/update
##### Input
```
{
    "email": "newmail@gmail.com",
    "password": "new pass"
}
```
##### Output
```
{
    "status": 200,
    "error": "",
    "data": ""
}
```

**POST** /user/logout

##### Output
```
{
    "status": 200,
    "error": "",
    "data": ""
}
```

**GET** /user - get user info
##### Output
```
{
    "status": 200,
    "error": "",
    "data": {
        "user_id": "5a28f8091967bcae95b1ca3c977032a0",
        "name": "name",
        "surname": "suraname",
        "email": "newmail@gmail.com",
        "phone_number": "+1234"
    }
}
```

**DELETE** /user
##### Output
```
{
    "status": 200,
    "error": "",
    "data": ""
}
```

**POST** /order/create
##### Input
```
{
    "books_id": ["27dbfc56ea8f9a107376faf2880333c9"],
    "delivery_date_time": "2023-05-01T20:00:00Z"
}
```
##### Output
```
{
    "status": 200,
    "error": "",
    "data": {
        "id": "09fb2a3b52fa0a19c88c64d1fa919924",
        "books_id": [
            "27dbfc56ea8f9a107376faf2880333c9"
        ],
        "delivery_date_time": "2023-05-01T20:00:00Z",
        "created_at": 1678547286,
        "delivered_at": 0
    }
}
```

**POST** /order/update
##### Input
```
{
    "order_id": "09fb2a3b52fa0a19c88c64d1fa919924",
    "delivered_at": 1682964000
}
```
##### Output
```
{
    "status": 200,
    "error": "",
    "data": {
        "id": "09fb2a3b52fa0a19c88c64d1fa919924",
        "books_id": [],
        "delivery_date_time": "2023-05-01T20:00:00Z",
        "created_at": 1678547286,
        "delivered_at": 1682964000
    }
}
```

**GET** /order/list
##### Output
```
{
    "status": 200,
    "error": "",
    "data": {
        "orders": [
            {
                "id": "09fb2a3b52fa0a19c88c64d1fa919924",
                "books_id": [
                    "27dbfc56ea8f9a107376faf2880333c9"
                ],
                "delivery_date_time": "2023-05-01T20:00:00Z",
                "created_at": 1678547286,
                "delivered_at": 1682964000
            }
        ]
    }
}
```

**DELETE** /order
##### Output
```
{
    "status": 200,
    "error": "",
    "data": ""
}
```

## Admin

**POST** /author/create

##### Input
```
{
    "name": "name",
    "surname": "surname"
}
```
##### Output
```
{
    "status": 0,
    "error": "",
    "data": {
        "author_id": "68444111a907302917f1abbcb867b6a7",
        "name": "name",
        "surname": "surname",
        "bio": ""
    }
}
```
**POST** /author/update
##### Input
```
{
    "author_id": "68444111a907302917f1abbcb867b6a7",
    "bio": "bio"
}
```
##### Output
```
{
    "status": 0,
    "error": "",
    "data": {
        "author_id": "68444111a907302917f1abbcb867b6a7",
        "name": "name",
        "surname": "surname",
        "bio": "bio"
    }
}
```

**DELETE** /author
##### Output
```
{
    "status": 200,
    "error": "",
    "data": ""
}
```

**POST** /book/create
##### Input
```
{
    "title": "title",
    "language": "eng",
    "description": "desc",
    "price": 100
}
```
##### Output
```
{
    "status": 200,
    "error": "",
    "data": {
        "book_id": "3c8838d2d7fc3bfa407084482d2f8c72",
        "title": "title",
        "series": "",
        "price": 100,
        "picture": "",
        "publisher": "",
        "language": "eng",
        "description": "desc",
        "count": 0,
        "authors_id": []
    }
}
```

**POST** /book/update
##### Input
```
{
    "book_id": "3c8838d2d7fc3bfa407084482d2f8c72",
    "count": 1000,
    "authors_id": ["68444111a907302917f1abbcb867b6a7"]
}
```
##### Output
```
{
    "status": 200,
    "error": "",
    "data": {
        "book_id": "3c8838d2d7fc3bfa407084482d2f8c72",
        "title": "title",
        "series": "",
        "price": 100,
        "picture": "",
        "publisher": "",
        "language": "eng",
        "description": "desc",
        "count": 1000,
        "authors_id": [
            "68444111a907302917f1abbcb867b6a7"
        ]
    }
}
```

**POST** /book/picture - add book cover picture

*Form-data* body in the request. 
##### Input
```
book_id (text): {book_id}
picture (file): {file path}
```
##### Output
```
{
    "status": 200,
    "error": "",
    "data": ""
}
```

**DELETE** /book
##### Output
```
{
    "status": 200,
    "error": "",
    "data": ""
}
```

**GET** /user/list - get the list of all users
##### Output
```
{
    "status": 200,
    "error": "",
    "data": {
        "users": [
            {
                "user_id": "5a28f8091967bcae95b1ca3c977032a0",
                "name": "name",
                "surname": "suraname",
                "email": "newmail@gmail.com",
                "phone_number": "+1234"
            }
        ]
    }
}
```

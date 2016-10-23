# User Manager

Endpoint: `/users`

# Request 

HTTP verbs: http://www.restapitutorial.com/lessons/httpmethods.html

## POST: create

### required fields

- name
- email
- password

### optional fields

none

## GET: read

- single user: `/users/{id}`
- users: `/users`

### required fields

none

### optional fields

- name
- email


## DELETE: delete

- single user: `/users/{id}`
- users: `/users`


### required fields

none

### optional fields

- name
- email

## PUT: update (replacement)

- single user: `/users/{id}`

### required fields

- id
- name
- email
- password

### optional fields

none

## PATCH: partial update 

- single user: `/users/{id}`

### required fields

- id

### optional fields

- name
- email
- password

## Response

**single resource**

```json
{
  "data": {
    "type": "users",
    "id": "1",
    "attributes": {
      "name": "john",
      "email": "john@thesmithcompany.com",
      "password": "ja8q\+Y!?FCD3":P"
    }
  }
}
```
**multiple resources**

```json
{
  "data": [{
    "type": "users",
    "id": "1",
    "attributes": {
      "name": "john",
      "email": "john@thesmithcompany.com",
      "password": "ja8q\+Y!?FCD3\":P"
    }
  }, 
  {
    "type": "users",
    "id": "2",
    "attributes": {
      "name": "bob",
      "email": "bob@hello.co.uk",
      "password": "ja8q\+Y!?FCD3(:P"
    }
  }]
}
```

[![Coverage Status](https://coveralls.io/repos/github/microlib/usermanager/badge.svg?branch=master)](https://coveralls.io/github/microlib/usermanager?branch=master)

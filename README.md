# User Manager

[![Build Status](https://travis-ci.org/microlib/usermanager.svg?branch=master)](https://travis-ci.org/microlib/usermanager)
[![codecov](https://codecov.io/gh/microlib/usermanager/branch/master/graph/badge.svg)](https://codecov.io/gh/microlib/usermanager)
[![Go Report Card](https://goreportcard.com/badge/github.com/microlib/usermanager)](https://goreportcard.com/report/github.com/microlib/usermanager)
[![GoDoc](https://godoc.org/github.com/microlib/usermanager?status.svg)](https://godoc.org/github.com/microlib/usermanager)

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

*** 

# Installation

The dependencies of this project are stored inside the `vendor/` folder. While we use [gvt](https://github.com/FiloSottile/gvt) 
 to manage the dependencies, we're also pushing them anyway, so that users can feel free to simply run `go get` to
 get the dependencies.
 
Keep in mind that the `vendor/` folder can only be used for dependencies out-of-the-box if you're running Go 1.6 or higher. 
 It is also available on Go 1.5 but you will need to enable the environment variable 
 [`GO15VENDOREXPERIMENT`](https://blog.filippo.io/most-go-tools-now-work-with-go15vendorexperiment/) and set it to `1`. 

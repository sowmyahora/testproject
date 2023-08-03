# testproject

## Description:

This project is a demonstration of CRUD (Create, Read, Update, Delete) operations using APIs in Golang and is dockerized for easy deployment. It provides a basic RESTful API server that allows users to interact with a database and perform CRUD operations on a collection of items.

## Features:

1. Create: Add new items to the collection with unique identifiers.
2. Read: Retrieve individual items or all items from the collection.
3. Update: Modify existing items in the collection by their identifiers.
4. Delete: Remove items from the collection using their identifiers.

## Usage:

Build and run the API server using Docker Compose: docker-compose up --build
The API server will be running at http://localhost:8080 by default.

## API Endpoints 

1. GET /users : Get all items from the collection.
2. GET /users/{id} : Get a specific item by its identifier.
3. POST /users/insert : Create a new item.
4. PUT /users/insert?{id}= : Update an existing item by its identifier.
5. DELETE /users/delete?{id}= : Delete an item by its identifier.

## API Examples:

1. **Create a new user :** 
curl -X POST -H "Content-Type: application/json" -d '{
    "name": "Dr. Ram Sharma",
    "phone": "8178860317",
    "address": {
        "street": "street 26",
        "city": "Mumbai",
        "state": "Maharashtra",
        "country": "India"
    },
    "hobbies":  ["Playing Cricket", "Reading Books", "Swimming"]
}' http://localhost:8080/users

2. **Update an existing user :**
curl -X PUT -H "Content-Type: application/json" -d '{
    "name": "Dr. Rakesh Singh",
    "phone": "908876554",
}' 'http://localhost:8080/users/{userid}

3. **Get all users:**
curl http://localhost:8080/users

4. **Get a specific user:**
curl http://localhost:8080/users/{user_id}

5. **Delete an user:**
curl -X DELETE http://localhost:8080/users/{user_id}
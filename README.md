# fiber-clean

Go Fiber Clean Architecture

# Description

CRUD Online Learning Platform with simplified Clean Architecture. Created with Go Fiber - MongoDB - JWT - Cloudinary.

# API Documentation

`root` = `/api/v1`

## Public APIs

### Register User

```
POST /users
{
    "email": string,
    "password": string
}
```

### Login

```
POST /auth
{
    "email": string,
    "password": string
}
```

## Private APIs

Provide `Bearer Token` with `JWT` for `Authorization` to access these APIs. Several APIs can only be accessed by user with `admin` role.

### Admin APIs

`root`: `/admins`

1. Register New Admin

```
POST /
{
    "email": string,
    "password": string
}
```

2. Get Statistics

```
GET /statistics
```

3. Soft-Delete User

```
DELETE /delete_user/:user_id
```

### Course APIs

`root`: `/courses`

1. Create New Course

Use multipart form to upload image for this. Upload the file with `image` as the key.

```
POST /
{
    "name": string,
    "category": string,
    "price": int,
    "details": string,
    "image": File
}
```

2. Update Course

Note: This update method hasn't supported changing image of the course.

```
PUT /
{
    "name": string,
    "category": string,
    "price": int,
    "details": string
}
```

3. Soft-Delete Course

```
DELETE /:course_id
```

4. Get All Courses

```
GET /
```

5. Search Courses

```
POST /search
{
    "sort": object,
    "filter": object,
    "projection": object,
}
```

This is a general endpoint that can be used to query courses. Here are several examples on how to use this endpoint.

```
Search all free courses

POST /search
{
    "sort": {},
    "filter": {
        "price": 0
    },
    "projection": {},
}
```

```
Searching certain course by name

POST /search
{
    "sort": {},
    "filter": {
        "name": "fiber-clean"
    },
    "projection": {},
}
```

```
Selecting or removing certain attributes

POST /search
{
    "sort": {},
    "filter": {
        "name": "fiber-clean"
    },
    "projection": {
        "_id": false
        "details": true,
        "price": true,
        "category": true
    },
}
```

```
Sorting based on certain attribute. Use `-1` to descending sort the key.

POST /search
{
    "sort": {
        "price": 1
    },
    "filter": {
        "name": "fiber-clean"
    },
    "projection": {
        "_id": false,
        "details": true,
        "price": true,
        "category": true
    }
}
```

# How to Run

## Dependencies

1. Add environment variables to `.env` as follows:

```
MONGO_DB=<db>
MONGO_URI=mongodb+srv://<user>:<password>@<cluster>.g8tknrr.mongodb.net/<db>?retryWrites=true&w=majority

CLOUDINARY_CLOUD_NAME=<cloud_name>
CLOUDINARY_API_KEY=<api_key>
CLOUDINARY_API_SECRET=<api_secret>
CLOUDINARY_UPLOAD_FOLDER=<upload_folder>

JWT_AUTH_SECRET=<secret_string>
```

2. Simply execute `go build` or `go run .` to install all dependencies from `go.mod` and run the server.

# Things to Improve

1. Deeper layer of abstraction if needed.
2. Clearer entities, API definition, and database definition.
3. Separation of environment variables.
4. Security things, like password encryption.
5. Deployment: Heroku now requires credit card to use. :(

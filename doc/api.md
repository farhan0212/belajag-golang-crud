# 📘 API Documentation – Belajar Golang CRUD

> RESTful API untuk manajemen user dengan autentikasi JWT, role-based access control, dan pagination.

---

## 📍 Base URL

```
http://localhost:8080/
```

---

## 🔐 Authentication

### 🔑 Login User

**Endpoint:**  
`POST /user/login`

**Request Body:**
```json
{
    "email": "farhan@example.com",
    "password": "123456"
}
```

**Response (Success):**
```json
{
    "success": true,
    "message": "berhasil login",
    "data": {
        "email": "farhan@example.com",
        "name": "Farhan",
        "token": "JWT_TOKEN_HERE"
    }
}
```

**Response (Failed):**
```json
{
    "success": false,
    "message": "Email atau password salah"
}
```

---

## 👤 User Registration

**Endpoint:**  
`POST /user`

**Request Body:**
```json
{
    "name": "farhan",
    "email": "farhan@example.com",
    "password": "123456"
}
```

**Response:**
```json
{
    "success": true,
    "message": "Sukses input user",
    "data": {
        "name": "farhan",
        "email": "farhan@example.com",
        "role": "user"
    }
}
```

---

## 📄 Get All Users

**Endpoint:**  
`GET /users?page=1&limit=10`

**Headers:**
```
Authorization: Bearer JWT_TOKEN_HERE
```

**Response:**
```json
{
    "page": 1,
    "limit": 10,
    "total": 25,
    "total_pages": 3,
    "data": [
        {
            "id": 1,
            "name": "Farhan",
            "email": "farhan@example.com",
            "role": "user",
            "created_at": "2025-07-10T00:00:00Z",
            "updated_at": "2025-07-10T00:00:00Z"
        }
    ]
}
```

---

## 📄 Get User by ID

**Endpoint:**  
`GET /user/{id}`

**Headers:**
```
Authorization: Bearer JWT_TOKEN_HERE
```

**Response:**
```json
{
    "success": true,
    "message": "Id ditemukan",
    "data": {
        "id": 1,
        "name": "Farhan",
        "email": "farhan@example.com",
        "role": "user",
        "created_at": "2025-07-10T00:00:00Z",
        "updated_at": "2025-07-10T00:00:00Z"
    }
}
```

---

## ✏️ Update User

**Endpoint:**  
`PUT /user/{id}`

**Headers:**
```
Authorization: Bearer JWT_TOKEN_HERE
```

**Request Body (form-data):**
```
name=Farhan Update
email=farhan_update@example.com
```

**Response:**
```json
{
    "success": true,
    "message": "User berhasil diupdate",
    "data": {
        "id": 1,
        "name": "Farhan Update",
        "email": "farhan_update@example.com"
    }
}
```

---

## ❌ Delete User

**Endpoint:**  
`DELETE /user/{id}`

**Headers:**
```
Authorization: Bearer JWT_TOKEN_HERE
```

**Response:**
```json
{
    "success": true,
    "message": "data 1 berhasil didelete"
}
```

---

## 🛡️ JWT Token Payload Example

```json
{
    "user_id": 1,
    "email": "farhan@example.com",
    "role": "admin",
    "exp": 1736538473
}
```

---

## 📝 Example Requests

### Register
```http
POST http://localhost:8080/user
Content-Type: application/json

{
    "name": "ahmad",
    "email": "ahmad@example.com",
    "password": "123456"
}
```

### Login
```http
POST http://localhost:8080/user/login
Content-Type: application/json

{
    "email": "ahmad@example.com",
    "password": "123456"
}
```

### Get all users
```http
GET http://localhost:8080/users?page=1&limit=10
Authorization: Bearer JWT_TOKEN_HERE
```

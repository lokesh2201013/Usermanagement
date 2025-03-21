# User Management API

## Overview
This is a **User Management API** built using **Golang Fiber** and **PostgreSQL**. It provides authentication, user role management, and CRUD operations for users and admins.

### Features:
- **Admin Signup & Login** with JWT authentication
- **User Creation, Retrieval, Update, and Deletion** (Admin only)
- **Admin Management**
- **PostgreSQL Database Integration**
- **Role-based Authorization**

## Tech Stack
- **Golang** (Fiber framework)
- **PostgreSQL** (Database)
- **JWT** (Authentication)
- **GORM** (ORM for PostgreSQL)

---
## API Endpoints

### 1️⃣ Admin Signup
Registers a new admin account.

#### **Request:**
```sh
curl --location --request POST 'http://localhost:3000/signup' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Admin User",
    "email": "admin@example.com",
    "password": "admin123"
}'
```

### 2️⃣ Admin Login
Authenticates an admin and returns a JWT token.

#### **Request:**
```sh
curl --location --request POST 'http://localhost:3000/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "admin@example.com",
    "password": "admin123"
}'
```

#### **Response:**
```json
{
    "token": "YOUR_JWT_TOKEN"
}
```
Use the received JWT token for protected routes.

---
### 3️⃣ Create User (Admin Only)
Creates a new user (Only admins can create users).

#### **Request:**
```sh
curl --location --request POST 'http://localhost:3000/createuser' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer YOUR_JWT_TOKEN' \
--data-raw '{
    "name": "User One",
    "email": "user1@example.com",
    "age": 25
}'
```

---
### 4️⃣ Get All Users (Admin Only)
Fetches all users from the database.

#### **Request:**
```sh
curl --location --request GET 'http://localhost:3000/getallusers' \
--header 'Authorization: Bearer YOUR_JWT_TOKEN'
```

---
### 5️⃣ Get Single User (Admin Only)
Fetches details of a specific user.

#### **Request:**
```sh
curl --location --request GET 'http://localhost:3000/getuser/{id}' \
--header 'Authorization: Bearer YOUR_JWT_TOKEN'
```
🔹 Replace `{id}` with the actual user ID.

---
### 6️⃣ Update User (Admin Only)
Updates a user's details.

#### **Request:**
```sh
curl --location --request PUT 'http://localhost:3000/updateuser/{id}' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer YOUR_JWT_TOKEN' \
--data-raw '{
    "name": "Updated Name",
    "email": "updated@example.com",
    "age": 30
}'
```
🔹 Replace `{id}` with the actual user ID.

---
### 7️⃣ Delete User (Admin Only)
Deletes a user from the database.

#### **Request:**
```sh
curl --location --request DELETE 'http://localhost:3000/deleteuser/{id}' \
--header 'Authorization: Bearer YOUR_JWT_TOKEN'
```
🔹 Replace `{id}` with the actual user ID.

---
### 8️⃣ Get All Admins
Fetches a list of all registered admins.

#### **Request:**
```sh
curl --location --request GET 'http://localhost:3000/getalladmins' \
--header 'Authorization: Bearer YOUR_JWT_TOKEN'
```

---
## 🔧 Setup Instructions

### 1️⃣ Install Dependencies
```sh
go mod tidy
```

### 2️⃣ Configure Database (PostgreSQL)
- Create a PostgreSQL database.
- Update the database connection details in `.env` file:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=yourusername
DB_PASSWORD=yourpassword
DB_NAME=yourdbname
JWT_SECRET=your_jwt_secret
```

### 3️⃣ Run the Application
```sh
go run main.go
```

### 4️⃣ Access the API
The API runs on `http://localhost:3000/`

---
## 📌 Authentication
This API uses **JWT authentication**. After logging in, include the token in your request headers:
```sh
--header 'Authorization: Bearer YOUR_JWT_TOKEN'
```

---
## 🚀 Future Enhancements
- Password hashing using **bcrypt**
- User roles and permissions
- Refresh token mechanism
- Rate limiting

---
## 📜 License
This project is open-source and available under the **MIT License**.

---
## 📩 Contact
For any questions, feel free to reach out!

---
🎯 **Built with Fiber, PostgreSQL, and JWT in Golang!** 🚀


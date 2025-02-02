# Authentication System In Go (Backend)

![Go](https://img.shields.io/badge/Go-1.23-blue)
![REST API](https://img.shields.io/badge/REST-API-brightgreen)
![PostgreSQL](https://img.shields.io/badge/Database-PostgreSQL-blue)

---

## API Endpoints

| Method | Endpoint                       | Description               |
| ------ | ------------------------------ | ------------------------- |
| GET    | `api/user`                     | Get all Users data        |
| GET    | `api/user/{id}`                | Get User by Id            |
| PUT    | `api/user/{id}`                | Update User by Id.        |
| DELETE | `api/user/{id}`                | Delete a User by Id.      |
| POST   | `api/auth/signup`              | Sign Up a user.           |
| POST   | `api/auth/login`               | Login User.               |
| POST   | `api/auth/refresh`             | Refresh JWT Token.        |
| POST   | `api/auth/resend-verification` | Resend Verification Email |
| POST   | `api/auth/forgot-password`     | Forgot Password.          |
| POST   | `api/auth/reset-password`      | Reset Password.           |
| POST   | `api/auth/change-password`     | Change Password.          |
| GET    | `api/auth/status`              | Get Auth/Login Status.    |

---

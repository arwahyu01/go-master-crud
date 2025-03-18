# Go Master CRUD

Go Master CRUD adalah aplikasi REST API sederhana menggunakan Golang dan GORM untuk mengelola operasi CRUD (Create, Read, Update, Delete) pada database.

## 🚀 Fitur
- Authentication (Login & Register) menggunakan JWT
- Middleware untuk proteksi API
- CRUD data pengguna dengan pagination, sorting, dan filtering
- Struktur proyek yang modular
- Menggunakan `mux` sebagai router

## 📂 Struktur Folder
```
/go-master-crud
│── app/
│   ├── controllers/      # Controller untuk request API
│   ├── models/           # Model database dengan GORM
│   ├── routes/           # Routing API
│   ├── middlewares/      # Middleware seperti authentication
│   ├── helpers/          # Helper functions seperti response JSON
│
│── config/
│   ├── database.go       # Konfigurasi database dan migrasi
│
│── .env                  # Konfigurasi environment
│── go.mod                # Go Modules
│── main.go               # Entry point aplikasi
```

## 🔧 Instalasi
### 1. Clone Repository
```sh
git clone https://github.com/username/go-master-crud.git
cd go-master-crud
```

### 2. Install Dependencies
```sh
go mod tidy
```

### 3. Konfigurasi Environment
Buat file `.env` dan isi dengan konfigurasi database:
```
DB_USER=root
DB_PASSWORD=
DB_HOST=localhost
DB_PORT=3306
DB_NAME=go_master_crud
JWT_SECRET=your_secret_key
```

### 4. Jalankan Aplikasi
```sh
go run main.go
```

## 📌 API Endpoints
### 1️⃣ Auth
- **`POST /api/register`** - Registrasi user baru
- **`POST /api/login`** - Login dan mendapatkan token JWT

### 2️⃣ User Management
- **`GET /api/users?page=1&page_size=10&sort=name&filter=admin`** - Get all users (Pagination, Sorting, Filtering)
- **`GET /api/users/{id}`** - Get user by ID
- **`POST /api/users`** - Create new user
- **`PUT /api/users/{id}`** - Update user
- **`DELETE /api/users/{id}`** - Delete user

## 🛠 Teknologi yang Digunakan
- **Golang** (Backend)
- **GORM** (ORM untuk database)
- **MySQL** (Database)
- **Mux** (Routing)
- **JWT** (Authentication)

## 📜 Lisensi
Proyek ini menggunakan lisensi MIT.

---
📌 Dibuat dengan ❤️ oleh ARWP

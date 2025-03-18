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
│   ├── controllers/        # Controller untuk request API
│   ├── middlewares/        # Middleware seperti authentication
│   ├── models/             # Model database dengan GORM
│── config/                 # Konfigurasi aplikasi
│── helpers/                # Helper functions seperti response JSON
│── routes/                 # Routing API
│── database/
│   ├── setup.go            # Konfigurasi database dan migrasi
|   |── migrations/         # File migrasi database
│
│── .env                    # Konfigurasi environment
│── go.mod                  # Go Modules
│── main.go                 # Entry point aplikasi
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
- **`GET /api/user?page=1&page_size=10&sort=name&filter=admin`** - Get all users (Pagination, Sorting, Filtering)
- **`GET /api/user/{id}`** - Get user by ID
- **`POST /api/user`** - Create new user
- **`PUT /api/user/{id}`** - Update user
- **`DELETE /api/user/{id}`** - Delete user

## 🛠 Teknologi yang Digunakan
- **Golang** (Backend)
- **GORM** (ORM untuk database)
- **MySQL** (Database)
- **Mux** (Routing)
- **JWT** (Authentication)

## 📜 Status Proyek
Dalam pengembangan

---
📌 Dibuat dengan ❤️ oleh ARWP

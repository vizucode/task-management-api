# 📌 Task Management API

Sebuah aplikasi sederhana untuk mengelola *tasks* menggunakan **Golang**, **PostgreSQL**, dan **Docker**.

---

## 🔗 API Endpoints

| Method | Endpoint      | Deskripsi                               |
|--------|---------------|-----------------------------------------|
| POST   | `/tasks`      | Membuat task baru untuk user yang login |
| GET    | `/tasks`      | Mengambil daftar semua task user        |
| GET    | `/tasks/:id`  | Mengambil detail task berdasarkan ID    |
| PUT    | `/tasks/:id`  | Mengupdate task berdasarkan ID          |
| DELETE | `/tasks/:id`  | Menghapus task berdasarkan ID           |
| POST   | `/signup`     | Register user baru                      |
| POST   | `/signin`     | Login user                              |

---

## ⚙️ Prerequisites

Sebelum menjalankan project, pastikan sudah menginstall software berikut:

### 1. Golang
- [Download Go](https://go.dev/dl/)  
- Verifikasi instalasi:

    go version

### 2. golang-migrate
- [Install Guide](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)  
- Install via Homebrew (Mac/Linux):

    brew install golang-migrate

- Verifikasi:

    migrate -version

### 3. Docker & Docker Compose
- [Install Docker](https://docs.docker.com/get-docker/)  
- Verifikasi instalasi:

    docker --version  
    docker-compose --version

### 4. PostgreSQL (psql client)
- Install PostgreSQL sesuai OS.  
- Verifikasi:

    psql --version

---

## 🔑 Environment Variables

Buat file `.env` di root project. Berikut contoh template:

    DB_CONNECTION=postgres://user:password@localhost:5432/taskdb?sslmode=disable
    ACCESS_SECRET_KEY=my-secret-key
    APP_PORT=8080
    APP_HOST=localhost

- **`DATABASE_URL`** → URL koneksi ke database PostgreSQL.  
- **`ACCESS_SECRET_KEY`** → Secret key untuk JWT authentication.  
- **`APP_PORT`** → digunakan untuk PORT
- **`APP_PORT`** → digunakan untuk menentukan HOST

---

## 🗄️ Database Migration

Gunakan **golang-migrate** untuk mengatur migrasi database.

### Menjalankan migration (up)

    migrate -path apps/migrations -database "$DATABASE_URL" up

### Rollback migration (down)

    migrate -path apps/migrations -database "$DATABASE_URL" down

### Membuat migration baru

    migrate create -ext sql -dir apps/migrations -seq create_tasks_table

---

## 🚀 Menjalankan Aplikasi

### 1. Menjalankan secara lokal

    go run main.go

### 2. Menjalankan dengan Docker Compose

    docker-compose up --build

Aplikasi akan tersedia di:  
👉 http://localhost:8080

---

## 🧪 Menjalankan Unit Tests

Jalankan semua unit test dengan:

    go test ./... -v

---

## 📖 Swagger Documentation

Swagger UI tersedia untuk dokumentasi API interaktif. Ikuti langkah-langkah berikut:

### 1. Install Swagger Dependencies

    make swagger-install

### 2. Generate Swagger Documentation

    make swagger-gen

### 3. Menjalankan Swagger UI

Setelah aplikasi berjalan, akses Swagger UI di:  
👉 http://localhost:8080/swagger/index.html

### 4. Menjalankan Swagger Server (Development)

Untuk development, Anda juga bisa menjalankan Swagger server terpisah:

    make swagger-serve

### 5. Menjalankan Aplikasi dengan Makefile

    make run

### Fitur Swagger UI:
- **Interactive API Testing** - Test endpoint langsung dari browser
- **Authentication** - Login menggunakan JWT token
- **Request/Response Examples** - Contoh payload dan response
- **Schema Documentation** - Detail struktur data

### Cara Menggunakan Authentication di Swagger:
1. Buka Swagger UI di browser
2. Gunakan endpoint `/signup` untuk register atau `/signin` untuk login
3. Copy `access_token` dari response
4. Klik tombol **"Authorize"** di Swagger UI
5. Masukkan token dengan format: `Bearer <your_access_token>`
6. Sekarang Anda bisa mengakses endpoint yang memerlukan authentication

---

## 📦 Dependencies

Daftar dependency utama dalam project ini:

### 1. [Fiber](https://github.com/gofiber/fiber)
Framework web **Go** yang cepat, ringan, dan mirip dengan Express.js pada Node.js. Digunakan untuk routing, middleware, dan request handling.

### 2. [Golang JWT](https://github.com/golang-jwt/jwt)
Library untuk membuat dan memverifikasi **JSON Web Token (JWT)**. Dipakai untuk otentikasi dan otorisasi user.

### 3. [GORM](https://gorm.io/index.html)
Library untuk melakukan koneksi dan query ke database dengan mudah.

### 4. [Gokit](https://github.com/vizucode/gokit)
Paket library yang ready-to-use tanpa perlu setting dll.

### 5. [Swaggo](https://github.com/swaggo/swag)
Library untuk generate dokumentasi Swagger dari komentar Go code.

---

## 📝 Catatan Tambahan
- Pastikan `.env` sudah dikonfigurasi dengan benar.  
- Jika menjalankan dengan Docker Compose, database akan otomatis berjalan dalam container `db`.  
- Gunakan `migrate create` untuk menambah migrasi baru sesuai kebutuhan.
- Swagger documentation akan otomatis ter-update setiap kali menjalankan `make swagger-gen`.

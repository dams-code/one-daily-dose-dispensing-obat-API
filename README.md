# One Daily Dose Dispensing Obat API

Backend REST API untuk Sistem Farmasi yang dibangun menggunakan **Go (Golang)**, **Gin Framework**, dan **PostgreSQL**.

API ini mendukung autentikasi JWT, CRUD pada pasien, obat, transaksi resep obat (header dan detail), dispensing, pembayaran, hingga laporan transaksi resep obat.

---

## Features

- Login dan Register
- JWT AUTH
- CRUD User (Soft delete)
- User Role (get, update, auth dengan JWT)
- CRUD Kategori Obat
- CRUD Obat
- CRUD Pasien
- CRUD Transaksi Resep Obat
- CRUD Detail Transaksi Resep Obat
- Dispensing Obat (Update dan Validasi)
- Pembayaran dan Validasi pada status transaksi (CASH dan QRIS)
- Riwayat Transaksi Resep Obat (rentang tanggal)
- Swagger UI API Documentation
- Validasi Input obat, dan transaksi pada resep.

---

## Tech Stack

| Technology | Description |
|------------|-------------|
| Go | Programming Language |
| Gin | HTTP Web Framework |
| PostgreSQL | Database |
| JWT | Authentication |
| bcrypt | Password Hashing |
| Swagger (Swaggo) | API Documentation |
| Postman / Thunder Client | API Testing |

---

## Struktur Project

```text
│   .gitignore                              # Daftar file/folder yang diabaikan Git
│   go.mod                                  # Daftar module dan dependency Go
│   go.sum                                  # Checksum dependency Go
│   main.go                                 # Entry point aplikasi
│   README.md                               # Dokumentasi project
│   Relasi-table-one-daily-dose-dispensing-obat-api.png   # Diagram ERD database table
│
├───config                                  # Konfigurasi aplikasi
│       .env                                # Environment variables (Database, JWT Secret, dll)
│
├───controllers                             # Layer HTTP Handler (Request dan Response)
│       detailtransaksiobat.go              # Controller Detail Transaksi Obat
│       dispensing.go                       # Controller Dispensing Obat
│       kategoriobat.go                     # Controller Kategori Obat
│       koneksidb.go                        # Inisialisasi koneksi database
│       laporantransaksi.go                 # Controller Laporan Transaksi
│       obat.go                             # Controller Obat
│       pasien.go                           # Controller Pasien
│       pembayaranobat.go                   # Controller Pembayaran Obat
│       transaksiobat.go                    # Controller Transaksi Obat
│       users.go                            # Controller User & Authentication
│
├───docs                                    # Dokumentasi Swagger UI
│       docs.go
│       swagger.json
│       swagger.yaml
│
├───middleware                              # Middleware aplikasi
│       middleware.go                       # Middleware autentikasi and otorisasi
│       setservicejwt.go                    # JWT Helper & Token Validation
│
├───migrations                              # Database migration and seeding
│   │   migrations.go                       # Menjalankan migration database
│   │
│   └───sql_migrations
│           data_obat.sql                   # Data awal (Seeder) obat ke database
│
├───repositories                            # Layer Database (SQL Query and Business Logic)
│       detailtransaksiobat.go              # Repository Detail Transaksi Obat
│       dispensing.go                       # Repository Dispensing
│       kategoriobat.go                     # Repository Kategori Obat
│       laporantransaksi.go                 # Repository Laporan
│       obat.go                             # Repository Obat
│       pasien.go                           # Repository Pasien
│       pembayaranobat.go                   # Repository Pembayaran
│       transaksiobat.go                    # Repository Transaksi Obat
│       users.go                            # Repository User
│
├───routers                                 # Routing endpoint API
│       router.go                           # Seluruh endpoint REST API
│
└───structs                                 # Struct Request, Response, & Model
        structsODDD.go                      # Seluruh model, request, dan response API
```

---

## Database

Database terdiri dari tabel:

- user_obat
- pasien
- kategori_obat
- obat
- transaksi_obat
- detail_transaksi_obat
- dispensing

---

## User Roles

| Role | Description |
|------|-------------|
| KEPALAFARMASI | Mengelola laporan transaksi resep obat dan monitoring user |
| ADMINFARMASI | Mengelola data obat dan transaksi resep obat |
| APOTEKER | Melakukan proses dispensing (pemberian obat ke pasien) |

---

## Main Features pada Project Ini

### Authentication
- Register
- Login
- JWT Authentication

### UserObat
- Get UserObat
- Get UserObat By ID
- Add UserObat
- Update UserObat
- Delete UserObat
- Activate UserObat

### Kategori Obat
- Get Kategori Obat
- Get Kategori Obat By ID
- Add Kategori Obat
- Update Kategori Obat
- Delete Kategori Obat

### Obat
- Get Obat
- Get Obat By ID
- Get Obat By Kategori Obat
- Tambah Obat
- Update Obat
- Delete Obat

### Pasien
- Get Pasien
- Get Pasien By ID
- Add Pasien
- Update Pasien
- Delete Pasien

### Transaksi Resep Obat (Header)
- Create Transaksi Resep Obat
- Update Transaksi Resep Obat
- Delete Transaksi Resep Obat
- Cancel Transaksi Resep Obat
- Get Transaksi Resep Obat
- Get Transaksi Resep Obat By ID

### Detail Transaksi Resep Obat (Detail)
- Add Detail Transaksi Resep Obat
- Update Detail Transaksi Resep Obat
- Delete Detail Transaksi Resep Obat
- Get Detail Transaksi Resep Obat
- Get Detail Transaksi Resep Obat By Transaksi Obat

### Dispensing
- Get Pending Transaksi Resep Obat
- Dispense Obat

### Pembayaran
- Pembayaran Tunai
- Pembayaran QRIS
- Perhitungan pembayaran (kembalian, grandtotal resep obat, dan total yang dibayar)

### Laporan transaksi resep obat
- Laporan transaksi resep obat dengan rentang tanggal.

---

## API Docs

Swagger UI tersedia di:

```
http://localhost:8080/swagger/index.html#/
```

---

## ⚙ Instalasi

Clone repository

```bash
git clone https://github.com/dams-code/one-daily-dose-dispensing-obat-api.git
```

Masuk ke project

```bash
cd one-daily-dose-dispensing-obat-api
```

Install dependency

```bash
go mod tidy
```

Jalankan aplikasi project

```bash
go run main.go
```

---

## API Testing

Gunakan Postman, Thunder Client atau Swagger UI.

---

## Copyright Personal Portfolio
* **Project Owner / Created By:** Damar Djati Wahyu Kemala
* **Role:** Final Project Sanbercode (One Daily Dose Dispensing Obat API)
* **Date Created:** Juli 2026
* **GitHub Portfolio:** [https://github.com/dams-code](https://github.com/dams-code)

---
*© 2026 Damar Djati Wahyu Kemala. This project is a part of my Final Project Sanbercode (One Daily Dose Dispensing Obat API). Authorization is required for commercial use or modification.*

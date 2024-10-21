# Event Management System API

Sistem ini memungkinkan pengguna untuk membuat, mengelola, dan mengikuti acara. Pengguna dapat mendaftar untuk acara tertentu, memberikan ulasan, serta melihat detail acara. Sistem juga mengelola kategori acara, dan partisipasi pengguna.

## Fitur

- **CRUD Pengguna**: Kelola pengguna, termasuk menambah, memperbarui, menghapus, dan mendapatkan detail pengguna.
- **CRUD Kategori**: Kelola kategori acara, termasuk menambah, memperbarui, menghapus, dan mendapatkan daftar kategori.
- **CRUD Acara**: Kelola acara, termasuk membuat, memperbarui, menghapus, dan melihat detail acara.
- **CRUD Ulasan**: Pengguna dapat memberikan ulasan untuk acara, serta mengelola ulasan.
- **CRUD Partisipan**: Kelola partisipasi pengguna dalam acara, seperti mendaftar dan memperbarui informasi partisipan.

## Teknologi yang Digunakan

- **Golang**: Bahasa pemrograman utama yang digunakan untuk membangun API.
- **Gin Framework**: Framework web untuk Golang yang digunakan untuk menangani request HTTP.
- **PostgreSQL**: Basis data yang digunakan untuk menyimpan data pengguna, acara, kategori, ulasan, dan partisipan.
- **dotenv**: Digunakan untuk manajemen konfigurasi melalui file `.env`.
- **SQL Migrate**: Untuk mengelola migrasi database, termasuk pembuatan tabel dan perubahan skema database.
- **Basic Auth**: Autentikasi dasar menggunakan username dan password untuk mengakses API.
- **Railway**: Menggunakan railway untuk mendeploy aplikasi.

## Spesifikasi Aplikasi

- Aplikasi ini menggunakan PostgreSQL sebagai basis data utama.
- Basic Authentication digunakan untuk mengamankan endpoint.
- Sistem migrasi database SQL memungkinkan kemudahan dalam pembaruan dan pengelolaan skema basis data.
- Menerapkan pattern RESTful API sehingga semua operasi pada data dapat diakses melalui endpoint HTTP yang spesifik.
- Menyediakan berbagai rute untuk mengelola pengguna, kategori, acara, ulasan, dan partisipasi acara.
- Aplikasi ini berjalan pada port `8080` secara default.

## Prasyarat

Pastikan Anda memiliki perangkat lunak berikut terpasang di sistem Anda sebelum menjalankan proyek:

- [Golang](https://golang.org/dl/) versi terbaru
- [PostgreSQL](https://www.postgresql.org/download/) sesuai dengan sistem operasi Anda

## Instalasi

1. **Clone repositori ini**

   ```bash
   git clone https://github.com/Hassanjadi/event-management-api.git
   ```

2. **Masuk ke direktori proyek**

   ```bash
   cd event-management-api
   ```

3. **Instal dependensi proyek**

   Jalankan perintah di bawah ini untuk mengunduh semua dependensi yang dibutuhkan oleh proyek.

   ```bash
   go mod tidy
   ```

4. **Salin file konfigurasi**

   Salin file `.env.example` menjadi `.env` dan sesuaikan pengaturan lingkungan (environment settings) sesuai dengan kebutuhan Anda, seperti konfigurasi database.

   ```bash
   cp .env.example .env
   ```

5. **Jalankan migrasi database**

   Gunakan perintah berikut untuk menjalankan migrasi database.

   ```bash
   go run main.go
   ```

6. **Menjalankan server**

   Setelah berhasil melakukan semua langkah di atas, jalankan server menggunakan perintah berikut:

   ```bash
   go run main.go
   ```

   Aplikasi ini dapat diakses di `http://localhost:8080/`.

## Diagram Entity Relationship (ERD)

![ERD](https://i.ibb.co.com/bKxdhc6/emsy.png)

## Endpoint API

Berikut adalah daftar endpoint yang dapat diakses dalam aplikasi ini:

### Pengguna

- **GET** `/api/users` : Mendapatkan semua pengguna
- **POST** `/api/users` : Menambahkan pengguna baru
- **GET** `/api/users/:id` : Mendapatkan detail pengguna berdasarkan ID
- **PUT** `/api/users/:id` : Memperbarui pengguna berdasarkan ID
- **DELETE** `/api/users/:id` : Menghapus pengguna berdasarkan ID

### Kategori

- **GET** `/api/categories` : Mendapatkan semua kategori
- **POST** `/api/categories` : Menambahkan kategori baru
- **GET** `/api/categories/:id` : Mendapatkan detail kategori berdasarkan ID
- **PUT** `/api/categories/:id` : Memperbarui kategori berdasarkan ID
- **DELETE** `/api/categories/:id` : Menghapus kategori berdasarkan ID

### Acara

- **GET** `/api/events` : Mendapatkan semua acara
- **POST** `/api/events` : Menambahkan acara baru
- **GET** `/api/events/:id` : Mendapatkan detail acara berdasarkan ID
- **PUT** `/api/events/:id` : Memperbarui acara berdasarkan ID
- **DELETE** `/api/events/:id` : Menghapus acara berdasarkan ID

### Ulasan

- **GET** `/api/reviews` : Mendapatkan semua ulasan
- **POST** `/api/reviews` : Menambahkan ulasan baru
- **GET** `/api/reviews/:id` : Mendapatkan detail ulasan berdasarkan ID
- **PUT** `/api/reviews/:id` : Memperbarui ulasan berdasarkan ID
- **DELETE** `/api/reviews/:id` : Menghapus ulasan berdasarkan ID

### Partisipan

- **GET** `/api/participants` : Mendapatkan semua partisipan
- **POST** `/api/participants` : Menambahkan partisipan baru
- **GET** `/api/participants/:id` : Mendapatkan detail partisipan berdasarkan ID
- **PUT** `/api/participants/:id` : Memperbarui partisipan berdasarkan ID
- **DELETE** `/api/participants/:id` : Menghapus partisipan berdasarkan ID

## Autentikasi

API ini menggunakan **Basic Authentication** untuk mengamankan endpoint. Pastikan Anda menggunakan username dan password yang valid (contoh: `admin` / `admin`) saat melakukan request ke API.

# Cinema Booking System API - MKP Skill Test

Proyek ini adalah implementasi dari seleksi teknis (Skill Test) untuk posisi Golang Developer di MKP. Sistem ini mencakup perancangan sistem pembelian tiket bioskop skala nasional, desain database, dan implementasi API menggunakan bahasa pemrograman Golang.

**Kandidat:** Kevin Majesta Ivano

---

## A. System Design Analysis

### 1. Flowchart Sistem

Flowchart yang menggambarkan alur proses dari pemilihan jadwal hingga pembayaran tiket dapat dilihat pada file berikut:

- [Flowchart Sistem (PNG)](./MKP.drawio.png)

### 2. Solusi Teknis & Alur Bisnis

#### **Sistem Pemilihan Tempat Duduk & Performa**

Untuk menangani akses bersamaan oleh banyak orang dan mencegah **double booking**, solusi yang saya tawarkan adalah:

- **Locking dengan Redis**: Menggunakan _locking_ dari redis, atau _locking_ db bebas tergantung kesepakatan, namun saya sendiri lebih suka dari redis karena redis memakan ram yang dimana lebih cepat dari pada menggunakan _locking_ db itu sendiri, jadi system membuat _key_ untuk Waktu kadaluasa pemesanan selama 10 menit untuk mencegah adanya _double booking_ dari user lain.

#### **Sistem Restok Tiket**

Sistem restok otomatis dipicu oleh dua kondisi:

1.  **Timeout Redis**: Jika waktu pemesanan 10 menit berakhir tanpa adanya konfirmasi sukses dari payment gateway, status kursi otomatis kembali dari _Reserved_ menjadi _Available_.
2.  **Kegagalan Pembayaran**: Jika sistem menerima notifikasi _failed_ dari payment gateway, sistem akan langsung mengembalikan status kursi menjadi _Available_ tanpa menunggu timeout.

#### **Alur Refund & Pembatalan (Pihak Bioskop)**

Jika bioskop melakukan pembatalan jadwal tayang (status _Cancelled_):

1.  **Block Payment**: Sistem secara otomatis memblokir semua proses pembayaran yang sedang berlangsung untuk jadwal tersebut.
2.  **Async Refund Processing**: Sistem memproses refund untuk transaksi yang sudah berstatus _Sold_.
    - **Teknis**: Menggunakan **Goroutine** atau **Worker Pool**(saya belum pernah nyoba ini) untuk menjalankan proses refund di latar belakang (asynchronous) sehingga tidak membebani performa API utama & Loading lama.
    - **Batching**: Proses refund dikirim secara batch (misal 20 transaksi per proses) ke payment gateway melalui API `transaction_id` jika pake worker pool, 1 1 jika pake gouroutine.
3.  **Notifikasi**: Setelah refund sukses, sistem mengirimkan email konfirmasi dan struk refund sebagai bukti kepada pelanggan.

---

## B. Database Design

Rancangan database telah disesuaikan untuk mendukung fungsionalitas di atas (Tabel User, Schedules, Seats, dan Transactions).
Script SQL untuk inisialisasi database dapat ditemukan di:

- [PostgreSQL Script](./db/migrations/) Untuk inisialisasi database & setup database soal C
- [Database Diagram](./relasi.drawio.png) Untuk memahami database yang sudah dibuat soal B

---

## C. Skill Test (API Implementation)

### **Tech Stack**

- **Language**: Golang
- **Web Framework**: Echo (v4)
- **ORM**: GORM
- **Database**: PostgreSQL
- **Authentication**: JWT (JSON Web Token)
- **Encryption**: Bcrypt for Password Hashing

### **API Endpoints**

| Method | Endpoint                | Description                  | Auth |
| :----- | :---------------------- | :--------------------------- | :--- |
| POST   | `/api/v1/register`      | Pendaftaran User Baru        | No   |
| POST   | `/api/v1/login`         | Login User & Mendapatkan JWT | No   |
| GET    | `/api/v1/schedules`     | Mengambil Semua Jadwal       | Yes  |
| POST   | `/api/v1/schedules`     | Membuat Jadwal Baru          | Yes  |
| GET    | `/api/v1/schedules/:id` | Detail Jadwal berdasarkan ID | Yes  |
| PUT    | `/api/v1/schedules/:id` | Update Data Jadwal           | Yes  |
| DELETE | `/api/v1/schedules/:id` | Menghapus Jadwal             | Yes  |

---

## Setup & Running

1. **Clone Repository**

   ```bash
   git clone <link-repository>
   ```

2. **Configuration**
   Sesuaikan file `.env` dengan kredensial database PostgreSQL kamu:

   ```env
   PORT=8080
   POSTGRES_HOST=localhost
   POSTGRES_USER=postgres
   POSTGRES_PASSWORD=password
   POSTGRES_DATABASE=db
   POSTGRES_PORT=5432
   JWT_SECRET=digit
   ```

3. **Run Migrations**

   ```bash
   migrate -path db/migrations -database "postgres://user:pass@localhost:5432/dbname?sslmode=disable" up
   ```

4. **Run Application**
   ```bash
   // jika belum punya library
   go mod tidy
   go run cmd/app/main.go
   ```

---

## Postman Collection

Ekspor Postman tersedia untuk mempermudah pengujian API:

- [Postman Collection JSON](./bioskopMKP_API.postman_collection.json)

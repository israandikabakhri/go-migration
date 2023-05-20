# go-migration
Proyek ini adalah implementasi Go Migration dengan menggunakan database MariaDB. Go Migration adalah pendekatan yang digunakan untuk mengelola versi dan struktur database secara terstruktur dan terkendali. Dalam proyek ini, kami menggunakan Go sebagai bahasa pemrograman utama dan MariaDB sebagai sistem manajemen basis data.


## Dalam project ini melibatkan 3 file yaitu:

```
go-migration/migration/migration.go
go-migration/app/Model/penduduk.go
go-migration/main.go
.env
```

## Cara Menjalankan Project:

1. Buat database di mariaDB setelah itu setting .env agar database bisa terhubung ke project **go-migration**
2. Buka **CMD**
3. Masuk ke Folder project melalui **CMD**
4. ketik **go run main.go** 
5. Buka databse anda maka anda dapat melihat table, field, type data yang di import dari **Modal/penduduk.go**


Enjoyy..

Jika ada kendala jangan sungkan email ke andikaisra7@gmail.com


## Catatan Penggunaan Tiap Folder:

- app/Function -> Fungsinya sebagai penyimpanan query kompleks atau fungsi query yang memiliki relasi yang kompleks yang nanti sisa dipanggil oleh controller
- app/Helpers -> Fungsinya sebagai penyimpanan file helper seperti mengubah format tanggal, waktu, uang dan lain-lain
- app/Http/Controllers -> Fungsinya sebagai penyimpanan seluruh controller
- app/Http/Requests -> Fungsinya sebagai penyimpanan semua validator requests yang nanti akan dipanggil di controller untuk validasi input data seperti simpan dan update
- app/Model -> Fungsinya sebagai penyimpanan folder-folder Model, Connect Database serta migration

## Catatan Batasan Penggunaan Komponen:

- Model: Komponen yang berinteraksi dengan database dan pengubah format data seperti mengubah format tanggal yang dipanggil dari Helper
- Helper: Komponen fungsi yang dapat dipanggil dari mana saja
- Controller: Komponen yang mengelola pemanggilan data dari model, pemanggilan validasi requests, pemanggilan Functions Query Kompleks dan Pengelolaan Rest API
- Request: Komponen yang mengelola validasi input dari user termasuk keamanan SQL Injection dan XSS Script
- Main.go: Komponen yang mengelola Routing dan Pengelolaan hak akses serta authentikasi menggunakan JWT auth

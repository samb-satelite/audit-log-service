# Audit Log (Service)

## Deskripsi
App Audit Log Service adalah layanan untuk mencatat dan mengelola log audit aplikasi. Dibangun menggunakan Golang dan framework Echo, layanan ini menyediakan API untuk menyimpan, mengakses, dan mengelola log audit dengan mudah.

## Fitur
- Pencatatan log audit dengan detil lengkap.
- API untuk menyimpan dan mengambil log audit.
- Konfigurasi yang fleksibel dan mudah disesuaikan.

## Persyaratan
- Go 1.16 atau versi lebih baru.

## Instalasi
1. Clone repositori:
    ```sh
    git clone https://github.com/irfandiricon/app-audit-log-service.git
    cd app-audit-log-service
    ```

2. Install dependensi:
    ```sh
    go mod download
    ```

3. Jalankan aplikasi:
    ```sh
    go run main.go
    ```
## Struktur project

```
├───src
│   ├───common
│   │   ├───http
│   │   └───utils
│   │       ├───httpresponse
│   │       └───middlewares
│   ├───domain
│   │   ├───models
│   │   └───repositories
│   ├───infrastructure
│   │   ├───api
│   │   ├───cache
│   │   ├───database
│   │   │   └───mysql
│   │   └───rabbitmq
│   ├───interfaces
│   │   ├───kafka
│   │   ├───mq
│   │   ├───rest
│   │   │   └───controller
│   │   └───rpc
│   ├───payload
│   │   ├───request
│   │   └───response
│   └───usecases
└───tmp
```

## Konfigurasi
Konfigurasi aplikasi dapat dilakukan melalui file `.env` atau menggunakan variabel lingkungan. Beberapa pengaturan penting meliputi:
- `port`: Port dimana server akan berjalan.
- `database`: Pengaturan koneksi ke database.

## Kontribusi
Kami menerima kontribusi dalam bentuk pull request. Silakan buka issue terlebih dahulu untuk mendiskusikan perubahan yang ingin Anda lakukan.

## Lisensi
Proyek ini dilisensikan di bawah lisensi MIT. Lihat file [LICENSE](LICENSE) untuk informasi lebih lanjut.

## Kontak
Jika Anda memiliki pertanyaan atau masukan, silakan hubungi kami melalui email: contact@auditlogservice.com.

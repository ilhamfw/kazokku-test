# Kazokku App

Kazokku App adalah aplikasi sederhana untuk registrasi pengguna dengan data kartu kredit dan pengunggahan foto. Aplikasi ini menggunakan Echo sebagai framework web dan GORM sebagai ORM untuk koneksi ke database.

## Menjalankan Aplikasi

1. Pastikan Go sudah terinstal di perangkat Anda. [Instal Go](https://golang.org/doc/install)

2. Clone repositori ini:

    ```bash
    gh repo clone ilhamfw/kazokku-test
    cd kazokku-app
    ```

3. Install dependensi:

    ```bash
    go mod tidy
    ```

4. Buat file konfigurasi `.env` (duplikat `.env.example` dan ubah sesuai kebutuhan):

    ```bash
    cp .env.example .env
    ```

5. Jalankan aplikasi:

    ```bash
    go run main.go
    ```

   Aplikasi akan berjalan di [http://localhost:3000](http://localhost:3000).

## Endpoint

### 1. Register User

**URL**: `POST /user/register`

**Contoh Request Body**:

```json
{
    "name": "John Doe",
    "address": "123 Main Street",
    "email": "john.doe@example.com",
    "password": "secretpass123",
    "photos": ["photo1.jpg", "photo2.jpg"],
    "creditcard_type": "VISA",
    "creditcard_number": "4111111111111111",
    "creditcard_name": "John Doe",
    "creditcard_expired": "12/25",
    "creditcard_cvv": "123"
}

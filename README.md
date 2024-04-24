# Aplikasi Enigma Laundry

## API Spec

## Autentikasi

Autentikasi diperlukan untuk mengakses beberapa endpoint API. Saat ini, kami menggunakan metode autentikasi [Basic Authentication] untuk melindungi akses ke data sensitif. Untuk mengakses endpoint yang memerlukan autentikasi, Anda perlu menyertakan username dan password yang valid dalam header permintaan Anda.

### Cara Mendapatkan Token JWT

Untuk mendapatkan kredensial, Anda perlu melakukan langkah-langkah berikut:

1. **Registrasi Akun Pengguna:** Kunjungi halaman registrasi di https://get-credential-api.yusharwz.my.id dan buatlah akun pengguna baru.
2. **Autentikasi Pengguna:** Setelah melakukan pendaftaran, anda akan menerima email yang memberitahukan kredensial yang anda ajukan sedang masuk daftar tunggu.
3. **Kredensial Siap Digunakan:** Jika kredensial yang anda ajukan di setujui, anda akan menrima email pemberitahuan kembali dan pada saat itu juga kredensial yang anda ajukan sudah bisa digunakan dengan batasan hit per hari sebanyak 1000x hit. Kami menyediakan opsi berlangganan untuk hit limit yang lebih banyak. Kunjungi https://www.yusharwz.my.id/ untuk info lebih lanjut.

Pastikan untuk menyimpan Username dan Password Authentikasi dengan aman dan tidak membagikannya dengan orang lain.

### Customer API

#### Create Customer

Request :

- Method : `POST`
- Endpoint : `/api/customers`

```json
{
  "Name": "string",
  "NoHp": "string"
}
```

Response :

- Status : 201 Created
- Body :

```json
{
  "id": "string",
  "Name": "string",
  "NoHp": "string"
}
```

#### Get Customer

Request :

- Method : GET
- Endpoint : `/api/customers/:id`
- Header :
  - Accept : application/json

Response :

- Status : 200 OK
- Body :

```json
{
  "Id": "string",
  "Name": "string",
  "NoHp": "string"
}
```

#### Update Customer

Request :

- Method : PUT
- Endpoint : `/api/customers/:id`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Body :

```json
{
  "Name": "string",
  "NoHp": "string"
}
```

Response :

- Status : 200 OK
- Body :

```json
{
 {
    "Id": "string",
    "Name": "string",
    "NoHp": "string"
  }
}
```

#### Delete Customer

Request :

- Method : DELETE
- Endpoint : `/api/customers/:id`
- Header :
  - Accept : application/json
- Body :

Response :

- Status : 200 OK
- Body :

```json
{
  "message": "string",
  "data": "OK"
}
```

### Transaction API

#### Create Transaction

Request :

- Method : POST
- Endpoint : `/api/transactions`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Body :

```json
{
  "IdPelanggan": int,
  "IdPegawai":  int,
  "TanggalKeluar": "stirng",
  "StatusPembayaran": "string",
  "DetailTransaksi": [
    {
      "IdLayanan": int,
      "Quantity": int
    }
  ]
}

```

Response :

- Status Code: 201 Created
- Body :

```json
{
    "transaksi": {
        "Id": int,
        "IdPelanggan": int,
        "IdPegawai": int,
        "TanggalMasuk": "string",
        "TanggalKeluar": "stirng",
        "StatusPembayaran": "string",
        "DetailTransaksi": [
            {
                "Id": int,
                "IdTransaksi": int,
                "IdLayanan": int,
                "Quantity": int
            }
        ]
    }
}
```

#### Get Transaction By Transaction ID

Request :

Pattern string date : `yyyy-mm-dd`

- Method : GET
- Endpoint : `/api/transactions/:id`
- Header :
  - Accept : application/json
- Body :

Response :

- Status Code: 200 OK
- Body :

```json
{
    "Detail Transaksi": [
        {
            "NamaPelanggan": "string",
            "NamaLayanan": "string",
            "Quantity": int,
            "NamaPegawai": "string",
            "TanggalMasuk": "string",
            "Harga": int
        },
        {
            "NamaPelanggan": "string",
            "NamaLayanan": "string",
            "Quantity": int,
            "NamaPegawai": "string",
            "TanggalMasuk": "string",
            "Harga": int
        }
    ],
    "Total Pembayaran": int
}


```

#### Get Transaction By Customer ID

Request :

- Method : GET
- Endpoint : `/api/transactions/customers/id/:id`
- Header :
  - Accept : application/json
- Body :

Response :

- Status Code: 200 OK
- Body :

```json
{
    "Detail Transaksi": [
        {
            "NamaPelanggan": "string",
            "NamaLayanan": "string",
            "Quantity": int,
            "NamaPegawai": "string",
            "TanggalMasuk": "string",
            "Harga": int
        },
        {
            "NamaPelanggan": "string",
            "NamaLayanan": "string",
            "Quantity": int,
            "NamaPegawai": "string",
            "TanggalMasuk": "string",
            "Harga": int
        }
    ],
    "Total Pembayaran": int
}
```

#### Get Transaction By Customer Name

Request :

- Method : GET
- Endpoint : `/api/transactions/customers/name/:name`
- Header :
  - Accept : application/json
- Body :

Response :

- Status Code: 200 OK
- Body :

```json
{
    "Detail Transaksi": [
        {
            "NamaPelanggan": "string",
            "NamaLayanan": "string",
            "Quantity": int,
            "NamaPegawai": "string",
            "TanggalMasuk": "string",
            "Harga": int
        },
        {
            "NamaPelanggan": "string",
            "NamaLayanan": "string",
            "Quantity": int,
            "NamaPegawai": "string",
            "TanggalMasuk": "string",
            "Harga": int
        }
    ],
    "Total Pembayaran": int
}
```

#### List Transaction

Pattern string date : `dd-MM-yyyy`

Request :

- Method : GET
- Endpoint : `/api/transactions`
- Header :
  - Accept : application/json
- Query Param :
  - startDate : string `optional`
  - endDate : string `optional`
  - productName : string `optional`
- Body :

Response :

- Status Code: 200 OK
- Body :

```json
{
    "All Transaksi": [
        {
            "NamaPelanggan": "string",
            "NamaLayanan": "string",
            "Quantity": int,
            "NamaPegawai": "string",
            "TanggalMasuk": "string",
            "Harga": int
        },
        {
            "NamaPelanggan": "string",
            "NamaLayanan": "string",
            "Quantity": int,
            "NamaPegawai": "string",
            "TanggalMasuk": "string",
            "Harga": int
        },
        {
            "NamaPelanggan": "string",
            "NamaLayanan": "string",
            "Quantity": int,
            "NamaPegawai": "string",
            "TanggalMasuk": "string",
            "Harga": int
        }
    ]
}
```

### Service API

#### Create Service

Request :

- Method : POST
- Endpoint : `/api/services`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Body :

```json
{
	"NamaLayanan": "string",
  "Satuan": "string",
  "Harga": int
}
```

Response :

- Status Code: 201 Created
- Body:

```json
{
    "Id": int,
    "NamaLayanan": "string",
    "Satuan": "string",
    "Harga": int
}
```

#### List Service

Request :

- Method : GET
- Endpoint : `/api/services`
  - Header :
  - Accept : application/json
- Query Param :
  - productName : string `optional`,

Response :

- Status Code : 200 OK
- Body:

```json
[
    {
        "Id": int,
        "NamaLayanan": "string",
        "Satuan": "string",
        "Harga": int
    },
    {
        "Id": int,
        "NamaLayanan": "string",
        "Satuan": "string",
        "Harga": int
    },
    {
        "Id": int,
        "NamaLayanan": "string",
        "Satuan": "string",
        "Harga": int
    },
    {
        "Id": int,
        "NamaLayanan": "string",
        "Satuan": "string",
        "Harga": int
    }
]
```

#### Service By Id

Request :

- Method : GET
- Endpoint : `/api/service/:id`
- Header :
  - Accept : application/json

Response :

- Status Code: 200 OK
- Body :

```json
[
    {
        "Id": int,
        "NamaLayanan": "string",
        "Satuan": "string",
        "Harga": int
    }
]
```

#### Update Service

Request :

- Method : PUT
- Endpoint : `/api/services/:id`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Body :

```json
{
    "NamaLayanan": "string",
    "Satuan": "string",
    "Harga": int
}
```

Response :

- Status Code: 200 OK
- Body :

```json
{
    "Id": int,
    "NamaLayanan": "string",
    "Satuan": "string",
    "Harga": int
}
```

#### Delete Service

Request :

- Method : DELETE
- Endpoint : `/api/services/:id`
- Header :
  - Accept : application/json
- Body :

Response :

- Status : 200 OK
- Body :

```json
{
  "message": "string"
}
```

### Employe API

#### Create Employe

Request :

- Method : POST
- Endpoint : `/api/employees`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Body :

```json
{
  "Name": "string"
}
```

Response :

- Status Code: 201 Created
- Body:

```json
{
    "Id": int,
    "Name": "string"
}
```

#### List Service

Request :

- Method : GET
- Endpoint : `/api/employees`
  - Header :
  - Accept : application/json

Response :

- Status Code : 200 OK
- Body:

```json
[
    {
        "Id": int,
        "Name": "string"
    },
    {
        "Id": int,
        "Name": "string"
    },
    {
        "Id": int,
        "Name": "string"
    },
    {
        "Id": int,
        "Name": "string"
    }
]
```

#### employe By Id

Request :

- Method : GET
- Endpoint : `/api/employees/:id`
- Header :
  - Accept : application/json

Response :

- Status Code: 200 OK
- Body :

```json
[
    {
        "Id": int,
        "Name": "string"
    }
]
```

#### Update Service

Request :

- Method : PUT
- Endpoint : `/api/employees/:id`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Body :

```json
{
  "Name": "string"
}
```

Response :

- Status Code: 200 OK
- Body :

```json
{
    "Id": int,
    "Name": "string"
}
```

#### Delete Service

Request :

- Method : DELETE
- Endpoint : `/api/employees/:id`
- Header :
  - Accept : application/json
- Body :

Response :

- Status : 200 OK
- Body :

```json
{
  "message": "string"
}
```

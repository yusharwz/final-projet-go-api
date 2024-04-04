# Aplikasi Enigma Laundry

## API Spec

### Customer API

#### Create Customer

Request :

- Method : `POST`
- Endpoint : `/api/customers/add`

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
- Endpoint : `/api/customers/search/:id`
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
- Endpoint : `/api/customers/update/:id`
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
- Endpoint : `/api/customers/delete/:id`
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
- Endpoint : `/api/transactions/add`
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

#### Get Transaction

Request :

Pattern string date : `yyyy-mm-dd`

- Method : GET
- Endpoint : `/api/transactions/search/id/:id`
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
        }
    ]
}


```

- Method : GET
- Endpoint : `/api/transactions/search/users/id/:id`
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
        }
    ]
}
```

- Method : GET
- Endpoint : `/api/transactions/search/users/name/:name`
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
        }
    ]
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
- Endpoint : `/api/services/add`
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
- Endpoint : `/api/service`
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
- Endpoint : `/api/service/search/:id`
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
- Endpoint : `/api/services/update/:id`
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
- Endpoint : `/api/services/delete/:id`
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
- Endpoint : `/api/employees/add`
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
- Endpoint : `/api/employees/search/:id`
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
- Endpoint : `/api/employees/update/:id`
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
- Endpoint : `/api/employees/delete/:id`
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

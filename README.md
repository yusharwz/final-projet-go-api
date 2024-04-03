# Aplikasi Enigma Laundry

## API Spec

### Customer API

#### Create Customer

Request :

- Method : `POST`
- Endpoint : `/customers/add`

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
  "data": {
    "id": "string",
    "Name": "string",
    "NoHp": "string"
  }
}
```

#### Get Customer

Request :

- Method : GET
- Endpoint : `/customers/:id`
- Header :
  - Accept : application/json

Response :

- Status : 200 OK
- Body :

```json
{
  "message": "string",
  "data": {
    "Id": "string",
    "Name": "string",
    "NoHp": "string"
  }
}
```

#### Update Customer

Request :

- Method : PUT
- Endpoint : `/customers/update/:id`
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
- Endpoint : `/customers/delete/:id`
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
- Endpoint : `/transactions`
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

Request :

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

- Method : GET
- Endpoint : `/transactions/:id_bill`
- Header :
  - Accept : application/json
- Body :

Response :

- Status Code: 200 OK
- Body :

```json
{
	"message": "string",
  "data": {
    "id": "string",
    "billDate": "string",
    "entryDate": "string",
    "finishDate": "string",
    "employee": {
      "id": "string",
      "name": "string",
      "phoneNumber": "string",
      "address": "string"
    },
    "customer": {
      "id": "string",
      "name": "string",
      "phoneNumber": "string",
      "address": "string"
    },
    "billDetails": [
      {
        "id": "string",
        "billId": "string",
        "product": {
          "id": "string",
          "name": "string",
          "price": int,
          "unit": "string" (satuan product,cth: Buah atau Kg)
        },
        "productPrice": int,
        "qty": int
      }
    ],
    "totalBill": int
  }
}
```

#### List Transaction

Pattern string date : `dd-MM-yyyy`

Request :

- Method : GET
- Endpoint : `/transactions`
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
	"message": "string",
  "data": [
    {
      "id": "string",
      "billDate": "string",
      "entryDate": "string",
      "finishDate": "string",
      "employee": {
        "id": "string",
        "name": "string",
        "phoneNumber": "string",
        "address": "string"
      },
      "customer": {
        "id": "string",
        "name": "string",
        "phoneNumber": "string",
        "address": "string"
      },
      "billDetails": [
        {
          "id": "string",
          "billId": "string",
          "product": {
            "id": "string",
            "name": "string",
            "price": int,
            "unit": "string" (satuan product,cth: Buah atau Kg)
          },
          "productPrice": int,
          "qty": int
        }
      ],
      "totalBill": int
    }
  ]
}
```

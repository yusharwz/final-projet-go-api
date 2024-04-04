package entity

type Customers struct {
	Id   int
	Name string
	NoHp string
}

type Pegawai struct {
	Id   int
	Name string
}

type Layanan struct {
	Id          int
	NamaLayanan string
	Satuan      string
	Harga       int
}

type Transaksi struct {
	Id               int
	IdPelanggan      int
	IdPegawai        int
	TanggalMasuk     string
	TanggalKeluar    string
	StatusPembayaran string
}

type TransaksiAndDetail struct {
	Id               int
	IdPelanggan      int
	IdPegawai        int
	TanggalMasuk     string
	TanggalKeluar    string
	StatusPembayaran string
	DetailTransaksi  []DetailTransaction
}

type DetailTransaction struct {
	Id          int
	IdTransaksi int
	IdLayanan   int
	Quantity    int
}

type DetailTransaksi struct {
	Id          int
	IdTransaksi int
	IdLayanan   int
	Quantity    int
	TotalHarga  int
}

type TransactionDetail struct {
	NamaPelanggan string
	NamaLayanan   string
	Quantity      int
	NamaPegawai   string
	TanggalMasuk  string
	Harga         int
}

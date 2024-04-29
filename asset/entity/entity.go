package entity

type Customers struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	NoHp string `json:"no_hp"`
}

type Pegawai struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Layanan struct {
	ID          int    `json:"id"`
	NamaLayanan string `json:"nama_layanan"`
	Satuan      string `json:"satuan"`
	Harga       int    `json:"harga"`
}

type Transaksi struct {
	ID               int    `json:"id"`
	IDPelanggan      int    `json:"id_pelanggan"`
	IDPegawai        int    `json:"id_pegawai"`
	TanggalMasuk     string `json:"tanggal_masuk"`
	TanggalKeluar    string `json:"tanggal_keluar"`
	StatusPembayaran string `json:"status_pembayaran"`
}

type DetailTransaksi struct {
	ID          int `json:"id"`
	IDTransaksi int `json:"id_transaksi"`
	IDLayanan   int `json:"id_layanan"`
	Quantity    int `json:"quantity"`
}

type TransactionDetail struct {
	NamaPelanggan string `json:"nama_pelanggan"`
	NamaLayanan   string `json:"nama_layanan"`
	Quantity      int    `json:"quantity"`
	NamaPegawai   string `json:"nama_pegawai"`
	TanggalMasuk  string `json:"tanggal_masuk"`
	Harga         int    `json:"harga"`
}

type TransaksiAndDetail struct {
	ID               int               `json:"id"`
	IDPelanggan      int               `json:"id_pelanggan"`
	IDPegawai        int               `json:"id_pegawai"`
	TanggalMasuk     string            `json:"tanggal_masuk"`
	TanggalKeluar    string            `json:"tanggal_keluar"`
	StatusPembayaran string            `json:"status_pembayaran"`
	DetailTransaksi  []DetailTransaksi `json:"detail_transaksi"`
}

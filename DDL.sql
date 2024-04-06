CREATE TABLE mst_pelanggan (
    id SERIAL PRIMARY KEY,
    nama_pelanggan VARCHAR(255),
    nomor_hp VARCHAR(15)
);
CREATE TABLE mst_pegawai (
    id SERIAL PRIMARY KEY,
    nama_pegawai VARCHAR(255)
);
CREATE TABLE layanan (
    id SERIAL PRIMARY KEY,
    nama_layanan VARCHAR(255),
    satuan VARCHAR(10),
    harga INTEGER
);
CREATE TABLE transaksi (
    id SERIAL PRIMARY KEY,
    id_pelanggan INTEGER REFERENCES mst_pelanggan(id),
    id_pegawai INTEGER REFERENCES mst_pegawai(id),
    tanggal_masuk TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    tanggal_keluar DATE,
    status_pembayaran VARCHAR(50)
);
CREATE TABLE detail_transaksi (
    id SERIAL PRIMARY KEY,
    id_transaksi INTEGER REFERENCES transaksi(id),
    id_layanan INTEGER REFERENCES layanan(id),
    quantity INTEGER
);

CREATE TABLE auth (
    username VARCHAR(255) PRIMARY KEY,
    password VARCHAR(255)
);
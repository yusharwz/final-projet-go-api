INSERT INTO mst_pelanggan (nama_pelanggan, nomor_hp)
VALUES 
    ('Yushar', '081234567891'),
    ('Dinda', '082345678902'),
    ('Annisa', '083456789013'),
    ('Rizal', '084567890124'),
    ('Tegar', '085678901235'),
    ('Alifia', '086789012346'),
    ('Jed', '087890123457'),
    ('Krisna', '088901234568'),
    ('Clarissa', '089012345679'),
    ('Farid', '080123456780');

INSERT INTO mst_pegawai (nama_pegawai)
VALUES 
    ('Diah'),
    ('Calista'),
    ('Kevin');

INSERT INTO layanan (nama_layanan, satuan, harga)
VALUES 
    ('Cuci', 'KG', 5000),
    ('Cuci dan setrika', 'KG', 7000),
    ('Cuci bedcover', 'Buah', 50000),
    ('Cuci boneka', 'Buah', 25000);

INSERT INTO transaksi (id_pelanggan, id_pegawai, tanggal_masuk, tanggal_keluar, status_pembayaran)
VALUES 
    (1, 1, '2024-03-25', '2024-03-27', 'Lunas'),
    (2, 2, '2024-03-26', '2024-03-28', 'Belum Lunas'),
    (3, 3, '2024-03-27', '2024-03-29', 'Belum Lunas');

INSERT INTO detail_transaksi (id_transaksi, id_layanan, quantity)
VALUES 
    (1, 1, 2),
    (1, 3, 1),
    (1, 4, 3),
    (2, 2, 1),
    (2, 3, 2),
    (2, 4, 1),
    (3, 1, 3),
    (3, 2, 2);
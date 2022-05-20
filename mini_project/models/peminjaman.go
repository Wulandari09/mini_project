package main

type peminjaman struct{
	no_pinjam int `json:"no_pinjam"`
	tanggal_pinjam string `json:"tanggal_pinjam"`
	tanggal_kembali string `json:"tanggal_kembali"`
	total_pinjam int `json:"total_pinjam"`
	id_anggota int `json:"id_anggota"`
	id_petugas int `json:"id_petugas"`
}

func table_peminjaman() string{
	return "peminjaman"
}
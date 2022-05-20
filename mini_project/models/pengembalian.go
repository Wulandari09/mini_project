package main

type pengembalian struct{
	no_kembali int `json:"no_kembali"`
	no_pinjam int `json:"no_pinjam"`
	tanggal_kembali string `json:"tanggal_kembali"`
	aktual_kemballi string `json:"aktual_kembali"`
	total_denda string `json:"total_denda"`
	id_anggota int `json:"id_anggota"`
	id_petugas int `json:"id_petugas"`
}

func table_pengembalian() string{
	return "pengembalian"
}
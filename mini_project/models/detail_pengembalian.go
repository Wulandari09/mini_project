package main

type detail_pengembalian struct{
	no_pinjam int `json:"no_pinjam"`
	jumlah_kembali int `json:"jumlah_pinjam"`
	lama_denda string `json:"lama_denda"`
	jumlah_denda string `json:"jumlah_denda"`
}

func table_detail_pengembalian() string{
	return "pengembalian"
}
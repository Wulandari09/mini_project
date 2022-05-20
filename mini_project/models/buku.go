package main

type buku struct{
	kode_buku	int `json:"kode_buku"`
	judul_buku  string `json:"judul_buku"`
	no_jenis_buku int `json:"no_jenis_buku"`
	pengarang string `json:"penerbit"`
	penerbit string `json:"penerbit"`
	tahun_terbit int `json:"tahun_terbit"`
	stock int `json:"stock"`
}

func table_buku() string {
	return "buku"
}
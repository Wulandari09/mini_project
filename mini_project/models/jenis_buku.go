package main

type jenis_buku struct{
	no_jenis_buku int `json:"no_jenis_buku"`
	jenis_buku string `json:"jenis_buku"`
	no_rak int `json:"no_rak"`
}

func table_jenis_buku() string {
	return "jenis_buku"
}
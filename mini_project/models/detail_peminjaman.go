package main


type detail_peminjaman struct{
	no_pinjam int `json:"no_pinjam"`
	kode_buku int `json:"kode_buku"`
	jumlah_pinjam string `json:"jumlah_pinjam"`
}

func table_detail_peminjaman() string{
	return "detail_peminjaman"
}
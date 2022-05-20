package main

type petugas struct{
	id_petugas 		int		`json:"id_petugas"`
	nama_petugas 	string `json:"nama_petugas"`
	alamat_petugas  string `json:"alamat_petugas"`
	no_telepon		int    `json:"no_telepon"`
	password		string  `json:"password"`
}

func table_petugas() string{
	return "petugas"
}
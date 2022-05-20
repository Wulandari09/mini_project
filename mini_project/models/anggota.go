package main

type anggota struct{
	id_anggota  	int	`json:"id_anggota"`
	nama_anggota 	string `json:"nama_anggota"`
	jenis_makanan 	string `json: "jenis_makanan"`
	tanggal_lahir 	string `json:"tanggal_lahir"`
	prodi 			string `json:"prodi"`
	tanggal_daftar 	string `json: "tanggal_daftar"`
	tanggal_expired string `json: "tangga_expired"`

}

func table_anggota() string {
	return "anggota"
}

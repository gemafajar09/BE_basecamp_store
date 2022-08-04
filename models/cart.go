package models

type Cart struct {
	Id        int64  `json:"id"`
	Id_produk int    `json:"id_produk"`
	Id_user   int    `json:"id_user"`
	Jumlah    int    `json:"jumlah"`
	Tanggal   string `json:"tanggal"`
}

type MinOrMax struct {
	Id     int64  `json:"id"`
	Jumlah int    `json:"jumlah"`
	Kode   string `json:"kode"`
}

type CartDetail struct {
	Id          int64   `json:"id"`
	Id_produk   int     `json:"id_produk"`
	Id_user     int     `json:"id_user"`
	Jumlah      int     `json:"jumlah"`
	Tanggal     string  `json:"tanggal"`
	Nama_produk string  `json:"nama_produk"`
	Deskripsi   string  `json:"deskripsi"`
	Harga       int     `json:"harga"`
	Stok        int     `json:"stok"`
	Gambar      string  `json:"gambar"`
	Raiting     float32 `json:"raiting"`
}

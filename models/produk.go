package models

type Produk struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	Nama_produk string  `json:"nama_produk"`
	Deskripsi   string  `json:"deskripsi"`
	Harga       int     `json:"harga"`
	Stok        int     `json:"stok"`
	Gambar      string  `json:"gambar"`
	Raiting     float32 `json:"raiting"`
}

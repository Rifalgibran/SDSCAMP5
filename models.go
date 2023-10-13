package models

type User struct {
	ID           uint   `json:"user_id" gorm:"primaryKey;autoIncrement:true"`
	Nama         string `json:"nama"`
	NomorTelepon string `json:"nomor_telepon"`
	AlamatEmail  string `json:"alamat_email"`
}

type Menu struct {
	ID        uint   `json:"menu_id" gorm:"primaryKey;autoIncrement:true"`
	NamaMenu  string `json:"nama_menu"`
	Harga     int    `json:"harga"`
	Deskripsi string `json:"deskripsi"`
}

type Pesanan struct {
	ID           uint   `json:"pesanan_id" gorm:"primaryKey;autoIncrement:true"`
	TanggalPesan string `json:"tanggal_pesan"`
	UserID       uint   `json:"user_id" gorm:"foreignKey:ID"`
}

type PemesananItem struct {
	ID           uint `json:"pemesanan_item_id" gorm:"primaryKey;autoIncrement:true"`
	PesananID    uint `json:"pesanan_id" gorm:"foreignKey:ID"`
	MenuID       uint `json:"menu_id" gorm:"foreignKey:ID"`
	Jumlah       int  `json:"jumlah"`
	HargaPerItem int  `json:"harga_per_item"`
}

type Pembayaran struct {
	ID                uint   `json:"pembayaran_id" gorm:"primaryKey;autoIncrement:true"`
	PesananID         uint   `json:"pesanan_id" gorm:"foreignKey:ID"`
	TanggalPembayaran string `json:"tanggal_pembayaran"`
	MetodePembayaran  string `json:"metode_pembayaran"`
	TotalPembayaran   int    `json:"total_pembayaran"`
}

type Stok struct {
	ID         uint `json:"stok_id" gorm:"primaryKey;autoIncrement:true"`
	MenuID     uint `json:"menu_id" gorm:"foreignKey:ID"`
	JumlahStok int  `json:"jumlah_stok"`
}

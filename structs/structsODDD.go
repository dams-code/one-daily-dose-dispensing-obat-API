package structs

import "time"

type UserObat struct {
	ID         int        `json:"id"`
	Nama       string     `json:"nama"`
	Username   string     `json:"username"`
	Password   string     `json:"-"`
	Role       string     `json:"role"`
	Status     bool       `json:"status"`
	CreatedAt  time.Time  `json:"created_at"`
	CreatedBy  string     `json:"created_by"`
	ModifiedAt *time.Time `json:"modified_at"`
	ModifiedBy *string    `json:"modified_by"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterUserRequest struct {
	Nama     string `json:"nama" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role" binding:"required,oneof=KEPALAFARMASI ADMINFARMASI APOTEKER"`
}

type UpdateUserRequest struct {
	Nama     *string `json:"nama"`
	Password *string `json:"password,omitempty"`
	Role     *string `json:"role" binding:"omitempty,oneof=KEPALAFARMASI ADMINFARMASI APOTEKER"`
}

type TambahUserRequest struct {
	Nama     string `json:"nama" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role" binding:"required,oneof=KEPALAFARMASI ADMINFARMASI APOTEKER"`
}

type Pasien struct {
	ID           int        `json:"id"`
	Rm           string     `json:"rm"`
	Nama         string     `json:"nama"`
	JenisKelamin string     `json:"jenis_kelamin"`
	Alamat       *string    `json:"alamat"`
	CreatedAt    time.Time  `json:"created_at"`
	CreatedBy    string     `json:"created_by"`
	ModifiedAt   *time.Time `json:"modified_at"`
	ModifiedBy   *string    `json:"modified_by"`
}

type TambahPasienRequest struct {
	Rm           string `json:"rm" binding:"required,len=8,numeric"`
	Nama         string `json:"nama" binding:"required"`
	JenisKelamin string `json:"jenis_kelamin" binding:"required"`
	Alamat       string `json:"alamat"`
}

type UpdatePasienRequest struct {
	Rm           string `json:"rm"`
	Nama         string `json:"nama"`
	JenisKelamin string `json:"jenis_kelamin"`
	Alamat       string `json:"alamat"`
}

type KategoriObat struct {
	ID           int        `json:"id"`
	NamaKategori string     `json:"nama_kategori"`
	Deskripsi    *string    `json:"deskripsi"`
	CreatedAt    time.Time  `json:"created_at"`
	CreatedBy    string     `json:"created_by"`
	ModifiedAt   *time.Time `json:"modified_at"`
	ModifiedBy   *string    `json:"modified_by"`
}

type TambahdanUpdateKategoriObatRequest struct {
	NamaKategori string  `json:"nama_kategori" binding:"required"`
	Deskripsi    *string `json:"deskripsi"`
}

type Obat struct {
	ID           int        `json:"id"`
	KategoriID   int        `json:"kategori_id"`
	NamaKategori *string    `json:"nama_kategori"`
	KodeObat     string     `json:"kode_obat"`
	NamaObat     string     `json:"nama_obat"`
	Stok         float64    `json:"stok"`
	Satuan       string     `json:"satuan"`
	Harga        float64    `json:"harga"`
	CreatedAt    time.Time  `json:"created_at"`
	CreatedBy    string     `json:"created_by"`
	ModifiedAt   *time.Time `json:"modified_at"`
	ModifiedBy   *string    `json:"modified_by"`
}

type TambahObatRequest struct {
	KategoriID int     `json:"kategori_id" binding:"required"`
	NamaObat   string  `json:"nama_obat" binding:"required"`
	Stok       float64 `json:"stok"`
	Satuan     string  `json:"satuan" binding:"required"`
	Harga      float64 `json:"harga"`
}

type UpdateObatRequest struct {
	KategoriID *int     `json:"kategori_id"`
	NamaObat   *string  `json:"nama_obat"`
	Stok       *float64 `json:"stok"`
	Satuan     *string  `json:"satuan"`
	Harga      *float64 `json:"harga"`
}

type ObatRes struct {
	Obat
	NamaKategori string `json:"nama_kategori"`
}

type TransaksiObat struct {
	ID              int        `json:"id"`
	NoResep         string     `json:"no_resep"`
	PasienID        int        `json:"pasien_id"`
	NamaPasien      string     `json:"nama_pasien"`
	Status          string     `json:"status"`
	TipePembayaran  *string    `json:"tipe_pembayaran"`
	GrandTotal      float64    `json:"grand_total"`
	TotalPembayaran *float64   `json:"total_pembayaran"`
	Kembalian       *float64   `json:"kembalian"`
	DibayarAt       *time.Time `json:"dibayar_at"`
	CreatedAt       time.Time  `json:"created_at"`
	CreatedBy       string     `json:"created_by"`
	ModifiedAt      *time.Time `json:"modified_at"`
	ModifiedBy      *string    `json:"modified_by"`
}

type TambahTransaksiObatRequest struct {
	PasienID int `json:"pasien_id" binding:"required"`
}

type PembayaranTransaksiObatRequest struct {
	TipePembayaran  string  `json:"tipe_pembayaran" binding:"required,oneof=TUNAI QRIS"`
	TotalPembayaran float64 `json:"total_pembayaran" binding:"required,gt=0"`
}

type UpdateTransaksiObatRequest struct {
	Status string `json:"status" binding:"omitempty,oneof=PENDING CANCELED"`
}

type DetailTransaksiObat struct {
	ID              int        `json:"id"`
	NoDetailResep   string     `json:"no_detail_resep"`
	TransaksiObatID int        `json:"transaksi_obat_id"`
	NoResep         string     `json:"no_resep"`
	NamaPasien      string     `json:"nama_pasien"`
	ObatID          int        `json:"obat_id"`
	KodeObat        string     `json:"kode_obat"`
	NamaObat        string     `json:"nama_obat"`
	Jumlah          float64    `json:"jumlah"`
	Satuan          string     `json:"satuan"`
	Harga           float64    `json:"harga"`
	Subtotal        float64    `json:"subtotal"`
	AturanPakai     string     `json:"aturan_pakai"`
	CreatedAt       time.Time  `json:"created_at"`
	CreatedBy       string     `json:"created_by"`
	ModifiedAt      *time.Time `json:"modified_at"`
	ModifiedBy      *string    `json:"modified_by"`
}

type DetailTransaksiObatResponse struct {
	ID              int                   `json:"id"`
	NoResep         string                `json:"no_resep"`
	NamaPasien      string                `json:"nama_pasien"`
	Status          string                `json:"status"`
	CreatedAt       time.Time             `json:"created_at"`
	TipePembayaran  *string               `json:"tipe_pembayaran"`
	TotalPembayaran *float64              `json:"total_pembayaran"`
	Kembalian       *float64              `json:"kembalian"`
	GrandTotal      float64               `json:"grand_total"`
	DetailObat      []DetailTransaksiObat `json:"detail_obat"`
}

type DetailTransaksiObatItem struct {
	ObatID      int     `json:"obat_id"`
	Jumlah      float64 `json:"jumlah"`
	AturanPakai string  `json:"aturan_pakai"`
}

type TambahDetailTransaksiObatRequest struct {
	ItemObat []DetailTransaksiObatItem `json:"itemObat" binding:"required,min=1"`
}

type UpdateDetailTransaksiObatRequest struct {
	Jumlah      *float64 `json:"jumlah"`
	AturanPakai *string  `json:"aturan_pakai"`
}

type LaporanTransaksiObat struct {
	ID              int       `json:"id"`
	NoResep         string    `json:"no_resep"`
	NamaPasien      string    `json:"nama_pasien"`
	Status          string    `json:"status"`
	TotalItem       int       `json:"total_item"`
	TotalHarga      float64   `json:"total_harga"`
	TipePembayaran  *string   `json:"tipe_pembayaran"`
	TotalPembayaran *float64  `json:"total_pembayaran"`
	Kembalian       *float64  `json:"kembalian"`
	CreatedAt       time.Time `json:"created_at"`
	CreatedBy       string    `json:"created_by"`
}

type Dispensing struct {
	ID              int       `json:"id"`
	TransaksiObatID int       `json:"transaksi_obat_id"`
	ApotekerID      int       `json:"apoteker_id"`
	TanggalDispense time.Time `json:"tanggal_dispense"`
	CreatedAt       time.Time `json:"created_at"`
}

//swagger
type ErrorResponse struct {
	Status string `json:"status" example:"error"`
	Pesan  string `json:"pesan" example:"Terjadi kesalahan"`
}

type LoginResponse struct {
	Status string `json:"status" example:"sukses"`
	Pesan  string `json:"pesan" example:"Hi, dodo Role: ADMINFARMASI"`
	Token  string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}

type RegisterUserResponse struct {
	Status string   `json:"status" example:"sukses"`
	Pesan  string   `json:"pesan" example:"registrasi user berhasil"`
	Data   UserObat `json:"data"`
}

type GetUserResponse struct {
	Status string     `json:"status" example:"sukses"`
	Pesan  string     `json:"pesan" example:"Get user berhasil"`
	Data   []UserObat `json:"data"`
}

type GetUserByIDResponse struct {
	Status string   `json:"status" example:"sukses"`
	Pesan  string   `json:"pesan" example:"Get user by ID berhasil"`
	Data   UserObat `json:"data"`
}

type TambahUserResponse struct {
	Status string   `json:"status" example:"sukses"`
	Pesan  string   `json:"pesan" example:"User berhasil ditambahkan"`
	Data   UserObat `json:"data"`
}

type DeleteUserResponse struct {
	Status string `json:"status" example:"sukses"`
	Pesan  string `json:"pesan" example:"User berhasil dihapus"`
}

type UpdateUserResponse struct {
	Status string   `json:"status" example:"sukses"`
	Pesan  string   `json:"pesan" example:"Update user berhasil"`
	Data   UserObat `json:"data"`
}

type AktivasiUserResponse struct {
	Status string   `json:"status" example:"sukses"`
	Pesan  string   `json:"pesan" example:"Aktivasi user berhasil"`
	Data   UserObat `json:"data"`
}

type GetObatResponse struct {
	Status string `json:"status" example:"sukses"`
	Pesan  string `json:"pesan" example:"Get data obat berhasil"`
	Data   []Obat `json:"data"`
}

type GetObatByIDResponse struct {
	Status string `json:"status" example:"sukses"`
	Pesan  string `json:"pesan" example:"Get data obat by ID berhasil"`
	Data   Obat   `json:"data"`
}

type TambahObatResponse struct {
	Status string `json:"status" example:"sukses"`
	Pesan  string `json:"pesan" example:"Tambah data obat berhasil"`
	Data   Obat   `json:"data"`
}

type UpdateObatResponse struct {
	Status string `json:"status" example:"sukses"`
	Pesan  string `json:"pesan" example:"Update data obat berhasil"`
	Data   Obat   `json:"data"`
}

type DeleteObatResponse struct {
	Status string `json:"status" example:"sukses"`
	Pesan  string `json:"pesan" example:"Obat berhasil dihapus"`
}

type GetKategoriObatResponse struct {
	Status string         `json:"status" example:"sukses"`
	Pesan  string         `json:"pesan" example:"Get kategori obat berhasil"`
	Data   []KategoriObat `json:"data"`
}

type GetKategoriObatByIDResponse struct {
	Status string       `json:"status" example:"sukses"`
	Pesan  string       `json:"pesan" example:"Get kategori obat by ID berhasil"`
	Data   KategoriObat `json:"data"`
}

type TambahKategoriObatResponse struct {
	Status string       `json:"status" example:"sukses"`
	Pesan  string       `json:"pesan" example:"Kategori obat berhasil ditambahkan"`
	Data   KategoriObat `json:"data"`
}

type UpdateKategoriObatResponse struct {
	Status string       `json:"status" example:"sukses"`
	Pesan  string       `json:"pesan" example:"Kategori obat berhasil diperbarui"`
	Data   KategoriObat `json:"data"`
}

type DeleteKategoriObatResponse struct {
	Status string `json:"status" example:"sukses"`
	Pesan  string `json:"pesan" example:"Kategori obat berhasil dihapus"`
}

type GetPasienResponse struct {
	Status string   `json:"status" example:"sukses"`
	Pesan  string   `json:"pesan" example:"Get pasien berhasil"`
	Data   []Pasien `json:"data"`
}

type GetPasienByIDResponse struct {
	Status string `json:"status" example:"sukses"`
	Pesan  string `json:"pesan" example:"Get pasien by ID berhasil"`
	Data   Pasien `json:"data"`
}

type TambahPasienResponse struct {
	Status string `json:"status" example:"sukses"`
	Pesan  string `json:"pesan" example:"Tambah pasien berhasil"`
	Data   Pasien `json:"data"`
}

type UpdatePasienResponse struct {
	Status string `json:"status" example:"sukses"`
	Pesan  string `json:"pesan" example:"Update pasien berhasil"`
	Data   Pasien `json:"data"`
}

type DeletePasienResponse struct {
	Status string `json:"status" example:"sukses"`
	Pesan  string `json:"pesan" example:"Pasien berhasil dihapus"`
}

type TambahTransaksiObatResponse struct {
	Status string        `json:"status" example:"sukses"`
	Pesan  string        `json:"pesan" example:"Transaksi obat berhasil ditambahkan"`
	Data   TransaksiObat `json:"data"`
}

type GetTransaksiObatResponse struct {
	Status string          `json:"status" example:"sukses"`
	Pesan  string          `json:"pesan" example:"Get transaksi obat berhasil"`
	Data   []TransaksiObat `json:"data"`
}

type GetTransaksiObatByIDResponse struct {
	Status string        `json:"status" example:"sukses"`
	Pesan  string        `json:"pesan" example:"Get transaksi obat by ID berhasil"`
	Data   TransaksiObat `json:"data"`
}

type DispensingObatResponse struct {
	Status string        `json:"status" example:"sukses"`
	Pesan  string        `json:"pesan" example:"Dispensing obat berhasil"`
	Data   TransaksiObat `json:"data"`
}

type DeleteTransaksiObatResponse struct {
	Status string        `json:"status" example:"sukses"`
	Pesan  string        `json:"pesan" example:"Transaksi resep obat berhasil dihapus."`
	Data   TransaksiObat `json:"data"`
}

type CancelTransaksiObatResponse struct {
	Status string        `json:"status" example:"sukses"`
	Pesan  string        `json:"pesan" example:"Transaksi obat berhasil dibatalkan"`
	Data   TransaksiObat `json:"data"`
}

type UpdateTransaksiObatResponse struct {
	Status string        `json:"status" example:"sukses"`
	Pesan  string        `json:"pesan" example:"Update data transaksi resep obat pasien berhasil"`
	Data   TransaksiObat `json:"data"`
}

type PembayaranTransaksiObatResponse struct {
	Status string        `json:"status" example:"sukses"`
	Pesan  string        `json:"pesan" example:"Pembayaran berhasil"`
	Data   TransaksiObat `json:"data"`
}

type TambahDetailTransaksiObatResponse struct {
	Status string                `json:"status" example:"sukses"`
	Pesan  string                `json:"pesan" example:"Detail transaksi obat berhasil ditambahkan"`
	Data   []DetailTransaksiObat `json:"data"`
}

type GetDetailTransaksiObatResponse struct {
	Status string                `json:"status" example:"sukses"`
	Pesan  string                `json:"pesan" example:"Get detail transaksi obat berhasil"`
	Data   []DetailTransaksiObat `json:"data"`
}

type GetDetailTransaksiObatByIDResponse struct {
	Status string              `json:"status" example:"sukses"`
	Pesan  string              `json:"pesan" example:"Get detail transaksi obat by ID berhasil"`
	Data   DetailTransaksiObat `json:"data"`
}

type UpdateDetailTransaksiObatResponse struct {
	Status string              `json:"status" example:"sukses"`
	Pesan  string              `json:"pesan" example:"Detail transaksi obat berhasil diperbarui"`
	Data   DetailTransaksiObat `json:"data"`
}

type DeleteDetailTransaksiObatResponse struct {
	Status string `json:"status" example:"sukses"`
	Pesan  string `json:"pesan" example:"Detail transaksi obat berhasil dihapus"`
}

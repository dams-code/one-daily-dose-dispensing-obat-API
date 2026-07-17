package controllers

import (
	"fmt"
	"net/http"
	"one-daily-dose-dispensing-obat-api/repositories"
	"one-daily-dose-dispensing-obat-api/structs"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetDetailTransaksiObat godoc
//
// @Summary Get Semua Detail Transaksi Resep Obat
// @Description Mengambil seluruh data detail transaksi resep obat
// @Tags 7. Detail Transaksi Obat
// @Produce json
// @Security BearerAuth
// @Success 200 {object} structs.GetDetailTransaksiObatResponse "Get detail transaksi resep obat pasien (berisikan list item obat) berhasil"
// @Failure 500 {object} structs.ErrorResponse "Internal Server Error"
// @Router /api/admin-farmasi/transaksi-obat/detail [get]
func GetDetailTransaksiObat(ctx *gin.Context) {

	IdTransaksiObat, _ := strconv.Atoi(ctx.Query("transaksi_obat_id"))

	NoDetailResep := ctx.Query("no_detail_resep")

	hasilGetData, err := repositories.GetDetailTransaksiObat(DBSqlConn, IdTransaksiObat, NoDetailResep)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "sukses",
		"pesan":  "Get detail transaksi resep obat pasien (berisikan list item obat) berhasil",
		"data":   hasilGetData,
	})

}

// GetDetailTransaksiObatByID godoc
//
// @Summary Get Detail Transaksi Obat By ID
// @Description Mengambil detail transaksi obat berdasarkan ID Detail
// @Tags 7. Detail Transaksi Obat
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID Detail Transaksi"
// @Success 200 {object} structs.GetDetailTransaksiObatByIDResponse "Get daftar item obat by ID pada detail transaksi resep obat pasien berhasil"
// @Failure 400 {object} structs.ErrorResponse "Gagal get daftar item obat pada detail transaksi obat"
// @Failure 404 {object} structs.ErrorResponse "Detail Transaksi resep obat tidak ditemukan"
// @Failure 500 {object} structs.ErrorResponse "Internal Server Error"
// @Router /api/admin-farmasi/transaksi-obat/detail/{id} [get]
// @Router /api/apoteker/transaksi-obat/detail/{id} [get]
func GetDetailTransaksiObatByID(ctx *gin.Context) {

	IdDetailTransaksiObat, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	hasilGetData, err := repositories.GetDetailTransaksiObatByID(
		DBSqlConn,
		IdDetailTransaksiObat,
	)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "sukses",
		"pesan":  "Get daftar item obat by ID pada detail transaksi resep obat pasien berhasil",
		"data":   hasilGetData,
	})
}

// GetDetailTransaksiObatByTransaksiID godoc
//
// @Summary Get Detail Transaksi Berdasarkan ID Transaksi resep obat (header)
// @Description Mengambil seluruh detail obat berdasarkan transaksi resep obat
// @Tags 7. Detail Transaksi Obat
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID Transaksi Obat"
// @Success 200 {object} structs.GetDetailTransaksiObatResponse "Get Id resep obat (header transaksi obat) beserta isi item obat-nya berhasil"
// @Failure 400 {object} structs.ErrorResponse "ID transaksi obat tidak ditemukan"
// @Failure 401 {object} structs.ErrorResponse "Transaksi obat dan detailnya tidak ada"
// @Failure 500 {object} structs.ErrorResponse "Internal Server Error"
// @Router /api/admin-farmasi/transaksi-obat/{id}/detail [get]
// @Router /api/apoteker/transaksi-obat/{id}/detail [get]
func GetDetailTransaksiObatByTransaksiID(ctx *gin.Context) {

	IdTransaksiObat, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	hasilGetData, err := repositories.GetDetailTransaksiObatByTransaksiID(DBSqlConn, IdTransaksiObat)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "sukses",
		"pesan":  "Get Id resep obat (header transaksi obat) beserta isi item obat-nya berhasil",
		"data":   hasilGetData,
	})
}

// UpdateDetailTransaksiObat godoc
//
// @Summary Update Detail Transaksi Obat
// @Description Mengubah jumlah obat atau aturan pakai. Stok obat akan disesuaikan secara otomatis.
// @Tags 7. Detail Transaksi Obat
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID Detail Transaksi"
// @Param request body structs.UpdateDetailTransaksiObatRequest true "Data Detail"
// @Success 200 {object} structs.UpdateDetailTransaksiObatResponse "Update informasi item obat pada detail transaksi kedalam transaksi resep obat pasien berhasil"
// @Failure 400 {object} structs.ErrorResponse "ID transaksi obat tidak ditemukan"
// @Failure 401 {object} structs.ErrorResponse "Transaksi obat dan detailnya tidak ada"
// @Failure 500 {object} structs.ErrorResponse "Internal Server Error"
// @Router /api/admin-farmasi/transaksi-obat/detail/{id} [put]
func UpdateDetailTransaksiObat(ctx *gin.Context) {

	IdDetailTransaksiObat, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	var setDetailTransaksiObat structs.UpdateDetailTransaksiObatRequest

	if err := ctx.ShouldBindJSON(&setDetailTransaksiObat); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	cekUsernameJWT := ctx.MustGet("username_sekarang").(string)

	if cekUsernameJWT == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status": "error",
			"pesan":  "Akses update detail transaksi resep obat ditolak, login dulu sebagai admin farmasi.",
		})
		return
	}

	hasilUpdateData, err := repositories.UpdateDetailTransaksiObat(
		DBSqlConn,
		IdDetailTransaksiObat,
		setDetailTransaksiObat,
		cekUsernameJWT,
	)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "sukses",
		"pesan":  "Update informasi item obat pada detail transaksi kedalam transaksi resep obat pasien berhasil",
		"data":   hasilUpdateData,
	})
}

// TambahDetailTransaksiObat godoc
//
// @Summary Tambah Detail Transaksi Obat
// @Description Menambahkan satu atau lebih obat ke transaksi dan otomatis mengurangi stok obat
// @Tags 7. Detail Transaksi Obat
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID Transaksi Obat"
// @Param request body structs.TambahDetailTransaksiObatRequest true "Daftar Obat"
// @Success 201 {object} structs.TambahDetailTransaksiObatResponse "Tambah item obat kedalam detail transaksi resep obat pasien berhasil"
// @Failure 400 {object} structs.ErrorResponse "Gagal Get data detail transaksi resep obat"
// @Failure 401 {object} structs.ErrorResponse "Detail transaksi resep obat tidak ada"
// @Failure 500 {object} structs.ErrorResponse "Internal server error"
// @Router /api/admin-farmasi/transaksi-obat/{id}/detail [post]
func TambahDetailTransaksiObat(ctx *gin.Context) {

	IdTransaksiObat, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	var setDetailTransaksiObat structs.TambahDetailTransaksiObatRequest

	if err := ctx.ShouldBindJSON(&setDetailTransaksiObat); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	cekUsernameJWT := ctx.MustGet("username_sekarang").(string)

	if cekUsernameJWT == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status": "error",
			"pesan":  "Akses tambah detail transaksi resep obat ditolak, login dulu sebagai admin farmasi.",
		})
		return
	}

	hasilTambahData, err := repositories.TambahDetailTransaksiObat(DBSqlConn, IdTransaksiObat, setDetailTransaksiObat, cekUsernameJWT)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "sukses",
		"pesan":  "Tambah item obat kedalam detail transaksi resep obat pasien berhasil",
		"data":   hasilTambahData,
	})
}

// DeleteDetailTransaksiObat godoc
//
// @Summary Delete Detail Transaksi Obat
// @Description Menghapus detail transaksi obat dan mengembalikan stok obat secara otomatis.
// @Tags 7. Detail Transaksi Obat
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID Detail Transaksi"
// @Success 200 {object} structs.DeleteDetailTransaksiObatResponse
// @Failure 400 {object} structs.ErrorResponse "Gagal Get data detail transaksi resep obat"
// @Failure 401 {object} structs.ErrorResponse "Detail transaksi resep obat tidak ada"
// @Failure 500 {object} structs.ErrorResponse "Internal server error"
// @Router /api/admin-farmasi/transaksi-obat/detail/{id} [delete]
func DeleteDetailTransaksiObat(ctx *gin.Context) {

	IdDetailTransaksiObat, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	cekUsernameJWT := ctx.MustGet("username_sekarang").(string)

	if cekUsernameJWT == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status": "error",
			"pesan":  "Akses delete detail transaksi resep obat ditolak, login dulu sebagai admin farmasi.",
		})
		return
	}

	err = repositories.DeleteDetailTransaksiObat(
		DBSqlConn,
		IdDetailTransaksiObat,
		cekUsernameJWT,
	)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "sukses",
		"pesan":  fmt.Sprintf("Detail transaksi obat ID %d berhasil dihapus.", IdDetailTransaksiObat),
	})
}

package controllers

import (
	"fmt"
	"net/http"
	"one-daily-dose-dispensing-obat-api/repositories"
	"one-daily-dose-dispensing-obat-api/structs"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetTransaksiObat godoc
//
// @Summary Get Semua Transaksi Obat
// @Description Menampilkan daftar transaksi obat
// @Tags 6. Transaksi Obat
// @Produce json
// @Security BearerAuth
// @Success 200 {object} structs.GetTransaksiObatResponse "Berhasil Get daftar transaksi resep obat"
// @Failure 400 {object} structs.ErrorResponse "Gagal Get data transaksi resep obat"
// @Failure 401 {object} structs.ErrorResponse "Transaksi resep obat tidak ada"
// @Failure 500 {object} structs.ErrorResponse "Internal server error"
// @Router /api/admin-farmasi/transaksi-obat [get]
func GetTransaksiObat(ctx *gin.Context) {

	IdTransaksiObat, _ := strconv.Atoi(ctx.Query("id"))

	NoResep := ctx.Query("no_resep")

	hasilGetData, err := repositories.GetTransaksiObat(DBSqlConn, IdTransaksiObat, NoResep)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "sukses",
		"pesan":  "Get data transaksi resep obat berhasil",
		"data":   hasilGetData,
	})

}

// GetTransaksiObatByID godoc
//
// @Summary Get Transaksi Obat By ID
// @Description Menampilkan detail transaksi obat
// @Tags 6. Transaksi Obat
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID Transaksi"
// @Success 200 {object} structs.GetTransaksiObatByIDResponse "Get transaksi resep obat by ID berhasil"
// @Failure 400 {object} structs.ErrorResponse "Gagal Get data transaksi resep obat by ID"
// @Failure 401 {object} structs.ErrorResponse "ID  Transaksi resep obat tidak ditemukan"
// @Failure 500 {object} structs.ErrorResponse "Internal server error"
// @Router /api/admin-farmasi/transaksi-obat/{id} [get]
func GetTransaksiObatByID(ctx *gin.Context) {

	IdTransaksiObat, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	hasilGetData, err := repositories.GetTransaksiObatByID(DBSqlConn, IdTransaksiObat)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "sukses",
		"pesan":  "Get data transaksi resep obat by ID berhasil",
		"data":   hasilGetData,
	})

}

// TambahTransaksiObat godoc
//
// @Summary Tambah Transaksi Obat
// @Description Membuat transaksi obat baru
// @Tags 6. Transaksi Obat
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body structs.TambahTransaksiObatRequest true "Data Transaksi Obat"
// @Success 201 {object} structs.TambahTransaksiObatResponse "Tambah data transaksi resep obat pasien berhasil"
// @Failure 400 {object} structs.ErrorResponse "Gagal tambah data transaksi resep obat"
// @Failure 401 {object} structs.ErrorResponse "Transaksi resep obat sudah dibuat"
// @Failure 500 {object} structs.ErrorResponse "Internal server error"
// @Router /api/admin-farmasi/transaksi-obat [post]
func TambahTransaksiObat(ctx *gin.Context) {

	var setTransaksiObat structs.TambahTransaksiObatRequest

	if err := ctx.ShouldBindJSON(&setTransaksiObat); err != nil {
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
			"pesan":  "Akses tambah transaksi resep obat ditolak, login dulu sebagai admin farmasi.",
		})
		return
	}

	hasilTambahData, err := repositories.TambahTransaksiObat(DBSqlConn, setTransaksiObat, cekUsernameJWT)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": "sukses",
		"pesan":  "Tambah data transaksi resep obat pasien berhasil",
		"data":   hasilTambahData,
	})
}

// UpdateTransaksiObat godoc
//
// @Summary      Update Transaksi Obat
// @Description  Mengubah data transaksi obat
// @Tags         6. Transaksi Obat
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "ID Transaksi Obat"
// @Param        request body structs.UpdateTransaksiObatRequest true "Data Transaksi resep obat"
// @Success      200 {object} structs.UpdateTransaksiObatResponse "Update data transaksi resep obat pasien berhasil"
// @Failure      400 {object} structs.ErrorResponse "Gagal update transaksi resep obat"
// @Failure      401 {object} structs.ErrorResponse "ID Transaksi resep obat tidak ditemukan"
// @Failure 500 {object} structs.ErrorResponse "Internal server error"
// @Router       /api/admin-farmasi/transaksi-obat/{id} [put]
func UpdateTransaksiObat(ctx *gin.Context) {
	IdTransaksiObat, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	var setTransaksiObat structs.UpdateTransaksiObatRequest

	if err := ctx.ShouldBindJSON(&setTransaksiObat); err != nil {
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
			"pesan":  "Akses update transaksi resep obat ditolak, login dulu sebagai admin farmasi.",
		})
		return
	}

	hasilUpdateData, err := repositories.UpdateTransaksiObat(DBSqlConn, IdTransaksiObat, setTransaksiObat, cekUsernameJWT)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "sukses",
		"pesan":  "Update data transaksi resep obat pasien berhasil",
		"data":   hasilUpdateData,
	})
}

// DeleteTransaksiObat godoc
//
// @Summary Delete Transaksi Obat
// @Description Menghapus transaksi obat
// @Tags 6. Transaksi Obat
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID Transaksi Obat"
// @Success 200 {object} structs.DeleteTransaksiObatResponse "Transaksi resep obat berhasil dihapus"
// @Failure 400 {object} structs.ErrorResponse "Gagal hapus transaksi resep obat"
// @Failure 401 {object} structs.ErrorResponse "ID Transaksi resep obat tidak ditemukan"
// @Failure 500 {object} structs.ErrorResponse "Internal server error"
// @Router /api/admin-farmasi/transaksi-obat/{id} [delete]
func DeleteTransaksiObat(ctx *gin.Context) {

	IdTransaksiObat, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	err = repositories.DeleteTransaksiObat(DBSqlConn, IdTransaksiObat)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "sukses",
		"pesan":  fmt.Sprintf("Transaksi resep obat ID %d berhasil dihapus.", IdTransaksiObat),
	})

}

// CancelTransaksiObat godoc
//
// @Summary Cancel Transaksi Obat
// @Description Membatalkan transaksi obat dengan mengembalikan stok obat jika status transaksi masih PENDING
// @Tags 6. Transaksi Obat
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID Transaksi Obat"
// @Success 200 {object} structs.CancelTransaksiObatResponse "Cancel data transaksi resep obat pasien berhasil"
// @Failure 400 {object} structs.ErrorResponse "Gagal hapus transaksi resep obat"
// @Failure 401 {object} structs.ErrorResponse "ID Transaksi resep obat tidak ditemukan"
// @Failure 500 {object} structs.ErrorResponse "Internal server error"
// @Router /api/admin-farmasi/transaksi-obat/{id}/cancel [patch]
func CancelTransaksiObat(ctx *gin.Context) {

	IdTransaksiObat, err := strconv.Atoi(ctx.Param("id"))

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
			"pesan":  "Akses cancel transaksi resep obat ditolak, login dulu sebagai admin farmasi.",
		})
		return
	}

	hasil, err := repositories.CancelTransaksiObat(DBSqlConn, IdTransaksiObat, cekUsernameJWT)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "sukses",
		"pesan":  "Cancel data transaksi resep obat pasien berhasil",
		"data":   hasil,
	})
}

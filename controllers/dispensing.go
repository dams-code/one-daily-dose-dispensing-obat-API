package controllers

import (
	"net/http"
	"one-daily-dose-dispensing-obat-api/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DispensingTransaksiObat godoc
//
// @Summary      Dispensing Obat
// @Description  Mengubah status transaksi obat dari PENDING menjadi DISPENSED. Transaksi harus memiliki minimal satu detail obat.
// @Tags         8. Dispensed
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "ID Transaksi Obat"
// @Success      200 {object} structs.DispensingObatResponse "Proses dispensing transaksi resep obat (pemberian obat) berhasil"
// @Failure      400 {object} structs.ErrorResponse "ID Transaksi obat tidak ditemukan"
// @Failure      401 {object} structs.ErrorResponse "Gagal dispense transaksi resep obat"
// @Failure      500 {object} structs.ErrorResponse "Internal Server Error"
// @Router       /api/apoteker/transaksi-obat/{id}/dispensing [put]
func DispensingTransaksiObat(ctx *gin.Context) {

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
			"pesan":  "Akses dispensing transaksi resep obat ditolak, login dulu sebagai apoteker.",
		})
		return
	}

	hasilDispensingObat, err := repositories.DispensingObat(DBSqlConn, IdTransaksiObat, cekUsernameJWT)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "sukses",
		"pesan":  "Proses dispensing transaksi resep obat (pemberian obat) berhasil",
		"data":   hasilDispensingObat,
	})

}

// GetPendingTransaksiObat godoc
//
// @Summary      Get Transaksi Obat Pending
// @Description  Menampilkan daftar transaksi obat dengan status PENDING yang belum didispensing.
// @Tags         8. Dispensed
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} structs.GetTransaksiObatResponse "Get daftar pending transaksi resep obat berhasil"
// @Failure      401 {object} structs.ErrorResponse "Data pending transaksi resep obat tidak ditemukan"
// @Failure      500 {object} structs.ErrorResponse "Internal Server Error"
// @Router       /api/apoteker/transaksi-obat/pending [get]
func GetPendingTransaksiObat(ctx *gin.Context) {

	cekUsernameJWT := ctx.MustGet("username_sekarang").(string)

	if cekUsernameJWT == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status": "error",
			"pesan":  "Akses cek pending transaksi resep obat ditolak, login dulu sebagai apoteker.",
		})
	}

	hasilGetPendingData, err := repositories.GetPendingTransaksiObat(DBSqlConn)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "sukses",
		"pesan":  "Get daftar pending transaksi resep obat berhasil",
		"data":   hasilGetPendingData,
	})

}

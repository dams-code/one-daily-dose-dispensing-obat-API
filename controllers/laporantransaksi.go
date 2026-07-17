package controllers

import (
	"net/http"
	"one-daily-dose-dispensing-obat-api/repositories"

	"github.com/gin-gonic/gin"
)

// GetLaporanTransaksiObat godoc
//
// @Summary      Get Laporan Transaksi Resep Obat
// @Description  Menampilkan laporan transaksi obat berdasarkan rentang tanggal.
// @Tags         10. Laporan
// @Produce      json
// @Security     BearerAuth
// @Param        tanggal_awal query string true "Tanggal awal (YYYY-MM-DD)" example(2026-07-01)
// @Param        tanggal_akhir query string true "Tanggal akhir (YYYY-MM-DD)" example(2026-07-31)
// @Success      200 {object} structs.LaporanTransaksiObat "Get daftar laporan transaksi resep obat pasien berhasil"
// @Failure      400 {object} structs.ErrorResponse "Transaksi obat tidak ditemukan pada rentang tanggal awal - akhir."
// @Failure      401 {object} structs.ErrorResponse "Transaksi obat tidak ada."
// @Failure      500 {object} structs.ErrorResponse "Internal Server Error"
// @Router       /api/kepala-farmasi/laporan/transaksi-obat [get]
func GetLaporanTransaksiObat(ctx *gin.Context) {

	tanggalAwal := ctx.Query("tanggal_awal")
	tanggalAkhir := ctx.Query("tanggal_akhir")

	cekUsernameJWT := ctx.MustGet("username_sekarang").(string)

	if cekUsernameJWT == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"pesan":  "Akses ke laporan transaksi resep obat ditolak, login dulu sebagai admin farmasi.",
		})
		return
	}

	if tanggalAwal == "" || tanggalAkhir == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"pesan":  "tanggal_awal dan tanggal_akhir wajib diisi",
		})
		return
	}

	hasil, err := repositories.GetLaporanTransaksiObat(DBSqlConn, tanggalAwal, tanggalAkhir)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "sukses",
		"pesan":  "Get daftar laporan transaksi resep obat pasien berhasil",
		"data":   hasil,
	})
}

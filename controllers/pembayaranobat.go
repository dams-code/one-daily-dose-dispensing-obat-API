package controllers

import (
	"fmt"
	"net/http"
	"one-daily-dose-dispensing-obat-api/repositories"
	"one-daily-dose-dispensing-obat-api/structs"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func PembayaranTransaksiObat(ctx *gin.Context) {
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
			"pesan":  "Akses pembayaran resep obat ditolak, login dulu sebagai admin farmasi",
		})
		return
	}

	var setPembayaranObat structs.PembayaranTransaksiObatRequest

	if err := ctx.ShouldBindJSON(&setPembayaranObat); err != nil {
		if cekErrorPembayaranObat, ok := err.(validator.ValidationErrors); ok {

			for _, fieldError := range cekErrorPembayaranObat {

				var pesan string
				switch fieldError.Field() {

				case "TipePembayaran":
					switch fieldError.Tag() {
					case "required":
						pesan = "Tipe Pembayaran wajib diisi"
					case "oneof":
						pesan = "Tipe Pembayaran harus TUNAI, QRIS"

					default:
						pesan = fmt.Sprintf("Tipe Pembayaran Tidak Valid %s", fieldError.Tag())
					}

				case "TotalPembayaran":
					switch fieldError.Tag() {
					case "required":
						pesan = "Total Pembayaran wajib diisi"
					case "gt":
						pesan = "Total Pembayaran harus lebih dari 0"

					}
				}

				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"status": "error",
					"pesan":  pesan,
				})
				return
			}
			return
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	hasilPembayranData, err := repositories.PembayaranTransaksiObat(DBSqlConn, IdTransaksiObat, setPembayaranObat, cekUsernameJWT)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "sukses",
		"pesan":  "Pembayaran resep obat Pasien berhasil",
		"data":   hasilPembayranData,
	})
}

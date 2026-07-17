package controllers

import (
	"fmt"
	"net/http"
	"one-daily-dose-dispensing-obat-api/repositories"
	"one-daily-dose-dispensing-obat-api/structs"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetObat godoc
//
// @Summary Get Semua Obat
// @Description Mengambil seluruh data obat
// @Tags 4. Obat
// @Produce json
// @Security BearerAuth
// @Success 200 {object} structs.GetObatResponse "Berhasil get data obat"
// @Failure 401 {object} structs.ErrorResponse "Gagal get data obat"
// @Failure 500 {object} structs.ErrorResponse "Internal server error"
// @Router /api/admin-farmasi/obat [get]
func GetObat(ctx *gin.Context) {

	IdObat, _ := strconv.Atoi(ctx.Query("id"))
	NamaObat := ctx.Query("nama")
	KodeObat := ctx.Query("kode_obat")

	hasilGetData, err := repositories.GetObat(DBSqlConn, IdObat, NamaObat, KodeObat)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "sukses",
		"pesan":  "Get data daftar obat berhasil",
		"data":   hasilGetData,
	})

}

// GetObatByID godoc
//
// @Summary Get Obat By ID
// @Description Mengambil detail obat berdasarkan ID
// @Tags 4. Obat
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID Obat"
// @Success 200 {object} structs.GetObatByIDResponse "Get data obat by ID berhasil"
// @Failure 400 {object} structs.ErrorResponse "ID Obat tidak valid"
// @Failure 404 {object} structs.ErrorResponse "Obat tidak ditemukan"
// @Failure 500 {object} structs.ErrorResponse "Internal server error"
// @Router /api/admin-farmasi/obat/{id} [get]
func GetObatByID(ctx *gin.Context) {

	IdObat, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	hasilGetData, err := repositories.GetObatByID(DBSqlConn, IdObat)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "sukses",
		"pesan":  "Get data obat by ID berhasil",
		"data":   hasilGetData,
	})

}

// GetObatByKategori godoc
//
//	@Summary Get Obat Berdasarkan Kategori
//	@Description Mengambil daftar obat berdasarkan ID kategori obat
//	@Tags 4. Obat
//	@Produce json
//	@Security BearerAuth
//	@Param id path int true "ID Kategori Obat"
//	@Success 200 {object} structs.GetObatResponse "Berhasil mengambil data obat by kategori"
//	@Failure 400 {object} structs.ErrorResponse "Gagal get obat by kategori"
//	@Failure 404 {object} structs.ErrorResponse "Kategori obat tidak ditemukan"
//	@Failure 500 {object} structs.ErrorResponse "Internal server error"
//	@Router /api/admin-farmasi/kategori-obat/{id}/obat [get]
func GetObatByKategori(ctx *gin.Context) {
	IdKategori, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	hasilGetData, err := repositories.GetObatByKategori(DBSqlConn, IdKategori)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": "sukses",
		"pesan":  "Get data obat by kategori berhasil",
		"data":   hasilGetData,
	})
}

// TambahObat godoc
//
// @Summary Tambah Obat
// @Description Menambahkan data obat baru
// @Tags 4. Obat
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body structs.TambahObatRequest true "Data Obat"
// @Success 201 {object} structs.TambahObatResponse "Tambah data obat berhasil"
// @Failure 400 {object} structs.ErrorResponse "Gagal tambah obat"
// @Failure 409 {object} structs.ErrorResponse "Kode obat sudah digunakan"
// @Failure 500 {object} structs.ErrorResponse "Internal server error"
// @Router /api/admin-farmasi/obat [post]
func TambahObat(ctx *gin.Context) {
	var setObat structs.TambahObatRequest

	if err := ctx.ShouldBindJSON(&setObat); err != nil {
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
			"pesan":  "Akses tambah obat ditolak, login dulu sebagai admin farmasi.",
		})
		return
	}

	hasilTambahData, err := repositories.TambahObat(DBSqlConn, setObat, cekUsernameJWT)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": "sukses",
		"pesan":  "Tambah data obat berhasil",
		"data":   hasilTambahData,
	})
}

// UpdateObat godoc
//
// @Summary Update Obat
// @Description Mengubah data obat
// @Tags 4. Obat
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID Obat"
// @Param request body structs.UpdateObatRequest true "Data Obat"
// @Success 200 {object} structs.UpdateObatResponse "Update data obat berhasil"
// @Failure 400 {object} structs.ErrorResponse "Gagal update obat"
// @Failure 404 {object} structs.ErrorResponse "Obat tidak ditemukan"
// @Failure 500 {object} structs.ErrorResponse "Internal server error"
// @Router /api/admin-farmasi/obat/{id} [put]
func UpdateObat(ctx *gin.Context) {

	IdObat, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	var setObat structs.UpdateObatRequest

	if err := ctx.ShouldBindJSON(&setObat); err != nil {
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
			"pesan":  "Akses update obat ditolak, login dulu sebagai admin farmasi.",
		})
		return
	}

	hasilUpdateData, err := repositories.UpdateObat(
		DBSqlConn,
		IdObat,
		setObat,
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
		"pesan":  "Update data obat berhasil",
		"data":   hasilUpdateData,
	})
}

// DeleteObat godoc
//
// @Summary Delete Obat
// @Description Menghapus data obat
// @Tags 4. Obat
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID Obat"
// @Success 200 {object} structs.DeleteObatResponse "Obat berhasil dihapus"
// @Failure 400 {object} structs.ErrorResponse "Gagal hapus obat"
// @Failure 404 {object} structs.ErrorResponse "Obat tidak ditemukan"
// @Failure 500 {object} structs.ErrorResponse "Internal server error"
// @Router /api/admin-farmasi/obat/{id} [delete]
func DeleteObat(ctx *gin.Context) {

	IdObat, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	err = repositories.DeleteObat(DBSqlConn, IdObat)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "sukses",
		"pesan":  fmt.Sprintf("Obat ID %d berhasil dihapus.", IdObat),
	})
}

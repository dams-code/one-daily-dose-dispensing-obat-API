package controllers

import (
	"fmt"
	"net/http"
	"one-daily-dose-dispensing-obat-api/repositories"
	"one-daily-dose-dispensing-obat-api/structs"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetKategoriObat godoc
//
// @Summary Get Semua Kategori Obat
// @Description Mengambil Semua data kategori obat
// @Tags 3. Kategori Obat
// @Produce json
// @Security BearerAuth
// @Success 200 {object} structs.GetKategoriObatResponse "Berhasil mengambil data kategori obat"
// @Failure 500 {object} structs.ErrorResponse "Internal server error"
// @Router /api/admin-farmasi/kategori-obat [get]
func GetKategoriObat(ctx *gin.Context) {

	IdKategoriObat, _ := strconv.Atoi(ctx.Query("id"))
	NamaKategoriObat := ctx.Query("nama")

	hasilGetData, err := repositories.GetKategoriObat(DBSqlConn, IdKategoriObat, NamaKategoriObat)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "sukses",
		"pesan":  "Get daftar kategori obat berhasil",
		"data":   hasilGetData,
	})
}

// GetKategoriObatByID godoc
//
// @Summary Get Kategori Obat By ID
// @Description Mengambil detail kategori obat berdasarkan ID
// @Tags 3. Kategori Obat
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID Kategori Obat"
// @Success 200 {object} structs.GetKategoriObatByIDResponse "Berhasil mengambil detail kategori obat"
// @Failure 400 {object} structs.ErrorResponse "ID Kategori Obat tidak valid"
// @Failure 404 {object} structs.ErrorResponse "Kategori obat tidak ditemukan"
// @Failure 500 {object} structs.ErrorResponse "Internal server error"
// @Router /api/admin-farmasi/kategori-obat/{id} [get]
func GetKategoriObatID(ctx *gin.Context) {
	IdKategoriObat, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	hasilGetData, err := repositories.GetKategoriObatID(DBSqlConn, IdKategoriObat)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "sukses",
		"pesan":  "Get kategori obat by ID berhasil",
		"data":   hasilGetData,
	})
}

// TambahKategoriObat godoc
//
// @Summary Tambah Kategori Obat
// @Description Menambahkan kategori obat baru
// @Tags 3. Kategori Obat
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body structs.TambahdanUpdateKategoriObatRequest true "Data Kategori Obat"
// @Success 201 {object} structs.TambahKategoriObatResponse "Kategori obat berhasil ditambahkan"
// @Failure 400 {object} structs.ErrorResponse "Gagal tambah kategori obat"
// @Failure 409 {object} structs.ErrorResponse "Kategori obat sudah ada"
// @Failure 500 {object} structs.ErrorResponse "Internal server error"
// @Router /api/admin-farmasi/kategori-obat [post]
func TambahKategoriObat(ctx *gin.Context) {

	var setKategoriObat structs.TambahdanUpdateKategoriObatRequest

	if err := ctx.ShouldBindJSON(&setKategoriObat); err != nil {
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
			"pesan":  "Akses tambah kategori obat ditolak, login dulu sebagai admin farmasi.",
		})
		return
	}

	hasilTambahData, err := repositories.TambahKategoriObat(DBSqlConn, setKategoriObat, cekUsernameJWT)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": "sukses",
		"pesan":  "Tambah data kategori obat berhasil",
		"data":   hasilTambahData,
	})

}

// UpdateKategoriObat godoc
//
// @Summary Update Kategori Obat
// @Description Mengubah data kategori obat
// @Tags 3. Kategori Obat
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID Kategori Obat"
// @Param request body structs.TambahdanUpdateKategoriObatRequest true "Data Kategori Obat"
// @Success 200 {object} structs.UpdateKategoriObatResponse "Kategori obat berhasil diUpdate"
// @Failure 400 {object} structs.ErrorResponse "Gagal Update kategori obat"
// @Failure 404 {object} structs.ErrorResponse "Kategori obat tidak ditemukan"
// @Failure 500 {object} structs.ErrorResponse "Internal server error"
// @Router /api/admin-farmasi/kategori-obat/{id} [put]
func UpdateKategoriObat(ctx *gin.Context) {

	IdKategoriObat, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	var setKategoriObat structs.TambahdanUpdateKategoriObatRequest

	if err := ctx.ShouldBindJSON(&setKategoriObat); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	cekUsernameJWT := ctx.MustGet("username_sekarang").(string)

	if cekUsernameJWT == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status": "error",
			"pesan":  "Akses update kategori obat ditolak, login dulu sebagai admin farmasi.",
		})
		return
	}

	hasilUpdateData, err := repositories.UpdateKategoriObat(DBSqlConn, IdKategoriObat, setKategoriObat, cekUsernameJWT)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "sukses",
		"pesan":  "Update data kategori obat berhasil",
		"data":   hasilUpdateData,
	})
}

// DeleteKategoriObat godoc
//
// @Summary Delete Kategori Obat
// @Description Menghapus kategori obat
// @Tags 3. Kategori Obat
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID Kategori Obat"
// @Success 200 {object} structs.DeleteKategoriObatResponse "Kategori obat berhasil dihapus"
// @Failure 400 {object} structs.ErrorResponse "ID Kategori obat tidak valid"
// @Failure 404 {object} structs.ErrorResponse "Kategori obat tidak ditemukan"
// @Failure 500 {object} structs.ErrorResponse "Internal server error"
// @Router /api/admin-farmasi/kategori-obat/{id} [delete]
func DeleteKategoriObat(ctx *gin.Context) {
	IdKategoriObat, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	err = repositories.DeleteKategoriObat(DBSqlConn, IdKategoriObat)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "sukses",
		"pesan":  fmt.Sprintf("Kategori Obat ID %d berhasil dihapus.", IdKategoriObat),
	})

}

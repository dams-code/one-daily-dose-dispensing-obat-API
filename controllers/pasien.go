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

// GetPasien godoc
//
// @Summary Get Semua Pasien
// @Description Mengambil Semua data pasien
// @Tags 5. Pasien
// @Produce json
// @Security BearerAuth
// @Success 200 {object} structs.GetPasienResponse "Berhasil Get data pasien"
// @Failure 500 {object} structs.ErrorResponse "Internal server error"
// @Router /api/admin-farmasi/pasien [get]
func GetPasien(ctx *gin.Context) {

	IdPasien, _ := strconv.Atoi(ctx.Query("id"))

	RmPasien := ctx.Query("rm")
	NamaPasien := ctx.Query("nama")

	hasilGetData, err := repositories.GetPasien(DBSqlConn, IdPasien, RmPasien, NamaPasien)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "sukses",
		"pesan":  "Get data daftar pasien berhasil",
		"data":   hasilGetData,
	})
}

// GetPasienByID godoc
//
// @Summary Get Pasien By ID
// @Description Mengambil detail pasien berdasarkan ID
// @Tags 5. Pasien
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID Pasien"
// @Success 200 {object} structs.GetPasienByIDResponse "Berhasil mengambil detail pasien"
// @Failure 400 {object} structs.ErrorResponse "ID Pasien tidak valid"
// @Failure 404 {object} structs.ErrorResponse "Pasien tidak ditemukan"
// @Failure 500 {object} structs.ErrorResponse "Internal server error"
// @Router /api/admin-farmasi/pasien/{id} [get]
func GetPasienByID(ctx *gin.Context) {
	IdPasien, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	hasilGetData, err := repositories.GetPasienByID(DBSqlConn, IdPasien)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "sukses",
		"pesan":  "Get data pasien by ID berhasil",
		"data":   hasilGetData,
	})
}

// TambahPasien godoc
//
// @Summary Tambah Pasien
// @Description Menambahkan data pasien baru
// @Tags 5. Pasien
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body structs.TambahPasienRequest true "Data Pasien"
// @Success 201 {object} structs.TambahPasienResponse "Pasien berhasil ditambahkan"
// @Failure 400 {object} structs.ErrorResponse "Gagal tambah pasien"
// @Failure 409 {object} structs.ErrorResponse "No RM Pasien sudah digunakan"
// @Failure 500 {object} structs.ErrorResponse "Internal server error"
// @Router /api/admin-farmasi/pasien [post]
func TambahPasien(ctx *gin.Context) {

	var setPasien structs.TambahPasienRequest

	if err := ctx.ShouldBindJSON(&setPasien); err != nil {
		if cekErrorTambahPasien, ok := err.(validator.ValidationErrors); ok {

			for _, fieldError := range cekErrorTambahPasien {

				var pesan string
				switch fieldError.Field() {

				case "Rm":
					switch fieldError.Tag() {
					case "required":
						pesan = "Rm wajib diisi"
					case "len":
						pesan = "Rm wajib 8 digit"
					case "numeric":
						pesan = "Rm hanya boleh berisi angka"
					default:
						pesan = fmt.Sprintf("Rm Tidak Valid %s", fieldError.Tag())
					}

				case "Nama":
					switch fieldError.Tag() {
					case "required":
						pesan = "Nama wajib diisi"
					}

				case "JenisKelamin":
					switch fieldError.Tag() {
					case "required":
						pesan = "Jenis Kelamin wajib diisi"
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

	cekUsernameJWT := ctx.MustGet("username_sekarang").(string)

	if cekUsernameJWT == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status": "error",
			"pesan":  "Akses tambah pasien ditolak, login dulu sebagai admin farmasi.",
		})
		return
	}

	hasilTambahData, err := repositories.TambahPasien(DBSqlConn, setPasien, cekUsernameJWT)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": "sukses",
		"pesan":  "Tambah data pasien berhasil",
		"data":   hasilTambahData,
	})
}

// UpdatePasien godoc
//
// @Summary Update Pasien
// @Description Mengubah data pasien
// @Tags 5. Pasien
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID Pasien"
// @Param request body structs.UpdatePasienRequest true "Data Pasien"
// @Success 200 {object} structs.UpdatePasienResponse "Update Pasien berhasil"
// @Failure 400 {object} structs.ErrorResponse "Gagal update pasien"
// @Failure 404 {object} structs.ErrorResponse "Pasien tidak ditemukan"
// @Failure 500 {object} structs.ErrorResponse "Internal server error"
// @Router /api/admin-farmasi/pasien/{id} [put]
func UpdatePasien(ctx *gin.Context) {
	IdPasien, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	var setPasien structs.UpdatePasienRequest

	if err := ctx.ShouldBindJSON(&setPasien); err != nil {
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
			"pesan":  "Akses update pasien ditolak, login dulu sebagai admin farmasi.",
		})
		return
	}

	hasilUpdateData, err := repositories.UpdatePasien(DBSqlConn, IdPasien, setPasien, cekUsernameJWT)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "sukses",
		"pesan":  "Update data pasien berhasil",
		"data":   hasilUpdateData,
	})
}

// DeletePasien godoc
//
// @Summary Delete Pasien
// @Description Menghapus data pasien
// @Tags 5. Pasien
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID Pasien"
// @Success 200 {object} structs.DeletePasienResponse "Pasien berhasil dihapus"
// @Failure 400 {object} structs.ErrorResponse "ID tidak valid"
// @Failure 404 {object} structs.ErrorResponse "Pasien tidak ditemukan"
// @Failure 500 {object} structs.ErrorResponse "Internal server error"
// @Router /api/admin-farmasi/pasien/{id} [delete]
func DeletePasien(ctx *gin.Context) {
	IdPasien, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	err = repositories.DeletePasien(DBSqlConn, IdPasien)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "sukses",
		"pesan":  fmt.Sprintf("Pasien ID %d Berhasil Terhapus", IdPasien),
	})

}

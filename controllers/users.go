package controllers

import (
	"fmt"
	"net/http"
	"one-daily-dose-dispensing-obat-api/middleware"
	"one-daily-dose-dispensing-obat-api/repositories"
	"one-daily-dose-dispensing-obat-api/structs"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

// GetUser godoc
// @Summary Get Semua User
// @Description Mengambil seluruh data user
// @Tags 2. User
// @Produce json
// @Security BearerAuth
// @Success 200 {object} structs.GetUserResponse "Get User"
// @Failure 401 {object} structs.ErrorResponse "Gagal get User"
// @Failure 500 {object} structs.ErrorResponse "Internal server error"
// @Router /kepala-farmasi/users [get]
func GetUser(ctx *gin.Context) {
	IdUserObat, _ := strconv.Atoi(ctx.Query("id"))
	NamaUser := ctx.Query("nama")

	hasilGetUser, err := repositories.GetUser(DBSqlConn, IdUserObat, NamaUser)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "sukses",
		"pesan":  "Get user berhasil",
		"data":   hasilGetUser,
	})

}

// GetUserByID godoc
// @Summary Get User By ID
// @Description Mengambil detail user berdasarkan ID
// @Tags 2. User
// @Produce json
// @Security BearerAuth
// @Param id path int true "User ID"
// @Success 200 {object} structs.GetUserByIDResponse "Get User by ID"
// @Failure 400 {object} structs.ErrorResponse "ID User tidak valid"
// @Failure 404 {object} structs.ErrorResponse "Gagal get User by ID"
// @Failure 500 {object} structs.ErrorResponse "Internal server error"
// @Router /kepala-farmasi/users/{id} [get]
func GetUserByID(ctx *gin.Context) {

	IdUser, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
	}

	hasilGetData, err := repositories.GetUserByID(DBSqlConn, IdUser)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "sukses",
		"pesan":  "Get user by ID berhasil",
		"data":   hasilGetData,
	})

}

// UpdateUser godoc
// @Summary Update User
// @Description Mengubah data user
// @Tags 2. User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "User ID"
// @Param request body structs.UpdateUserRequest true "Data User"
// @Success 200 {object} structs.UpdateUserResponse "Update User"
// @Failure 400 {object} structs.ErrorResponse "Gagal update User"
// @Failure 404 {object} structs.ErrorResponse "User tidak ditemukan"
// @Failure 500 {object} structs.ErrorResponse "Internal server error"
// @Router /kepala-farmasi/users/{id} [put]
func UpdateUser(ctx *gin.Context) {
	var setUserObat structs.UpdateUserRequest

	if err := ctx.ShouldBindJSON(&setUserObat); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	IdUserObat, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	cekUsernameJWT := ctx.MustGet("username_sekarang").(string)

	if cekUsernameJWT == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status": "error",
			"pesan":  "Akses update user ditolak, login dulu sebagai kepala farmasi.",
		})
		return
	}

	if setUserObat.Password != nil && *setUserObat.Password != "" {
		hashPasswordUser, err := bcrypt.GenerateFromPassword([]byte(*setUserObat.Password), bcrypt.DefaultCost)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status": "error",
				"pesan":  err.Error(),
			})
			return
		}

		*setUserObat.Password = string(hashPasswordUser)
	}

	hasilUpdateData, err := repositories.UpdateUser(DBSqlConn, IdUserObat, setUserObat, cekUsernameJWT)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": "sukses",
		"pesan":  "Update user berhasil",
		"data":   hasilUpdateData,
	})

}

// DeleteUser godoc
// @Summary Delete User
// @Description Menghapus user
// @Tags 2. User
// @Produce json
// @Security BearerAuth
// @Param id path int true "User ID"
// @Success 200 {object} structs.DeleteUserResponse "delete user"
// @Failure 400 {object} structs.ErrorResponse "ID User tidak valid"
// @Failure 404 {object} structs.ErrorResponse "User tidak ditemukan"
// @Failure 500 {object} structs.ErrorResponse "Internal server error"
// @Router /kepala-farmasi/users/{id} [delete]
func DeleteUser(ctx *gin.Context) {
	IdUserObat, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	cekUsernameJWT := ctx.MustGet("username_sekarang").(string)

	if cekUsernameJWT == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status": "error",
			"pesan":  "Akses delete user ditolak, login dulu sebagai kepala farmasi.",
		})
		return
	}

	err = repositories.DeleteUser(DBSqlConn, IdUserObat, cekUsernameJWT)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "sukses",
		"pesan":  fmt.Sprintf("User ID %d berhasil dinonaktifkan", IdUserObat),
	})

}

// AktivasiUser godoc
// @Summary Aktivasi User
// @Description Mengaktifkan atau menonaktifkan user
// @Tags 2. User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "User ID"
// @Param request body structs.AktivasiUserResponse true "Status User"
// @Success 200 {object} structs.AktivasiUserResponse "aktivasi user"
// @Failure 400 {object} structs.ErrorResponse "gagal aktivasi user"
// @Failure 404 {object} structs.ErrorResponse "ID user tidak ditemukan"
// @Failure 500 {object} structs.ErrorResponse "Internal server error"
// @Router /kepala-farmasi/users/{id}/aktivasi [patch]
func AktivasiUser(ctx *gin.Context) {

	IdUserObat, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	cekUsernameJWT := ctx.MustGet("username_sekarang").(string)

	if cekUsernameJWT == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status": "error",
			"pesan":  "Akses aktivasi user ditolak, login dulu sebagai kepala farmasi.",
		})
		return
	}

	err = repositories.AktivasiUser(DBSqlConn, IdUserObat, cekUsernameJWT)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "sukses",
		"pesan":  "Aktivasi user berhasil",
		"data ":  fmt.Sprintf("User ID %d aktif kembali", IdUserObat),
	})
}

// TambahUser godoc
// @Summary Tambah User
// @Description Menambahkan user baru
// @Tags 2. User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body structs.TambahUserRequest true "Data User"
// @Success 201 {object} structs.TambahUserResponse "Tambah user berhasil"
// @Failure 400 {object} structs.ErrorResponse "Gagal tambah user"
// @Failure 500 {object} structs.ErrorResponse "Internal server error"
// @Router /kepala-farmasi/users [post]
func TambahUser(ctx *gin.Context) {

	var setUserObat structs.TambahUserRequest

	if err := ctx.ShouldBindJSON(&setUserObat); err != nil {
		if cekErrorTambahUser, ok := err.(validator.ValidationErrors); ok {

			for _, fieldError := range cekErrorTambahUser {

				var pesan string
				switch fieldError.Field() {

				case "Username":
					switch fieldError.Tag() {
					case "required":
						pesan = "Username wajib diisi"
					}

				case "Password":
					switch fieldError.Tag() {
					case "required":
						pesan = "Password wajib diisi"
					case "min":
						pesan = "Password minimal 6 karakter"
					}

				case "Role":
					switch fieldError.Tag() {
					case "required":
						pesan = "Role wajib diisi"
					case "oneof":
						pesan = "Role harus ADMINFARMASI atau APOTEKER"
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
			"pesan":  "Format request tidak valid",
		})
		return
	}

	cekUsernameJWT := ctx.MustGet("username_sekarang").(string)

	if cekUsernameJWT == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status": "error",
			"pesan":  "Akses tambah user ditolak, login dulu sebagai kepala farmasi.",
		})
		return
	}

	hashPasswordUser, err := bcrypt.GenerateFromPassword(
		[]byte(setUserObat.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
	}

	setUserObat.Password = string(hashPasswordUser)

	hasilTambahData, err := repositories.TambahUser(DBSqlConn, setUserObat, cekUsernameJWT)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": "sukses",
		"pesan":  "User berhasil ditambahkan",
		"data":   hasilTambahData,
	})

}

// RegisterUser godoc
// @Summary Register User
// @Description Menambahkan user baru
// @Tags 1. Authentication
// @Accept json
// @Produce json
// @Param request body structs.TambahUserRequest true "Registrasi Data User"
// @Success 201 {object} structs.RegisterUserResponse "Registrasi user berhasil"
// @Failure 400 {object} structs.ErrorResponse "Request Register tidak valid"
// @Failure 500 {object} structs.ErrorResponse "Internal server error"
// @Router /auth/register [post]
func RegisterUser(ctx *gin.Context) {

	var setUserObat structs.RegisterUserRequest

	if err := ctx.ShouldBindJSON(&setUserObat); err != nil {
		if cekErrorRegistrasiUser, ok := err.(validator.ValidationErrors); ok {

			for _, fieldError := range cekErrorRegistrasiUser {

				var pesan string
				switch fieldError.Field() {

				case "Username":
					switch fieldError.Tag() {
					case "required":
						pesan = "Username wajib diisi"
					}

				case "Password":
					switch fieldError.Tag() {
					case "required":
						pesan = "Password wajib diisi"
					case "min":
						pesan = "Password minimal 6 karakter"
					}

				case "Role":
					switch fieldError.Tag() {
					case "required":
						pesan = "Role wajib diisi"
					case "oneof":
						pesan = "Role harus KEPALAFARMASI atau ADMINFARMASI atau APOTEKER"
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

	if setUserObat.Password != "" && len(setUserObat.Password) < 6 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"pesan":  "Password minimal 6 karakter",
		})
		return
	}

	_, err := repositories.GetUsername(DBSqlConn, setUserObat.Username)

	if err == nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"pesan":  "Username sudah terdaftar, bisa isi dengan username lain.",
		})
		return
	}

	hashPasswordUser, err := bcrypt.GenerateFromPassword(
		[]byte(setUserObat.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	setUserObat.Password = string(hashPasswordUser)

	if setUserObat.Role == "" {
		setUserObat.Role = "ADMIN"
	}

	var CreatedBy = "Sistem"

	hasilTambahData, err := repositories.RegisterUser(DBSqlConn, setUserObat, CreatedBy)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": "sukses",
		"pesan":  "registrasi user berhasil",
		"data":   hasilTambahData,
	})

}

// Login godoc
// @Summary Login User
// @Description Login menggunakan username dan password untuk mendapatkan JWT Token
// @Tags 1. Authentication
// @Accept json
// @Produce json
// @Param request body structs.LoginRequest true "Data Login"
// @Success 200 {object} structs.LoginResponse "Login berhasil"
// @Failure 400 {object} structs.ErrorResponse "Request login tidak valid"
// @Failure 401 {object} structs.ErrorResponse "Username atau password salah"
// @Failure 500 {object} structs.ErrorResponse "Internal server error"
// @Router /auth/login [post]
func LoginUser(ctx *gin.Context) {
	var setUserObat structs.LoginRequest

	if err := ctx.ShouldBindJSON(&setUserObat); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
	}

	UserFromDB, err := repositories.GetUsername(DBSqlConn, setUserObat.Username)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	token, err := middleware.GenerateJWT(UserFromDB.ID, UserFromDB.Username, UserFromDB.Role)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"pesan":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "sukses",
		"pesan":  "Hi, " + UserFromDB.Username + " Role: " + UserFromDB.Role,
		"token":  token,
	})

}

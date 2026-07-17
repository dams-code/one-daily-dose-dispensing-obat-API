package routers

import (
	"one-daily-dose-dispensing-obat-api/controllers"
	"one-daily-dose-dispensing-obat-api/middleware"

	"github.com/gin-gonic/gin"

	_ "one-daily-dose-dispensing-obat-api/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func StartServer() *gin.Engine {

	setRouter := gin.Default()

	setRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	setRouter.POST("/api/auth/register", controllers.RegisterUser)
	setRouter.POST("/api/auth/login", controllers.LoginUser)

	kepalaFarmasi := setRouter.Group("/api/kepala-farmasi")
	kepalaFarmasi.Use(middleware.MiddlewareJWT(), middleware.RoleUserObat("KEPALAFARMASI"))
	{
		kepalaFarmasi.GET("/users", controllers.GetUser)
		kepalaFarmasi.GET("/users/:id", controllers.GetUserByID)
		kepalaFarmasi.POST("/users", controllers.TambahUser)
		kepalaFarmasi.PUT("/users/:id", controllers.UpdateUser)
		kepalaFarmasi.DELETE("/users/:id", controllers.DeleteUser)
		kepalaFarmasi.PATCH("/users/:id/aktivasi", controllers.AktivasiUser)

		kepalaFarmasi.GET("/transaksi-obat/pending", controllers.GetPendingTransaksiObat)
		kepalaFarmasi.GET("/laporan/transaksi-obat", controllers.GetLaporanTransaksiObat)
	}

	adminFarmasi := setRouter.Group("/api/admin-farmasi")
	adminFarmasi.Use(middleware.MiddlewareJWT(), middleware.RoleUserObat("ADMINFARMASI"))
	{
		adminFarmasi.GET("/pasien", controllers.GetPasien)
		adminFarmasi.GET("/pasien/:id", controllers.GetPasienByID)
		adminFarmasi.POST("/pasien", controllers.TambahPasien)
		adminFarmasi.PUT("/pasien/:id", controllers.UpdatePasien)
		adminFarmasi.DELETE("/pasien/:id", controllers.DeletePasien)

		adminFarmasi.GET("/kategori-obat", controllers.GetKategoriObat)
		adminFarmasi.GET("/kategori-obat/:id", controllers.GetKategoriObatID)
		adminFarmasi.POST("/kategori-obat", controllers.TambahKategoriObat)
		adminFarmasi.PUT("/kategori-obat/:id", controllers.UpdateKategoriObat)
		adminFarmasi.DELETE("/kategori-obat/:id", controllers.DeleteKategoriObat)

		adminFarmasi.GET("/kategori-obat/:id/obat", controllers.GetObatByKategori)
		adminFarmasi.GET("/obat", controllers.GetObat)
		adminFarmasi.GET("/obat/:id", controllers.GetObatByID)
		adminFarmasi.POST("/obat", controllers.TambahObat)
		adminFarmasi.PUT("/obat/:id", controllers.UpdateObat)
		adminFarmasi.DELETE("/obat/:id", controllers.DeleteObat)

		adminFarmasi.GET("/transaksi-obat", controllers.GetTransaksiObat)
		adminFarmasi.GET("/transaksi-obat/:id", controllers.GetTransaksiObatByID)
		adminFarmasi.POST("/transaksi-obat", controllers.TambahTransaksiObat)
		adminFarmasi.PUT("/transaksi-obat/:id", controllers.UpdateTransaksiObat)
		adminFarmasi.DELETE("/transaksi-obat/:id", controllers.DeleteTransaksiObat)
		adminFarmasi.POST("/transaksi-obat/:id/pembayaran", controllers.PembayaranTransaksiObat)

		adminFarmasi.GET("/transaksi-obat/detail", controllers.GetDetailTransaksiObat)
		adminFarmasi.GET("/transaksi-obat/:id/detail", controllers.GetDetailTransaksiObatByTransaksiID)
		adminFarmasi.GET("/transaksi-obat/detail/:id", controllers.GetDetailTransaksiObatByID)
		adminFarmasi.POST("/transaksi-obat/:id/detail", controllers.TambahDetailTransaksiObat)
		adminFarmasi.PUT("/transaksi-obat/detail/:id", controllers.UpdateDetailTransaksiObat)
		adminFarmasi.DELETE("/transaksi-obat/detail/:id", controllers.DeleteDetailTransaksiObat)
		adminFarmasi.PATCH("/transaksi-obat/:id/cancel", controllers.CancelTransaksiObat)
	}

	apoteker := setRouter.Group("/api/apoteker")
	apoteker.Use(middleware.MiddlewareJWT(), middleware.RoleUserObat("APOTEKER"))
	{
		apoteker.GET("/transaksi-obat/pending", controllers.GetPendingTransaksiObat)
		apoteker.GET("/transaksi-obat/:id/detail", controllers.GetDetailTransaksiObatByTransaksiID)
		apoteker.GET("/transaksi-obat/detail/:id", controllers.GetDetailTransaksiObatByID)
		apoteker.PUT("/transaksi-obat/:id/dispense", controllers.DispensingTransaksiObat)

	}

	return setRouter
}

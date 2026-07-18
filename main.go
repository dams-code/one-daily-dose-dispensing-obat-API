// @title           Final Project REST API Golang
// @version         1.0
// @description     Dokumentasi REST API Alur Obat Farmasi menggunakan Golang (Framework : Gin dan Database : PostgreSQL).
// @termsOfService  http://swagger.io/terms/

// @contact.name    Damar Djati Wahyu Kemala
// @contact.url     https://github.com/dams-code

// @license.name    MIT
// @license.url     https://opensource.org/licenses/MIT

// @host            localhost:8080
// @BasePath        /api

// @tag.name 1. Authentication
// @tag.description Endpoint autentikasi

// @tag.name 2. User
// @tag.description Manajemen user

// @tag.name 3. Kategori Obat
// @tag.description Manajemen kategori obat

// @tag.name 4. Obat
// @tag.description Manajemen obat

// @tag.name 5. Pasien
// @tag.description Manajemen pasien

// @tag.name 6. Transaksi Obat
// @tag.description Manajemen transaksi obat

// @tag.name 7. Detail Transaksi Obat
// @tag.description Manajemen detail transaksi obat

// @tag.name 8. Dispensed
// @tag.description Dispensed (obat diberikan ke pasien / pembeli di apotik )

// @tag.name 9. Laporan
// @tag.description Laporan farmasi

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
package main

import (
	"log"
	"one-daily-dose-dispensing-obat-api/controllers"
	router "one-daily-dose-dispensing-obat-api/routers"
	"os"

	_ "one-daily-dose-dispensing-obat-api/docs"
)

func main() {

	sqlCon, err := controllers.KoneksiDB()

	if err != nil {
		log.Fatal("Gagal tersambung ke postgres ", err)
	}

	defer sqlCon.Close()

	controllers.DBSqlConn = sqlCon

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	err = router.StartServer().Run(":" + PORT)

	if err != nil {
		log.Fatal(err)
	}

}

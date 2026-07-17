package repositories

import (
	"database/sql"
	"fmt"
	"one-daily-dose-dispensing-obat-api/structs"
)

func PembayaranTransaksiObat(db *sql.DB, IdTransaksiObat int, setPembayaranObat structs.PembayaranTransaksiObatRequest, ModifiedBy string) (HasilPembayaranObat structs.TransaksiObat, err error) {

	execPembayaranObat, err := db.Begin()
	if err != nil {
		return structs.TransaksiObat{}, err
	}

	defer func() {
		if err != nil {
			_ = execPembayaranObat.Rollback()
		}
	}()

	queryCekPembayaran := `SELECT status, grand_total, total_pembayaran FROM transaksi_obat WHERE id = $1`

	var (
		status            string
		grandTotal        float64
		totalPembayaranDB sql.NullFloat64
	)

	err = execPembayaranObat.QueryRow(queryCekPembayaran, IdTransaksiObat).Scan(&status, &grandTotal, &totalPembayaranDB)

	if err == sql.ErrNoRows {
		return structs.TransaksiObat{}, fmt.Errorf(
			"Transaksi Obat ID %d tidak ditemukan",
			IdTransaksiObat,
		)
	}

	if err != nil {
		return structs.TransaksiObat{}, err
	}

	if status != "DISPENSED" {
		return structs.TransaksiObat{}, fmt.Errorf("Transaksi obat belum dapat dibayar karena status %s", status)
	}

	if totalPembayaranDB.Valid {
		return structs.TransaksiObat{}, fmt.Errorf("Transaksi obat sudah pernah dibayar")
	}

	if setPembayaranObat.TotalPembayaran < grandTotal {
		return structs.TransaksiObat{}, fmt.Errorf("Total pembayaran kurang dari grand total obat")
	}

	kembalian := setPembayaranObat.TotalPembayaran - grandTotal

	queryUpdate := `
		UPDATE transaksi_obat
		SET
			tipe_pembayaran = $1,
			total_pembayaran = $2,
			kembalian = $3,
			dibayar_at = CURRENT_TIMESTAMP,
			modified_by = $4
		WHERE id = $5
		RETURNING
			id, no_resep, pasien_id, status, tipe_pembayaran, grand_total, total_pembayaran, kembalian, dibayar_at, created_at, created_by, modified_at, modified_by
	`

	err = execPembayaranObat.QueryRow(
		queryUpdate, setPembayaranObat.TipePembayaran, setPembayaranObat.TotalPembayaran, kembalian, ModifiedBy, IdTransaksiObat).Scan(
		&HasilPembayaranObat.ID, &HasilPembayaranObat.NoResep, &HasilPembayaranObat.PasienID, &HasilPembayaranObat.Status,
		&HasilPembayaranObat.TipePembayaran, &HasilPembayaranObat.GrandTotal, &HasilPembayaranObat.TotalPembayaran, &HasilPembayaranObat.Kembalian,
		&HasilPembayaranObat.DibayarAt, &HasilPembayaranObat.CreatedAt, &HasilPembayaranObat.CreatedBy, &HasilPembayaranObat.ModifiedAt, &HasilPembayaranObat.ModifiedBy,
	)

	if err != nil {
		return structs.TransaksiObat{}, err
	}

	err = execPembayaranObat.QueryRow(`SELECT nama FROM pasien WHERE id = $1`, HasilPembayaranObat.PasienID).Scan(&HasilPembayaranObat.NamaPasien)

	if err != nil {
		return structs.TransaksiObat{}, err
	}

	err = execPembayaranObat.Commit()

	if err != nil {
		return structs.TransaksiObat{}, err
	}

	return HasilPembayaranObat, nil
}

package repositories

import (
	"database/sql"
	"fmt"
	"one-daily-dose-dispensing-obat-api/structs"
)

func DispensingObat(db *sql.DB, IdTransaksiObat int, ModifiedBy string) (HasilDispensingObat structs.TransaksiObat, err error) {

	execDispensingObat, err := db.Begin()

	if err != nil {
		return structs.TransaksiObat{}, err
	}

	defer func() {
		if err != nil {
			_ = execDispensingObat.Rollback()
		}
	}()

	queryStatusTransaksiObat := `SELECT no_resep, status FROM transaksi_obat WHERE id = $1`

	var (
		cekNoResep          string
		statusTransaksiObat string
	)
	err = execDispensingObat.QueryRow(queryStatusTransaksiObat, IdTransaksiObat).Scan(&cekNoResep, &statusTransaksiObat)

	if err == sql.ErrNoRows {
		return structs.TransaksiObat{}, fmt.Errorf(
			"Transaksi Obat ID %d tidak ditemukan",
			IdTransaksiObat,
		)
	}

	if err != nil {
		return structs.TransaksiObat{}, err
	}

	if statusTransaksiObat != "PENDING" {
		return structs.TransaksiObat{}, fmt.Errorf("transaksi tidak dapat didispensed karena status saat ini adalah %s", statusTransaksiObat)
	}

	queryCekDetailTransaksiObat := `SELECT COUNT(*) FROM detail_transaksi_obat WHERE transaksi_obat_id = $1`

	var jumlahDetailObat int

	err = execDispensingObat.QueryRow(queryCekDetailTransaksiObat, IdTransaksiObat).Scan(
		&jumlahDetailObat,
	)

	if err != nil {
		return structs.TransaksiObat{}, err
	}

	if jumlahDetailObat == 0 {
		return structs.TransaksiObat{}, fmt.Errorf("Transaksi Obat belum memiliki detail obat")
	}

	queryGrandTotal := `SELECT COALESCE(SUM(dto.jumlah * o.harga),0) FROM detail_transaksi_obat dto JOIN obat o ON dto.obat_id = o.id WHERE dto.transaksi_obat_id = $1`

	var grandTotal float64

	err = execDispensingObat.QueryRow(queryGrandTotal, IdTransaksiObat).Scan(&grandTotal)

	if err != nil {
		return structs.TransaksiObat{}, err
	}

	queryUbahStatusTransaksiObat := `
		UPDATE transaksi_obat
		SET status = 'DISPENSED',
			grand_total = $1,
			modified_by = $2
		WHERE id = $3
		RETURNING id
	`

	var updatedID int

	err = execDispensingObat.QueryRow(queryUbahStatusTransaksiObat, grandTotal, ModifiedBy, IdTransaksiObat).Scan(&updatedID)

	if err == sql.ErrNoRows {
		return structs.TransaksiObat{}, fmt.Errorf("Status Dispensing pada Transaksi Obat tidak berhasil di-update")
	}

	if err != nil {
		return structs.TransaksiObat{}, err
	}

	queryHasilDispensingObat := `
		SELECT 	t.id, t.no_resep, t.pasien_id, p.nama,
				t.status, t.grand_total, t.tipe_pembayaran, t.total_pembayaran,
				t.kembalian, t.created_at, t.created_by, t.modified_at, t.modified_by
		FROM transaksi_obat t
		JOIN pasien p ON p.id = t.pasien_id
		WHERE t.id = $1
	`

	err = execDispensingObat.QueryRow(queryHasilDispensingObat, IdTransaksiObat).Scan(
		&HasilDispensingObat.ID, &HasilDispensingObat.NoResep, &HasilDispensingObat.PasienID,
		&HasilDispensingObat.NamaPasien, &HasilDispensingObat.Status, &HasilDispensingObat.GrandTotal,
		&HasilDispensingObat.TipePembayaran, &HasilDispensingObat.TotalPembayaran,
		&HasilDispensingObat.Kembalian, &HasilDispensingObat.CreatedAt,
		&HasilDispensingObat.CreatedBy, &HasilDispensingObat.ModifiedAt, &HasilDispensingObat.ModifiedBy,
	)

	if err != nil {
		return structs.TransaksiObat{}, err
	}

	err = execDispensingObat.Commit()

	if err != nil {
		return structs.TransaksiObat{}, err
	}

	return HasilDispensingObat, nil

}

func GetPendingTransaksiObat(db *sql.DB) (HasilGetPendingTransaksiObat []structs.TransaksiObat, err error) {
	query := `
		SELECT  t.id, t.no_resep, t.pasien_id, p.nama,
				t.status, t.created_at, t.created_by,
				t.modified_at, t.modified_by
		FROM transaksi_obat t
		INNER JOIN pasien p ON p.id = t.pasien_id
		WHERE t.status = 'PENDING'
		ORDER BY t.created_at ASC
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {

		var setTransaksiObat structs.TransaksiObat

		err = rows.Scan(
			&setTransaksiObat.ID, &setTransaksiObat.NoResep, &setTransaksiObat.PasienID,
			&setTransaksiObat.NamaPasien, &setTransaksiObat.Status, &setTransaksiObat.CreatedAt,
			&setTransaksiObat.CreatedBy, &setTransaksiObat.ModifiedAt, &setTransaksiObat.ModifiedBy,
		)

		if err != nil {
			return nil, err
		}

		HasilGetPendingTransaksiObat = append(HasilGetPendingTransaksiObat, setTransaksiObat)
	}

	if HasilGetPendingTransaksiObat == nil {
		HasilGetPendingTransaksiObat = []structs.TransaksiObat{}
	}

	return HasilGetPendingTransaksiObat, nil
}

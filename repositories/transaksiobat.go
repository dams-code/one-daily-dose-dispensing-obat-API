package repositories

import (
	"database/sql"
	"fmt"
	"one-daily-dose-dispensing-obat-api/structs"
)

func GetTransaksiObat(db *sql.DB, IdTransaksiObat int, NoResep string) (HasilGetTransaksiObat []structs.TransaksiObat, err error) {

	querySelect := `
		SELECT 	t.id, t.no_resep, t.pasien_id, p.nama, t.status, 
				t.tipe_pembayaran, t.grand_total, t.total_pembayaran, t.kembalian, t.dibayar_at,
				t.created_at, t.created_by, t.modified_at, t.modified_by
		FROM transaksi_obat t
		LEFT JOIN pasien p ON t.pasien_id = p.id
		WHERE 1=1
	`

	var paramTransaksi []interface{}

	jumlahParam := 1

	if IdTransaksiObat > 0 {
		querySelect += fmt.Sprintf(" AND t.id = $%d", jumlahParam)
		paramTransaksi = append(paramTransaksi, IdTransaksiObat)
		jumlahParam++
	}

	if NoResep != "" {
		querySelect += fmt.Sprintf(" AND t.no_resep ILIKE $%d", jumlahParam)
		paramTransaksi = append(paramTransaksi, "%"+NoResep+"%")
		jumlahParam++
	}

	rows, err := db.Query(querySelect, paramTransaksi...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		var setTransaksiObat structs.TransaksiObat

		err = rows.Scan(&setTransaksiObat.ID, &setTransaksiObat.NoResep, &setTransaksiObat.PasienID, &setTransaksiObat.NamaPasien, &setTransaksiObat.Status,
			&setTransaksiObat.TipePembayaran, &setTransaksiObat.GrandTotal, &setTransaksiObat.TotalPembayaran, &setTransaksiObat.Kembalian, &setTransaksiObat.DibayarAt,
			&setTransaksiObat.CreatedAt, &setTransaksiObat.CreatedBy,
			&setTransaksiObat.ModifiedAt, &setTransaksiObat.ModifiedBy,
		)

		if err != nil {
			return nil, err
		}

		HasilGetTransaksiObat = append(HasilGetTransaksiObat, setTransaksiObat)
	}

	if HasilGetTransaksiObat == nil {
		HasilGetTransaksiObat = []structs.TransaksiObat{}
	}

	return HasilGetTransaksiObat, nil

}

func GetTransaksiObatByID(db *sql.DB, IdTransaksiObat int) (HasilGetTransaksiObat structs.TransaksiObat, err error) {

	querySelect := `
		SELECT 	t.id, t.no_resep, t.pasien_id, p.nama, t.status, 
				t.tipe_pembayaran, t.grand_total, t.total_pembayaran, t.kembalian, t.dibayar_at,
				t.created_at, t.created_by, t.modified_at, t.modified_by
		FROM transaksi_obat t
		LEFT JOIN pasien p ON t.pasien_id = p.id
		WHERE t.id = $1
	`

	err = db.QueryRow(querySelect, IdTransaksiObat).Scan(
		&HasilGetTransaksiObat.ID, &HasilGetTransaksiObat.NoResep, &HasilGetTransaksiObat.PasienID, &HasilGetTransaksiObat.NamaPasien, &HasilGetTransaksiObat.Status,
		&HasilGetTransaksiObat.TipePembayaran, &HasilGetTransaksiObat.GrandTotal, &HasilGetTransaksiObat.TotalPembayaran, &HasilGetTransaksiObat.Kembalian, &HasilGetTransaksiObat.DibayarAt,
		&HasilGetTransaksiObat.CreatedAt, &HasilGetTransaksiObat.CreatedBy,
		&HasilGetTransaksiObat.ModifiedAt, &HasilGetTransaksiObat.ModifiedBy,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return structs.TransaksiObat{}, fmt.Errorf("Transaksi Obat ID %d tidak ditemukan", IdTransaksiObat)
		}
		return structs.TransaksiObat{}, err
	}

	return HasilGetTransaksiObat, nil

}

func TambahTransaksiObat(db *sql.DB, setTransaksiObat structs.TambahTransaksiObatRequest, CreatedBy string) (HasilTambahTransaksiObat structs.TransaksiObat, err error) {

	queryInsert := `
		INSERT INTO transaksi_obat (
			pasien_id, created_by
		)
		VALUES ($1, $2)
		RETURNING 	id, no_resep, pasien_id, status, 
					tipe_pembayaran, grand_total, total_pembayaran, kembalian, dibayar_at,
					created_at, created_by, modified_at, modified_by
	`

	err = db.QueryRow(queryInsert, setTransaksiObat.PasienID, CreatedBy).Scan(
		&HasilTambahTransaksiObat.ID, &HasilTambahTransaksiObat.NoResep, &HasilTambahTransaksiObat.PasienID, &HasilTambahTransaksiObat.Status,
		&HasilTambahTransaksiObat.TipePembayaran, &HasilTambahTransaksiObat.GrandTotal, &HasilTambahTransaksiObat.TotalPembayaran, &HasilTambahTransaksiObat.Kembalian, &HasilTambahTransaksiObat.DibayarAt,
		&HasilTambahTransaksiObat.CreatedAt, &HasilTambahTransaksiObat.CreatedBy, &HasilTambahTransaksiObat.ModifiedAt, &HasilTambahTransaksiObat.ModifiedBy,
	)

	if err != nil {
		return structs.TransaksiObat{}, err
	}

	queryPasien := `
		SELECT nama FROM pasien WHERE id = $1
	`

	err = db.QueryRow(queryPasien, HasilTambahTransaksiObat.PasienID).
		Scan(&HasilTambahTransaksiObat.NamaPasien)

	if err != nil {
		return structs.TransaksiObat{}, err
	}

	return HasilTambahTransaksiObat, nil

}

func UpdateTransaksiObat(db *sql.DB, IdTransaksiObat int, setTransaksiObat structs.UpdateTransaksiObatRequest, ModifiedBy string) (HasilUpdateTransaksiObat structs.TransaksiObat, err error) {
	queryUpdate := `
		UPDATE transaksi_obat
		SET
			status = CASE WHEN $1 = '' THEN status ELSE $1 END,
			modified_by = $2
		WHERE id = $3
		RETURNING 	id, no_resep, pasien_id, status, 
					tipe_pembayaran, grand_total, total_pembayaran, kembalian, dibayar_at,
					created_at, created_by, modified_at, modified_by
	`

	err = db.QueryRow(
		queryUpdate,
		setTransaksiObat.Status,
		ModifiedBy,
		IdTransaksiObat,
	).Scan(
		&HasilUpdateTransaksiObat.ID, &HasilUpdateTransaksiObat.NoResep, &HasilUpdateTransaksiObat.PasienID, &HasilUpdateTransaksiObat.Status,
		&HasilUpdateTransaksiObat.TipePembayaran, &HasilUpdateTransaksiObat.GrandTotal, &HasilUpdateTransaksiObat.TotalPembayaran, &HasilUpdateTransaksiObat.Kembalian, &HasilUpdateTransaksiObat.DibayarAt,
		&HasilUpdateTransaksiObat.CreatedAt, &HasilUpdateTransaksiObat.CreatedBy,
		&HasilUpdateTransaksiObat.ModifiedAt, &HasilUpdateTransaksiObat.ModifiedBy,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return structs.TransaksiObat{}, fmt.Errorf("Transaksi Obat ID %d tidak ditemukan", IdTransaksiObat)
		}
		return structs.TransaksiObat{}, err
	}

	queryPasien := `
		SELECT nama FROM pasien WHERE id = $1
	`

	err = db.QueryRow(queryPasien, HasilUpdateTransaksiObat.PasienID).
		Scan(&HasilUpdateTransaksiObat.NamaPasien)

	if err != nil {
		if err == sql.ErrNoRows {
			return structs.TransaksiObat{}, fmt.Errorf("Transaksi Obat ID %d tidak ditemukan", IdTransaksiObat)
		}
		return structs.TransaksiObat{}, err
	}

	return HasilUpdateTransaksiObat, nil
}

func DeleteTransaksiObat(db *sql.DB, IdTransaksiObat int) (err error) {
	queryDelete := `
		DELETE FROM transaksi_obat WHERE id = $1
	`

	hasilDelete, err := db.Exec(queryDelete, IdTransaksiObat)

	if err != nil {
		return err
	}

	affectRow, err := hasilDelete.RowsAffected()

	if err != nil {
		return err
	}

	if affectRow == 0 {
		return fmt.Errorf("Transaksi Obat ID %d tidak ditemukan", IdTransaksiObat)
	}

	fmt.Printf("Transaksi Obat ID %d berhasil dihapus\n", IdTransaksiObat)

	return nil
}

func CancelTransaksiObat(db *sql.DB, IdTransaksiObat int, ModifiedBy string) (HasilCancelTransaksiObat structs.TransaksiObat, err error) {

	execCancelTransaksiObat, err := db.Begin()

	if err != nil {
		return structs.TransaksiObat{}, err
	}

	defer func() {
		if err != nil {
			_ = execCancelTransaksiObat.Rollback()
		}
	}()

	queryCekTransaksi := `
		SELECT no_resep, status
		FROM transaksi_obat
		WHERE id = $1
	`

	var (
		noResep string
		status  string
	)

	err = execCancelTransaksiObat.QueryRow(queryCekTransaksi, IdTransaksiObat).Scan(
		&noResep,
		&status,
	)

	if err == sql.ErrNoRows {
		return structs.TransaksiObat{}, fmt.Errorf("Transaksi Obat ID %d tidak ditemukan", IdTransaksiObat)
	}

	if err != nil {
		return structs.TransaksiObat{}, err
	}

	if status != "PENDING" {
		return structs.TransaksiObat{}, fmt.Errorf("Transaksi tidak dapat dibatalkan karena status saat ini %s", status)
	}

	queryDetail := `
		SELECT obat_id, jumlah FROM detail_transaksi_obat WHERE transaksi_obat_id = $1
	`

	rows, err := execCancelTransaksiObat.Query(queryDetail, IdTransaksiObat)

	if err != nil {
		return structs.TransaksiObat{}, err
	}

	type detailTransaksiObat struct {
		obatID int
		jumlah float64
	}

	var listDetailTransaksiObat []detailTransaksiObat

	for rows.Next() {
		var detailTransaksi detailTransaksiObat
		err = rows.Scan(&detailTransaksi.obatID, &detailTransaksi.jumlah)
		if err != nil {
			rows.Close()
			return structs.TransaksiObat{}, err
		}
		listDetailTransaksiObat = append(listDetailTransaksiObat, detailTransaksi)
	}

	rows.Close()

	if err = rows.Err(); err != nil {
		return structs.TransaksiObat{}, err
	}

	if len(listDetailTransaksiObat) == 0 {
		return structs.TransaksiObat{}, fmt.Errorf("Transaksi obat ID %d belum memiliki detail obat", IdTransaksiObat)
	}

	queryTambahStok := `
		UPDATE obat
		SET stok = stok + $1,
			modified_by = $2
		WHERE id = $3
	`

	for _, detail := range listDetailTransaksiObat {
		hasilCancelData, err := execCancelTransaksiObat.Exec(
			queryTambahStok,
			detail.jumlah,
			ModifiedBy,
			detail.obatID,
		)
		if err != nil {
			return structs.TransaksiObat{}, err
		}

		affectRow, err := hasilCancelData.RowsAffected()
		if err != nil {
			return structs.TransaksiObat{}, err
		}

		if affectRow == 0 {
			return structs.TransaksiObat{}, fmt.Errorf("Obat ID %d tidak ditemukan", detail.obatID)
		}
	}

	queryCancel := `
		UPDATE transaksi_obat
		SET status = 'CANCELED',
			modified_by = $1
		WHERE id = $2
		RETURNING id
	`

	var idCancelTransaksiObat int

	err = execCancelTransaksiObat.QueryRow(queryCancel, ModifiedBy, IdTransaksiObat).Scan(&idCancelTransaksiObat)

	if err != nil {
		return structs.TransaksiObat{}, err
	}

	queryCekDetailTransaksiObat := `SELECT COUNT(*) FROM detail_transaksi_obat WHERE transaksi_obat_id = $1`

	var jumlahDetailTransaksiObat int

	err = execCancelTransaksiObat.QueryRow(queryCekDetailTransaksiObat, IdTransaksiObat).Scan(&jumlahDetailTransaksiObat)

	if err != nil {
		return structs.TransaksiObat{}, err
	}

	if jumlahDetailTransaksiObat == 0 {
		return structs.TransaksiObat{}, fmt.Errorf("transaksi ID %d belum memiliki detail obat", IdTransaksiObat)
	}

	queryHasilCancelTransaksiObat := `
		SELECT 	t.id, t.no_resep, t.pasien_id, p.nama,
				t.status, t.created_at, t.created_by, t.modified_at, t.modified_by
		FROM transaksi_obat t
		INNER JOIN pasien p ON p.id = t.pasien_id
		WHERE t.id = $1
	`

	err = execCancelTransaksiObat.QueryRow(queryHasilCancelTransaksiObat, idCancelTransaksiObat).Scan(
		&HasilCancelTransaksiObat.ID, &HasilCancelTransaksiObat.NoResep, &HasilCancelTransaksiObat.PasienID,
		&HasilCancelTransaksiObat.NamaPasien, &HasilCancelTransaksiObat.Status, &HasilCancelTransaksiObat.CreatedAt,
		&HasilCancelTransaksiObat.CreatedBy, &HasilCancelTransaksiObat.ModifiedAt, &HasilCancelTransaksiObat.ModifiedBy,
	)

	if err != nil {
		return structs.TransaksiObat{}, err
	}

	err = execCancelTransaksiObat.Commit()

	if err != nil {
		return structs.TransaksiObat{}, err
	}

	return HasilCancelTransaksiObat, nil
}

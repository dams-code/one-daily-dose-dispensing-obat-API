package repositories

import (
	"database/sql"
	"fmt"
	"one-daily-dose-dispensing-obat-api/structs"
)

func GetDetailTransaksiObat(db *sql.DB, IdDetailTransaksiObat int, NoDetailResep string) (HasilGetDetailTransaksiObat []structs.DetailTransaksiObat, err error) {

	querySelect := `
		SELECT 	dto.id, dto.no_detail_resep, dto.transaksi_obat_id, t.no_resep,
				p.nama AS nama_pasien, o.id, o.kode_obat, o.nama_obat, dto.jumlah, o.satuan, o.harga, (dto.jumlah * o.harga) AS subtotal,
				dto.aturan_pakai, dto.created_at, dto.created_by, dto.modified_at, dto.modified_by
		FROM detail_transaksi_obat dto
		INNER JOIN transaksi_obat t ON dto.transaksi_obat_id = t.id
		INNER JOIN pasien p ON t.pasien_id = p.id
		LEFT JOIN obat o ON dto.obat_id = o.id
		WHERE 1=1
	`

	var paramDetailObat []interface{}
	jumlahParam := 1

	if IdDetailTransaksiObat > 0 {
		querySelect += fmt.Sprintf(" AND dto.id = $%d", jumlahParam)
		paramDetailObat = append(paramDetailObat, IdDetailTransaksiObat)
		jumlahParam++
	}

	if NoDetailResep != "" {
		querySelect += fmt.Sprintf(" AND dto.no_detail_resep ILIKE $%d", jumlahParam)
		paramDetailObat = append(paramDetailObat, "%"+NoDetailResep+"%")
		jumlahParam++
	}

	querySelect += " ORDER BY dto.id"

	rows, err := db.Query(querySelect, paramDetailObat...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {

		var detail structs.DetailTransaksiObat

		err = rows.Scan(
			&detail.ID, &detail.NoDetailResep, &detail.TransaksiObatID,
			&detail.NoResep,
			&detail.NamaPasien, &detail.ObatID, &detail.KodeObat,
			&detail.NamaObat, &detail.Jumlah, &detail.Satuan, &detail.Harga,
			&detail.Subtotal, &detail.AturanPakai, &detail.CreatedAt,
			&detail.CreatedBy, &detail.ModifiedAt, &detail.ModifiedBy,
		)

		if err != nil {
			return nil, err
		}

		HasilGetDetailTransaksiObat = append(HasilGetDetailTransaksiObat, detail)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if HasilGetDetailTransaksiObat == nil {
		HasilGetDetailTransaksiObat = []structs.DetailTransaksiObat{}
	}

	return HasilGetDetailTransaksiObat, nil
}

func GetDetailTransaksiObatByID(db *sql.DB, IdDetailTransaksiObat int) (HasilGetDetailTransaksiObat structs.DetailTransaksiObat, err error) {

	querySelect := `
		SELECT
			dto.id, dto.no_detail_resep, dto.transaksi_obat_id, t.no_resep,
			p.nama, dto.obat_id, o.kode_obat, o.nama_obat,
			dto.jumlah, o.satuan, o.harga, (dto.jumlah * o.harga) AS subtotal,
			dto.aturan_pakai, dto.created_at, dto.created_by, dto.modified_at, dto.modified_by
		FROM detail_transaksi_obat dto
		INNER JOIN transaksi_obat t ON dto.transaksi_obat_id = t.id
		INNER JOIN pasien p ON t.pasien_id = p.id
		INNER JOIN obat o ON dto.obat_id = o.id
		WHERE dto.id = $1
	`

	err = db.QueryRow(querySelect, IdDetailTransaksiObat).Scan(
		&HasilGetDetailTransaksiObat.ID, &HasilGetDetailTransaksiObat.NoDetailResep,
		&HasilGetDetailTransaksiObat.TransaksiObatID, &HasilGetDetailTransaksiObat.NoResep,
		&HasilGetDetailTransaksiObat.NamaPasien, &HasilGetDetailTransaksiObat.ObatID,
		&HasilGetDetailTransaksiObat.KodeObat, &HasilGetDetailTransaksiObat.NamaObat, &HasilGetDetailTransaksiObat.Jumlah,
		&HasilGetDetailTransaksiObat.Satuan, &HasilGetDetailTransaksiObat.Harga, &HasilGetDetailTransaksiObat.Subtotal,
		&HasilGetDetailTransaksiObat.AturanPakai, &HasilGetDetailTransaksiObat.CreatedAt, &HasilGetDetailTransaksiObat.CreatedBy,
		&HasilGetDetailTransaksiObat.ModifiedAt, &HasilGetDetailTransaksiObat.ModifiedBy,
	)

	if err == sql.ErrNoRows {
		return structs.DetailTransaksiObat{}, fmt.Errorf("Detail transaksi obat ID %d tidak ditemukan", IdDetailTransaksiObat)
	}

	if err != nil {
		return structs.DetailTransaksiObat{}, err
	}

	return HasilGetDetailTransaksiObat, nil
}

func GetDetailTransaksiObatByTransaksiID(db *sql.DB, IdTransaksiObat int) (HasilGetDetailTransaksiObat structs.DetailTransaksiObatResponse, err error) {

	queryHeader := `
		SELECT 	t.id, t.no_resep,
				p.nama, t.status, t.created_at, t.tipe_pembayaran, t.total_pembayaran, t.kembalian
		FROM transaksi_obat t
		INNER JOIN pasien p ON t.pasien_id = p.id
		WHERE t.id = $1
	`

	err = db.QueryRow(queryHeader, IdTransaksiObat).Scan(
		&HasilGetDetailTransaksiObat.ID,
		&HasilGetDetailTransaksiObat.NoResep, &HasilGetDetailTransaksiObat.NamaPasien, &HasilGetDetailTransaksiObat.Status,
		&HasilGetDetailTransaksiObat.CreatedAt, &HasilGetDetailTransaksiObat.TipePembayaran, &HasilGetDetailTransaksiObat.TotalPembayaran, &HasilGetDetailTransaksiObat.Kembalian,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return structs.DetailTransaksiObatResponse{}, fmt.Errorf("transaksi obat tidak ditemukan")
		}
		return structs.DetailTransaksiObatResponse{}, err
	}

	queryDetail := `
		SELECT 	dto.id, dto.no_detail_resep, dto.transaksi_obat_id, t.no_resep,
				p.nama, dto.obat_id, o.kode_obat,
				o.nama_obat, dto.jumlah, o.satuan, o.harga,
				(dto.jumlah * o.harga) AS subtotal,
				dto.aturan_pakai, dto.created_at, dto.created_by, 
				dto.modified_at, dto.modified_by
		FROM detail_transaksi_obat dto
		INNER JOIN obat o ON dto.obat_id = o.id
		JOIN transaksi_obat t ON t.id = dto.transaksi_obat_id
		JOIN pasien p ON p.id = t.pasien_id
		WHERE dto.transaksi_obat_id = $1
		ORDER BY dto.id
	`

	rows, err := db.Query(queryDetail, IdTransaksiObat)

	if err != nil {
		return structs.DetailTransaksiObatResponse{}, err
	}

	defer rows.Close()

	for rows.Next() {

		var detail structs.DetailTransaksiObat

		err = rows.Scan(
			&detail.ID, &detail.NoDetailResep, &detail.TransaksiObatID, &detail.NoResep,
			&detail.NamaPasien, &detail.ObatID, &detail.KodeObat,
			&detail.NamaObat, &detail.Jumlah, &detail.Satuan, &detail.Harga, &detail.Subtotal,
			&detail.AturanPakai, &detail.CreatedAt, &detail.CreatedBy,
			&detail.ModifiedAt, &detail.ModifiedBy,
		)

		if err != nil {
			return structs.DetailTransaksiObatResponse{}, err
		}

		HasilGetDetailTransaksiObat.GrandTotal += detail.Subtotal
		HasilGetDetailTransaksiObat.DetailObat = append(HasilGetDetailTransaksiObat.DetailObat, detail)
	}

	if err = rows.Err(); err != nil {
		return structs.DetailTransaksiObatResponse{}, err
	}

	if HasilGetDetailTransaksiObat.DetailObat == nil {
		HasilGetDetailTransaksiObat.DetailObat = []structs.DetailTransaksiObat{}
	}

	return HasilGetDetailTransaksiObat, nil
}

func TambahDetailTransaksiObat(db *sql.DB, IdTransaksiObat int, setDetailTransaksiObat structs.TambahDetailTransaksiObatRequest, CreatedBy string) (HasiTambahDetailTransaksiObat []structs.DetailTransaksiObat, err error) {

	execDetailTransaksi, err := db.Begin()

	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			_ = execDetailTransaksi.Rollback()
		}
	}()

	querySelect := `SELECT status FROM transaksi_obat WHERE id = $1`

	var status string

	err = execDetailTransaksi.QueryRow(querySelect, IdTransaksiObat).Scan(&status)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Transaksi obat ID %d tidak ditemukan", IdTransaksiObat)
		}
		return nil, err
	}

	if status != "PENDING" {
		return nil, fmt.Errorf("Transaksi sudah di-Dispensed (diberikan ke pasien) atau ter-Canceled")
	}

	queryCekObat := `SELECT nama_obat, stok FROM obat WHERE id = $1`

	for _, item := range setDetailTransaksiObat.ItemObat {
		var namaObat string
		var stok float64

		err = execDetailTransaksi.QueryRow(queryCekObat, item.ObatID).Scan(&namaObat, &stok)

		if err != nil {
			if err == sql.ErrNoRows {
				return nil, fmt.Errorf("Obat ID %d tidak ditemukan", item.ObatID)
			}
			return nil, err
		}

		if stok < item.Jumlah {
			return nil, fmt.Errorf("stok obat %s tidak mencukupi. Stok tersedia %.2f, diminta %.2f", namaObat, stok, item.Jumlah)
		}
	}

	queryInsert := `
		INSERT INTO detail_transaksi_obat (transaksi_obat_id, obat_id, jumlah, aturan_pakai, created_by)
		VALUES($1, $2, $3, $4, $5)
		RETURNING id
	`

	queryUpdateStok := `UPDATE obat SET stok = stok - $1, modified_by = $2 WHERE id = $3`

	querySelectDetail := ` 
		SELECT 	dto.id, dto.no_detail_resep, dto.transaksi_obat_id, t.no_resep,
				p.nama, dto.obat_id, o.kode_obat, o.nama_obat, dto.jumlah,
				o.satuan, o.harga, (dto.jumlah * o.harga) AS subtotal, dto.aturan_pakai, dto.created_at,
				dto.created_by, dto.modified_at, dto.modified_by
		FROM detail_transaksi_obat dto
		INNER JOIN transaksi_obat t ON dto.transaksi_obat_id = t.id
		INNER JOIN pasien p ON t.pasien_id = p.id
		INNER JOIN obat o ON dto.obat_id = o.id
		WHERE dto.id = $1
	`

	for _, item := range setDetailTransaksiObat.ItemObat {

		var idDetailTransaksiObat int

		err = execDetailTransaksi.QueryRow(queryInsert, IdTransaksiObat, item.ObatID, item.Jumlah, item.AturanPakai, CreatedBy).Scan(
			&idDetailTransaksiObat,
		)

		if err != nil {
			return nil, err
		}

		_, err = execDetailTransaksi.Exec(queryUpdateStok, item.Jumlah, CreatedBy, item.ObatID)

		if err != nil {
			return nil, err
		}

		var hasilTempDetailTransaksiObat structs.DetailTransaksiObat

		err = execDetailTransaksi.QueryRow(querySelectDetail, idDetailTransaksiObat).Scan(
			&hasilTempDetailTransaksiObat.ID,
			&hasilTempDetailTransaksiObat.NoDetailResep, &hasilTempDetailTransaksiObat.TransaksiObatID, &hasilTempDetailTransaksiObat.NoResep, &hasilTempDetailTransaksiObat.NamaPasien,
			&hasilTempDetailTransaksiObat.ObatID, &hasilTempDetailTransaksiObat.KodeObat, &hasilTempDetailTransaksiObat.NamaObat, &hasilTempDetailTransaksiObat.Jumlah,
			&hasilTempDetailTransaksiObat.Satuan, &hasilTempDetailTransaksiObat.Harga, &hasilTempDetailTransaksiObat.Subtotal, &hasilTempDetailTransaksiObat.AturanPakai,
			&hasilTempDetailTransaksiObat.CreatedAt, &hasilTempDetailTransaksiObat.CreatedBy, &hasilTempDetailTransaksiObat.ModifiedAt, &hasilTempDetailTransaksiObat.ModifiedBy,
		)

		if err != nil {
			return nil, err
		}

		HasiTambahDetailTransaksiObat = append(HasiTambahDetailTransaksiObat, hasilTempDetailTransaksiObat)
	}

	err = execDetailTransaksi.Commit()

	if err != nil {
		return nil, err
	}

	if HasiTambahDetailTransaksiObat == nil {
		HasiTambahDetailTransaksiObat = []structs.DetailTransaksiObat{}
	}

	return HasiTambahDetailTransaksiObat, nil

}

func UpdateDetailTransaksiObat(db *sql.DB, IdDetailTransaksiObat int, setDetailTransaksiObat structs.UpdateDetailTransaksiObatRequest, ModifiedBy string,
) (HasilUpdateDetailTransaksiObat structs.DetailTransaksiObat, err error) {

	execUpdateDetailTransaksiObat, err := db.Begin()
	if err != nil {
		return HasilUpdateDetailTransaksiObat, err
	}

	defer func() {
		if err != nil {
			_ = execUpdateDetailTransaksiObat.Rollback()
		}
	}()

	if setDetailTransaksiObat.Jumlah == nil && setDetailTransaksiObat.AturanPakai == nil {
		return HasilUpdateDetailTransaksiObat,
			fmt.Errorf("minimal jumlah atau aturan_pakai harus diisi")
	}

	queryStokObatLama := `
		SELECT dto.transaksi_obat_id, dto.obat_id, dto.jumlah, t.status
		FROM detail_transaksi_obat dto
		JOIN transaksi_obat t
			ON dto.transaksi_obat_id = t.id
		WHERE dto.id = $1
	`

	var (
		transaksiID int
		obatID      int
		jumlahLama  float64
		status      string
	)

	err = execUpdateDetailTransaksiObat.QueryRow(queryStokObatLama, IdDetailTransaksiObat).Scan(&transaksiID, &obatID, &jumlahLama, &status)

	if err == sql.ErrNoRows {
		return HasilUpdateDetailTransaksiObat, fmt.Errorf("detail transaksi tidak ditemukan")
	}

	if err != nil {
		return HasilUpdateDetailTransaksiObat, err
	}

	if status != "PENDING" {
		return HasilUpdateDetailTransaksiObat,
			fmt.Errorf("transaksi sudah tidak dapat diubah")
	}

	if setDetailTransaksiObat.Jumlah != nil {

		jumlahBaru := *setDetailTransaksiObat.Jumlah

		var stok float64
		var namaObat string

		err = execUpdateDetailTransaksiObat.QueryRow(`SELECT stok,nama_obat FROM obat WHERE id=$1`, obatID).Scan(&stok, &namaObat)

		if err != nil {
			return HasilUpdateDetailTransaksiObat, err
		}

		delta := jumlahBaru - jumlahLama

		if delta > 0 {

			if stok < delta {
				return HasilUpdateDetailTransaksiObat,
					fmt.Errorf("stok obat %s tidak mencukupi. tersedia %.2f diminta %.2f", namaObat, stok, delta)
			}

			_, err = execUpdateDetailTransaksiObat.Exec(`
				UPDATE obat
				SET stok = stok - $1,
					modified_by=$2
				WHERE id=$3
			`, delta, ModifiedBy, obatID)

		} else if delta < 0 {

			_, err = execUpdateDetailTransaksiObat.Exec(`
				UPDATE obat
				SET stok = stok + $1,
					modified_by=$2
				WHERE id=$3
			`, -delta, ModifiedBy, obatID)

		}

		if err != nil {
			return HasilUpdateDetailTransaksiObat, err
		}
	}

	queryUpdate := `
		UPDATE detail_transaksi_obat
		SET jumlah = COALESCE($1,jumlah),
			aturan_pakai = COALESCE($2,aturan_pakai),
			modified_at = CURRENT_TIMESTAMP,
			modified_by = $3
		WHERE id = $4
		RETURNING id
	`

	var idDetail int

	err = execUpdateDetailTransaksiObat.QueryRow(queryUpdate, setDetailTransaksiObat.Jumlah, setDetailTransaksiObat.AturanPakai, ModifiedBy, IdDetailTransaksiObat).Scan(&idDetail)

	if err != nil {
		return HasilUpdateDetailTransaksiObat, err
	}

	_, err = execUpdateDetailTransaksiObat.Exec(`
		UPDATE transaksi_obat
		SET grand_total = (
				SELECT COALESCE(SUM(dto.jumlah * o.harga),0)
				FROM detail_transaksi_obat dto
				JOIN obat o ON dto.obat_id=o.id
				WHERE dto.transaksi_obat_id=$1
			),
		modified_at = CURRENT_TIMESTAMP,
		modified_by = $2
		WHERE id = $1
	`, transaksiID, ModifiedBy)

	if err != nil {
		return HasilUpdateDetailTransaksiObat, err
	}

	queryResult := `
	SELECT
		dto.id, dto.no_detail_resep, dto.transaksi_obat_id, t.no_resep,
		p.nama, dto.obat_id, o.kode_obat, o.nama_obat,
		dto.jumlah, o.satuan, o.harga, (dto.jumlah*o.harga) as subtotal,
		dto.aturan_pakai, dto.created_at, dto.created_by, dto.modified_at, dto.modified_by
	FROM detail_transaksi_obat dto
	JOIN transaksi_obat t ON dto.transaksi_obat_id=t.id
	JOIN pasien p ON t.pasien_id=p.id
	JOIN obat o ON dto.obat_id=o.id
	WHERE dto.id=$1
	`

	err = execUpdateDetailTransaksiObat.QueryRow(queryResult, idDetail).Scan(
		&HasilUpdateDetailTransaksiObat.ID, &HasilUpdateDetailTransaksiObat.NoDetailResep, &HasilUpdateDetailTransaksiObat.TransaksiObatID, &HasilUpdateDetailTransaksiObat.NoResep,
		&HasilUpdateDetailTransaksiObat.NamaPasien, &HasilUpdateDetailTransaksiObat.ObatID, &HasilUpdateDetailTransaksiObat.KodeObat, &HasilUpdateDetailTransaksiObat.NamaObat,
		&HasilUpdateDetailTransaksiObat.Jumlah, &HasilUpdateDetailTransaksiObat.Satuan, &HasilUpdateDetailTransaksiObat.Harga, &HasilUpdateDetailTransaksiObat.Subtotal,
		&HasilUpdateDetailTransaksiObat.AturanPakai, &HasilUpdateDetailTransaksiObat.CreatedAt, &HasilUpdateDetailTransaksiObat.CreatedBy,
		&HasilUpdateDetailTransaksiObat.ModifiedAt, &HasilUpdateDetailTransaksiObat.ModifiedBy,
	)

	if err != nil {
		return HasilUpdateDetailTransaksiObat, err
	}

	err = execUpdateDetailTransaksiObat.Commit()
	if err != nil {
		return HasilUpdateDetailTransaksiObat, err
	}

	return HasilUpdateDetailTransaksiObat, nil
}

func DeleteDetailTransaksiObat(db *sql.DB, IdDetailTransaksiObat int, ModifiedBy string) (err error) {
	execDeleteDetail, err := db.Begin()

	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = execDeleteDetail.Rollback()
		}
	}()

	queryGetDetail := `
		SELECT
			dto.obat_id,
			dto.jumlah,
			t.status
		FROM detail_transaksi_obat dto
		INNER JOIN transaksi_obat t
			ON dto.transaksi_obat_id = t.id
		WHERE dto.id = $1
	`

	var (
		obatID int
		jumlah float64
		status string
	)

	err = execDeleteDetail.QueryRow(queryGetDetail, IdDetailTransaksiObat).Scan(
		&obatID,
		&jumlah,
		&status,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("Detail transaksi obat ID %d tidak ditemukan", IdDetailTransaksiObat)
		}

		return err
	}

	if status != "PENDING" {
		return fmt.Errorf("transaksi sudah tidak dapat dihapus")
	}

	queryKembalikanStok := `
		UPDATE obat
		SET stok = stok + $1, modified_by = $2
		WHERE id = $3
	`

	result, err := execDeleteDetail.Exec(
		queryKembalikanStok,
		jumlah,
		ModifiedBy,
		obatID,
	)

	if err != nil {
		return err
	}

	affectedRow, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if affectedRow == 0 {
		return fmt.Errorf("obat ID %d tidak ditemukan", obatID)
	}

	queryDelete := `DELETE FROM detail_transaksi_obat WHERE id = $1`

	result, err = execDeleteDetail.Exec(queryDelete, IdDetailTransaksiObat)

	if err != nil {
		return err
	}

	affectedRow, err = result.RowsAffected()

	if err != nil {
		return err
	}

	if affectedRow == 0 {
		return fmt.Errorf("Detail transaksi obat ID %d tidak ditemukan", IdDetailTransaksiObat)
	}

	err = execDeleteDetail.Commit()

	if err != nil {
		return err
	}

	return nil

}

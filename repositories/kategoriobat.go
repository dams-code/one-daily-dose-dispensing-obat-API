package repositories

import (
	"database/sql"
	"fmt"
	"one-daily-dose-dispensing-obat-api/structs"
)

func GetKategoriObat(db *sql.DB, IdKategoriObat int, NamaKategori string) (HasilGetKategoriObat []structs.KategoriObat, err error) {

	querySelect := `
		SELECT 	id, nama_kategori, deskripsi, 
				created_at, created_by, modified_at, modified_by
		FROM kategori_obat
		WHERE 1=1
	`
	var paramKategoriObat []interface{}

	jumlahParam := 1

	if IdKategoriObat > 0 {
		querySelect += fmt.Sprintf(" AND id = $%d", jumlahParam)
		paramKategoriObat = append(paramKategoriObat, IdKategoriObat)
		jumlahParam++
	}

	if NamaKategori != "" {
		querySelect += fmt.Sprintf(" AND nama_kategori ILIKE $%d", jumlahParam)
		paramKategoriObat = append(paramKategoriObat, "%"+NamaKategori+"%")
		jumlahParam++
	}

	rows, err := db.Query(querySelect, paramKategoriObat...)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var setKategoriObat structs.KategoriObat

		err = rows.Scan(&setKategoriObat.ID, &setKategoriObat.NamaKategori, &setKategoriObat.Deskripsi,
			&setKategoriObat.CreatedAt, &setKategoriObat.CreatedBy,
			&setKategoriObat.ModifiedAt, &setKategoriObat.ModifiedBy)

		if err != nil {

			return nil, err
		}

		HasilGetKategoriObat = append(HasilGetKategoriObat, setKategoriObat)
	}

	if HasilGetKategoriObat == nil {
		HasilGetKategoriObat = []structs.KategoriObat{}
	}

	return HasilGetKategoriObat, nil
}

func GetKategoriObatID(db *sql.DB, IdKategoriObat int) (HasilGetKategoriObat structs.KategoriObat, err error) {

	querySelect := `
		SELECT 	id, nama_kategori, deskripsi, 
				created_at, created_by, modified_at, modified_by
		FROM kategori_obat
		WHERE id = $1
	`

	err = db.QueryRow(querySelect, IdKategoriObat).Scan(&HasilGetKategoriObat.ID, &HasilGetKategoriObat.NamaKategori, &HasilGetKategoriObat.Deskripsi,
		&HasilGetKategoriObat.CreatedAt, &HasilGetKategoriObat.CreatedBy,
		&HasilGetKategoriObat.ModifiedAt, &HasilGetKategoriObat.ModifiedBy,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return structs.KategoriObat{}, fmt.Errorf("Kategori obat ID %d tidak ditemukan", IdKategoriObat)
		}
		return structs.KategoriObat{}, err
	}

	return HasilGetKategoriObat, nil
}

func TambahKategoriObat(db *sql.DB, setKategoriObat structs.TambahdanUpdateKategoriObatRequest, CreatedBy string) (HasilTambahKategoriObat structs.KategoriObat, err error) {

	queryInsert := `
		INSERT INTO kategori_obat (nama_kategori, deskripsi, created_by)
		VALUES ($1, $2, $3)
		RETURNING id, nama_kategori, deskripsi, created_at, created_by, modified_at, modified_by
	`

	err = db.QueryRow(queryInsert, setKategoriObat.NamaKategori, setKategoriObat.Deskripsi, CreatedBy).Scan(
		&HasilTambahKategoriObat.ID, &HasilTambahKategoriObat.NamaKategori, &HasilTambahKategoriObat.Deskripsi,
		&HasilTambahKategoriObat.CreatedAt, &HasilTambahKategoriObat.CreatedBy,
		&HasilTambahKategoriObat.ModifiedAt, &HasilTambahKategoriObat.ModifiedBy,
	)

	if err != nil {
		return structs.KategoriObat{}, err
	}

	return HasilTambahKategoriObat, nil
}

func UpdateKategoriObat(db *sql.DB, IdKategoriObat int, setKategoriObat structs.TambahdanUpdateKategoriObatRequest, ModifiedBy string) (HasilUpdateKategoriObat structs.KategoriObat, err error) {

	queryUpdate := `
		UPDATE kategori_obat
		SET nama_kategori = CASE WHEN $1 = '' THEN nama_kategori ELSE $1 END,
		 	deskripsi = CASE WHEN $2 = '' THEN deskripsi ELSE $2 END,
			modified_by = $3
		WHERE id = $4
		RETURNING id, nama_kategori, deskripsi, created_at, created_by, modified_at, modified_by
	`

	err = db.QueryRow(queryUpdate, setKategoriObat.NamaKategori, setKategoriObat.Deskripsi, ModifiedBy, IdKategoriObat).Scan(
		&HasilUpdateKategoriObat.ID, &HasilUpdateKategoriObat.NamaKategori, &HasilUpdateKategoriObat.Deskripsi,
		&HasilUpdateKategoriObat.CreatedAt, &HasilUpdateKategoriObat.CreatedBy,
		&HasilUpdateKategoriObat.ModifiedAt, &HasilUpdateKategoriObat.ModifiedBy,
	)

	if err != nil {
		return structs.KategoriObat{}, err
	}

	return HasilUpdateKategoriObat, nil
}

func DeleteKategoriObat(db *sql.DB, IdKategoriObat int) (err error) {

	queryDelete := `
		DELETE FROM kategori_obat
		WHERE id = $1
	`

	HasilDeleteData, err := db.Exec(queryDelete, IdKategoriObat)

	if err != nil {
		return err
	}

	affectRow, err := HasilDeleteData.RowsAffected()

	if err != nil {
		return err
	}

	if affectRow == 0 {
		return fmt.Errorf("Kategori Obat ID %d Tidak ditemukan", IdKategoriObat)
	}

	fmt.Printf("Kategori Obat ID %d berhasil dihapus.", IdKategoriObat)
	return nil
}

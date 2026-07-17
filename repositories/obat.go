package repositories

import (
	"database/sql"
	"fmt"
	"one-daily-dose-dispensing-obat-api/structs"
)

func GetObat(db *sql.DB, IdObat int, NamaObat, KodeObat string) ([]structs.Obat, error) {

	querySelect := `
		SELECT o.id, o.kategori_id, k.nama_kategori,
			o.kode_obat, o.nama_obat, o.stok,
			o.satuan, o.harga,
			o.created_at, o.created_by,
			o.modified_at, o.modified_by
		FROM obat o
		LEFT JOIN kategori_obat k
			ON o.kategori_id = k.id
		WHERE 1=1
	`

	var paramObat []interface{}
	jumlahParam := 1

	if IdObat > 0 {
		querySelect += fmt.Sprintf(" AND o.id = $%d", jumlahParam)
		paramObat = append(paramObat, IdObat)
		jumlahParam++
	}

	if NamaObat != "" {
		querySelect += fmt.Sprintf(" AND o.nama_obat ILIKE $%d", jumlahParam)
		paramObat = append(paramObat, "%"+NamaObat+"%")
		jumlahParam++
	}

	if KodeObat != "" {
		querySelect += fmt.Sprintf(" AND o.kode_obat ILIKE $%d", jumlahParam)
		paramObat = append(paramObat, "%"+KodeObat+"%")
		jumlahParam++
	}

	rows, err := db.Query(querySelect, paramObat...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hasil []structs.Obat

	for rows.Next() {

		var obat structs.Obat

		err = rows.Scan(
			&obat.ID, &obat.KategoriID, &obat.NamaKategori, &obat.KodeObat,
			&obat.NamaObat, &obat.Stok, &obat.Satuan, &obat.Harga,
			&obat.CreatedAt, &obat.CreatedBy, &obat.ModifiedAt, &obat.ModifiedBy,
		)

		if err != nil {
			return nil, err
		}

		hasil = append(hasil, obat)
	}

	if hasil == nil {
		hasil = []structs.Obat{}
	}

	return hasil, nil
}

func GetObatByID(db *sql.DB, IdObat int) (structs.Obat, error) {

	querySelect := `
		SELECT o.id, o.kategori_id, k.nama_kategori,
			o.kode_obat, o.nama_obat, o.stok,
			o.satuan, o.harga, 
			o.created_at, o.created_by, o.modified_at, o.modified_by
		FROM obat o
		LEFT JOIN kategori_obat k
			ON o.kategori_id = k.id
		WHERE o.id = $1
	`

	var hasil structs.Obat

	err := db.QueryRow(querySelect, IdObat).Scan(
		&hasil.ID, &hasil.KategoriID, &hasil.NamaKategori, &hasil.KodeObat,
		&hasil.NamaObat, &hasil.Stok, &hasil.Satuan, &hasil.Harga,
		&hasil.CreatedAt, &hasil.CreatedBy, &hasil.ModifiedAt, &hasil.ModifiedBy,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return structs.Obat{}, fmt.Errorf("Obat ID %d tidak ditemukan", IdObat)
		}
		return structs.Obat{}, err
	}

	return hasil, nil
}

func GetObatByKategori(db *sql.DB, IdKategori int) ([]structs.Obat, error) {

	querySelect := `
		SELECT
			o.id, o.kategori_id, k.nama_kategori, o.kode_obat,
			o.nama_obat, o.stok, o.satuan, o.harga,
			o.created_at, o.created_by, o.modified_at, o.modified_by
		FROM obat o
		LEFT JOIN kategori_obat k
			ON o.kategori_id = k.id
		WHERE o.kategori_id = $1
	`

	rows, err := db.Query(querySelect, IdKategori)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hasil []structs.Obat

	for rows.Next() {

		var setObat structs.Obat

		err = rows.Scan(
			&setObat.ID, &setObat.KategoriID, &setObat.NamaKategori, &setObat.KodeObat,
			&setObat.NamaObat, &setObat.Stok, &setObat.Satuan, &setObat.Harga,
			&setObat.CreatedAt, &setObat.CreatedBy, &setObat.ModifiedAt, &setObat.ModifiedBy,
		)

		if err != nil {
			return nil, err
		}

		hasil = append(hasil, setObat)
	}

	if hasil == nil {
		hasil = []structs.Obat{}
	}

	return hasil, nil
}

func TambahObat(db *sql.DB, setObat structs.TambahObatRequest, CreatedBy string) (HasilTambahObat structs.Obat, err error) {

	queryInsert := `
		INSERT INTO obat ( kategori_id, nama_obat, stok, satuan, harga, created_by)
		VALUES ($1,$2,$3,$4,$5,$6)
		RETURNING
			id, kategori_id, kode_obat, nama_obat,
			stok, satuan, harga, created_at,
			created_by, modified_at, modified_by
	`

	err = db.QueryRow(
		queryInsert,
		setObat.KategoriID, setObat.NamaObat, setObat.Stok,
		setObat.Satuan, setObat.Harga, CreatedBy,
	).Scan(
		&HasilTambahObat.ID, &HasilTambahObat.KategoriID, &HasilTambahObat.KodeObat, &HasilTambahObat.NamaObat,
		&HasilTambahObat.Stok, &HasilTambahObat.Satuan, &HasilTambahObat.Harga, &HasilTambahObat.CreatedAt,
		&HasilTambahObat.CreatedBy, &HasilTambahObat.ModifiedAt, &HasilTambahObat.ModifiedBy,
	)

	if err != nil {
		return structs.Obat{}, err
	}

	err = db.QueryRow(`
		SELECT nama_kategori
		FROM kategori_obat
		WHERE id = $1
	`, HasilTambahObat.KategoriID).Scan(&HasilTambahObat.NamaKategori)

	if err != nil && err != sql.ErrNoRows {
		return structs.Obat{}, err
	}

	return HasilTambahObat, nil
}

func UpdateObat(db *sql.DB, IdObat int, setObat structs.UpdateObatRequest, ModifiedBy string) (HasilUpdateObat structs.Obat, err error) {

	queryUpdate := `
		WITH updated AS (
			UPDATE obat
			SET
				kategori_id = COALESCE($1, kategori_id),
				nama_obat   = COALESCE($2, nama_obat),
				stok        = COALESCE($3, stok),
				satuan      = COALESCE($4, satuan),
				harga       = COALESCE($5, harga),
				modified_by = $6
			WHERE id = $7
			RETURNING
				id, kategori_id, kode_obat,
				nama_obat, stok, satuan,
				harga, created_at, created_by,
				modified_at, modified_by
		)

		SELECT
			u.id, u.kategori_id, k.nama_kategori, u.kode_obat,
			u.nama_obat, u.stok, u.satuan, u.harga,
			u.created_at, u.created_by, u.modified_at, u.modified_by
		FROM updated u
		LEFT JOIN kategori_obat k ON u.kategori_id = k.id
	`

	err = db.QueryRow(queryUpdate, setObat.KategoriID,
		setObat.NamaObat, setObat.Stok, setObat.Satuan,
		setObat.Harga, ModifiedBy, IdObat,
	).Scan(
		&HasilUpdateObat.ID, &HasilUpdateObat.KategoriID, &HasilUpdateObat.NamaKategori, &HasilUpdateObat.KodeObat,
		&HasilUpdateObat.NamaObat, &HasilUpdateObat.Stok, &HasilUpdateObat.Satuan, &HasilUpdateObat.Harga,
		&HasilUpdateObat.CreatedAt, &HasilUpdateObat.CreatedBy, &HasilUpdateObat.ModifiedAt, &HasilUpdateObat.ModifiedBy,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return structs.Obat{}, fmt.Errorf("Obat ID %d tidak ditemukan", IdObat)
		}
		return structs.Obat{}, err
	}

	return HasilUpdateObat, nil
}

func DeleteObat(db *sql.DB, IdObat int) error {

	queryDelete := `
		DELETE FROM obat
		WHERE id = $1
	`

	hasilDelete, err := db.Exec(queryDelete, IdObat)

	if err != nil {
		return err
	}

	affectRow, err := hasilDelete.RowsAffected()

	if err != nil {
		return err
	}

	if affectRow == 0 {
		return fmt.Errorf("Obat ID %d tidak ditemukan", IdObat)
	}

	fmt.Printf("Obat ID %d berhasil dihapus.\n", IdObat)

	return nil
}

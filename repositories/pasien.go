package repositories

import (
	"database/sql"
	"fmt"
	"one-daily-dose-dispensing-obat-api/structs"
)

func GetPasien(db *sql.DB, IdPasien int, Rm string, NamaPasien string) (HasilGetPasien []structs.Pasien, err error) {

	querySelect := `
		SELECT 	id, rm, nama, jenis_kelamin, alamat,
				created_at, created_by, modified_at, modified_by
		FROM pasien
		WHERE 1=1
	`

	var paramPasien []interface{}

	jumlahParam := 1

	if IdPasien > 0 {
		querySelect += fmt.Sprintf(" AND id = $%d", jumlahParam)
		paramPasien = append(paramPasien, IdPasien)
		jumlahParam++
	}

	if Rm != "" {
		querySelect += fmt.Sprintf(" AND rm ILIKE $%d", jumlahParam)
		paramPasien = append(paramPasien, "%"+Rm+"%")
		jumlahParam++
	}

	if NamaPasien != "" {
		querySelect += fmt.Sprintf(" AND nama ILIKE $%d", jumlahParam)
		paramPasien = append(paramPasien, "%"+NamaPasien+"%")
		jumlahParam++
	}

	rows, err := db.Query(querySelect, paramPasien...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var setPasien structs.Pasien

		err = rows.Scan(&setPasien.ID, &setPasien.Rm, &setPasien.Nama,
			&setPasien.JenisKelamin, &setPasien.Alamat,
			&setPasien.CreatedAt, &setPasien.CreatedBy,
			&setPasien.ModifiedAt, &setPasien.ModifiedBy,
		)

		if err != nil {
			return nil, err
		}

		HasilGetPasien = append(HasilGetPasien, setPasien)
	}

	if HasilGetPasien == nil {
		HasilGetPasien = []structs.Pasien{}
	}

	return HasilGetPasien, nil

}

func GetPasienByID(db *sql.DB, IdPasien int) (HasilGetPasien structs.Pasien, err error) {

	querySelect := `SELECT id, rm, nama, jenis_kelamin, alamat,
				created_at, created_by, modified_at, modified_by
		FROM pasien
		WHERE id = $1
	`

	err = db.QueryRow(querySelect, IdPasien).Scan(&HasilGetPasien.ID, &HasilGetPasien.Rm, &HasilGetPasien.Nama,
		&HasilGetPasien.JenisKelamin, &HasilGetPasien.Alamat,
		&HasilGetPasien.CreatedAt, &HasilGetPasien.CreatedBy,
		&HasilGetPasien.ModifiedAt, &HasilGetPasien.ModifiedBy,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return structs.Pasien{}, fmt.Errorf("Pasien ID %d tidak ditemukan", IdPasien)
		}
		return structs.Pasien{}, err
	}

	return HasilGetPasien, nil
}

func TambahPasien(db *sql.DB, setPasien structs.TambahPasienRequest, CreatedBy string) (HasilTambahPasien structs.Pasien, err error) {

	queryInsert := `
		INSERT INTO Pasien (rm, nama, jenis_kelamin, alamat, created_by)
		VALUES($1, $2, $3, $4, $5)
		RETURNING id, rm, nama, jenis_kelamin, alamat, created_at,created_by, modified_at, modified_by
	`

	err = db.QueryRow(queryInsert, setPasien.Rm, setPasien.Nama, setPasien.JenisKelamin, setPasien.Alamat, CreatedBy).Scan(
		&HasilTambahPasien.ID,
		&HasilTambahPasien.Rm,
		&HasilTambahPasien.Nama,
		&HasilTambahPasien.JenisKelamin, &HasilTambahPasien.Alamat,
		&HasilTambahPasien.CreatedAt, &HasilTambahPasien.CreatedBy,
		&HasilTambahPasien.ModifiedAt, &HasilTambahPasien.ModifiedBy,
	)

	if err != nil {
		return HasilTambahPasien, err
	}

	return HasilTambahPasien, nil

}

func UpdatePasien(db *sql.DB, IdPasien int, setPasien structs.UpdatePasienRequest, ModifiedBy string) (HasilUpdatePasien structs.Pasien, err error) {

	queryUpdate := `
		UPDATE pasien
		SET rm = CASE WHEN $1 = '' THEN rm ELSE $1 END,
			nama = CASE WHEN $2 = '' THEN nama ELSE $2 END, 
			jenis_kelamin = CASE WHEN $3 = '' THEN jenis_kelamin ELSE $3 END, 
			alamat = CASE WHEN $4 = '' THEN alamat ELSE $4 END,
			modified_by = $5
		WHERE id = $6
		RETURNING id, rm, nama, jenis_kelamin, alamat, created_at, created_by, modified_at, modified_by
	`

	err = db.QueryRow(queryUpdate, setPasien.Rm, setPasien.Nama, setPasien.JenisKelamin, setPasien.Alamat, ModifiedBy, IdPasien).Scan(
		&HasilUpdatePasien.ID, &HasilUpdatePasien.Rm, &HasilUpdatePasien.Nama, &HasilUpdatePasien.JenisKelamin,
		&HasilUpdatePasien.Alamat, &HasilUpdatePasien.CreatedAt, &HasilUpdatePasien.CreatedBy,
		&HasilUpdatePasien.ModifiedAt, &HasilUpdatePasien.ModifiedBy,
	)

	if err != nil {
		return HasilUpdatePasien, err
	}

	return HasilUpdatePasien, nil
}

func DeletePasien(db *sql.DB, IdPasien int) (err error) {

	queryDelete := `
		DELETE FROM pasien
		WHERE id = $1
	`

	hasilDelete, err := db.Exec(queryDelete, IdPasien)

	if err != nil {
		return fmt.Errorf("Delete Pasien ID %d gagal, %v", IdPasien, err.Error())
	}

	affectRow, err := hasilDelete.RowsAffected()

	if err != nil {
		return fmt.Errorf("Delete Pasien ID %d gagal, %v", IdPasien, err.Error())
	}

	if affectRow == 0 {
		return fmt.Errorf("Pasien ID %d tidak ditemukan", IdPasien)
	}

	fmt.Printf("Pasien ID %d Berhasil dihapus", IdPasien)

	return nil
}

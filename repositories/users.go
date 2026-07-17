package repositories

import (
	"database/sql"
	"fmt"
	"one-daily-dose-dispensing-obat-api/structs"
)

func RegisterUser(db *sql.DB, setUserObat structs.RegisterUserRequest, CreatedBy string) (HasilRegisterUserObat structs.UserObat, err error) {

	queryInsert := `
		INSERT INTO user_obat (nama, username, password, role, created_by)
		VALUES($1, $2, $3, $4, $5)
		RETURNING id, nama, username, role, status, created_at, created_by, modified_at, modified_by
	`

	err = db.QueryRow(queryInsert, setUserObat.Nama, setUserObat.Username, setUserObat.Password, setUserObat.Role, CreatedBy).Scan(
		&HasilRegisterUserObat.ID, &HasilRegisterUserObat.Nama, &HasilRegisterUserObat.Username,
		&HasilRegisterUserObat.Role, &HasilRegisterUserObat.Status, &HasilRegisterUserObat.CreatedAt, &HasilRegisterUserObat.CreatedBy,
		&HasilRegisterUserObat.ModifiedAt, &HasilRegisterUserObat.ModifiedBy,
	)

	if err != nil {
		return HasilRegisterUserObat, err
	}

	return HasilRegisterUserObat, nil

}

func TambahUser(db *sql.DB, setUserObat structs.TambahUserRequest, CreatedBy string) (HasilTambahUserObat structs.UserObat, err error) {

	queryInsert := `
		INSERT INTO user_obat (nama, username, password, role, created_by)
		VALUES($1, $2, $3, $4, $5)
		RETURNING id, nama, username, role, status, created_at, created_by, modified_at, modified_by
	`

	err = db.QueryRow(queryInsert, setUserObat.Nama, setUserObat.Username, setUserObat.Password, setUserObat.Role, CreatedBy).Scan(
		&HasilTambahUserObat.ID, &HasilTambahUserObat.Nama, &HasilTambahUserObat.Username,
		&HasilTambahUserObat.Role, &HasilTambahUserObat.Status, &HasilTambahUserObat.CreatedAt, &HasilTambahUserObat.CreatedBy,
		&HasilTambahUserObat.ModifiedAt, &HasilTambahUserObat.ModifiedBy,
	)

	if err != nil {
		return HasilTambahUserObat, err
	}

	return HasilTambahUserObat, nil

}

func UpdateUser(db *sql.DB, IdUserObat int, setUserObat structs.UpdateUserRequest, ModifiedBy string) (HasilUpdateUserObat structs.UserObat, err error) {

	queryUpdate := `
		UPDATE user_obat
		SET nama = COALESCE($1, nama),
			password = COALESCE(NULLIF($2::text, ''), password),
			role = COALESCE($3, role),
			modified_by = $4
		WHERE id = $5
		RETURNING 	id, nama, username, role, status,
					created_at, created_by, modified_at, modified_by
	`

	err = db.QueryRow(queryUpdate, setUserObat.Nama, setUserObat.Password, setUserObat.Role, ModifiedBy, IdUserObat).Scan(
		&HasilUpdateUserObat.ID, &HasilUpdateUserObat.Nama, &HasilUpdateUserObat.Username,
		&HasilUpdateUserObat.Role, &HasilUpdateUserObat.Status, &HasilUpdateUserObat.CreatedAt,
		&HasilUpdateUserObat.CreatedBy, &HasilUpdateUserObat.ModifiedAt, &HasilUpdateUserObat.ModifiedBy,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return structs.UserObat{}, fmt.Errorf("User ID %d tidak ditemukan", IdUserObat)
		}
		return HasilUpdateUserObat, err
	}

	return HasilUpdateUserObat, nil
}

func DeleteUser(db *sql.DB, IdUserObat int, Username string) (err error) {
	queryDelete := `UPDATE user_obat
		SET status = false, modified_by = $2
		WHERE id = $1
	`

	hasilDeleteUser, err := db.Exec(queryDelete, IdUserObat, Username)

	if err != nil {
		return err
	}

	affectRow, err := hasilDeleteUser.RowsAffected()

	if err != nil {
		return err
	}

	if affectRow == 0 {
		return fmt.Errorf("User ID %d tidak ditemukan", IdUserObat)
	}

	fmt.Printf("User ID %d berhasil dinonaktifkan\n", IdUserObat)

	return nil
}

func AktivasiUser(db *sql.DB, idUser int, modifiedBy string) error {

	query := `
		UPDATE user_obat
		SET
			status = TRUE,
			modified_by = $1
		WHERE id = $2
	`

	result, err := db.Exec(query, modifiedBy, idUser)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("User ID %d tidak ditemukan", idUser)
	}

	return nil
}

func GetUsername(db *sql.DB, Username string) (HasilGetUser structs.UserObat, err error) {
	querySelect := `
		SELECT id, username, password, role
		FROM user_obat
		WHERE username = $1;
	`

	err = db.QueryRow(querySelect, Username).Scan(
		&HasilGetUser.ID,
		&HasilGetUser.Username,
		&HasilGetUser.Password,
		&HasilGetUser.Role,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return structs.UserObat{}, fmt.Errorf("Username %s tidak ditemukan", Username)
		}
		return HasilGetUser, err
	}

	return HasilGetUser, nil
}

func GetUser(db *sql.DB, IdUser int, NamaUser string) (HasilGetUserObat []structs.UserObat, err error) {

	querySelect := `
		SELECT 	id, nama, username, role, status, created_at, created_by,
				modified_at, modified_by
		FROM user_obat WHERE 1=1
	`

	var paramUser []interface{}

	jumlahParam := 1

	if IdUser > 0 {
		querySelect += fmt.Sprintf(" AND id = $%d", jumlahParam)
		paramUser = append(paramUser, IdUser)
		jumlahParam++
	}

	if NamaUser != "" {
		querySelect += fmt.Sprintf(" AND nama ILIKE $%d", jumlahParam)
		paramUser = append(paramUser, "%"+NamaUser+"%")
		jumlahParam++
	}

	rows, err := db.Query(querySelect, paramUser...)

	if err != nil {
		return HasilGetUserObat, err
	}

	defer rows.Close()

	for rows.Next() {
		var setUserObat structs.UserObat

		err = rows.Scan(&setUserObat.ID,
			&setUserObat.Nama, &setUserObat.Username, &setUserObat.Role, &setUserObat.Status,
			&setUserObat.CreatedAt, &setUserObat.CreatedBy, &setUserObat.ModifiedAt, &setUserObat.ModifiedBy,
		)

		if err != nil {
			return nil, err
		}

		HasilGetUserObat = append(HasilGetUserObat, setUserObat)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if HasilGetUserObat == nil {
		HasilGetUserObat = []structs.UserObat{}
	}

	return HasilGetUserObat, nil
}

func GetUserByID(db *sql.DB, IdUser int) (HasilGetUserObat structs.UserObat, err error) {
	querySelect := `
		SELECT 	id, nama, username, role, status, created_at, created_by,
				modified_at, modified_by
		FROM user_obat WHERE id = $1
	`

	err = db.QueryRow(querySelect, IdUser).Scan(
		&HasilGetUserObat.ID, &HasilGetUserObat.Nama, &HasilGetUserObat.Username, &HasilGetUserObat.Role,
		&HasilGetUserObat.Status, &HasilGetUserObat.CreatedAt, &HasilGetUserObat.CreatedBy,
		&HasilGetUserObat.ModifiedAt, &HasilGetUserObat.ModifiedBy,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return structs.UserObat{}, fmt.Errorf("User ID %d tidak ditemukan", IdUser)
		}
		return structs.UserObat{}, err
	}

	return HasilGetUserObat, nil
}

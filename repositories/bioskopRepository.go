package repositories

import (
	"database/sql"
	"tugas13/structs"
)

func GetAllBioskop(db *sql.DB) (bioskops []structs.Bioskop, err error) {
	sqlStatement := `
	SELECT * FROM tugas_13.bioskop
	`

	res, err := db.Query(sqlStatement)
	if err != nil {
		return
	}

	defer res.Close()
	for res.Next() {
		var bioskop structs.Bioskop

		err = res.Scan(&bioskop.ID, &bioskop.Nama, &bioskop.Lokasi, &bioskop.Rating)
		if err != nil {
			return
		}

		bioskops = append(bioskops, bioskop)
	}

	return
}

func GetBioskop(db *sql.DB, id int) (bioskop structs.Bioskop, err error) {
	sqlStatement := `
	SELECT * FROM tugas_13.bioskop
	WHERE id = $1
	`

	res := db.QueryRow(sqlStatement, id)
	err = res.Err()
	if err != nil {
		return
	}

	err = res.Scan(&bioskop.ID, &bioskop.Nama, &bioskop.Lokasi, &bioskop.Rating)
	if err != nil {
		return
	}

	return
}

func CreateBioskop(db *sql.DB, bioskop structs.Bioskop) error {
	sqlStatement := `
	INSERT INTO tugas_13.bioskop(nama, lokasi, rating)
	VALUES ($1, $2, $3)
	`

	errs := db.QueryRow(sqlStatement, bioskop.Nama, bioskop.Lokasi, bioskop.Rating)
	return errs.Err()
}

func UpdateBioskop(db *sql.DB, bioskop structs.Bioskop) error {
	sqlStatement := `
	UPDATE tugas_13.bioskop
	SET nama = $2, lokasi = $3, rating = $4
	WHERE id = $1
	`

	errs := db.QueryRow(sqlStatement, bioskop.ID, bioskop.Nama, bioskop.Lokasi, bioskop.Rating)
	return errs.Err()
}

func DeleteBioskop(db *sql.DB, id int) error {
	sqlStatement := `
	DELETE FROM tugas_13.bioskop
	WHERE id = $1
	`

	errs := db.QueryRow(sqlStatement, id)
	return errs.Err()
}

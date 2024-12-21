package repository

import (
	"database/sql"
	"quiz-golang/structs"
)

func GetAllKategori(db *sql.DB) (result []structs.Kategori, err error) {
	sql := "SELECT * FROM kategori"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var kategori structs.Kategori
		err = rows.Scan(&kategori.ID, &kategori.Name, &kategori.Created_at, &kategori.Created_by, &kategori.Modified_at, &kategori.Modified_by)
		if err != nil {
			panic(err)
		}
		result = append(result, kategori)
	}
	return
}

func GetKategoriById(db *sql.DB, kategori structs.Kategori) (result structs.Kategori, err error) {
	sql := "SELECT * FROM kategori WHERE id = $1"
	rows, err := db.Query(sql, kategori.ID)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	rows.Next()
	var data structs.Kategori
	err = rows.Scan(
		&data.ID,
		&data.Name,
		&data.Created_at,
		&data.Created_by,
		&data.Modified_at,
		&data.Modified_by,
	)
	if err != nil {
		return
	}
	result = data
	return
}

func GetBukuByKategoriId(db *sql.DB, kategori structs.Kategori) (result []structs.Buku, err error) {
	sql := "SELECT * FROM buku WHERE category_id = $1"
	rows, err := db.Query(sql, kategori.ID)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var buku structs.Buku
		err = rows.Scan(
			&buku.ID,
			&buku.Title,
			&buku.Description,
			&buku.Image_url,
			&buku.Release_year,
			&buku.Price,
			&buku.Total_page,
			&buku.Thickness,
			&buku.Category_id,
			&buku.Created_at,
			&buku.Created_by,
			&buku.Modified_at,
			&buku.Modified_by,
		)
		if err != nil {
			return
		}
		result = append(result, buku)
	}
	return
}

func AddNewKategori(db *sql.DB, kategori structs.Kategori) (err error) {
	sql := "INSERT INTO kategori(name,created_by,modified_by) VALUES ($1, $2, $3)"
	errs := db.QueryRow(sql, kategori.Name, kategori.Created_by, kategori.Modified_by)
	return errs.Err()
}

func DeleteKategori(db *sql.DB, kategori structs.Kategori) (err error) {
	sql := "DELETE FROM kategori WHERE id = $1"
	errs := db.QueryRow(sql, kategori.ID)
	return errs.Err()
}

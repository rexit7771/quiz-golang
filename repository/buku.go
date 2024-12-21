package repository

import (
	"database/sql"
	"quiz-golang/structs"
)

func GetAllBuku(db *sql.DB) (result []structs.Buku, err error) {
	sql := "SELECT * FROM buku"
	rows, err := db.Query(sql)
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
			panic(err)
		}
		result = append(result, buku)
	}
	return
}

func GetBukuById(db *sql.DB, buku structs.Buku) (result structs.Buku, err error) {
	sql := "SELECT * FROM buku WHERE id = $1"
	row, err := db.Query(sql, buku.ID)
	if err != nil {
		panic(err)
	}
	defer row.Close()
	row.Next()
	var data structs.Buku
	err = row.Scan(
		&data.ID,
		&data.Title,
		&data.Description,
		&data.Image_url,
		&data.Release_year,
		&data.Price,
		&data.Total_page,
		&data.Thickness,
		&data.Category_id,
		&data.Created_at,
		&data.Created_by,
		&data.Modified_at,
		&data.Modified_by,
	)
	if err != nil {
		panic(err)
	}
	result = data
	return
}

func AddNewBuku(db *sql.DB, buku structs.Buku) (err error) {
	sql := "INSERT INTO buku(title,description, image_url, release_year, price, total_page, thickness, category_id,  created_by,  modified_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"
	if buku.Total_page > 100 {
		buku.Thickness = "tebal"
	} else {
		buku.Thickness = "tipis"
	}
	errs := db.QueryRow(sql, buku.Title, buku.Description, buku.Image_url, buku.Release_year, buku.Price, buku.Total_page, buku.Thickness, buku.Category_id, buku.Created_by, buku.Modified_by)
	return errs.Err()
}

func DeleteBuku(db *sql.DB, buku structs.Buku) (err error) {
	sql := "DELETE FROM buku WHERE id = $1"
	errs := db.QueryRow(sql, buku.ID)
	return errs.Err()
}

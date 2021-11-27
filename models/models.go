package models

import (
	"database/sql"
	"fmt"
	"go-postgres/config"
	"log"

	_ "github.com/lib/pq" // postgres golang driver
)

// books scheme from books table
// we try if data is null
// if data return is any of null, let using nullstring

type Buku struct {
	ID            int64  `json:"id"`
	Judul_Buku    string `json:"judul_buku"`
	Penulis       string `json:"penulis"`
	Tgl_publikasi string `json:"tgl_publikasi"`
}

// Func addbook with int64
func AddBook(buku Buku) int64 {

	// connect to postgres db
	db := config.CreateConnection()

	// close connection in end process
	defer db.Close()

	// make query insert
	// return id value will return id from book of insert to db
	sqlStatement := `INSERT INTO buku (judul_buku, penulis, tgl_publikasi) VALUES ($1, $2, $3) RETURNING id`

	// insert id will store in this id
	var id int64

	// func scan will store id insert in id
	err := db.QueryRow(sqlStatement, buku.Judul_Buku, buku.Penulis, buku.Tgl_publikasi).Scan(&id)

	if err != nil {
		log.Fatalf("Tidak bisa mengakses query. %v", err)
	}

	fmt.Printf("Insert data single record %v", id)

	// return insert id
	return id
}

// Func get all book
func GetAllBooks() ([]Buku, error) {

	// connect to postgres db
	db := config.CreateConnection()

	// close connection in end process
	defer db.Close()

	var bukus []Buku

	// make query select
	sqlStatement := `SELECT * FROM buku`

	// execute query sql
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Tidak bisa mengakses query. %v", err)
	}

	// close execute query process
	defer rows.Close()

	// iterate to get the data
	for rows.Next() {
		var buku Buku

		// get data and unmarshal to struct
		err = rows.Scan(&buku.ID, &buku.Judul_Buku, &buku.Penulis, &buku.Tgl_publikasi)

		if err != nil {
			log.Fatalf("Tidak bisa mengambil data. %v", err)
		}

		// insert it to bukus slice
		bukus = append(bukus, buku)
	}

	// return if error found in emtpy book
	return bukus, err
}

// Func to get one book
func GetOneBook(id int64) (Buku, error) {

	// connect to postgres db
	db := config.CreateConnection()

	// close connection in end process
	defer db.Close()

	var buku Buku

	// make query sql
	sqlStatement := `SELECT * FROM buku WHERE id=$1`

	// execute sql statement
	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&buku.ID, &buku.Judul_Buku, &buku.Penulis, &buku.Tgl_publikasi)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("Tidak ada data yang dicari")
		return buku, nil
	case nil:
		return buku, nil
	default:
		log.Fatalf("Tidak bisa mengambil data. %v", err)
	}
	return buku, err
}

// Update book in db
func UpdateBook(id int64, buku Buku) int64 {

	// connect to db progress
	db := config.CreateConnection()

	// close the connection in end process
	defer db.Close()

	// make query create sql
	sqlStatement := `UPDATE buku SET judul_buku=$2, penulis=$3, tgl_publikasi=$4 WHERE id=$1`

	// exec statment sql
	res, err := db.Exec(sqlStatement, id, buku.Judul_Buku, buku.Penulis, buku.Tgl_publikasi)

	if err != nil {
		log.Fatalf("Tidak bisa mengakses query. %v", err)
	}

	// check how many row/data was uploaded
	RowsAffected, err := res.RowsAffected()

	// check it
	if err != nil {
		log.Fatalf("Error ketika mengechek rows/data. %v", err)
	}
	fmt.Printf("Total rows/record yang diupdate %v\n", RowsAffected)
	return RowsAffected
}

// Func to delete book
func DeleteBook(id int64) int64 {

	// connect to posgres db
	db := config.CreateConnection()

	// close the connection in end process
	defer db.Close()

	// make sql query
	sqlStatement := `DELETE FROM buku WHERE id=$1`

	// exec statment sql
	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Tidak bisa mengakses query. %v", err)
	}

	// check how many row/data was deleted
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Tidak bisa mencari data. %v", err)
	}

	fmt.Printf("Total data yang terhapus %v", rowsAffected)
	return rowsAffected
}

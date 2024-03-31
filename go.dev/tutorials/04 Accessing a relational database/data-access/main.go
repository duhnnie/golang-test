package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Album struct {
	ID uint64
	Title string
	Artist string
	Price float32
}

func main() {
	// Capture connection properties.
	cfg := mysql.Config{
		User: os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net: "tcp",
		Addr: "127.0.0.1:3306",
		DBName: "recordings",
	}

	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()

	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected!")

	albums, err := albumsByArtist("John Coltrane")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Albums found: %v\n", albums)

	album, err := albumById(2)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Album found: %v\n", album)

	albumToAdd := Album{
		Title: "In Keeping Secrets Of Silent Earth: 3",
		Artist: "Coheed and Cambria",
		Price: 10.99,
	}

	id, err := addAlbum(albumToAdd)

	if err != nil {
		log.Fatal(err)
	}

	albumToAdd.ID = id

	fmt.Printf("Album %+v added", albumToAdd)
}

// albumsByArtist queries for albums that have the specified artist name.
func albumsByArtist(name string) ([]Album, error) {
	// An albums slice to hold data from returned rows.
	var albums []Album

	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)

	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}

	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var album Album

		if err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}

		albums = append(albums, album)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}

	return albums, nil
}

func albumById(id uint64) (Album, error) {
	var album Album

	row := db.QueryRow("SELECT * FROM album where id = ?", id)

	if err := row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
		if err == sql.ErrNoRows {
			return album, fmt.Errorf("albumById(%v) = no such album", id)
		}

		return album, fmt.Errorf("albumById(%v) = %q", id, err)
	}

	return album, nil
}

func addAlbum(album Album) (uint64, error) {
	result, err :=  db.Exec("INSERT into album (title, artist, price) VALUES (?, ? ,?)", album.Title, album.Artist, album.Price)

	if err != nil {
		return 0, fmt.Errorf("addAlbum(%v): %v", album, err)
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("addAlbum(%v): %v", album, err)
	}

	return uint64(id), nil

}

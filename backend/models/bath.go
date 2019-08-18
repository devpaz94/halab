package models

import (
	"fmt"

	dto "github.com/halab/backend/dto"

	_ "github.com/lib/pq"
)

func ReadBaths() (*[]dto.Bath, error) {
	var b dto.Bath
	var baths []dto.Bath
	rows, err := db.Query(`SELECT bath_id, name, url FROM baths`)
	if err != nil {
		fmt.Println("error querying")
	}
	for rows.Next() {
		rows.Scan(&b.ID, &b.Name, &b.URL)
		baths = append(baths, b)
	}
	return &baths, nil
}

func CreateBath(b *dto.Bath) error {
	sqlStatement := `
	INSERT INTO baths (bath_id, name, size, url)
	VALUES ($1, $2, $3, $4)
	RETURNING bath_id`
	id := 0
	err := db.QueryRow(sqlStatement, b.ID, b.Name, b.Size, b.URL).Scan(&id)
	if err != nil {
		return err
	}
	return nil
}

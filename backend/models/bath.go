package models

import (
	dto "github.com/halab/backend/dto"

	_ "github.com/lib/pq"
)

func ReadBath(id int) (*dto.Bath, error) {
	var b dto.Bath
	row, err := db.Query(`SELECT id, name, url FROM baths WHERE id = $1`, id)
	if row.Scan(&b.ID, &b.Name, &b.URL); err != nil {
		return nil, err
	}
	return &b, nil
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

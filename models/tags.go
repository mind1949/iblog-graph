package models

import (
	"time"
	"database/sql"
)

type Tag struct {
	ID        int
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t *Tag) Find() error {
	statement := `select id, title, createdAt, updatedAt from tags  where id = %1`
	err := sqldb.QueryRow(statement, t.ID).Scan(&t.ID, &t.Title, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (t *Tag) FindOrCreateByTitle() error {
	statement := `select id from tags where title = $1`
	err := sqldb.QueryRow(statement, t.Title).Scan(&t.ID)
	if err == sql.ErrNoRows {
		statement = `insert into tags(title, createdAt, updatedAt) values($1, $2) returning id, title, createdAt`
		err = sqldb.QueryRow(statement, t.Title, time.Now(), time.Now()).Scan(&t.ID, &t.Title, &t.CreatedAt)
		if err != nil {
			return err
		}
	}
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	return nil
}

func Tags() ([]*Tag, error) {
	statement := `select id, title, createdAt, updatedAt from tags order by id`
	rows, err := sqldb.Query(statement)
	if err !=nil {
		return nil, err
	}
	defer rows.Close()

	tags := []*Tag{}
	if rows.Next() {
		tag := &Tag{}
		err := rows.Scan(&tag.ID, &tag.Title, &tag.CreatedAt, &tag.UpdatedAt)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}
	return tags, nil
}

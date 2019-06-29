package models

import (
	"time"
	"github.com/lib/pq"
	"fmt"
)

type Post struct {
	ID        int
	Title     string
	Content   string
	TagsID    []int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p *Post) Save() error {
	statement := `insert into posts(title, content, tagsid, createdAt) values($1, $2, $3, $4) returning id`
	row := sqldb.QueryRow(statement, p.Title, p.Content, pq.Array(p.TagsID), p.CreatedAt)
	err := row.Scan(&(p.ID))
	if err != nil {
		return err
	}
	return nil
}

func (p *Post) Find() error {
	statement := "select id, title, content, createdAt, updatedAt from posts where id = $1"
	row := sqldb.QueryRow(statement, p.ID)
	err := row.Scan(&p.ID, &p.Title, &p.Content, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func Posts() ([]*Post, error) {
	statement := `select id, title, content, createdAt, updatedAt from posts order by id`
	rows, err := sqldb.Query(statement)
	defer rows.Close()

	posts := []*Post{}
	for rows.Next() {
		post := &Post{}
		err = rows.Scan(&post.ID, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			return nil, err
		}
		//debug
		fmt.Printf("%#v", *post)
		posts = append(posts, post)
	}
	return posts, nil
}

func (p *Post) Tags() ([]*Tag, error) {
	tags := make([]*Tag, 0, len(p.TagsID))
	statement := `select id, title, createdAt, updatedAt from tags where id = $1`
	for i, tagid := range p.TagsID {
		tag := &Tag{}
		err := sqldb.QueryRow(statement, tagid).Scan(&tag.ID, &tag.Title, &tag.CreatedAt, &tag.UpdatedAt)
		if err != nil {
			return nil, err
		}
		tags[i] = tag
	}
	return tags, nil
}

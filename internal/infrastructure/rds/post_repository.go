package rds

import (
	"database/sql"
	"github.com/keitam913/airlog/internal/domain"
	"time"
)

type PostRepository struct {
	DB *sql.DB
}

func (repo *PostRepository) Filename(filename string) (*domain.Post, error) {
	rs, err := repo.DB.Query("select filename, timestamp, hash, title, body from post where filename = ?", filename)
	if err != nil {
		return nil, err
	}
	defer rs.Close()
	if !rs.Next() {
		return nil, domain.ErrNotFound{}
	}
	post := &domain.Post{}
	var tss string
	if err := rs.Scan(
		&post.Filename,
		&tss,
		&post.Hash,
		&post.Title,
		&post.Body,
	); err != nil {
		return nil, err
	}
	tst, err := time.Parse("2006-01-02 15:04:05", tss)
	if err != nil {
		return nil, err
	}
	post.Timestamp = tst

	rs, err = repo.DB.Query("select label from post_label where filename = ?", filename)
	if err != nil {
		return nil, err
	}
	defer rs.Close()
	for rs.Next() {
		var label string
		if err := rs.Scan(&label); err != nil {
			return nil, err
		}
		post.Labels = append(post.Labels, label)
	}

	return post, nil
}

func (repo *PostRepository) All() ([]*domain.Post, error) {
	labels := map[string][]string{}
	rs, err := repo.DB.Query(`select filename, label from post_label`)
	if err != nil {
		return nil, err
	}
	defer rs.Close()
	for rs.Next() {
		var filename, label string
		if err := rs.Scan(&filename, &label); err != nil {
			return nil, err
		}
		labels[filename] = append(labels[filename], label)
	}
	rs, err = repo.DB.Query("select filename, timestamp, hash, title, body from post order by timestamp desc")
	if err != nil {
		return nil, err
	}
	var posts []*domain.Post
	for rs.Next() {
		post := &domain.Post{}
		var tss string
		if err := rs.Scan(
			&post.Filename,
			&tss,
			&post.Hash,
			&post.Title,
			&post.Body,
		); err != nil {
			return nil, err
		}
		tst, err := time.Parse("2006-01-02 15:04:05", tss)
		if err != nil {
			return nil, err
		}
		post.Timestamp = tst
		post.Labels = labels[post.Filename]
		posts = append(posts, post)
	}
	return posts, nil
}

func (repo *PostRepository) Put(post *domain.Post) error {
	if _, err := repo.DB.Exec(
		`insert into post (filename, timestamp, hash, title, body) values (?, ?, ?, ?, ?) 
on duplicate key update title = ?, body = ?`,
		post.Filename,
		post.Timestamp.Format("2006-01-02 15:04:05"),
		"",
		post.Title,
		post.Body,
		post.Title,
		post.Body,
	); err != nil {
		return err
	}
	if _, err := repo.DB.Exec(`delete from post_label where filename = ?`, post.Filename); err != nil {
		return err
	}
	for _, label := range post.Labels {
		if _, err := repo.DB.Exec(`insert into post_label (filename, label) values (?, ?)`, post.Filename, label);
			err != nil {
			return err
		}
	}
	return nil
}

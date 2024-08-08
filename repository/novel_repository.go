package repository

import (
	"database/sql"
	"go-novel-api/model"
	"go-novel-api/utils/constant"
)

type NovelRepository interface {
	BaseRepository[model.Novel]
}

// Struct ini memiliki satu field db yang merupakan pointer ke objek sql.DB. Objek ini digunakan untuk berinteraksi dengan database SQL.
type novelRepository struct {
	db *sql.DB
}

func (e *novelRepository) Create(payload model.Novel) error {
	_, err := e.db.Exec(constant.NOVEL_INSERT, payload.Id, payload.Judul, payload.Penerbit, payload.TahunTerbit, payload.Penulis)
	if err != nil {
		return err
	}
	return nil
}

func NewNovelRepository(db *sql.DB) NovelRepository {
	return &novelRepository{
		db: db,
	}
}

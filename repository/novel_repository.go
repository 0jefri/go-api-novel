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

func (e *novelRepository) List() ([]model.Novel, error) {
	rows, err := e.db.Query(constant.LIST_NOVEL)
	if err != nil {
		return nil, err
	}

	var novels []model.Novel
	for rows.Next() {
		var novel model.Novel
		err = rows.Scan(&novel.Id, &novel.Judul, &novel.Penerbit, &novel.TahunTerbit, &novel.Penulis)
		if err != nil {
			return nil, err
		}
		novels = append(novels, novel)
	}
	return novels, nil
}

func NewNovelRepository(db *sql.DB) NovelRepository {
	return &novelRepository{
		db: db,
	}
}

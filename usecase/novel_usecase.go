package usecase

import (
	"fmt"
	"go-novel-api/model"
	"go-novel-api/repository"
)

type NovelUsecase interface {
	RegisterNewNovel(payload model.Novel) error
	FindAllNovels() ([]model.Novel, error)
	GetById(id string) (model.Novel, error)
}

type novelUsecase struct {
	repo repository.NovelRepository
}

func (e *novelUsecase) RegisterNewNovel(payload model.Novel) error {
	if payload.Judul == "" || payload.Penerbit == "" || payload.TahunTerbit == "" || payload.Penulis == "" {
		return fmt.Errorf("Judul, Penerbit, Tahun Terbit, Penulis is required")
	}
	err := e.repo.Create(payload)
	if err != nil {
		return fmt.Errorf("Failed to create novel: %s", err.Error())
	}
	return nil
}

func (e *novelUsecase) FindAllNovels() ([]model.Novel, error) {
	novels, err := e.repo.List()
	if err != nil {
		return nil, err
	}
	return novels, nil
}

func (e *novelUsecase) GetById(id string) (model.Novel, error) {
	return e.repo.Get(id)
}

func NewNovelUsecase(nvlRepo repository.NovelRepository) NovelUsecase {
	return &novelUsecase{repo: nvlRepo}
}

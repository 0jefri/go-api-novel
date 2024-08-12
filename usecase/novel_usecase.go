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
	UpdateNovel(payload model.Novel) error
	DeleteNovel(id string) error
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

func (e *novelUsecase) UpdateNovel(payload model.Novel) error {
	_, err := e.GetById(payload.Id)
	if err != nil {
		return err
	}
	return e.repo.Update(payload)
}

func (e *novelUsecase) DeleteNovel(id string) error {
	_, err := e.GetById(id)
	if err != nil {
		return err
	}
	return e.repo.Delete(id)
}

func NewNovelUsecase(nvlRepo repository.NovelRepository) NovelUsecase {
	return &novelUsecase{repo: nvlRepo}
}

package repository

import "go.mod/model"

type MangaRepositoryInterface interface {
	InsertManga(mange model.PostManga)
	GetAllManga() []model.Manga
	GetOneManga(uint) model.Manga
	UpdateManga(uint, model.PostManga) model.Manga
	DeleteManga(uint) bool
	//
}

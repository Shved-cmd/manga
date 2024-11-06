package repository

import (
	"database/sql"
	"go.mod/model"
	"log"
)

type MangaRepository struct {
	Db *sql.DB
}

func (m *MangaRepository) InsertManga(mange model.PostManga) {
	//TODO implement me
	panic("implement me")
}

func (m *MangaRepository) GetAllManga() []model.Manga {
	query, err := m.Db.Query("select * from manga")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	var mangas []model.Manga
	if query != nil {
		for query.Next() {
			var (
				id       uint
				title    string
				genre    string
				volumes  uint8
				chapters uint16
				author   string
			)
			err := query.Scan(&id, &title, &genre, &volumes, &chapters, &author)
			if err != nil {
				log.Println(err)
			}
			manga := model.Manga{Id: id, Title: title, Genre: genre, Volumes: volumes, Chapters: chapters, Author: author}
			mangas = append(mangas, manga)
		}
		return mangas
	}

}

func (m *MangaRepository) GetOneManga(u uint) model.Manga {
	//TODO implement me
	panic("implement me")
}

func (m *MangaRepository) UpdateManga(u uint, manga model.PostManga) model.Manga {
	//TODO implement me
	panic("implement me")
}

func (m *MangaRepository) DeleteManga(u uint) bool {
	//TODO implement me
	panic("implement me")
}

func NewMangaRepository(db *sql.DB) MangaRepositoryInterface {
	return &MangaRepository{Db: db}
}

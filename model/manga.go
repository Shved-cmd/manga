package model

// Модель Manga будет хранить результат ответа из таблицы manga
type Manga struct {
	Id       uint   `json:"id"`
	Title    string `json:"title"`
	Genre    string `json:"genres"`
	Volumes  uint8  `json:"volumes"`
	Chapters uint16 `json:"chapters"`
	Author   string `json:"author"`
	//
}

//  PostManga будет хранить данные запроса при вставке или обновлении данных.

type PostManga struct {
	Title    string `json:"title"`
	Genre    string `json:"genres"`
	Volumes  uint8  `json:"volumes"`
	Chapters uint16 `json:"chapters"`
	Author   string `json:"author"`
}

// MangaUri будет хранить запрос uri id при получении одних данных, обновлении данных и удалении данных
type MangaUri struct {
	ID uint `uri:"id" binding:"required, number"`
	//
}

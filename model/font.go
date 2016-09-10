package model

import (
	"github.com/gocraft/dbr"
	"time"
)

type Font struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	DownloadUrl string    `json:"download_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewFont() *Font {
	return &Font{}
}

type Fonts []Font

func (m *Fonts) Load(tx *dbr.Tx) error {
	return tx.Select("*").
		From("fonts").
		OrderBy("id").
		LoadStruct(m)
}

func (m *Font) LoadColumnById(tx *dbr.Tx, id int, col string) *dbr.SelectBuilder {
	return tx.Select(col).
		From("fonts").
		Where("id = ?", id)
}

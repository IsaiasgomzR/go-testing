package albums

import (
	"errors"
)

type Album struct{
	ID int `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Year string `json:"year"`
}

type Storage interface{
	GetAlbums () (al *[]Album, err error)
	Create (a *Album) (err error)
}

var (
	ErrStorageInternal = errors.New("internal storage error")
	ErrStorageAtrbiutenotUnique = errors.New("atrbute not unique")
)

type Service interface {
	GetAlbums () (al *[]Album, err error)
	Create (a *Album) (err error)
}

var (
	ErrServiceInternal = errors.New("internal service error")
	ErrServiceAtributeNotUnique = errors.New("atribute not unique")
	ErrServiceInvalidAlbum = errors.New("album invalid")
)
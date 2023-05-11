package movies

import (
	"errors"
)

type Movie struct{
	ID int `json:"id"`
	Title string `json:"title"`
	RelaseYear string `json:"relase_year"`
}

type Repository interface{
	GetMovie(id int) (mv Movie, err error)
}

var (
	ErrRepointernal = errors.New("internal error")
	ErrRepoNotFound = errors.New("movie not found")
)

type Service interface{
	GetMovie(id int) (mv Movie, err error)
}

var (
	ErrServiceInternal = errors.New("internal error")
	ErrServiceNotFound = errors.New("movie not found")
)
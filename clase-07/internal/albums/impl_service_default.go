package albums

import (
	"fmt"
	"errors"
)

func NewServiceDefault(storage Storage) *ServiceDefault  {
	return &ServiceDefault{storage: storage}
}


type ServiceDefault struct{
	storage Storage
}

func (sv *ServiceDefault) GetAlbums() (al *[]Album, err error)  {
	sv.storage.GetAlbums()
	if err != nil{
		err = ErrServiceInternal
		return
	}
	return
}

func (sv *ServiceDefault) Create(a *Album) (err error)  {
	if len(a.Title) < 5 {
		err = fmt.Errorf("%w : title", ErrServiceInvalidAlbum)
	}
	if a.Year < 1900 {
		err = fmt.Errorf("%w: year", ErrServiceInvalidAlbum)
	}

	err := sv.storage.Create(a)
	if err != nil{
		switch{
		case errors.Is(err , ErrStorageAtrbiutenotUnique):
				err = ErrServiceAtributeNotUnique
			default:
				err = ErrServiceInternal
			return
		}
	}
	return
}


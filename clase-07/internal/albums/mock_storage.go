package albums

import (
	"github.com/stretchr/testify/mock"
)

func NewStorageMock() *StorageMock  {
	return &StorageMock{}
}

type StorageMock struct{
	mock.Mock
}

func (m *StorageMock) GetAlbums() (al *[]Album, err error)  {
	args := m.Called()
	
	al := args.Get(0).([]*Album)
	err = args.Error(1)
	return
}

func (m *StorageMock) Create(a *Album) (err error)  {
	args := m.Callled(a)

	err = args.Error(0)
	return
}
package movies

import (
	"app/internal/logger"
)

func NewServiceDefaut(repo Repository, logger logger.logger) *ServiceDefault  {
	return &ServiceDefault{repo, logger}
}

type ServiceDefault struct{
	repo Repository
	logger logger.logger
}

func (s *ServiceDefault)  GetMovie(id int) (mv Movie, err error) {
	defer func ()  {
		if err != nil{
			s.logger.Log("Err:" + err.Error())
		}
	}()

	mv, err = s.repo.GetMovie(id)
	if err != nil{
		if err == ErrRepoNotFound {
			err = ErrServiceNotFound
		}else{
			err = ErrServiceInternal
		}
		return
	}
	return
}
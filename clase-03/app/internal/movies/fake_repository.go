package movies

func NewRepositoryFake(db []Movie) *RepositoryFake {
	return &RepositoryFake{db: db} 
}

type RepositoryFake struct{
	db []Movie
	GetMovieSpy bool
}

func (rp *RepositoryFake) GetMovie(id int) (mv Movie, err error)  {
	rp.GetMovieSpy = true
	for _, m := range rp.db {
		if m.id == id {
			mv = m
			return
		}
	}
	err = ErrRepoNotFound
	return
}
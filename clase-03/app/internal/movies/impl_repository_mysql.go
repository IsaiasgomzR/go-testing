package movies

import (
	"database/sql"
)

const(
	queryGetmovie = `SELECT id, title, relase_year FROM movies WHERE id = ?`
)

func MyNewRepositorySQL(db *sql.DB) *RepositoryMySQL {
	return &RepositoryMySQL{db}
}

type RepositoryMySQL struct{
	db *sql.DB
}

func (rp *RepositoryMySQL) GetMovie(id int) (mv Movie, err error)  {
	var smt sql.Stmt
	smt, err := rp.db.Prepare(queryGetmovie)
	if err != nil{
		err = ErrRepointernal
		return
	}

	err =smt.QueryRow(id).Scan(&mv.ID, &mv.Title, &mv.RelaseYear)
	if err != nil{
		if err == sql.ErrNoRows{
			err= ErrRepoNotFound
		}else{
			err = ErrRepointernal
		}
		return
	}
	return
}
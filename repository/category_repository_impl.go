package repository

import (
	"belajar-go-api/model/domain"
	"context"
	"database/sql"
)

type CategoryRepositoryImpl struct {
	 
}


func (c CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "insert into category(name) values (?)"
	result, err := tx.ExecContext(ctx, SQL, category.Name)
	if err != nil {
		panic(err)
	}

	result.LastInsertId()
	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	category.Id = int(id)
	return category
}

func (c CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	panic("unimplemented")
}

func (c CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	panic("unimplemented")
}

func (c CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) domain.Category {
	panic("unimplemented")
}

func (c CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, category domain.Category) []domain.Category {
	panic("unimplemented")
}
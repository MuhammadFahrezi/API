package repository

import (
	"context"
	"database/sql"
	"project-workshop/go-api-ecom/helper"
	"project-workshop/go-api-ecom/model/domain"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "insert into category(category) values (?)"
	result, err := tx.ExecContext(ctx, SQL, category.Category)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.Id = int(id)
	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "update category set category = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Category, category.Id)
	helper.PanicIfError(err)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	SQL := "delete from category where id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Id)
	helper.PanicIfError(err)
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	SQL := `
		SELECT c.id, c.category, p.id AS product_id, p.name AS product_name
		FROM category c
		LEFT JOIN product p ON c.id = p.category_id
		WHERE c.id = ?
	`

	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicIfError(err)

	var category domain.Category
	category.Products = make([]domain.Product, 0)
	for rows.Next() {
		product := domain.Product{}
		err := rows.Scan(&category.Id, &category.Category, &product.Id, &product.Name)
		helper.PanicIfError(err)
		category.Products = append(category.Products, product)
	}
	return category, nil
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "select id, category from category"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Category)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}
	return categories
}

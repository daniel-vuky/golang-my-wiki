package repository

import (
	"database/sql"
	"github.com/daniel-vuky/golang-my-wiki/model"
)

type CategoryRepository struct {
	Db *sql.DB
}

func (repository *CategoryRepository) GetListCategory(userId uint64) ([]model.Category, error) {
	exec := "SELECT * FROM category WHERE user_id = ? ORDER BY name ASC, updated_at DESC"
	categories, queryErr := repository.Db.Query(exec, userId)
	if queryErr != nil {
		return nil, queryErr
	}
	defer categories.Close()

	listCategory := []model.Category{}
	for categories.Next() {
		var item model.Category
		scanErr := categories.Scan(
			&item.CategoryId,
			&item.Name,
			&item.UserId,
			&item.ShortDescription,
			&item.CreatedAt,
			&item.UpdatedAt,
		)
		if scanErr != nil {
			return nil, scanErr
		}
		listCategory = append(listCategory, item)
	}

	return listCategory, nil
}

func (repository *CategoryRepository) GetCategoryById(userId, categoryId uint64) (model.Category, error) {
	exec := "SELECT * FROM category where user_id = ? and category_id = ?"
	var category model.Category
	queryErr := repository.Db.QueryRow(exec, userId, categoryId).Scan(
		&category.CategoryId,
		&category.Name,
		&category.UserId,
		&category.ShortDescription,
		&category.CreatedAt,
		&category.UpdatedAt,
	)
	if queryErr != nil {
		return category, queryErr
	}
	return category, nil
}

func (repository *CategoryRepository) CreateCategory(category *model.Category) error {
	exec := "INSERT INTO category (user_id, name, short_description) VALUES (?, ?, ?)"
	execResult, execErr := repository.Db.Exec(
		exec,
		category.UserId,
		category.Name,
		category.ShortDescription,
	)
	if execErr != nil {
		return execErr
	}
	lastInsertId, _ := execResult.LastInsertId()
	category.CategoryId = uint64(lastInsertId)

	return nil
}

func (repository *CategoryRepository) UpdateCategory(category *model.Category) error {
	exec := "UPDATE category set name = ?, short_description = ? WHERE category_id = ?"
	_, execErr := repository.Db.Exec(
		exec,
		category.Name,
		category.ShortDescription,
		category.CategoryId,
	)

	return execErr
}

func (repository *CategoryRepository) DeleteCategory(category *model.Category) error {
	exec := "DELETE FROM category WHERE category_id = ?"
	_, execErr := repository.Db.Exec(
		exec,
		category.CategoryId,
	)

	return execErr
}

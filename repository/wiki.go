package repository

import (
	"database/sql"
	"github.com/daniel-vuky/golang-my-wiki/model"
)

type WikiRepository struct {
	Db *sql.DB
}

func (repository *WikiRepository) GetListWiki(userId, categoryId uint64) ([]model.Wiki, error) {
	exec := "SELECT * FROM wiki WHERE user_id = ? and category_id = ? ORDER BY title ASC, updated_at DESC"
	wikis, queryErr := repository.Db.Query(exec, userId, categoryId)
	if queryErr != nil {
		return nil, queryErr
	}
	defer wikis.Close()

	listWiki := []model.Wiki{}
	for wikis.Next() {
		var item model.Wiki
		scanErr := wikis.Scan(
			&item.WikiId,
			&item.Title,
			&item.CategoryId,
			&item.UserId,
			&item.Body,
			&item.CreatedAt,
			&item.UpdatedAt,
		)
		if scanErr != nil {
			return nil, scanErr
		}
		listWiki = append(listWiki, item)
	}

	return listWiki, nil
}

func (repository *WikiRepository) GetWikiById(userId, wikiId uint64) (model.Wiki, error) {
	exec := "SELECT * FROM wiki where user_id = ? and wiki_id = ?"
	var wiki model.Wiki
	queryErr := repository.Db.QueryRow(exec, userId, wikiId).Scan(
		&wiki.WikiId,
		&wiki.Title,
		&wiki.CategoryId,
		&wiki.UserId,
		&wiki.Body,
		&wiki.CreatedAt,
		&wiki.UpdatedAt,
	)
	if queryErr != nil {
		return wiki, queryErr
	}
	return wiki, nil
}

func (repository *WikiRepository) CreateWiki(wiki *model.Wiki) error {
	exec := "INSERT INTO wiki (user_id, category_id, title, body) VALUES (?, ?, ?, ?)"
	execResult, execErr := repository.Db.Exec(
		exec,
		wiki.UserId,
		wiki.CategoryId,
		wiki.Title,
		wiki.Body,
	)
	if execErr != nil {
		return execErr
	}
	lastInsertId, _ := execResult.LastInsertId()
	wiki.WikiId = uint64(lastInsertId)

	return nil
}

func (repository *WikiRepository) UpdateWiki(wiki *model.Wiki) error {
	exec := "UPDATE wiki set title = ?, body = ? WHERE wiki_id = ?"
	_, execErr := repository.Db.Exec(
		exec,
		wiki.Title,
		wiki.Body,
		wiki.WikiId,
	)

	return execErr
}

func (repository *WikiRepository) DeleteWiki(wiki *model.Wiki) error {
	exec := "DELETE FROM wiki WHERE wiki_id = ?"
	_, execErr := repository.Db.Exec(
		exec,
		wiki.WikiId,
	)

	return execErr
}

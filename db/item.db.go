package db

import (
	"database/sql"
	"errors"
)

type Item struct {
	Id   int    `param:"id"`
	Name string `query:"name" form:"name"`
	Type string `query:"type" param:"type"`
}

func GetItem(db *sql.DB, id int, t string) (Item, error) {
	var i Item
	var stmt string
	if t == "recipe" {
		stmt = `select id, name, 'recipe' as type 
              from recipe
              where id = $1`
	} else if t == "list" {
		stmt = `select id, name, 'list' as type 
              from list
              where id = $1`
	} else {
		return i, errors.New("type is not list or recipe")
	}
	row := db.QueryRow(stmt, id)

	err := row.Scan(&i.Id, &i.Name, &i.Type)
	if err != nil {
		return i, err
	}

	return i, nil
}

func UpdateItem(db *sql.DB, i Item) (Item, error) {
	var stmt string
	if i.Type == "recipe" {
		stmt = `update recipe set name = $1
              where id = $2`
	} else if i.Type == "list" {
		stmt = `update list set name = $1
              where id = $2`
	} else {
		return i, errors.New("type is not list or recipe")
	}

	_, err := db.Exec(stmt, i.Name, i.Id)
	if err != nil {
		return i, err
	}

	return i, nil
}

func DeleteItem(db *sql.DB, i Item) error {
	var stmt string
	if i.Type == "recipe" {
		stmt = `delete from link where child_id = $1 and child_type = $2`
	} else if i.Type == "list" {
		stmt = `delete from link where child_id = $1 and child_type = $2`
	} else {
		return errors.New("type is not list or recipe")
	}

	_, err := db.Exec(stmt, i.Id, i.Type)
	if err != nil {
		return err
	}

	return nil
}

type ItemSearchResults = []Item

func ItemSearch(db *sql.DB, search_term string) (ItemSearchResults, error) {
	var res ItemSearchResults

	rows, err := db.Query(`
        select id, name, 'recipe' as type
        from recipe
        where name ~* $1
        union all
        select id, name, 'list' as type
        from list
        where name ~* $1
    `, search_term)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		var i Item
		err := rows.Scan(&i.Id, &i.Name, &i.Type)
		if err != nil {
			return res, err
		}
		res = append(res, i)
	}
	return res, nil
}

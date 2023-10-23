package api

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

type Recipe struct {
	Id           int      `json:"id"`
	Name         string   `json:"name"`
	Ingredients  []string `json:"ingredients"`
	Instructions []string `json:"instructions"`
	Url          string   `form:"url"`
	ListId       int      `form:"list_id"`
	List         string   `form:"list"`
}

type List struct {
	Id       int
	Name     string
	Children []Item
}

type Item struct {
	Id   int    `param:"id"`
	Name string `query:"name" form:"name"`
	Type string `query:"type"`
}

func GetTopLevelLists(db *sql.DB) ([]Item, error) {
	var l []Item
	rows, err := db.Query(`select distinct parent_id, name, 'list'
                        from link 
                        left join list on id = parent_id
                        where parent_id not in (select child_id 
                                                from link 
                                                where child_type != 'recipe')`)
	if err != nil {
		return l, err
	}

	for rows.Next() {
		var i Item
		err = rows.Scan(&i.Id, &i.Name, &i.Type)
		if err != nil {
			return l, err
		}
		l = append(l, i)
	}

	return l, nil
}

func GetList(db *sql.DB, id int) (List, error) {
	var l List
	row := db.QueryRow(`SELECT id, name 
                      FROM list 
                      WHERE id = $1`, id)

	err := row.Scan(&l.Id, &l.Name)
	if err != nil {
		return l, err
	}

	rows, err := db.Query(`select coalesce(list.id, recipe.id) as id, 
                    coalesce(list.name, recipe.name) as name,
                    link.child_type
                    from link 
                    left join list on link.child_id = list.id and child_type = 'list' 
                    left join recipe on link.child_id = recipe.id and child_type = 'recipe' 
                    where parent_id = $1`, id)
	if err != nil {
		return l, err
	}

	for rows.Next() {
		var i Item
		err = rows.Scan(&i.Id, &i.Name, &i.Type)
		if err != nil {
			return l, err
		}
		l.Children = append(l.Children, i)
	}

	return l, nil
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

func PutItem(db *sql.DB, i Item) (Item, error) {
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

func GetRecipe(db *sql.DB, id int) (Recipe, error) {
	var r Recipe
	row := db.QueryRow(`select id, name, ingredients, instructions, 'recipe' as type
                      from recipe 
                      where id = $1`, id)
	err := row.Scan(&r.Id, &r.Name, pq.Array(&r.Ingredients), pq.Array(&r.Instructions))
	if err != nil {
		return r, err
	}
	return r, nil
}

func PostRecipe(db *sql.DB, r *Recipe) error {
	r.Ingredients = []string{"test", "test2"}
	r.Instructions = []string{"test", "test2"}
	r.Name = "test"

	row, err := db.Query(`insert into recipe
                        (name, url, ingredients, instructions)
                        values ($1, $2, $3, $4)
                        returning id`, r.Name, r.Url, pq.Array(r.Ingredients), pq.Array(r.Instructions))
	if err != nil {
		return err
	}

	row.Next()
	err = row.Scan(&r.Id)
	if err != nil {
		return err
	}

	_, err = db.Exec(`insert into link
                    values ($1, $2, 'recipe')`, r.ListId, r.Id)
	if err != nil {
		return err
	}

	return nil
}

func SearchList(db *sql.DB, list string) ([]Item, error) {
	var items []Item

	rows, err := db.Query(`select id, name
                        from list
                        where position($1 in name)>0`, list)
	if err != nil {
		return items, err
	}

	for rows.Next() {
		var i Item
		err = rows.Scan(&i.Id, &i.Name)
		if err != nil {
			return items, err
		}
		items = append(items, i)
	}

	return items, nil
}

package db

import (
	"database/sql"

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

func GetRecipe(db *sql.DB, id int) (Recipe, error) {
	var r Recipe
	row := db.QueryRow(`select id, name, ingredients, instructions
                      from recipe 
                      where id = $1`, id)
	err := row.Scan(&r.Id, &r.Name, pq.Array(&r.Ingredients), pq.Array(&r.Instructions))
	if err != nil {
		return r, err
	}
	return r, nil
}

func PostRecipe(db *sql.DB, r *Recipe) error {
	row := db.QueryRow(`insert into recipe
                        (name, url, ingredients, instructions)
                        values ($1, $2, $3, $4)
                        returning id`, r.Name, r.Url, pq.Array(r.Ingredients), pq.Array(r.Instructions))

    err := row.Scan(&r.Id)
	if err != nil {
		return err
	}

	_, err = db.Exec(`insert into link values($1, $2, 'recipe')`, r.ListId, r.Id)
	if err != nil {
		return err
	}

	return nil
}

func GetRecipeFromURL(db *sql.DB, url string) (Recipe, error) {
	var r Recipe
	row := db.QueryRow(`select 
                            recipe.id, 
                            recipe.name, 
                            ingredients, 
                            instructions, 
                            list.name as list
                      from recipe 
                      left join link on (recipe.id = child_id and link.child_type = 'recipe')
                      left join list on (parent_id = list.id)
                      where recipe.url = $1`, url)
	err := row.Scan(
		&r.Id,
		&r.Name,
		pq.Array(&r.Ingredients),
		pq.Array(&r.Instructions),
		&r.List,
	)
	if err != nil {
		return r, err
	}
	return r, nil
}

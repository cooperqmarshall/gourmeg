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
	ImageUrl     string
	ThumbnailUrl string
	Lists        []List
}

func GetRecipe(db *sql.DB, id int) (*Recipe, error) {
	r := new(Recipe)
	row := db.QueryRow(`select 
						 id,
						 name,
						 url,
						 ingredients,
						 instructions,
						 image_url,
						 thumbnail_url
                      from recipe
                      where id = $1`, id)
	err := row.Scan(
		&r.Id,
		&r.Name,
		&r.Url,
		pq.Array(&r.Ingredients),
		pq.Array(&r.Instructions),
		&r.ImageUrl,
		&r.ThumbnailUrl,
	)
	if err != nil {
		return r, err
	}

	rows, err := db.Query(`select id, name
                        from link 
						left join list on (link.parent_id = list.id)
						where child_type = 'recipe' and child_id=$1`, id)
	if err != nil {
		return r, err
	}
	defer rows.Close()

	for rows.Next() {
		var l List
		err = rows.Scan(&l.Id, &l.Name)
		if err != nil {
			return r, err
		}
		r.Lists = append(r.Lists, l)
	}

	return r, nil
}

func PostRecipe(db *sql.DB, r *Recipe) error {
	row := db.QueryRow(`insert into recipe
                        (name, url, ingredients, instructions, image_url, thumbnail_url)
                        values ($1, $2, $3, $4, $5, $6)
                        returning id`, r.Name, r.Url, pq.Array(r.Ingredients), pq.Array(r.Instructions), r.ImageUrl, r.ThumbnailUrl)

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

func GetRecipeFromURL(db *sql.DB, url string) (*Recipe, error) {
	r := new(Recipe)
	row := db.QueryRow(`select 
                            recipe.id, 
                            recipe.name, 
							url,
                            ingredients, 
                            instructions, 
                            list.name as list,
                            image_url,
                            thumbnail_url
                      from recipe 
                      left join link on (recipe.id = child_id and link.child_type = 'recipe')
                      left join list on (parent_id = list.id)
                      where recipe.url = $1`, url)
	err := row.Scan(
		&r.Id,
		&r.Name,
		&r.Url,
		pq.Array(&r.Ingredients),
		pq.Array(&r.Instructions),
		&r.List,
		&r.ImageUrl,
		&r.ThumbnailUrl,
	)
	if err != nil {
		return r, err
	}
	return r, nil
}

func UpdateRecipe(db *sql.DB, r *Recipe) error {
	_, err := db.Exec(`update recipe set 
                            name = $2,
                            ingredients = $3,
                            instructions = $4,
                            image_url = $5,
                            thumbnail_url = $6
                        where id = $1`,
		r.Id, r.Name, pq.Array(r.Ingredients), pq.Array(r.Instructions), r.ImageUrl, r.ThumbnailUrl)
	if err != nil {
		return err
	}

	return nil
}

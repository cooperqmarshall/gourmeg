package db

import (
	"database/sql"

	"github.com/lib/pq"
)

type List struct {
	Id       int
	Name     string
	Children []Item
}

func GetList(db *sql.DB, id int) (List, error) {
	var l List
	row := db.QueryRow(`select id, name 
                      from list 
                      where id = $1`, id)

	err := row.Scan(&l.Id, &l.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			// no matches found
		} else {
			return l, err
		}
	}

	rows, err := db.Query(`select 
                    coalesce(list.id, recipe.id) as id, 
                    coalesce(list.name, recipe.name) as name,
                    link.child_type,
                    coalesce(recipe.domain, ''),
                    coalesce(recipe.thumbnail_url, ''),
					(
						select array_agg(parent_id) 
						from link l 
						where 
							l.child_id = coalesce(list.id, recipe.id) 
							and l.child_type = link.child_type
					)
                    from link 
                    left join list on link.child_id = list.id and child_type = 'list' 
                    left join recipe on link.child_id = recipe.id and child_type = 'recipe' 
                    where parent_id = $1
                    order by coalesce(list.created_at, recipe.created_at) desc`, id)
	if err != nil {
		return l, err
	}
	defer rows.Close()

	for rows.Next() {
		var item Item
		var list_ids []int64
		err = rows.Scan(&item.Id,
			&item.Name,
			&item.Type,
			&item.Domain,
			&item.ThumbnailUrl,
			pq.Array(&list_ids))
		if err != nil {
			return l, err
		}
		for _, b := range list_ids {
			item.ListIds = append(item.ListIds, int(b))
		}
		l.Children = append(l.Children, item)
	}

	return l, nil
}

func SearchList(db *sql.DB, list string) ([]Item, error) {
	var items []Item

	rows, err := db.Query(`select id, name
                        from list
                        where position($1 in name)>0`, list)
	if err != nil {
		return items, err
	}
	defer rows.Close()

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

func DeleteList(db *sql.DB, id int) error {
	_, err := db.Exec(`delete from link 
                    where child_id = $1 
                    and child_type = 'list'`, id)
	if err != nil {
		return err
	}

	return nil
}

func PostList(db *sql.DB, name string, parent_id int) (Item, error) {
	var i Item
	i.Type = "list"
	row := db.QueryRow(`insert into list (name) values ($1) returning id, name`, name)

	err := row.Scan(&i.Id, &i.Name)
	if err != nil {
		return i, err
	}

	// TODO: add default for parent_id as root list for user
	_, err = db.Exec(`insert into link (parent_id, child_id, child_type) values ($1, $2, 'list')`, parent_id, i.Id)
	if err != nil {
		return i, err
	}

	return i, nil
}

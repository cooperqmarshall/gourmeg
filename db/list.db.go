package db

import (
	"database/sql"
)

type List struct {
	Id       int
	Name     string
	Children []Item
}

func GetTopLevelLists(db *sql.DB) ([]Item, error) {
	var l []Item
	rows, err := db.Query(`
        select id, name, 'list'
        from list
        where id not in (select child_id 
                            from link 
                            where child_type = 'list')
    `)

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

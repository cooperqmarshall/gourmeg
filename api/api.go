package api

type Recipe struct {
	Id           int
	Name         string
	Ingredients   []string
	Instructions []string
}

type Chapter struct {
	Id      int
	Name    string
	Recipes []Recipe
}

type Book struct {
	Id       int
	Name     string
	Chapters []Chapter
}

func GetBook(id int) Book {
  r := Recipe{
		1,
		"recipe",
		[]string{"egg", "bacon", "milk"},
		[]string{"make egg", "cook bacon", "pour milk"},
	}

  c := Chapter{1, "pasta", []Recipe{r}}

	return Book{1, "Italian", []Chapter{c}}
}

func GetChapter(id int) Chapter {
  r := Recipe{
		1,
		"recipe",
		[]string{"egg", "bacon", "milk"},
		[]string{"make egg", "cook bacon", "pour milk"},
	}

  return Chapter{1, "pasta", []Recipe{r}}

}

func GetRecipe(id int) Recipe {
  return Recipe{
		1,
		"recipe",
		[]string{"egg", "bacon", "milk"},
		[]string{"make egg", "cook bacon", "pour milk"},
	}
}

func PostRecipe(id int) Recipe {
	dish := Recipe{
		1,
		"recipe",
		[]string{"egg", "bacon", "milk"},
		[]string{"make egg", "cook bacon", "pour milk"},
	}

  return dish
}

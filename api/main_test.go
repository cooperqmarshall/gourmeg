package api

import (
	"bytes"
	"database/sql"
	"testing"
)

func TestPostRecipe(t *testing.T) {
	db, err := sql.Open("postgres", "host=localhost user=root password=secret dbname=gourmeg_2 sslmode=disable")
	defer db.Close()
	if err != nil {
		t.Fatalf("unable to open database connection: %b", err)
	}
	if err := db.Ping(); err != nil {
		t.Fatalf("unable to connect to database %b", err)
	}
	_ = &Handler{DB: db}

    // h.PostRecipe()
}

func TestRecipeHTMLExtract(t *testing.T) {
    b := []byte(`
<!DOCTYPE html>
<html lang="en-US">
<head>
<script type="application/ld+json">
{
  "@context": "https://schema.org",
  "@graph": [
    {
      "@type": "Recipe",
      "name": "Authentic New Orleans Style Gumbo",
      "recipeIngredient": [
        "1 heaping cup all-purpose flour",
        "2/3 cup oil ((vegetable or canola oil))",
        "1 bunch celery (, diced, leaves and all)",
        "1  green bell pepper (, diced)",
        "1 large yellow onion (, diced)",
        "1 bunch green onion (, finely chopped)",
        "1 bunch fresh chopped parsley (, finely chopped)",
        "2-3 cloves garlic",
        "1-2 Tablespoons cajun seasoning (*)",
        "6-8 cups Chicken broth (*)",
        "12 ounce package andouille sausages (, sliced into 'coins' (substitute Polska Kielbasa if you can't find a good Andouille))",
        "Meat from 1 Rotisserie Chicken*",
        "2 cups Shrimps (, pre cooked)",
        "cooked white rice (for serving)"
      ],
      "recipeInstructions": [
        {
          "@type": "HowToStep",
          "text": "Make the Roux*: In a large, heavy bottom stock pot combine flour and oil. Cook on medium-low heat, stirring constantly for 30-45 minutes. This part takes patience--when it&#x27;s finished it should be as dark as chocolate and have a soft, &quot;cookie dough&quot; like consistency. Be careful not to let it burn! Feel free to add a little more flour or oil as needed to reach this consistency.",
          "name": "Make the Roux*: In a large, heavy bottom stock pot combine flour and oil. Cook on medium-low heat, stirring constantly for 30-45 minutes. This part takes patience--when it&#x27;s finished it should be as dark as chocolate and have a soft, &quot;cookie dough&quot; like consistency. Be careful not to let it burn! Feel free to add a little more flour or oil as needed to reach this consistency.",
          "url": "https://tastesbetterfromscratch.com/authentic-new-orleans-style-gumbo/#wprm-recipe-10785-step-0-0"
        },
        {
          "@type": "HowToStep",
          "text": "Brown the sausage.  In a separate skillet on medium-high heat place the sausage slices in one layer in the pan. Brown them well on one side (2-3 minutes) and then use a fork to flip each over onto the other side to brown. Remove to a plate.",
          "name": "Brown the sausage.  In a separate skillet on medium-high heat place the sausage slices in one layer in the pan. Brown them well on one side (2-3 minutes) and then use a fork to flip each over onto the other side to brown. Remove to a plate.",
          "url": "https://tastesbetterfromscratch.com/authentic-new-orleans-style-gumbo/#wprm-recipe-10785-step-0-1"
        },
        {
          "@type": "HowToStep",
          "text": "Cook the vegetables in broth. Add 1/2 cup of the chicken broth to the hot skillet that had the sausage to deglaze the pan. Pour the broth and drippings into your large soup pot. ",
          "name": "Cook the vegetables in broth. Add 1/2 cup of the chicken broth to the hot skillet that had the sausage to deglaze the pan. Pour the broth and drippings into your large soup pot. ",
          "url": "https://tastesbetterfromscratch.com/authentic-new-orleans-style-gumbo/#wprm-recipe-10785-step-0-2"
        },
        {
          "@type": "HowToStep",
          "text": "Add remaining 5 1/2 cups of chicken broth. Add veggies, parsley, garlic and roux to the pot and stir well. ",
          "name": "Add remaining 5 1/2 cups of chicken broth. Add veggies, parsley, garlic and roux to the pot and stir well. ",
          "url": "https://tastesbetterfromscratch.com/authentic-new-orleans-style-gumbo/#wprm-recipe-10785-step-0-3"
        },
        {
          "@type": "HowToStep",
          "text": "Bring to a boil over medium heat and boil for 5-7 minutes, or until the vegetables are slightly tender. (Skim off any foam that may rise to the top of the pot.) Stir in cajun seasoning, to taste.",
          "name": "Bring to a boil over medium heat and boil for 5-7 minutes, or until the vegetables are slightly tender. (Skim off any foam that may rise to the top of the pot.) Stir in cajun seasoning, to taste.",
          "url": "https://tastesbetterfromscratch.com/authentic-new-orleans-style-gumbo/#wprm-recipe-10785-step-0-4"
        },
        {
          "@type": "HowToStep",
          "text": "Add meat. Add chicken, sausage, and shrimp.",
          "name": "Add meat. Add chicken, sausage, and shrimp.",
          "url": "https://tastesbetterfromscratch.com/authentic-new-orleans-style-gumbo/#wprm-recipe-10785-step-0-5"
        },
        {
          "@type": "HowToStep",
          "text": "Taste and serve. At this point taste it and add more seasonings to your liking--salt, pepper, chicken bullion paste, garlic, more Joe&#x27;s stuff or more chicken broth--until you reach the perfect flavor. Serve warm over rice.  (Tastes even better the next day!)",
          "name": "Taste and serve. At this point taste it and add more seasonings to your liking--salt, pepper, chicken bullion paste, garlic, more Joe&#x27;s stuff or more chicken broth--until you reach the perfect flavor. Serve warm over rice.  (Tastes even better the next day!)",
          "url": "https://tastesbetterfromscratch.com/authentic-new-orleans-style-gumbo/#wprm-recipe-10785-step-0-6"
        }
      ],
    }
  ]
}
</script>
</head>
</html>
    `)
    r := bytes.NewReader(b)
    ldjson, err := extract_recipe_ldjson(r)
	if err != nil {
        t.Fatalf("unable to extract ldjson: %b", err)
	}
    t.Logf("%s", ldjson)
}

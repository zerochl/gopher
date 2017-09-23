package gopher

import (
	"net/http"

	"github.com/jimmykuu/wtforms"
	"gopkg.in/mgo.v2/bson"
)

// URL: /admin/article_categories
// 列出所有的文章分类
func adminListArticleCategoriesHandler(handler *Handler) {
	var categories []SiteCategory
	c := handler.DB.C(ARTICLE_CATEGORIES)
	c.Find(nil).All(&categories)

	handler.renderTemplate("admin/article_categories.html", ADMIN, map[string]interface{}{"categories": categories})
}

// URL: /admin/article_category/new
// 新建文章分类
func adminNewArticleCategoryHandler(handler *Handler) {
	form := wtforms.NewForm(
		wtforms.NewTextField("name", "名称", "", wtforms.Required{}),
	)

	if handler.Request.Method == "POST" {
		if !form.Validate(handler.Request) {
			handler.renderTemplate("article_category/new.html", ADMIN, map[string]interface{}{"form": form})
			return
		}

		c := handler.DB.C(ARTICLE_CATEGORIES)
		var category ArticleCategory
		err := c.Find(bson.M{"name": form.Value("name")}).One(&category)

		if err == nil {
			form.AddError("name", "该名称已经有了")
			handler.renderTemplate("article_category/new.html", ADMIN, map[string]interface{}{"form": form})
			return
		}

		err = c.Insert(&ArticleCategory{
			Id_:  bson.NewObjectId(),
			Name: form.Value("name"),
		})

		if err != nil {
			panic(err)
		}

		http.Redirect(handler.ResponseWriter, handler.Request, "/admin/article_category/new", http.StatusFound)
	}

	handler.renderTemplate("article_category/new.html", ADMIN, map[string]interface{}{"form": form})
}

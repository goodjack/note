package controllers

import (
	"blog/pkg/database"
	"blog/pkg/logger"
	"blog/pkg/route"
	"blog/pkg/types"

	"database/sql"
	"fmt"
	"net/http"
	"text/template"
)

// ArticlesController 文章相关页面
type ArticlesController struct {
}

// Article  对应一条文章数据
type Article struct {
	Title, Content string
	ID             int64
}

// Show 文章详情页面
func (*ArticlesController) Show(w http.ResponseWriter, r *http.Request) {
	// 1. 获取 URL 参数
	id := route.GetRouteVariable("id", r)

	// 2. 读取对应的文章数据
	article, err := getArticleByID(id)

	// 3. 如果出现错误
	if err != nil {
		// 当 Scan() 发现没有返回数据的话，会返回 sql.ErrNoRows 类型的错误
		if err == sql.ErrNoRows {
			// 3.1 数据未找到
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 文章未找到")
		} else {
			// 3.2 数据库错误
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		}
	} else {
		// 4. 读取成功
		// fmt.Fprint(w, "读取成功，文章标题 —— "+article.Title)

		// 4. 读取成功，显示文章
		// Funcs() 方法的传参是 template.FuncMap 类型的 Map 对象。键为模板里调用的函数名称，值为当前上下文的函数名称
		tmpl, err := template.New("show.gohtml").
			Funcs(template.FuncMap{
				"RouteName2URL": route.Name2URL,
				"Int64ToString": types.Int64ToString,
			}).ParseFiles("resources/views/articles/show.gohtml")
		logger.LogError(err)

		err = tmpl.Execute(w, article)
		logger.LogError(err)
	}
}

func getArticleByID(id string) (Article, error) {
	article := Article{}
	query := "SELECT * FROM articles WHERE id = ?"
	err := database.DB.QueryRow(query, id).Scan(&article.ID, &article.Title, &article.Content)
	return article, err
}

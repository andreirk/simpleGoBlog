package models

type Article struct {
	BlogTitle   string
	BlogArticle string

	// 29 Need user_id and id
	UserID int
	ID     int
}

// 29 Used to get 3 articles from database
type ArticleList struct {
	ID      []int
	UserID  []int
	Title   []string
	Content []string
}

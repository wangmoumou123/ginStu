package modules

type ArticleCate struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Cate        int       `json:"cate"`
	ArticleList []Article `gorm:"foreignKey:ArticleCateID"`
}

func (ArticleCate) TableName() string {
	return "article_cate"
}

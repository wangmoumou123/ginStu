package modules

type Article struct {
	Id            int         `json:"id"`
	Name          string      `json:"name"`
	Content       string      `json:"content"`
	AddTime       int         `json:"add_time"`
	ArticleCateID int         `json:"cate"`
	ArticleCate   ArticleCate `gorm:"foreignKey:ArticleCateID"`
}

func (Article) TableName() string {
	return "article"
}

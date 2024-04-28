package domain

type Links struct {
	ID        uint   `json:"id" gorm:"unique;not null"`
	Link      string `json:"link"`
	ShortLink string `json:"short_link" gorm:"column:short_link"`
}

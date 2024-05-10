package model

type Post struct {
	ID      int    `gorm:"primaryKey"`
	Title   string `gorm:"notNull;default:null"`
	Content string `gorm:"notNull;default:null"`
	Author  string `gorm:"notNull;default:null"`
}

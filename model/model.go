package model

type Model struct{
	Id string `gorm:"primaryKey" json:"id"`
}
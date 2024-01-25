package models

type PersonDB struct {
	ID         int    `gorm:"primaryKey;autoIncrement:true"`
	Name       string `gorm:"column:name"`
	Surname    string `gorm:"column:surname"`
	Patronymic string `gorm:"column:patronymic"`
	Info       InfoDB `gorm:"foreignKey:Name"`
}

type InfoDB struct {
	Age    int    `gorm:"column:age"`
	Name   string `gorm:"primaryKey;column:name"`
	Gender string `gorm:"column:gender"`
	Nation string `gorm:"column:nation"`
}

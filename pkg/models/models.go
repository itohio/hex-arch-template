package models

type Greeting struct {
	BaseModel
	Greeting string `gorm:"type:varchar(255)"`
}

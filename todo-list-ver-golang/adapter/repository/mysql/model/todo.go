package model

type Todo struct {
    ID      int     `gorm:"column:id; primaryKey"`
    Title   *string `gorm:"column:title"`
    Content *string `gorm:"column:content"`
}

func (Todo) TableName() string {
    return "todos"
}

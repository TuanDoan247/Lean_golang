package connection

import (
    "fmt"
    "go.uber.org/fx"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
)

type DatasourceOut struct {
    fx.Out
    Connection *gorm.DB
}

func NewConnection() DatasourceOut {
    out := DatasourceOut{}
    dsn := "root:123456@tcp(localhost:3306)/todo-list?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

    if err != nil {
        log.Fatalln(err)
    }
    fmt.Print(db)
    out.Connection = db
    return out
}

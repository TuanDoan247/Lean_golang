package mysql

import "gorm.io/gorm"

type base struct {
    db *gorm.DB
}

func (b base) BeginTransaction() *gorm.DB {
    tx := b.db.Begin()
    return tx
}

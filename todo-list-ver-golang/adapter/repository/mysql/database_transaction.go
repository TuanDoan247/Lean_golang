package mysql

import (
    "gorm.io/gorm"
    "todo-list-ver-golang/core/port"
)

type DatabaseTransaction struct {
    base
}

func NewDatabaseTransaction(db *gorm.DB) port.DatabaseTransactionPort {
    return &DatabaseTransaction{base: base{db: db}}
}

func (dt *DatabaseTransaction) StartTransaction() *gorm.DB {
    return dt.base.BeginTransaction()
}

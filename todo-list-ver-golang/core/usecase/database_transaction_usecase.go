package usecase

import (
    "gorm.io/gorm"
    "todo-list-ver-golang/core/port"
)

type IDatabaseTransactionUseCase interface {
    StartTransaction() *gorm.DB
}

type DatabaseTransactionUseCase struct {
    databaseTransaction port.DatabaseTransactionPort
}

func NewDatabaseTransactionUseCase(databaseTransaction port.DatabaseTransactionPort) IDatabaseTransactionUseCase {
    return &DatabaseTransactionUseCase{databaseTransaction: databaseTransaction}
}

func (dt *DatabaseTransactionUseCase) StartTransaction() *gorm.DB {
    return dt.databaseTransaction.StartTransaction()
}

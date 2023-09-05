package port

import "gorm.io/gorm"

type DatabaseTransactionPort interface {
    StartTransaction() *gorm.DB
}

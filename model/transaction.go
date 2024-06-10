package model

import "time"

type Transaction struct {
	Id               int     `gorm:"primarykey;autoIncrement" json:id`
	AccountID       string  `gorm:"foreignkey" json:"account_id"`
	BankID          string  `gorm:"foreignkey" json:"bank_id"`
	Amount           float64 `gorm:"column:amount" json:amount`
	Transaction_date *time.Time
}

// type Transaction struct {
// 	ID              int    ` gorm:"primarykey" ` 
// 	AccountID       string ` gorm:"foreignkey" `
// 	BankID          string ` gorm:"foreignkey" `  
// 	Amount          float64    ` gorm:"column:amount" `
// 	TransactionDate *time.Time
// }

func (c *Transaction) TableName() string {
	return ("transaction")
}

package model

type Bank struct {
	BankCode string ` gorm:"primarykey" `  //account_id
	Name     string ` gorm:"column:name" ` //mempertegas kolom
	Address  string
}

func (b *Bank) TableName() string {
	return "bank"
}

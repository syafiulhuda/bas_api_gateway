package model

type Accounts struct {
	AccountID string ` gorm:"primarykey" `      //account_id
	Username  string ` gorm:"column:username" ` //mempertegas kolom
	Password  string //ini otomatis membaca huruf kecil
	Name      string
}

func (a *Accounts) TableName() string {
	return "account"
}

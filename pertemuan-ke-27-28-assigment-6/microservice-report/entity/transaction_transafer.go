package entity

type TransactionTransfer struct {
	Id              string `gorm:"type:string;not null"`
	UserIdFrom      string `gorm:"type:string;not null"`
	UserIdTo        string `gorm:"type:string;not null"`
	TypeTransaction string `gorm:"type:string;not null"`
	TypeCredit      string `gorm:"type:string;not null"`
	Total           int64  `gorm:"type:int;not null"`
}

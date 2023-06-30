package dao

type Token struct {
	Id      uint64 `gorm:"primaryKey;type:bigint(20);not null"`
	Name    string `gorm:"type:varchar(20);not null"`
	Address string `gorm:"type:varchar(66);not null"`
	Decimal uint64 `gorm:"type:bigint(29);not null"`
}

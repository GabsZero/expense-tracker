package models

type Expense struct {
	Id                  uint
	Uid                 byte
	Name                string
	Value               float64
	Total_installments  uint
	Current_installment uint
}

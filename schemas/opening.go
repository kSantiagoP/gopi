package schemas

//Lógica para a configuração dos schemas e dos dados organizados no DB

import (
	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Role     string
	Company  string
	Location string
	Remote   bool
	Link     string
	Salary   int64
}

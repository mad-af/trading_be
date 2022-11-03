package repositories

import "trading_be/bin/repositories/models"

var ModelTables []interface{} = []interface{}{
	&models.Roles{},
	&models.Users{},
	&models.Banks{},
	&models.Grades{},
	&models.TransactionTypes{},
	&models.UserGrades{},
	&models.Balances{},
	&models.TransactionStatus{},
	&models.Transactions{},
	&models.TransactionUserGrades{},
}

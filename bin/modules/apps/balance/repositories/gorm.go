package repositories

import (
	"trading_be/bin/modules/apps/balance/models"

	"gorm.io/gorm"
)

type res struct {
	Data  interface{}
	Row   int64
	Error error
}

type Payload struct {
	Table    string
	Where    map[string]interface{}
	WhereRaw string
	Join     string
	Select   string
	Order    string
	Output   interface{}
}

func (o *Options) UpdateBalanceIncrease(data *models.Balances) <-chan res {
	var output = make(chan res)

	go func() {
		defer close(output)

		var db = o.DB.Table("balances").Where("user_id = ?", data.UserID).Update("value", gorm.Expr("value + ?", data.Value))
		if db.Error != nil {
			output <- res{Error: db.Error}
			return
		}

		output <- res{Data: data}
	}()

	return output
}

func (o *Options) UpdateBalanceSubtraction(data *models.Balances) <-chan res {
	var output = make(chan res)

	go func() {
		defer close(output)

		var db = o.DB.Table("balances").Where("user_id = ?", data.UserID).Update("value", gorm.Expr("value - ?", data.Value))
		if db.Error != nil {
			output <- res{Error: db.Error}
			return
		}

		output <- res{Data: data}
	}()

	return output
}

func (o *Options) Find(p *Payload) <-chan res {
	var output = make(chan res)

	go func() {
		defer close(output)

		var db = o.DB.Table(p.Table).Select(p.Select).Where(p.Where).Where(p.WhereRaw).Order(p.Order).Joins(p.Join).Find(&p.Output)
		if db.Error != nil {
			output <- res{Error: db.Error}
			return
		}

		output <- res{Data: p.Output, Row: db.RowsAffected}
	}()

	return output
}

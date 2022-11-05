package repositories

import (
	"trading_be/bin/modules/apps/grade/models"
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

func (o *Options) CreateTransactionUserGrade(data *models.TransactionUserGrades) <-chan res {
	var output = make(chan res)

	go func() {
		defer close(output)

		var db = o.DB.Create(&data)
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

func (o *Options) Count(p *Payload) <-chan res {
	var output = make(chan res)

	go func() {
		defer close(output)

		var count int64
		var db = o.DB.Table(p.Table).Select(p.Select).Where(p.Where).Order(p.Order).Joins(p.Join).Count(&count)
		if db.Error != nil {
			output <- res{Error: db.Error}
			return
		}

		output <- res{Data: count, Row: db.RowsAffected}
	}()

	return output
}

package repositories

import (
	"trading_be/bin/modules/apps/user/models"

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

func (o *Options) CreateUser(user *models.Users) <-chan res {
	var output = make(chan res)

	go func() {
		defer close(output)

		var err = o.DB.Transaction(func(tx *gorm.DB) error {
			if err := tx.Create(&user).Error; err != nil {
				return err
			}

			if err := tx.Create(&models.Balances{UserID: user.ID}).Error; err != nil {
				return err
			}

			if err := tx.Create(&models.UserGrades{UserID: user.ID, GradeID: 1}).Error; err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			output <- res{Error: err}
			return
		}

		output <- res{Data: user.ID}
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

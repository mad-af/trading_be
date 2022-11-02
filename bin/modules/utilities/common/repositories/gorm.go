package repositories

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
	Document map[string]interface{}
	Output   interface{}
}

func (o *Options) CreateMap(p *Payload) <-chan res {
	var output = make(chan res)

	go func() {
		defer close(output)

		var data []map[string]interface{}
		var db = o.DB.Table(p.Table).Create(&p.Document)
		if db.Error != nil {
			output <- res{Error: db.Error}
			return
		}

		output <- res{Data: data, Row: db.RowsAffected}
	}()

	return output
}

func (o *Options) FindMap(p *Payload) <-chan res {
	var output = make(chan res)

	go func() {
		defer close(output)

		var data []map[string]interface{}
		var db = o.DB.Table(p.Table).Select(p.Select).Where(p.Where).Where(p.WhereRaw).Order(p.Order).Joins(p.Join).Find(&data)
		if db.Error != nil {
			output <- res{Error: db.Error}
			return
		}

		output <- res{Data: data, Row: db.RowsAffected}
	}()

	return output
}

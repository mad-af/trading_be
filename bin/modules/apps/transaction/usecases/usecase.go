package usecases

import (
	"context"
	"encoding/json"
	"log"
	"math"
	"net/http"
	"strings"
	h "trading_be/bin/modules/apps/transaction/helpers"
	"trading_be/bin/modules/apps/transaction/models"
	rep "trading_be/bin/modules/apps/transaction/repositories"
	r "trading_be/bin/pkg/response"
	"trading_be/bin/pkg/utils"
)

func (s *Services) Create(c context.Context, p *models.ReqCreate) (result r.SendData, err error) {
	var res = new(map[string]interface{})
	result.Data = &res

	var transaction = <-s.Repository.CreateTransaction(&models.Transactions{
		UserID:            p.Options.UserID,
		BankID:            p.BankID,
		TransactionTypeID: p.TransactionTypeID,
		Status:            h.Status.Pending,
		Value:             p.Value,
	}, p.Options.UserID)
	if transaction.Error != nil {
		return result, transaction.Error
	}

	res = &map[string]interface{}{"id": transaction.Data}
	return result, nil
}

func (s *Services) Update(c context.Context, p *models.ReqUpdate) (result r.SendData, err error) {
	var res = new(map[string]interface{})
	result.Data = &res

	var transactions = <-s.Repository.Find(&rep.Payload{
		Table:  "transactions t",
		Where:  map[string]interface{}{"id": p.Param.ID},
		Select: "*",
		Output: &models.Transactions{},
	})
	if transactions.Error != nil || transactions.Row == 0 {
		return result, r.ReplyError("Transaction not found", http.StatusNotFound)
	}
	var transactionData = transactions.Data.(*models.Transactions)

	if p.Type == "status" {
		var transaction = models.Transactions{
			ID:     p.Param.ID,
			Status: p.Status,
		}
		switch p.Status {
		case h.Status.Rejected, h.Status.Canceled:
			transaction.Description = p.Description
		}

		var status = <-s.Repository.UpdateStatusTransaction(&transaction, p.Options.UserID)
		if status.Error != nil {
			return result, status.Error
		}

		if p.Status == h.Status.Finalized {
			switch transactionData.TransactionTypeID {
			case 1:
				go func() {
					transactionFetch, err := utils.FetchModule(&utils.FetchRequest{
						Method:        http.MethodPost,
						Url:           "/apps/v1/grade/upgrade",
						Authorization: "Bearer " + p.Options.Token,
						Body:          map[string]interface{}{"transaction_id": p.Param.ID},
					})
					if err != nil || transactionFetch.Err {
						log.Println("transaction-update-status:upgrade: " + transactionFetch.Message)
					}
				}()
			case 2:
				go func() {
					transactionFetch, err := utils.FetchModule(&utils.FetchRequest{
						Method:        http.MethodPost,
						Url:           "/apps/v1/balance/topup",
						Authorization: "Bearer " + p.Options.Token,
						Body:          map[string]interface{}{"transaction_id": p.Param.ID},
					})
					if err != nil || transactionFetch.Err {
						log.Println("transaction-update-status:topup: " + transactionFetch.Message)
					}
				}()
			}
		}

		res = &map[string]interface{}{"id": p.Param.ID, "status": p.Status}
	}

	return result, nil
}

func (s *Services) GetList(c context.Context, p *models.ReqGetList) (result r.SendData, err error) {
	var res = make([]models.UserList, 0)
	var meta = r.Meta{Page: p.Query.Page}
	result.Data = &res
	result.Meta = &meta

	// FILTER
	var filter = map[string]interface{}{}
	json.Unmarshal([]byte(p.Query.Query), &filter)
	var search = "u.name ilike '%" + p.Query.Search + "%'"

	// USECASE
	var counting = <-s.Repository.Count(&rep.Payload{
		Table:    "users u",
		Join:     "inner join user_grades ug on ug.user_id = u.id",
		Where:    filter,
		WhereRaw: search,
	})
	if counting.Error != nil || counting.Data.(int64) < 0 {
		return result, nil
	}
	meta.TotalData = int(counting.Data.(int64))

	var users = <-s.Repository.Find(&rep.Payload{
		Table: "users u",
		Join: `
			inner join user_grades ug on ug.user_id = u.id
			left join grades g on g.id = ug.grade_id
			left join roles r on r.id = u.role_id`,
		Where:    filter,
		WhereRaw: search,
		Select:   "u.*, g.name as grade_name, r.name as role_name",
		Order:    strings.Join(p.Query.Sort, " "),
		Output:   []models.UserList{},
	})
	if users.Error != nil {
		return result, nil
	}
	meta.Quantity = int(users.Row)
	meta.TotalPage = int(math.Ceil(float64((meta.TotalData + p.Query.Quantity - 1) / p.Query.Quantity)))

	res = users.Data.([]models.UserList)
	return result, nil
}

func (s *Services) GetDetail(c context.Context, p *models.ReqGetDetail) (result r.SendData, err error) {
	var res = new(models.UserDetail)
	result.Data = &res

	var user = <-s.Repository.Find(&rep.Payload{
		Table: "users u",
		Join: `
			inner join user_grades ug on ug.user_id = u.id
			left join grades g on g.id = ug.grade_id
			left join roles r on r.id = u.role_id
			left join balances b on b.user_id = u.id`,
		Where:  map[string]interface{}{"u.id": p.Param.ID},
		Select: "u.*, g.id as grade_id, g.name as grade_name, r.name as role_name, b.value as balance_value",
		Output: &models.UserDetail{},
	})
	if user.Error != nil || user.Row == 0 {
		return result, r.ReplyError("User not found", http.StatusNotFound)
	}

	res = user.Data.(*models.UserDetail)

	return result, nil
}

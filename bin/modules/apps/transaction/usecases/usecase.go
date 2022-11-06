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
		if h.StatusMap[transactionData.Status] > h.StatusMap[p.Status] {
			return result, r.ReplyError("Cannot update status transaction", http.StatusConflict)
		}

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
	var res = make([]models.TransactionData, 0)
	var meta = r.Meta{Page: p.Query.Page}
	result.Data = &res
	result.Meta = &meta

	// FILTER
	var filter = map[string]interface{}{}
	json.Unmarshal([]byte(p.Query.Query), &filter)
	var search = "u.name ilike '%" + p.Query.Search + "%'"

	// USECASE
	var counting = <-s.Repository.Count(&rep.Payload{
		Table: "transactions t",
		Join: `
			inner join users u on u.id = t.user_id
			inner join transaction_status ts on ts.transaction_id = t.id and ts.status = 'pending'
			inner join transaction_status ts1 on ts1.transaction_id = t.id and ts1.status = t.status`,
		Where:    filter,
		WhereRaw: search,
	})
	if counting.Error != nil || counting.Data.(int64) < 0 {
		return result, nil
	}
	meta.TotalData = int(counting.Data.(int64))

	var transaction = <-s.Repository.Find(&rep.Payload{
		Table: "transactions t",
		Join: `
			inner join users u on u.id = t.user_id
			inner join banks b on b.id = t.bank_id
			inner join transaction_types tt on tt.id = t.transaction_type_id
			inner join transaction_status ts on ts.transaction_id = t.id and ts.status = 'pending'
			inner join transaction_status ts1 on ts1.transaction_id = t.id and ts1.status = t.status`,
		Where:    filter,
		WhereRaw: search,
		Select: `t.*, u.name as user_name, b.name as bank_name, tt.name as transaction_type_name, 
			ts.created_at, ts1.created_at as updated_at`,
		Order:  strings.Join(p.Query.Sort, " "),
		Offset: p.Query.Quantity * (p.Query.Page - 1),
		Limit:  p.Query.Quantity,
		Output: []models.TransactionData{},
	})
	if transaction.Error != nil {
		return result, nil
	}
	meta.Quantity = int(transaction.Row)
	meta.TotalPage = int(math.Ceil(float64((meta.TotalData + p.Query.Quantity - 1) / p.Query.Quantity)))

	res = transaction.Data.([]models.TransactionData)
	return result, nil
}

func (s *Services) GetDetail(c context.Context, p *models.ReqGetDetail) (result r.SendData, err error) {
	var res = new(models.TransactionDetail)
	result.Data = &res

	var transaction = <-s.Repository.Find(&rep.Payload{
		Table: "transactions t",
		Join: `
			inner join users u on u.id = t.user_id
			inner join banks b on b.id = t.bank_id
			inner join transaction_types tt on tt.id = t.transaction_type_id
			inner join transaction_status ts on ts.transaction_id = t.id and ts.status = 'pending'
			inner join transaction_status ts1 on ts1.transaction_id = t.id and ts1.status = t.status`,
		Where:  map[string]interface{}{"t.id": p.Param.ID},
		Select: `t.*, u.name as user_name, b.name as bank_name, tt.name as transaction_type_name, 
			ts.created_at, ts1.created_at as updated_at`,
		Output: &models.TransactionData{},
	})
	if transaction.Error != nil || transaction.Row == 0 {
		return result, r.ReplyError("Transaction not found", http.StatusNotFound)
	}
	var bt, _ = json.Marshal(transaction.Data)
	json.Unmarshal(bt, &res)

	var transactionStatus = <-s.Repository.Find(&rep.Payload{
		Table: "transaction_status ts",
		Where:  map[string]interface{}{"ts.transaction_id": p.Param.ID},
		Select: `ts.*`,
		Output: []models.TransactionStatus{},
	})

	res.TransactionStatus = transactionStatus.Data.([]models.TransactionStatus)
	return result, nil
}

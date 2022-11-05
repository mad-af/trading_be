package usecases

import (
	"context"
	"log"
	"net/http"

	h "trading_be/bin/modules/apps/balance/helpers"
	"trading_be/bin/modules/apps/balance/models"
	rep "trading_be/bin/modules/apps/balance/repositories"
	r "trading_be/bin/pkg/response"
	"trading_be/bin/pkg/utils"
)

func (s *Services) TopUp(c context.Context, p *models.ReqTopUp) (result r.SendData, err error) {
	var res interface{}
	result.Data = &res

	var transaction = <-s.Repository.Find(&rep.Payload{
		Table: "transactions t",
		Where: map[string]interface{}{
			"t.id":                  p.TransactionID,
			"t.status":              h.Status.Finalized,
			"t.transaction_type_id": 2},
		Select: "t.*",
		Output: &models.Transactions{},
	})
	if transaction.Error != nil || transaction.Row == 0 {
		return result, r.ReplyError("Transaction not found", http.StatusNotFound)
	}
	var transactionData = transaction.Data.(*models.Transactions)

	var balance = <-s.Repository.UpdateBalance(&models.Balances{
		UserID: transactionData.UserID,
		Value:  transactionData.Value,
	})
	if balance.Error != nil {
		return result, r.ReplyError("Failed to top up balance", http.StatusInternalServerError)
	}

	go func() {
		transactionFetch, err := utils.FetchModule(&utils.FetchRequest{
			Method:        http.MethodPut,
			Url:           "/apps/v1/transaction/" + transactionData.ID,
			Authorization: "Bearer " + p.Options.Token,
			Body:          map[string]interface{}{"type": "status", "status": h.Status.Used},
		})
		if err != nil || transactionFetch.Err {
			log.Println("balance-topup:update-status-transaction: " + transactionFetch.Message)
		}
	}()

	res = balance.Data
	return result, nil
}

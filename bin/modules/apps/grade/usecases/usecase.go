package usecases

import (
	"context"
	"net/http"

	h "trading_be/bin/modules/apps/grade/helpers"
	"trading_be/bin/modules/apps/grade/models"
	rep "trading_be/bin/modules/apps/grade/repositories"
	r "trading_be/bin/pkg/response"
	"trading_be/bin/pkg/utils"
)

func (s *Services) Upgrade(c context.Context, p *models.ReqUpgrade) (result r.SendData, err error) {
	var res = new(models.TransactionUserGrades)
	result.Data = &res

	var userGrade = <-s.Repository.Find(&rep.Payload{
		Table:  "user_grades",
		Where:  map[string]interface{}{"user_id": p.Options.UserID},
		Select: "*",
		Output: &models.UserGrades{},
	})
	if userGrade.Error != nil {
		return result, r.ReplyError("User not found", http.StatusNotFound)
	}

	var UserGradeStatus = <-s.Repository.Count(&rep.Payload{
		Table: "transaction_user_grades tug",
		Join: "inner join transactions t on t.id = tug.transaction_id",
		Where: map[string]interface{}{
			"tug.user_grade_id": userGrade.Data.(*models.UserGrades).ID,
			"t.status":   []string{h.Status.Pending, h.Status.Transfered, h.Status.Checked}},
		Select: "*",
	})
	if UserGradeStatus.Data.(int64) > 0 {
		return result, r.ReplyError("There are still active grade upgrade transactions", http.StatusInternalServerError)
	}

	var grade = <-s.Repository.Find(&rep.Payload{
		Table:  "grades",
		Where:  map[string]interface{}{"id": p.GradeID},
		Select: "*",
		Output: &models.Grades{},
	})
	if grade.Error != nil {
		return result, r.ReplyError("Grade not found", http.StatusNotFound)
	}

	transactionFetch, err := utils.FetchModule(&utils.FetchRequest{
		Method:        http.MethodPost,
		Url:           "/apps/v1/transaction",
		Authorization: "Bearer " + p.Options.Token,
		Body: map[string]interface{}{
			"bank_id":             p.BankID,
			"transaction_type_id": 1,
			"value":               grade.Data.(*models.Grades).Price,
		},
	})
	if err != nil || transactionFetch.Err {
		return result, r.ReplyError("Failed to create transaction", http.StatusConflict)
	}

	var grades = <-s.Repository.CreateTransactionUserGrade(&models.TransactionUserGrades{
		TransactionID: transactionFetch.Data.(map[string]interface{})["id"].(string),
		UserGradeID:   userGrade.Data.(*models.UserGrades).ID,
		GradeID:       p.GradeID,
	})
	if grades.Error != nil {
		return result, grade.Error
	}

	res = grades.Data.(*models.TransactionUserGrades)
	return result, nil
}

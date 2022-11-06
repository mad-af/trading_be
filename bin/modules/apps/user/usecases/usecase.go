package usecases

import (
	"context"
	"encoding/json"
	"math"
	"net/http"
	"strings"
	"trading_be/bin/middleware"
	"trading_be/bin/modules/apps/user/models"
	rep "trading_be/bin/modules/apps/user/repositories"
	r "trading_be/bin/pkg/response"

	"golang.org/x/crypto/bcrypt"
)

func (s *Services) Create(c context.Context, p *models.ReqCreate) (result r.SendData, err error) {
	var res = new(map[string]interface{})
	result.Data = &res

	var password, perr = bcrypt.GenerateFromPassword([]byte(p.Password), 11)
	if perr != nil {
		return result, perr
	}

	var user = <-s.Repository.CreateUser(&models.Users{
		RoleID:   p.RoleID,
		Name:     p.Name,
		Username: p.Username,
		Email:    p.Email,
		Phone:    p.Phone,
		Password: string(password),
	})
	if user.Error != nil {
		return result, user.Error
	}

	res = &map[string]interface{}{"id": user.Data}
	return result, nil
}

func (s *Services) Login(c context.Context, p *models.ReqLogin) (result r.SendData, err error) {
	var res = new(models.ResLogin)
	result.Data = &res

	err = r.ReplyError("Invalid username and password Please try again", http.StatusUnauthorized)
	var user = <-s.Repository.Find(&rep.Payload{
		Table:  "users u",
		Where:  map[string]interface{}{"username": p.Username},
		Select: "u.*",
		Output: &models.UserDetail{},
	})
	if user.Error != nil {
		return result, user.Error
	} else if user.Row == 0 {
		return result, err
	}

	var userData = user.Data.(*models.UserDetail)
	if berr := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(p.Password)); berr != nil {
		return result, err
	}

	var token, terr = middleware.GenerateToken(middleware.JwtClaim{
		RoleID:  userData.RoleID,
		UserID:  userData.ID,
	})
	if terr != nil {
		return result, terr
	}

	res = &models.ResLogin{
		ID:    userData.ID,
		Token: token,
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
		Offset: p.Query.Quantity * (p.Query.Page - 1),
		Limit:  p.Query.Quantity,
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

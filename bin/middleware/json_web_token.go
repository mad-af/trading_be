package middleware

import (
	"net/http"
	"strings"
	"time"
	cf "trading_be/config"

	r "trading_be/bin/pkg/response"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type JwtClaim struct {
	RoleID  int    `json:"role_id"`
	UserID  string `json:"user_id"`
	GradeID int    `json:"grade_id"`
}

func GenerateToken(claim JwtClaim) (string, error) {
	type jwtCustomClaims struct {
		JwtClaim
		jwt.StandardClaims
	}

	var now = time.Now()
	var claims = jwtCustomClaims{
		claim,
		jwt.StandardClaims{
			ExpiresAt: now.Add(time.Hour * 72).Unix(),
			IssuedAt:  now.Unix(),
			Audience:  "671bab75-9832-4f09-956c-b3d891d0f0fc",
		},
	}

	var token = jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	return token.SignedString(cf.Env.PrivateKey)
}

func VerifyBearerToken() echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenString := strings.TrimPrefix(c.Request().Header.Get(echo.HeaderAuthorization), "Bearer ")

			if len(tokenString) == 0 {
				return r.ReplyError("Invalid token!", http.StatusUnauthorized)
			}

			var tokenParse, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				return cf.Env.PublicKey, nil
			})

			var errToken string
			switch ve := err.(type) {
			case *jwt.ValidationError:
				if ve.Errors == jwt.ValidationErrorExpired {
					errToken = "token has been expired"
				} else {
					errToken = "token parsing error"
				}
			}

			if len(errToken) > 0 {
				return r.ReplyError(errToken, http.StatusUnauthorized)
			}

			if !tokenParse.Valid {
				return r.ReplyError("token parsing error", http.StatusUnauthorized)
			}

			return next(c)
		}
	}
}



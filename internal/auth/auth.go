package auth

import (
	"context"
	"github.com/dimaskiddo/go-whatsapp-multidevice-rest/pkg/auth"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"

	"github.com/dimaskiddo/go-whatsapp-multidevice-rest/internal/auth/types"
	"github.com/dimaskiddo/go-whatsapp-multidevice-rest/internal/database"
)

// Login godoc
// @Summary User login
// @Description Logs in user using MongoDB and returns JWT token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param credentials body types.LoginRequest true "Username and Password"
// @Success 200 {object} map[string]string "token"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /auth/login [post]
func Login(c echo.Context) error {
	var req types.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	collection := database.GetMongoCollection("auth", "users")

	var user types.User
	err := collection.FindOne(context.TODO(), map[string]interface{}{
		"username": req.Username,
		"password": req.Password,
	}).Decode(&user)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid credentials"})
	}

	jid := user.Phone + "@s.whatsapp.net"

	claims := &types.AuthJWTClaims{
		JID: jid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(auth.AuthJWTSecret)) // set in .env
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Token generation failed"})
	}

	return c.JSON(http.StatusOK, echo.Map{"token": signedToken})
}

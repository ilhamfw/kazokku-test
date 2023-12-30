// app/handlers/user_handler.go

package handlers

import (
	"fmt"
	"kazokku-app/app/models"
	"kazokku-app/database"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// var validate = validator.New()

// UserRegisterRequest adalah struktur untuk menyimpan data permintaan registrasi pengguna
type UserRegisterRequest struct {
	Name              string   `json:"name" form:"name" validate:"required"`
	Address           string   `json:"address" form:"address" validate:"required"`
	Email             string   `json:"email" form:"email" validate:"required,email"`
	Password          string   `json:"password" form:"password" validate:"required"`
	Photos            []string `json:"photos" form:"photos" validate:"required"`
	CreditCardType    string   `json:"creditcard_type" form:"creditcard_type" validate:"required,creditcardtype"`
	CreditCardNumber  string   `json:"creditcard_number" form:"creditcard_number" validate:"required,creditcardnumber"`
	CreditCardName    string   `json:"creditcard_name" form:"creditcard_name" validate:"required"`
	CreditCardExpired string   `json:"creditcard_expired" form:"creditcard_expired" validate:"required,creditcardexpired"`
	CreditCardCVV     string   `json:"creditcard_cvv" form:"creditcard_cvv" validate:"required,creditcardcvv"`
}

// UserRegisterResponse adalah struktur untuk menyimpan data respons registrasi pengguna
type UserRegisterResponse struct {
	UserID uint `json:"user_id"`
}

// ErrorResponse adalah struktur untuk menyimpan data respons kesalahan
type ErrorResponse struct {
	Error string `json:"error"`
}

// UserRegisterHandler menangani permintaan registrasi pengguna
func UserRegisterHandler(c echo.Context) error {
	// Bind request data ke struct UserRegisterRequest
	req := new(UserRegisterRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid request payload"})
	}

	// Validasi request data menggunakan library validasi
	var v = CustomValidator{validator: validator.New()}
	SetCreditCardValidators(v.validator)
	c.Echo().Validator = &v

	// Log detail nilai yang sedang divalidasi
	fmt.Printf("Validating Credit Card Data: %+v\n", req)

	if err := c.Echo().Validator.Validate(req); err != nil {
		// Log detail kesalahan validasi
		fmt.Printf("Validation Error: %+v\n", err)

		return c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Credit card data invalid."})
	}

	// Lakukan logika bisnis untuk registrasi pengguna di sini
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to hash password"})
	}

	newUser := models.User{
		Name:     req.Name,
		Address:  req.Address,
		Email:    req.Email,
		Password: string(hashedPassword),
		Photos:   make([]models.Photo, len(req.Photos)),
		CreditCards: []models.CreditCard{{
			Type:    req.CreditCardType,
			Number:  req.CreditCardNumber,
			Name:    req.CreditCardName,
			Expired: req.CreditCardExpired,
			CVV:     req.CreditCardCVV,
		}},
	}

	// Simulasikan penyimpanan data ke database
	if err := database.DB.Create(&newUser).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return c.JSON(http.StatusNotFound, ErrorResponse{Error: "Record not found"})
		case gorm.ErrInvalidDB:
			return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Invalid DB"})
		default:
			return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to create user"})
		}
	}

	// Kirim respons sukses dengan ID pengguna
	res := UserRegisterResponse{UserID: newUser.ID}
	response := map[string]interface{}{
		"user_id": res.UserID,
	}

	return c.JSON(http.StatusOK, response)
}

// UserResponse struct untuk membentuk respons yang diharapkan
type UserResponse struct {
	UserID     uint       `json:"user_id"`
	Name       string     `json:"name"`
	Email      string     `json:"email"`
	Address    string     `json:"address"`
	Photos     Photos     `json:"photos"`
	CreditCard CreditCard `json:"creditcard"`
}

// Photos struct untuk membentuk struktur data photo pada respons
type Photos struct {
	FileList []string `json:"file_list"`
}

// CreditCard struct untuk membentuk struktur data credit card pada respons
type CreditCard struct {
	Type    string `json:"type"`
	Number  string `json:"number"`
	Name    string `json:"name"`
	Expired string `json:"expired"`
}

// GetAllUser handler untuk mendapatkan semua pengguna
func GetAllUser(c echo.Context) error {
	// Lakukan logika bisnis untuk mendapatkan semua pengguna dari database
	var users []models.User
	if err := database.DB.Preload("CreditCards").Preload("Photos").Find(&users).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return c.JSON(http.StatusNotFound, ErrorResponse{Error: "No users found"})
		case gorm.ErrInvalidDB:
			return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Invalid DB"})
		default:
			return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to get users"})
		}
	}

	// Format respons dengan data pengguna
	var userResponses []UserResponse
	for _, user := range users {
		photoFileList := make([]string, len(user.Photos))
		for i, photo := range user.Photos {
			photoFileList[i] = photo.FileName
		}

		creditCardInfo := CreditCard{}
		// Pengecekan panjang slice CreditCards
		if len(user.CreditCards) > 0 {
			creditCardInfo = CreditCard{
				Type:    user.CreditCards[0].Type,
				Number:  user.CreditCards[0].Number[len(user.CreditCards[0].Number)-4:],
				Name:    user.CreditCards[0].Name,
				Expired: user.CreditCards[0].Expired,
			}
		}

		userResponse := UserResponse{
			UserID:     user.ID,
			Name:       user.Name,
			Email:      user.Email,
			Address:    user.Address,
			Photos:     Photos{FileList: photoFileList},
			CreditCard: creditCardInfo,
		}
		userResponses = append(userResponses, userResponse)
	}

	response := map[string]interface{}{
		"count": len(userResponses),
		"rows":  userResponses,
	}

	return c.JSON(http.StatusOK, response)
}

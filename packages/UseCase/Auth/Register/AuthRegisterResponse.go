package Register

type AuthRegisterResponse struct {
  AccessToken string `json:"access_token" gorm:"not null;type:varchar(255)"`
}

func NewAuthRegisterResponse(tokenString string) AuthRegisterResponse {
  return AuthRegisterResponse{AccessToken: tokenString}
}
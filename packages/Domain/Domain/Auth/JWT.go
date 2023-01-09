package Auth

import (
  "fmt"
  "net/http"
  "strings"
  jwt "github.com/dgrijalva/jwt-go"
  "birthday-reminder/config"
)

var cfg = config.NewConfig()
// validateToken Validate JWT token from HTTP Header
func ValidateToken(w http.ResponseWriter, r *http.Request) error {
  tokenString := getAccessToken(w, r)
  token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
  if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
    return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
  }
    return []byte(cfg.JWT_SECRET_KEY), nil
  })
  if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
    fmt.Println(claims["user_id"])
    return nil
  } else {
    return err
  }
}

// getAccessToken Get JWT AccessToken from header. and remove "Bearer " string from Token
func getAccessToken(w http.ResponseWriter, r *http.Request) string {
  token_string_poor := r.Header["Authorization"][0]
  return strings.Split(token_string_poor, "Bearer ")[1]
}

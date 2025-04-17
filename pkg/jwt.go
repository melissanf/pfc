package utils
import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
	"net/http"
)
type JwtToken struct {
	Token string `json:"token"`
}
type Exception struct {
	Message string `json:"message"`
}



func CreateTokenEndpoint(w http.ResponseWriter, r *http.Request) {
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Name,
	})
	tokenString, err := token.SignedString([]byte("secret"))

	if err != nil {
		fmt.Println(err)
	}

	json.NewEncoder(w).Encode(JwtToken{Token: tokenString})
}

func ProtectedEndpoint(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	token, _ := jwt.Parse(params["token"][0], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return []byte("secret"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var user User

		mapstructure.Decode(claims, &user)
		json.NewEncoder(w).Encode(user)
	} else {
		json.NewEncoder(w).Encode(Exception{Message: "Invalid authorization token"})
	}
}
package main
import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

func Refresh(w http.ResponseWriter, r *http.Request){
	
	cookie, err := r.Cookie("token")
	if err != nil{
		if err == http.ErrNoCookie{
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
	}

	tokenStr := cookie.Value
	
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
	func(t *jwt.Token) (interface{}, error){
		return jwtKey, nil
	})
	if err != nil{
		if err == jwt.ErrSignatureInvalid{
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if!tkn.Valid{
        w.WriteHeader(http.StatusUnauthorized)
        return
    }

	// if time.unix(claims.ExpiresAt,0).Sub(time.Now()) > 30 * time.Second{
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	expirationTime := time.Now().Add(time.Minute*5)

	claims.ExpiresAt = expirationTime.Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenString, err := token.SignedString(jwtKey)

	if  err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name: "refresh_token",
		Value: tokenString,
		Expires: expirationTime,
	})

}
package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"hexarch/pkg/config"
	"io"
	"net/http"
	"strings"

	"github.com/form3tech-oss/jwt-go"
	"github.com/go-chi/chi/v5"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/auth0"
	"golang.org/x/crypto/bcrypt"
)

func New(cfg *config.Config) http.Handler {
	goth.UseProviders(
		auth0.New(cfg.Auth.ClientID, cfg.Auth.Secret, fmt.Sprintf("http://%s/auth/auth0/callback", cfg.Server.Address), cfg.Auth.Domain),
	)

	router := chi.NewRouter()

	router.HandleFunc("/{provider}/callback", func(res http.ResponseWriter, req *http.Request) {
		_, err := gothic.CompleteUserAuth(res, req)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		res.Header().Set("Location", "/")
		res.WriteHeader(http.StatusTemporaryRedirect)
	})

	router.HandleFunc("/logout/{provider}", func(res http.ResponseWriter, req *http.Request) {
		gothic.Logout(res, req)
		res.Header().Set("Location", "/")
		res.WriteHeader(http.StatusTemporaryRedirect)
	})

	router.HandleFunc("/{provider}", func(res http.ResponseWriter, req *http.Request) {
		// try to get the user without re-authenticating
		if _, err := gothic.CompleteUserAuth(res, req); err == nil {
			res.Header().Set("Location", "/")
			res.WriteHeader(http.StatusTemporaryRedirect)
		} else {
			gothic.BeginAuthHandler(res, req)
		}
	})

	return router
}

func SecureToken() string {
	b := make([]byte, 16)
	// rand should never fail, if it does we have bigger problems
	_, _ = io.ReadFull(rand.Reader, b)
	return strings.TrimRight(base64.URLEncoding.EncodeToString(b), "=")
}

func HashPassword(password string) []byte {
	// we can safely ignore any error because we control the cost
	pw, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return pw
}

// CheckPassword checks to see if the password matches the hashed password.
func CheckPassword(hash []byte, password string) error {
	return bcrypt.CompareHashAndPassword(hash, []byte(password))
}

func GetAuthorization(ctx context.Context) (*jwt.Token, error) {
	userVal := ctx.Value("user")
	if userVal == nil {
		return nil, fmt.Errorf("not authorized")
	}
	token, ok := userVal.(*jwt.Token)
	if !ok {
		return nil, fmt.Errorf("bad token")
	}
	if !token.Valid {
		return nil, fmt.Errorf("bad token")
	}

	return token, nil
}

type CustomClaims struct {
	Scope string `json:"scope"`
	jwt.StandardClaims
}

func GetScopes(token *jwt.Token) ([]string, error) {
	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, fmt.Errorf("bad claims")
	}

	return strings.Split(claims.Scope, " "), nil
}

func GetSubject(token *jwt.Token) (string, error) {
	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return "", fmt.Errorf("bad claims")
	}

	return claims.Subject, nil
}

func GetTokenID(token *jwt.Token) (string, error) {
	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return "", fmt.Errorf("bad claims")
	}

	return claims.Id, nil
}

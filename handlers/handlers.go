package handlers

import (
	"errors"
	"fmt"
	"gameon-twotwentyk-api/models"
	"gameon-twotwentyk-api/store"
	"net/http"

	"github.com/go-chi/jwtauth"
	"github.com/pandoratoolbox/json"
)

func ServeError(w http.ResponseWriter, message string, code int) {
	fmt.Printf("Http error: %s\n", message)
	http.Error(w, message, code)
}

func ServeJSON(w http.ResponseWriter, data interface{}) {
	js, err := json.Marshal(data)

	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

var (
	TokenAuth = jwtauth.New("HS256", []byte("h1l32b"), nil)
)

func Login(w http.ResponseWriter, r *http.Request) {
	input := models.User{}
	ctx := r.Context()

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&input)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	user, err := store.GetUserByEmail(ctx, *input.Username)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	if user.Password == nil {
		ServeError(w, "Please log in with your social account", http.StatusUnauthorized)
		return
	}

	if *user.Password != *input.Password {
		ServeError(w, errors.New("Wrong password").Error(), 400)
		return
	}

	_, jwtstring, err := TokenAuth.Encode(map[string]interface{}{
		"id":       *user.Id,
		"role_ids": *user.RoleIds,
	})
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	out := struct {
		Token string
	}{
		Token: jwtstring,
	}

	js, err := json.Marshal(out)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	w.Write(js)
}

func Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	data := struct {
		models.User
		Sign string `json:"sign"`
	}{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	roles := models.Ints{1}

	user := models.User{
		UserData: models.UserData{
			Username:      data.Username,
			Password:      data.Password,
			Email:         data.Email,
			RoleIds:       &roles,
			WalletAddress: data.WalletAddress,
		},
	}

	err = registerNewUser(ctx, user)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	_, jwtstring, err := TokenAuth.Encode(map[string]interface{}{
		"id":       *user.Id,
		"role_ids": *user.RoleIds,
	})

	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	response := struct {
		User  models.User
		Token string
	}{
		User:  user,
		Token: jwtstring,
	}

	ServeJSON(w, response)
}

func GetMyUserData(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	mid := ctx.Value(models.CTX_user_id).(int64)

	user, err := store.GetUser(ctx, mid)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.Password = nil

	ServeJSON(w, user)
}

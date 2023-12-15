package handlers

import (
	"context"
	"gameon-twotwentyk-api/models"
	"net/http"

	"github.com/go-chi/jwtauth"
)

const (
	ROLE_ID_ADMIN = 999
)

func RestrictAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		is_admin := false

		rids := r.Context().Value(models.CTX_user_role_ids).([]int64)

		for _, rid := range rids {
			if rid == ROLE_ID_ADMIN {
				is_admin = true
			}
		}

		if !is_admin {
			http.Error(w, "Restricted, access for admins only", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func RestrictAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !r.Context().Value(models.CTX_is_auth).(bool) {
			http.Error(w, "Restricted access, please log in.", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		_, claims, err := jwtauth.FromContext(ctx)
		if err != nil {
			ctx = context.WithValue(ctx, models.CTX_is_auth, false)
		} else {
			claimsRole := claims["role_ids"]
			if claimsRole == nil {
				ctx = context.WithValue(ctx, models.CTX_is_auth, false)
			} else {
				var rids []int64
				rrids := claimsRole.([]interface{})
				for _, rrid := range rrids {
					rids = append(rids, int64(rrid.(float64)))
				}

				ctx = context.WithValue(ctx, models.CTX_user_role_ids, rids)
			}

			claimsId := claims["id"]
			if claimsId == nil {
				ctx = context.WithValue(ctx, models.CTX_is_auth, false)
			} else {
				ctx = context.WithValue(ctx, models.CTX_user_id, int64(claimsId.(float64)))
				ctx = context.WithValue(ctx, models.CTX_is_auth, true)
			}

		}

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

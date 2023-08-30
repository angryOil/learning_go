package main

import (
	"context"
	"net/http"
)

func main() {

}

func extractUser(r *http.Request) (string, error) {
	userCookie, err := r.Cookie("user")
	if err != nil {
		return "", err
	}
	return userCookie.Value, nil
}

func middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, err := extractUser(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx := r.Context()
		ctx = ContextWithUser(ctx, user)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	})

}

func ContextWithUser(ctx context.Context, userId string) context.Context {
	return context.WithValue(ctx, "key", userId)
}

func UserFromContext(ctx context.Context) (string, bool) {
	user, ok := ctx.Value("key").(string)
	return user, ok
}

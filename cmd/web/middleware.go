package main

import (
	"fmt"
	"github.com/justinas/nosurf"
	"net/http"
)

func secureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("X-XSS-Protection", "1; mode=block")
		rw.Header().Set("X-Frame-Options", "deny")
		next.ServeHTTP(rw, req)
	})
}

func (app *application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		app.infoLog.Printf("%s - %s %s %s", req.RemoteAddr, req.Proto, req.Method, req.URL)
		next.ServeHTTP(rw, req)
	})
}

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				rw.Header().Set("Connection", "close")
				app.serverError(rw, fmt.Errorf("%s", err))
			}
		}()
		next.ServeHTTP(rw, req)
	})
}

func (app *application) requireAuthenticatedUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if app.authenticatedUser(req) == 0 {
			http.Redirect(rw, req, "/user/login", 302)
			return
		}

		next.ServeHTTP(rw, req)
	})
}

func noSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   true,
	})
	return csrfHandler
}

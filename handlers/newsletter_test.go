package handlers_test

import (
	"canvas/handlers"
	"canvas/model"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/matryer/is"
)

type signupperMock struct {
	email model.Email
}

func (s *signupperMock) SignupForNewsletter(ctx context.Context, email model.Email) (string, error) {
	s.email = email
	return "", nil
}

func TestNewsletterSignup(t *testing.T) {
	mux := chi.NewMux()
	s := &signupperMock{}
	handlers.NewsletterSignup(mux, s)

	t.Run("signs up a valid email address", func(t *testing.T) {
		is := is.New(t)

		code, _, _ := makePostRequest(
			mux,
			"/newsletter/signup",
			createFormHeader(),
			strings.NewReader("email=me%40example.com"),
		)
		is.Equal(code, http.StatusFound)
		is.Equal(s.email, model.Email("me@example.com"))
	})

	t.Run("rejects an invalid email address", func(t *testing.T) {
		is := is.New(t)

		code, _, _ := makePostRequest(
			mux,
			"/newsletter/signup",
			createFormHeader(),
			strings.NewReader("email=invalidemail"),
		)
		is.Equal(code, http.StatusBadRequest)
	})
}

// makePostRequest and returns the status code, response header, and body.
func makePostRequest(handler http.Handler, target string, header http.Header, body io.Reader) (int, http.Header, []byte) {
	r := httptest.NewRequest(http.MethodPost, target, body)
	r.Header = header
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, r)
	result := w.Result()
	bodyBytes, err := io.ReadAll(result.Body)
	if err != nil {
		panic(err)
	}
	return result.StatusCode, result.Header, bodyBytes
}

func createFormHeader() http.Header {
	header := http.Header{}
	header.Set("Content-Type", "application/x-www-form-urlencoded")
	return header
}

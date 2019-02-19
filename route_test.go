package bon

import (
	"net/http"
	"testing"
)

func TestRouteMiddleware(t *testing.T) {
	r := NewRouter()

	r.Use(WriteMiddleware("A"))
	r.Get("/a", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("a"))
	})

	rt := r.Route()
	rt.Get("/b", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("b"))
	})

	if err := Verify(r,
		[]*Want{
			{"/a", 200, "Aa"},
			{"/b", 200, "b"},
		},
	); err != nil {
		t.Fatal(err)
	}
}

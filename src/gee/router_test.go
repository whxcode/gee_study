package gee

import (
	"reflect"
	"testing"
)

func newTestRouter() *router {
	r := newRouter()
	r.addRoute("GET", "/", nil, nil)
	r.addRoute("GET", "/hello/:name", nil, nil)

	r.addRoute("GET", "/hello/b/c", nil, nil)
	r.addRoute("GET", "/hi/:name", nil, nil)
	r.addRoute("GET", "/assets/*filepath", nil, nil)
	return r
}

func TestParsePattern(t *testing.T) {
	ok := reflect.DeepEqual(parsePattern("/p/:name"), []string{"p", ":name"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*"), []string{"p", "*"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*name/*"), []string{"p", "*name"})

	if !ok {
		t.Fatal("test parsePattern failed")
	}
}

func TestGetRoute(t *testing.T) {
	r := newTestRouter()

	n, ps := r.getRoute("GET", "/hello/geektutu")

	if n == nil {
		t.Fatal("nil shouldn't be returned")
	}

	if n.pattern != "/hello/:name" {
		t.Fatal("should match /hello/:name")
	}

	if ps["name"] != "geektutu" {
		t.Fatal("name should be equal to 'geektutu'")
	}

	{
		n, ps := r.getRoute("GET", "/hi/whx")

		if n == nil {
			t.Fatal("nil shouldn't be returned")
		}

		if n.pattern != "/hi/:name" {
			t.Fatal("should match /hello/:name")
		}

		if ps["name"] != "whx" {
			t.Fatal("name should be equal to 'geektutu'")
		}

		t.Logf("matched path: %s, params['name']: %s\n", n.pattern, ps["name"])

	}

	t.Logf("matched path: %s, params['name']: %s\n", n.pattern, ps["name"])
}

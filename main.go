package main

import "net/http"

type Router struct {
	RouterEntry []RouterEntry
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	// Iterate throught the router's router entry and if path matches process the request.
	// If not return 404
	for _, r := range r.RouterEntry {
		if r.Match(req) {
			r.Handler.ServeHTTP(w, req)
			return
		}
	}
	http.NotFound(w, req)

}

// We need to get 2 information from the request, 1. HTTP method, 2. Path
// then invoke the right handler which will handle the request
// This info needs to be stored in the router.
type RouterEntry struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

// on strating the server we need to register the route to the router
func (r *Router) Register(method, path string, handler http.HandlerFunc) {
	// create a route entry and add it to the router entry list
	e := RouterEntry{
		Method:  method,
		Path:    path,
		Handler: handler,
	}
	r.RouterEntry = append(r.RouterEntry, e)
}

// Match the request on RouterEntry
func (r *RouterEntry) Match(req *http.Request) bool {

	if r.Method == req.Method && r.Path == req.URL.Path {
		return true
	}
	return false

}

func main() {
	router := &Router{}

	router.Register("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("The Best Router!"))
	})

	err := http.ListenAndServe(":8000", router)
	if err != nil {
		panic(err)
	}
}

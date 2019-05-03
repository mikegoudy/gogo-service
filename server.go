package service

import (
    "net/http"
    
    "github.com/codegangsta/negroni"
    "github.com/gorilla/mux"
    "github.com/unrolled/render"
)

// NewServer configures and returns a server
func NewServer() *negroni.Negroni {
    
    formatter := render.New(render.Options{
        IndentJSON: true,
    })
    
    n := negroni.Classic()
    mx := mux.NewRouter()
    
    initRoutes(mx, formatter)
    
    n.UseHandler(mx)
    
    return n
}

// InitRoutes
func initRoutes(mx *mux.Router, fomatter *render.Render) {
    mx.HandleFunc("/test", testHandler(formatter)).Methods("GET")
}

// testHandler
func testHandler(formatter *render.Render) http.HandlerFunc {
    
    return func( w http.Response, req *http.Request) {
        formatter.JSON(w, http.StatusOK, struct{ Test string } {"This is a test"} )
    }
}
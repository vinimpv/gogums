package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"vinimpv/gogums/controlers"
	"vinimpv/gogums/services"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func SiteContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siteId, err := strconv.ParseInt(chi.URLParam(r, "siteId"), 10, 64)

		site, err := services.Sites.GetSite(siteId)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}
		ctx := context.WithValue(r.Context(), "site", site)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func fileServer(router *chi.Mux) {
	router.Get("/preview", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "repos/cde/exampleSite/public/index.html")
	})

	router.Get("/preview/*", func(w http.ResponseWriter, r *http.Request) {
		arquivo := "repos/cde/exampleSite/public" + strings.ReplaceAll(r.RequestURI, "/preview", "")
		fmt.Println(arquivo)
		if _, err := os.Stat(arquivo); os.IsNotExist(err) {
			http.ServeFile(w, r, arquivo)
		} else {
			http.ServeFile(w, r, arquivo+"index.html")
		}
	})
}

func main() {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	r.Use(middleware.Logger)
	r.Route("/sites", func(r chi.Router) {
		r.Get("/", controlers.GetSites)

		r.Put("/", controlers.CreateSite)
		r.Route("/{siteId}", func(r chi.Router) {
			r.Use(SiteContext)
			r.Get("/", controlers.GetSite)
			r.Put("/resources", controlers.ResourceControler)
		})
	})
	r.Route("/repositories", func(r chi.Router) {
		r.Get("/", controlers.GetRepositories)
		r.Put("/", controlers.CreateRepository)

	})
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "repos/cde/exampleSite/public"))
	FileServer(r, "/preview", filesDir)

	http.ListenAndServe(":3001", r)

}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}

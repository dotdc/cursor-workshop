package app

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"path/filepath"
	"runtime"
	"time"

	"html/template"
)

// Server wraps the HTTP router and template state for the Gopher Bookshop.
type Server struct {
	mux       *http.ServeMux
	basePath  string
	templates map[string]*template.Template
}

// NewServer constructs a Server with all routes and assets ready to serve.
func NewServer() (*Server, error) {
	basePath, err := projectRoot()
	if err != nil {
		return nil, err
	}

	srv := &Server{
		mux:      http.NewServeMux(),
		basePath: basePath,
	}
	if err := srv.loadTemplates(); err != nil {
		return nil, err
	}

	srv.registerRoutes()
	return srv, nil
}

// ServeHTTP satisfies http.Handler.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

// ListenAndServe starts the HTTP server on the provided address.
func (s *Server) ListenAndServe(addr string) error {
	httpServer := &http.Server{
		Addr:              addr,
		Handler:           s,
		ReadHeaderTimeout: 5 * time.Second,
	}
	log.Printf("Gopher Bookshop listening on %s", addr)
	return httpServer.ListenAndServe()
}

func (s *Server) registerRoutes() {
	staticDir := http.Dir(filepath.Join(s.basePath, "app", "static"))
	s.mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(staticDir)))

	s.mux.HandleFunc("/", s.handleIndex)
	s.mux.HandleFunc("/books", s.handleBooks)
	s.mux.HandleFunc("/api/healthz", s.handleHealth)
	s.mux.HandleFunc("/api/books", s.handleBooksAPI)
	s.mux.HandleFunc("/ee", s.handleErrorEndpoint)
}

func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"WorkshopGoalsHTML": LoadWorkshopGoalsHTML(s.basePath),
	}
	s.render(w, "index", data)
}

func (s *Server) handleBooks(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"Books": GetBooks(),
	}
	s.render(w, "books", data)
}

func (s *Server) handleErrorEndpoint(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"Books":           GetBooks(),
		"ShowCursorHero":  true,
		"CursorHeroImage": "img/cursor-hero.svg",
	}
	s.render(w, "books", data)
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	payload := map[string]string{
		"status": GetHealthStatus(),
	}
	s.respondJSON(w, payload, http.StatusOK)
}

func (s *Server) handleBooksAPI(w http.ResponseWriter, r *http.Request) {
	s.respondJSON(w, GetBooks(), http.StatusOK)
}

func (s *Server) respondJSON(w http.ResponseWriter, payload any, status int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(payload); err != nil {
		log.Printf("failed to encode JSON: %v", err)
	}
}

func (s *Server) render(w http.ResponseWriter, templateName string, data map[string]any) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tmpl, ok := s.templates[templateName]
	if !ok {
		http.Error(w, "template not found", http.StatusInternalServerError)
		return
	}

	if err := tmpl.ExecuteTemplate(w, "base", data); err != nil {
		log.Printf("render error: %v", err)
		http.Error(w, "template rendering error", http.StatusInternalServerError)
	}
}

func projectRoot() (string, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "", errors.New("unable to determine project root")
	}
	return filepath.Clean(filepath.Join(filepath.Dir(filename), "..")), nil
}

func (s *Server) loadTemplates() error {
	funcs := template.FuncMap{
		"formatRarity": FormatRarity,
	}

	base := filepath.Join(s.basePath, "app", "templates", "base.html")
	pages := map[string]string{
		"index": "index.html",
		"books": "items.html",
	}

	s.templates = make(map[string]*template.Template, len(pages))

	for key, file := range pages {
		pagePath := filepath.Join(s.basePath, "app", "templates", file)
		tmpl, err := template.New("base").Funcs(funcs).ParseFiles(base, pagePath)
		if err != nil {
			return err
		}
		s.templates[key] = tmpl
	}

	return nil
}

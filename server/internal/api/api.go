package api

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/HyphaGroup/armor/server/internal/db"
	"github.com/HyphaGroup/armor/server/internal/validator"
)

type Server struct {
	db        *db.DB
	validator *validator.Validator
	password  string
	mux       *http.ServeMux
}

func NewServer(database *db.DB, val *validator.Validator) *Server {
	password := os.Getenv("ARMOR_PASSWORD")
	if password == "" {
		password = "armor" // default for development
		log.Println("Warning: ARMOR_PASSWORD not set, using default 'armor'")
	}

	s := &Server{
		db:        database,
		validator: val,
		password:  password,
		mux:       http.NewServeMux(),
	}

	s.setupRoutes()
	return s
}

func (s *Server) setupRoutes() {
	s.mux.HandleFunc("/api/profiles", s.handleProfiles)
	s.mux.HandleFunc("/api/profiles/", s.handleProfileRoutes)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Auth check
	if !s.checkAuth(r) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	s.mux.ServeHTTP(w, r)
}

func (s *Server) checkAuth(r *http.Request) bool {
	auth := r.Header.Get("Authorization")
	if auth == "" {
		return false
	}

	parts := strings.SplitN(auth, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		return false
	}

	return parts[1] == s.password
}

func (s *Server) handleProfiles(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		s.listProfiles(w, r)
	case "POST":
		s.createProfile(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (s *Server) handleProfileRoutes(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api/profiles/")
	parts := strings.Split(path, "/")

	if len(parts) == 0 || parts[0] == "" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	profileID := parts[0]

	if len(parts) == 1 {
		switch r.Method {
		case "GET":
			s.getProfile(w, r, profileID)
		case "DELETE":
			s.deleteProfile(w, r, profileID)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
		return
	}

	section := parts[1]
	if !db.IsValidSection(section) {
		http.Error(w, "Invalid section", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		s.getSection(w, r, profileID, section)
	case "PUT":
		s.updateSection(w, r, profileID, section)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (s *Server) listProfiles(w http.ResponseWriter, r *http.Request) {
	profiles, err := s.db.ListProfiles()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var summaries []map[string]interface{}
	for _, p := range profiles {
		sections := map[string]*string{
			"mission":     p.Mission,
			"assets":      p.Assets,
			"adversaries": p.Adversaries,
			"threats":     p.Threats,
			"risks":       p.Risks,
			"mitigations": p.Mitigations,
		}
		completeness := validator.CalculateProfileCompleteness(sections)

		summaries = append(summaries, map[string]interface{}{
			"id":           p.ID,
			"name":         p.Name,
			"description":  p.Description,
			"completeness": completeness.Overall,
			"created_at":   p.CreatedAt,
			"updated_at":   p.UpdatedAt,
		})
	}

	if summaries == nil {
		summaries = []map[string]interface{}{}
	}

	writeJSON(w, summaries)
}

func (s *Server) createProfile(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	profile, err := s.db.CreateProfile(req.Name, req.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	writeJSON(w, profile)
}

func (s *Server) getProfile(w http.ResponseWriter, r *http.Request, id string) {
	profile, err := s.db.GetProfile(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if profile == nil {
		http.Error(w, "Profile not found", http.StatusNotFound)
		return
	}

	sections := map[string]*string{
		"mission":     profile.Mission,
		"assets":      profile.Assets,
		"adversaries": profile.Adversaries,
		"threats":     profile.Threats,
		"risks":       profile.Risks,
		"mitigations": profile.Mitigations,
	}
	completeness := validator.CalculateProfileCompleteness(sections)

	response := map[string]interface{}{
		"id":           profile.ID,
		"name":         profile.Name,
		"description":  profile.Description,
		"mission":      parseJSON(profile.Mission),
		"assets":       parseJSON(profile.Assets),
		"adversaries":  parseJSON(profile.Adversaries),
		"threats":      parseJSON(profile.Threats),
		"risks":        parseJSON(profile.Risks),
		"mitigations":  parseJSON(profile.Mitigations),
		"completeness": completeness,
		"created_at":   profile.CreatedAt,
		"updated_at":   profile.UpdatedAt,
	}

	writeJSON(w, response)
}

func (s *Server) deleteProfile(w http.ResponseWriter, r *http.Request, id string) {
	err := s.db.DeleteProfile(id)
	if err != nil {
		http.Error(w, "Profile not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) getSection(w http.ResponseWriter, r *http.Request, profileID, section string) {
	profile, err := s.db.GetProfile(profileID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if profile == nil {
		http.Error(w, "Profile not found", http.StatusNotFound)
		return
	}

	var data *string
	switch section {
	case "mission":
		data = profile.Mission
	case "assets":
		data = profile.Assets
	case "adversaries":
		data = profile.Adversaries
	case "threats":
		data = profile.Threats
	case "risks":
		data = profile.Risks
	case "mitigations":
		data = profile.Mitigations
	}

	writeJSON(w, map[string]interface{}{
		"data": parseJSON(data),
	})
}

func (s *Server) updateSection(w http.ResponseWriter, r *http.Request, profileID, section string) {
	profile, err := s.db.GetProfile(profileID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if profile == nil {
		http.Error(w, "Profile not found", http.StatusNotFound)
		return
	}

	var body json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	dataStr := string(body)

	if s.validator.HasSchema(section) {
		errors, err := s.validator.Validate(section, dataStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if len(errors) > 0 {
			w.WriteHeader(http.StatusBadRequest)
			writeJSON(w, map[string]interface{}{
				"error":  "Validation failed",
				"errors": errors,
			})
			return
		}
	}

	if err := s.db.UpdateSection(profileID, section, dataStr); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, map[string]interface{}{
		"success": true,
		"data":    parseJSON(&dataStr),
	})
}

func writeJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func parseJSON(data *string) interface{} {
	if data == nil {
		return nil
	}

	var result interface{}
	if err := json.Unmarshal([]byte(*data), &result); err != nil {
		return nil
	}

	return result
}

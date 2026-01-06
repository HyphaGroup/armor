package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	conn *sql.DB
}

type Profile struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description,omitempty"`
	Mission     *string    `json:"mission,omitempty"`
	Assets      *string    `json:"assets,omitempty"`
	Adversaries *string    `json:"adversaries,omitempty"`
	Threats     *string    `json:"threats,omitempty"`
	Risks       *string    `json:"risks,omitempty"`
	Mitigations *string    `json:"mitigations,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type ProfileSummary struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description,omitempty"`
	Completeness float64 `json:"completeness"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}

func Open(path string) (*DB, error) {
	conn, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := conn.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	db := &DB{conn: conn}
	if err := db.migrate(); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	return db, nil
}

func (db *DB) Close() error {
	return db.conn.Close()
}

func (db *DB) migrate() error {
	schema := `
	CREATE TABLE IF NOT EXISTS profiles (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT,
		mission TEXT,
		assets TEXT,
		adversaries TEXT,
		threats TEXT,
		risks TEXT,
		mitigations TEXT,
		created_at TEXT NOT NULL,
		updated_at TEXT NOT NULL
	);
	`
	_, err := db.conn.Exec(schema)
	return err
}

func (db *DB) CreateProfile(name, description string) (*Profile, error) {
	id := uuid.New().String()
	now := time.Now().UTC()
	nowStr := now.Format(time.RFC3339)

	_, err := db.conn.Exec(`
		INSERT INTO profiles (id, name, description, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?)
	`, id, name, description, nowStr, nowStr)
	if err != nil {
		return nil, fmt.Errorf("failed to create profile: %w", err)
	}

	return &Profile{
		ID:          id,
		Name:        name,
		Description: description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

func (db *DB) GetProfile(id string) (*Profile, error) {
	var p Profile
	var createdAt, updatedAt string
	var description, mission, assets, adversaries, threats, risks, mitigations sql.NullString

	err := db.conn.QueryRow(`
		SELECT id, name, description, mission, assets, adversaries, threats, risks, mitigations, created_at, updated_at
		FROM profiles WHERE id = ?
	`, id).Scan(&p.ID, &p.Name, &description, &mission, &assets, &adversaries, &threats, &risks, &mitigations, &createdAt, &updatedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get profile: %w", err)
	}

	p.Description = description.String
	if mission.Valid {
		p.Mission = &mission.String
	}
	if assets.Valid {
		p.Assets = &assets.String
	}
	if adversaries.Valid {
		p.Adversaries = &adversaries.String
	}
	if threats.Valid {
		p.Threats = &threats.String
	}
	if risks.Valid {
		p.Risks = &risks.String
	}
	if mitigations.Valid {
		p.Mitigations = &mitigations.String
	}

	p.CreatedAt, _ = time.Parse(time.RFC3339, createdAt)
	p.UpdatedAt, _ = time.Parse(time.RFC3339, updatedAt)

	return &p, nil
}

func (db *DB) ListProfiles() ([]Profile, error) {
	rows, err := db.conn.Query(`
		SELECT id, name, description, mission, assets, adversaries, threats, risks, mitigations, created_at, updated_at
		FROM profiles ORDER BY updated_at DESC
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to list profiles: %w", err)
	}
	defer rows.Close()

	var profiles []Profile
	for rows.Next() {
		var p Profile
		var createdAt, updatedAt string
		var description, mission, assets, adversaries, threats, risks, mitigations sql.NullString

		err := rows.Scan(&p.ID, &p.Name, &description, &mission, &assets, &adversaries, &threats, &risks, &mitigations, &createdAt, &updatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan profile: %w", err)
		}

		p.Description = description.String
		if mission.Valid {
			p.Mission = &mission.String
		}
		if assets.Valid {
			p.Assets = &assets.String
		}
		if adversaries.Valid {
			p.Adversaries = &adversaries.String
		}
		if threats.Valid {
			p.Threats = &threats.String
		}
		if risks.Valid {
			p.Risks = &risks.String
		}
		if mitigations.Valid {
			p.Mitigations = &mitigations.String
		}

		p.CreatedAt, _ = time.Parse(time.RFC3339, createdAt)
		p.UpdatedAt, _ = time.Parse(time.RFC3339, updatedAt)

		profiles = append(profiles, p)
	}

	return profiles, nil
}

func (db *DB) DeleteProfile(id string) error {
	result, err := db.conn.Exec(`DELETE FROM profiles WHERE id = ?`, id)
	if err != nil {
		return fmt.Errorf("failed to delete profile: %w", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (db *DB) GetSection(profileID, section string) (*string, error) {
	var value sql.NullString
	query := fmt.Sprintf(`SELECT %s FROM profiles WHERE id = ?`, section)

	err := db.conn.QueryRow(query, profileID).Scan(&value)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get section: %w", err)
	}

	if value.Valid {
		return &value.String, nil
	}
	return nil, nil
}

func (db *DB) UpdateSection(profileID, section, data string) error {
	now := time.Now().UTC().Format(time.RFC3339)
	query := fmt.Sprintf(`UPDATE profiles SET %s = ?, updated_at = ? WHERE id = ?`, section)

	result, err := db.conn.Exec(query, data, now, profileID)
	if err != nil {
		return fmt.Errorf("failed to update section: %w", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

var ValidSections = map[string]bool{
	"mission":     true,
	"assets":      true,
	"adversaries": true,
	"threats":     true,
	"risks":       true,
	"mitigations": true,
}

func IsValidSection(section string) bool {
	return ValidSections[section]
}

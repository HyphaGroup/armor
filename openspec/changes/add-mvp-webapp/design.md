# Design: MVP Web Application

## Context

This is a proof-of-concept for the ARMOR platform. The goal is to validate the core workflow (create/edit threat model profiles) with minimal infrastructure. We're intentionally deferring auth, encryption, multi-org, and agent integration to keep scope tight.

**Stakeholders:** Civil society organizations who need accessible threat modeling tools.

**Constraints:**
- Single developer building MVP
- Must work locally without external services
- Must validate against existing JSON schemas in `schemas/`

## Goals / Non-Goals

**Goals:**
- Prove end-to-end profile creation and editing works
- Validate form UX for complex nested data (assets, risks, etc.)
- Test JSON schema validation in practice
- Establish project structure for future expansion

**Non-Goals:**
- Production-ready security (no encryption, simple password)
- Multi-user or multi-org support
- Agent/AI integration
- Mobile optimization

## Architecture

```
┌─────────────────────────────────────────────────────────┐
│                    Browser                              │
│  ┌───────────────────────────────────────────────────┐ │
│  │              SvelteKit Frontend                    │ │
│  │  - Password gate (client-side check)              │ │
│  │  - Section forms with validation                  │ │
│  │  - Completeness display                           │ │
│  └───────────────────────┬───────────────────────────┘ │
└──────────────────────────┼──────────────────────────────┘
                           │ HTTP/JSON
                           ▼
┌─────────────────────────────────────────────────────────┐
│                  Go Backend (:8080)                     │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────────┐ │
│  │  API Router │  │  Validator  │  │  Profile Store  │ │
│  │  (net/http) │  │  (jsonschema)│  │  (SQLite)      │ │
│  └─────────────┘  └─────────────┘  └─────────────────┘ │
└─────────────────────────────────────────────────────────┘
```

## Decisions

### Decision 1: Single-file SQLite, no encryption

**What:** Use plain SQLite without SQLCipher for MVP.

**Why:** Simplifies build process (no CGO/SQLCipher dependencies), faster iteration. Encryption is important for production but not for local POC.

**Alternatives considered:**
- SQLCipher: Adds build complexity, deferred to post-MVP
- In-memory: Loses data on restart, not useful for testing workflows

### Decision 2: Runtime password via environment variable

**What:** Single password set via `ARMOR_PASSWORD` env var, checked on API requests.

**Why:** Minimal protection for local use without full auth system. Frontend stores password in sessionStorage after initial entry.

**Implementation:**
- Backend: Middleware checks `Authorization: Bearer <password>` header
- Frontend: Password prompt on load, stores in sessionStorage, includes in all API calls

### Decision 3: Multiple profiles with JSON columns

**What:** A `profiles` table supporting multiple profiles, each section is a JSON column.

**Why:** Allows creating and managing multiple threat model profiles. Dashboard shows list of all profiles with completeness.

**Schema:**
```sql
CREATE TABLE profiles (
  id TEXT PRIMARY KEY,       -- UUID
  name TEXT NOT NULL,        -- Display name for the profile
  description TEXT,          -- Optional description
  mission TEXT,              -- JSON
  assets TEXT,               -- JSON
  adversaries TEXT,          -- JSON
  threats TEXT,              -- JSON
  risks TEXT,                -- JSON
  mitigations TEXT,          -- JSON
  created_at TEXT NOT NULL,
  updated_at TEXT NOT NULL
);
```

### Decision 4: Completeness as field-fill percentage

**What:** Count non-null/non-empty required fields divided by total required fields per section.

**Why:** Simple to implement, gives users directional feedback. Semantic completeness (cross-references) deferred to post-MVP.

### Decision 5: Validation on backend only

**What:** Frontend does basic required checks, backend does full JSON schema validation.

**Why:** Single source of truth for validation. Frontend validation is UX convenience, backend is authoritative.

## API Design

```
# Profile list
GET  /api/profiles                      # List all profiles with completeness
POST /api/profiles                      # Create new profile (name, description)

# Single profile
GET  /api/profiles/:id                  # Get full profile (all sections)
DELETE /api/profiles/:id                # Delete profile

# Profile sections
GET  /api/profiles/:id/:section         # Get single section
PUT  /api/profiles/:id/:section         # Update single section (validated)
```

All endpoints require `Authorization: Bearer <password>` header.

## Frontend Routes

```
/                              # Password gate → redirect to /profiles
/profiles                      # Dashboard: list of profiles with completeness
/profiles/new                  # Create new profile form
/profiles/:id                  # Profile overview with section cards
/profiles/:id/mission          # Mission form
/profiles/:id/assets           # Assets form (list + item editing)
/profiles/:id/adversaries      # Adversaries form
/profiles/:id/threats          # Threats form
/profiles/:id/risks            # Risks form
/profiles/:id/mitigations      # Mitigations form
```

## Risks / Trade-offs

| Risk | Mitigation |
|------|------------|
| Password in sessionStorage is not secure | Acceptable for local MVP; real auth in future |
| No data backup/export | Add export endpoint if time permits |
| Complex nested forms may be slow to build | Start with Mission (simplest), iterate |
| Schema validation errors may be unclear | Return detailed error paths from backend |

## Open Questions

1. Should we seed with example data for demo purposes?
2. Do we need a "save draft" vs "save final" distinction, or just auto-save?
3. Should completeness show on dashboard only or inline in forms?

# ARMOR - Adversary Risk Modeling for Organizational Resilience

A threat modeling platform for civil society organizations.

## Quick Start

### Prerequisites

- Go 1.21+
- Node.js 22+
- npm

### Backend

```bash
cd server

# Build
go build -o armor-server ./cmd/armor-server

# Run (default password: "armor")
ARMOR_PASSWORD=your-password ./armor-server -schemas ../schemas

# Server runs on http://localhost:8080
```

### Frontend

```bash
cd web

# Install dependencies
npm install

# Run development server
npm run dev

# Frontend runs on http://localhost:5173
```

### Configuration

| Environment Variable | Description | Default |
|---------------------|-------------|---------|
| `ARMOR_PASSWORD` | Access password | `armor` |
| `ARMOR_PORT` | Server port | `8080` |
| `ARMOR_DB_PATH` | SQLite database path | `./armor.db` |
| `ARMOR_SCHEMAS_DIR` | JSON schemas directory | `../schemas` |

## Project Structure

```
├── server/           # Go backend
│   ├── cmd/          # Entry points
│   └── internal/     # Internal packages
│       ├── api/      # HTTP handlers
│       ├── db/       # Database layer
│       └── validator/# JSON schema validation
├── web/              # SvelteKit frontend
│   └── src/
│       ├── lib/      # Shared code
│       └── routes/   # Pages
├── schemas/          # JSON schemas for profile sections
└── docs/             # Documentation
```

## API Endpoints

```
GET    /api/profiles              # List all profiles
POST   /api/profiles              # Create profile
GET    /api/profiles/:id          # Get profile
DELETE /api/profiles/:id          # Delete profile
GET    /api/profiles/:id/:section # Get section
PUT    /api/profiles/:id/:section # Update section
```

All endpoints require `Authorization: Bearer <password>` header.

## Profile Sections

1. **Mission** - Organization mission and impact areas
2. **Assets** - Information assets requiring protection
3. **Adversaries** - Potential threat actors
4. **Threats** - Specific threats to the organization
5. **Risks** - Risk scenarios with scoring
6. **Mitigations** - Planned security improvements

## License

MIT

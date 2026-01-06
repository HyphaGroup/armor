# Project Context

## Purpose

ARMOR (Adversary Risk Modeling for Organizational Resilience) is a threat modeling platform for civil society organizations - human rights groups, journalists, advocacy organizations, and legal aid providers who face sophisticated threats from well-resourced adversaries while operating with limited security budgets.

The platform enables organizations to:
- Build and maintain structured threat model profiles through guided workflows
- Track security incidents and link them to profile updates
- Receive AI agent assistance for profile completion and analysis
- Review and approve agent-proposed changes before they're applied

### Core Methodology

ARMOR follows a logical chain where each element informs the next:
```
MISSION → ASSETS → ADVERSARIES → THREATS → RISKS → MITIGATIONS
```

**Six Core Components:**
1. Mission & Impact Framework - Why security matters for THIS organization
2. Asset Identification - What needs protection and where it lives
3. Adversary Profiling - Who might target you and why
4. Threat Mapping - Specific threats relevant to your situation
5. Risk Assessment - Prioritized risk scenarios with scoring
6. Mitigation Planning - Actionable improvements

**Optional Modules:**
- Deep Adversary Profiling
- Information Operations
- OPSEC Analysis
- Incident Response
- Technical Deep-Dive

## Tech Stack

### Backend (armor-server)
- **Language:** Go
- **Database:** SQLite with SQLCipher (encrypted at rest)
- **MCP Server:** github.com/mark3labs/mcp-go
- **SQLite:** github.com/mutecomm/go-sqlcipher

### Frontend (armor-web)
- **Framework:** SvelteKit
- **Language:** TypeScript
- **Styling:** TailwindCSS
- **Forms:** Superforms + Zod validation
- **Types:** Generated from JSON schemas

### Infrastructure
- **Tunnel:** cloudflared for exposure
- **Single binary:** Go server serves both REST API and MCP

### Key Files
- `platform-spec.md` - Detailed technical specification
- `framework.md` - Conceptual framework and principles
- `methodology.md` - Step-by-step threat modeling process
- `schemas/*.json` - JSON schemas for all profile sections

## Project Conventions

### Code Style

**Go:**
- Standard `gofmt` formatting
- Package names: lowercase, single word (e.g., `core`, `api`, `mcp`)
- Interfaces: verb-noun (e.g., `ProfileStore`, `ProposalService`)
- Error handling: Return errors, don't panic; wrap with context
- Comments: Godoc style for exported functions

**TypeScript/Svelte:**
- Prettier formatting
- Types over interfaces for data shapes
- Svelte stores for global state
- Component naming: PascalCase (e.g., `MissionForm.svelte`)

**General:**
- No abbreviations in names unless domain-standard
- Explicit over implicit
- Prefer composition over inheritance

### Architecture Patterns

**Backend:**
- **Function Layer:** Core business logic in `internal/core/`, called by both REST and MCP
- **Repository Pattern:** Database access abstracted behind interfaces
- **Handler Pattern:** HTTP handlers are thin wrappers calling core functions
- **JSON Columns:** Profile sections stored as JSON, validated in application layer

**Frontend:**
- **Route-based Pages:** SvelteKit file-based routing
- **Component Composition:** Small, focused components composed into pages
- **Store Pattern:** Svelte stores for cross-component state (auth, org, profile)
- **API Client:** Single typed client class for all REST calls

**Data Flow:**
```
User/Agent → REST API/MCP → Function Layer → SQLite
                              ↓
                    JSON Schema Validation
```

### Testing Strategy

**Backend:**
- Unit tests for core functions with mocked DB
- Integration tests for API handlers with test database
- Table-driven tests for validation logic

**Frontend:**
- Component tests with Svelte Testing Library
- E2E tests with Playwright for critical flows
- Visual regression tests for form components

**Coverage Priorities:**
1. Risk score calculations (must be correct)
2. Proposal workflow (accept/reject/preview)
3. Schema validation (all sections)
4. Cross-reference integrity checks

### Git Workflow

- **Main branch:** `main` - always deployable
- **Feature branches:** `feature/[change-id]` matching openspec change IDs
- **Commit messages:** Imperative mood, reference change ID
  - `Add mission form validation (add-mission-form)`
  - `Fix risk score calculation`
- **PRs:** Reference openspec proposal, include test plan

## Domain Context

### Threat Modeling Concepts

**Assets:** Information or systems that need protection. Have value ratings (Critical/High/Medium/Low) that map to risk scores (3/2/2/1).

**Adversaries:** Who might target the organization. Categories:
- Nation-State / Intelligence
- Ideological Opposition  
- Cybercriminal
- Insider Threat
- Competitor / Opposing Org
- Opportunistic

**Threats:** What adversaries could do. Domains:
- Account & Access (phishing, account takeover)
- Data & Information (breach, tampering, surveillance)
- Disruption (DDoS, infrastructure attacks)
- Information & Reputation (narrative attacks, harassment)
- Physical (intrusion, surveillance, violence)
- Operational (human error, technical failure)

**Risk Scoring:** Three-factor model
```
RISK SCORE = ASSET VALUE (1-3) × LIKELIHOOD (1-3) × VULNERABILITY (1-3)
```
Priority bands: Critical (18-27), High (10-17), Moderate (4-9), Low (1-3)

### Proposal System

Agents propose changes, humans approve. This ensures:
- No autonomous modification of sensitive security data
- Audit trail of all changes
- Human review of AI suggestions

Proposal states: `pending` → `accepted`/`rejected`

### Profile Completeness

Completeness is semantic, not just field fill:
- Each critical asset has at least one threat mapped
- Each threat has likelihood assessed
- Each adversary has at least one linked threat
- Top risks have mitigations planned
- Support network contacts are established (not just identified)

## Important Constraints

### Security
- All data encrypted at rest (SQLCipher)
- Database key via environment variable, never in code
- Sensitive fields (source data, beneficiary info) flagged in schemas
- Audit log for all profile changes
- Agent proposals require human approval

### Privacy
- Organizations own their data
- No cross-organization data sharing without explicit consent
- Personal safety info (staff addresses, etc.) requires extra protection

### Usability
- Organizations have limited security expertise
- Interface must be accessible to non-technical users
- Agent guidance should use plain language, not jargon
- Methodology content embedded in UI (help text, examples)

### Scale
- Single-org deployments supported (not multi-tenant SaaS initially)
- Typical org: 1-50 users, 1 profile, dozens of assets/threats/risks
- SQLite sufficient for expected scale

## External Dependencies

### Go Packages
- `github.com/mutecomm/go-sqlcipher` - Encrypted SQLite
- `github.com/mark3labs/mcp-go` - MCP server SDK
- `github.com/santhosh-tekuri/jsonschema` - JSON Schema validation (or similar)

### Frontend Packages
- `@sveltejs/kit` - Application framework
- `sveltekit-superforms` - Form handling
- `zod` - Runtime validation
- `tailwindcss` - Styling

### External Services
- None required for core functionality
- Optional: cloudflared for tunnel exposure

### Reference Materials
- JSON Schemas in `schemas/` directory
- Adversary templates in `schemas/adversary-templates.json`
- DISARM framework subset in `schemas/disarm-civil-society-subset.json`

# ARMOR Platform Specification

Technical specification for the ARMOR digital platform - a system for capturing, maintaining, and completing civil society threat model profiles.

---

## Overview

The ARMOR platform enables organizations to build and maintain threat model profiles through two interfaces:

1. **Web Application** - Forms-based UI for direct viewing and editing of profile data
2. **MCP Server** - Tool interface for AI agents to read, analyze, and propose profile updates

Both interfaces share a common backend function layer and data store.

```
┌────────────────────┐     ┌────────────────────┐
│   Web Application  │     │   AI Agent (LLM)   │
│   (Browser)        │     │   (Claude, etc.)   │
└─────────┬──────────┘     └─────────┬──────────┘
          │                          │
          ▼                          ▼
┌─────────────────────────────────────────────────┐
│                 armor-server (Go)               │
│  ┌──────────────────┐  ┌──────────────────────┐ │
│  │   REST API       │  │   MCP Server         │ │
│  │   (HTTP/JSON)    │  │   (mcp-go SDK)       │ │
│  └────────┬─────────┘  └───────────┬──────────┘ │
│           │                        │            │
│           └───────────┬────────────┘            │
│                       ▼                         │
│          ┌────────────────────────┐             │
│          │   Function Layer       │             │
│          │   (Core Operations)    │             │
│          └───────────┬────────────┘             │
│                      │                          │
│                      ▼                          │
│          ┌────────────────────────┐             │
│          │   SQLite + SQLCipher   │             │
│          └────────────────────────┘             │
└─────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────┐
│                 armor-web                       │
│            (static SPA / frontend)              │
│              talks to REST API                  │
└─────────────────────────────────────────────────┘

             ┌─────────────────────┐
             │    cloudflared      │
             │  (tunnel exposure)  │
             └─────────────────────┘
```

---

## Project Structure

```
armor/
├── server/                     # Go backend
│   ├── cmd/
│   │   └── armor-server/
│   │       └── main.go         # Entry point
│   ├── internal/
│   │   ├── db/                 # SQLite/SQLCipher wrapper
│   │   │   ├── db.go
│   │   │   ├── migrations.go
│   │   │   └── schema.sql
│   │   ├── core/               # Function layer (business logic)
│   │   │   ├── organizations.go
│   │   │   ├── users.go
│   │   │   ├── profiles.go
│   │   │   ├── proposals.go
│   │   │   └── analysis.go
│   │   ├── api/                # REST API handlers
│   │   │   ├── router.go
│   │   │   ├── auth.go
│   │   │   ├── organizations.go
│   │   │   ├── profiles.go
│   │   │   └── middleware.go
│   │   ├── mcp/                # MCP server implementation
│   │   │   ├── server.go
│   │   │   ├── tools.go
│   │   │   └── resources.go
│   │   └── validation/         # JSON schema validation
│   │       └── validator.go
│   ├── go.mod
│   └── go.sum
│
├── web/                        # Frontend (SvelteKit)
│   ├── src/
│   │   ├── lib/
│   │   │   ├── components/     # Reusable UI components
│   │   │   │   ├── forms/      # Form components per section
│   │   │   │   ├── ui/         # Generic UI (buttons, modals, etc.)
│   │   │   │   └── layout/     # Nav, sidebar, etc.
│   │   │   ├── api.ts          # REST API client
│   │   │   ├── stores/         # Svelte stores (auth, org context)
│   │   │   └── types.ts        # TypeScript types (from schemas)
│   │   ├── routes/
│   │   │   ├── +layout.svelte
│   │   │   ├── +page.svelte                    # Landing / login
│   │   │   ├── [org]/
│   │   │   │   ├── +layout.svelte              # Org context wrapper
│   │   │   │   ├── +page.svelte                # Dashboard
│   │   │   │   ├── profile/
│   │   │   │   │   ├── +page.svelte            # Profile overview
│   │   │   │   │   ├── mission/+page.svelte
│   │   │   │   │   ├── assets/+page.svelte
│   │   │   │   │   ├── adversaries/+page.svelte
│   │   │   │   │   ├── threats/+page.svelte
│   │   │   │   │   ├── risks/+page.svelte
│   │   │   │   │   └── mitigations/+page.svelte
│   │   │   │   ├── modules/
│   │   │   │   │   └── [module]/+page.svelte
│   │   │   │   ├── proposals/+page.svelte      # Review agent proposals
│   │   │   │   ├── incidents/
│   │   │   │   │   ├── +page.svelte            # Incident log
│   │   │   │   │   └── [id]/+page.svelte       # Incident detail
│   │   │   │   ├── conversations/
│   │   │   │   │   ├── +page.svelte            # Agent conversation history
│   │   │   │   │   └── [id]/+page.svelte       # Conversation detail
│   │   │   │   ├── history/+page.svelte        # Profile change log
│   │   │   │   └── settings/+page.svelte       # Org settings, members
│   │   │   └── api/                            # API routes (if needed for BFF)
│   │   └── app.html
│   ├── static/
│   ├── package.json
│   ├── svelte.config.js
│   ├── tsconfig.json
│   └── vite.config.ts
│
├── schemas/                    # JSON schemas (existing)
│   ├── mission.schema.json
│   ├── assets.schema.json
│   └── ...
│
└── docs/                       # Documentation
    ├── framework.md
    ├── methodology.md
    └── platform-spec.md
```

### Build & Run

```bash
# Build server
cd server
go build -o armor-server ./cmd/armor-server

# Run (development)
ARMOR_DB_PATH=./data/armor.db \
ARMOR_DB_KEY=<key> \
./armor-server

# Expose via cloudflared (if needed)
cloudflared tunnel --url http://localhost:8080
```

### Configuration

Environment variables:

| Variable | Description | Default |
|----------|-------------|---------|
| `ARMOR_DB_PATH` | Path to SQLite database file | `./armor.db` |
| `ARMOR_DB_KEY` | SQLCipher encryption key (base64) | Required |
| `ARMOR_API_PORT` | REST API port | `8080` |
| `ARMOR_MCP_PORT` | MCP server port | `8081` |
| `ARMOR_LOG_LEVEL` | Log level (debug/info/warn/error) | `info` |

---

## Data Model

### Entities

#### Organization

The top-level entity. Each organization has one threat model profile.

```typescript
interface Organization {
  id: string;                    // UUID
  name: string;
  slug: string;                  // URL-friendly identifier
  createdAt: DateTime;
  updatedAt: DateTime;
}
```

#### User

Users belong to one or more organizations with specific roles.

```typescript
interface User {
  id: string;                    // UUID
  email: string;
  name: string;
  createdAt: DateTime;
  lastLoginAt: DateTime;
}

interface OrganizationMembership {
  userId: string;
  organizationId: string;
  role: 'owner' | 'admin' | 'editor' | 'viewer';
  createdAt: DateTime;
}
```

#### Profile

The threat model profile, composed of sections matching ARMOR methodology components.

```typescript
interface Profile {
  organizationId: string;        // Foreign key
  version: number;               // Increments on each change
  updatedAt: DateTime;
  updatedBy: string;             // User ID or 'agent'
  
  // Core components (each validated against respective schema)
  mission: MissionData | null;
  assets: AssetsData | null;
  adversaries: AdversariesData | null;
  threats: ThreatsData | null;
  risks: RisksData | null;
  mitigations: MitigationsData | null;
  
  // Optional modules
  deepAdversaryProfiling: DeepAdversaryProfilingData | null;
  informationOperations: InformationOperationsData | null;
  opsec: OpsecData | null;
  responseCapability: ResponseCapabilityData | null;
  technicalDeepDive: TechnicalDeepDiveData | null;
}
```

#### Profile History

Audit log of all profile changes.

```typescript
interface ProfileChange {
  id: string;
  organizationId: string;
  version: number;               // Profile version after this change
  timestamp: DateTime;
  userId: string | null;         // Null if agent-initiated
  agentSessionId: string | null; // Null if user-initiated
  changeType: 'create' | 'update' | 'delete';
  section: string;               // e.g., 'mission', 'assets'
  path: string;                  // JSON path within section
  previousValue: any;
  newValue: any;
  source: 'web' | 'agent' | 'api' | 'import';
}
```

#### Incident

Security incidents logged by the organization.

```typescript
interface Incident {
  id: string;
  organizationId: string;
  title: string;
  description: string;
  type: 'phishing' | 'account_compromise' | 'data_breach' | 'malware' | 
        'physical' | 'harassment' | 'disinformation' | 'other';
  severity: 'critical' | 'high' | 'medium' | 'low';
  status: 'open' | 'investigating' | 'resolved' | 'closed';
  occurredAt: DateTime;
  discoveredAt: DateTime;
  resolvedAt: DateTime | null;
  createdAt: DateTime;
  createdBy: string;             // User ID
  
  // Optional links
  affectedAssets: string[];      // Asset IDs from profile
  relatedThreats: string[];      // Threat IDs from profile
  linkedProposals: string[];     // Proposals generated from this incident
  
  // Response tracking
  responseActions: ResponseAction[];
  lessonsLearned: string | null;
}

interface ResponseAction {
  id: string;
  description: string;
  assignedTo: string | null;
  status: 'pending' | 'in_progress' | 'completed';
  completedAt: DateTime | null;
}
```

---

## Function Layer

The function layer provides all operations on the data store. Both the REST API and MCP server call these functions.

### Organization Functions

```typescript
// Create a new organization
createOrganization(params: {
  name: string;
  slug: string;
  createdBy: string;  // User ID of creator (becomes owner)
}): Promise<Organization>

// Get organization by ID or slug
getOrganization(params: {
  id?: string;
  slug?: string;
}): Promise<Organization | null>

// List organizations for a user
listUserOrganizations(params: {
  userId: string;
}): Promise<Array<Organization & { role: Role }>>

// Update organization metadata
updateOrganization(params: {
  id: string;
  name?: string;
  slug?: string;
}): Promise<Organization>
```

### User Functions

```typescript
// Create user (typically via auth provider)
createUser(params: {
  email: string;
  name: string;
}): Promise<User>

// Get user by ID or email
getUser(params: {
  id?: string;
  email?: string;
}): Promise<User | null>

// Add user to organization
addOrganizationMember(params: {
  organizationId: string;
  userId: string;
  role: Role;
}): Promise<OrganizationMembership>

// Update member role
updateOrganizationMember(params: {
  organizationId: string;
  userId: string;
  role: Role;
}): Promise<OrganizationMembership>

// Remove user from organization
removeOrganizationMember(params: {
  organizationId: string;
  userId: string;
}): Promise<void>

// List organization members
listOrganizationMembers(params: {
  organizationId: string;
}): Promise<Array<User & { role: Role }>>
```

### Profile Functions

#### Read Operations

```typescript
// Get full profile for an organization
getProfile(params: {
  organizationId: string;
}): Promise<Profile>

// Get specific section of profile
getProfileSection(params: {
  organizationId: string;
  section: ProfileSection;
}): Promise<SectionData | null>

// Get profile completeness metrics
getProfileCompleteness(params: {
  organizationId: string;
}): Promise<{
  overall: number;                    // 0-100 percentage
  sections: Record<ProfileSection, {
    complete: number;                 // 0-100 percentage
    requiredFields: number;
    completedFields: number;
    missingFields: string[];          // Field paths that are empty
  }>;
}>

// Get profile history
getProfileHistory(params: {
  organizationId: string;
  section?: ProfileSection;           // Filter by section
  since?: DateTime;
  limit?: number;
}): Promise<ProfileChange[]>
```

#### Write Operations

```typescript
// Update a profile section (full replacement)
updateProfileSection(params: {
  organizationId: string;
  section: ProfileSection;
  data: SectionData;
  updatedBy: string;                  // User ID
  source: 'web' | 'api' | 'import';
}): Promise<{
  profile: Profile;
  validationErrors: ValidationError[];  // Empty if valid
}>

// Patch a profile section (partial update)
patchProfileSection(params: {
  organizationId: string;
  section: ProfileSection;
  patch: Partial<SectionData>;        // Only fields to update
  updatedBy: string;
  source: 'web' | 'api' | 'import';
}): Promise<{
  profile: Profile;
  validationErrors: ValidationError[];
}>

// Propose a profile change (for agent use - not auto-committed)
proposeProfileChange(params: {
  organizationId: string;
  section: ProfileSection;
  patch: Partial<SectionData>;
  agentSessionId: string;
  rationale: string;                  // Agent's explanation
}): Promise<{
  proposalId: string;
  preview: SectionData;               // What section would look like
  diff: Diff;                         // Visual diff
  validationErrors: ValidationError[];
}>

// Accept a proposed change
acceptProposal(params: {
  proposalId: string;
  acceptedBy: string;                 // User ID
}): Promise<Profile>

// Reject a proposed change
rejectProposal(params: {
  proposalId: string;
  rejectedBy: string;
  reason?: string;
}): Promise<void>

// List pending proposals
listPendingProposals(params: {
  organizationId: string;
}): Promise<Proposal[]>
```

#### Validation

```typescript
// Validate data against schema without saving
validateSection(params: {
  section: ProfileSection;
  data: SectionData;
}): Promise<{
  valid: boolean;
  errors: ValidationError[];
}>

interface ValidationError {
  path: string;                       // JSON path to invalid field
  message: string;
  constraint: string;                 // Schema constraint that failed
}
```

### Incident Functions

```typescript
// Create a new incident
createIncident(params: {
  organizationId: string;
  title: string;
  description: string;
  type: IncidentType;
  severity: Severity;
  occurredAt: DateTime;
  discoveredAt: DateTime;
  createdBy: string;
  affectedAssets?: string[];
  relatedThreats?: string[];
}): Promise<Incident>

// Get incident by ID
getIncident(params: {
  id: string;
}): Promise<Incident | null>

// List incidents for organization
listIncidents(params: {
  organizationId: string;
  status?: IncidentStatus[];
  severity?: Severity[];
  since?: DateTime;
  limit?: number;
}): Promise<Incident[]>

// Update incident
updateIncident(params: {
  id: string;
  title?: string;
  description?: string;
  severity?: Severity;
  status?: IncidentStatus;
  resolvedAt?: DateTime;
  lessonsLearned?: string;
  affectedAssets?: string[];
  relatedThreats?: string[];
}): Promise<Incident>

// Add response action to incident
addResponseAction(params: {
  incidentId: string;
  description: string;
  assignedTo?: string;
}): Promise<ResponseAction>

// Update response action
updateResponseAction(params: {
  incidentId: string;
  actionId: string;
  status?: ActionStatus;
  completedAt?: DateTime;
}): Promise<ResponseAction>

// Link proposal to incident
linkProposalToIncident(params: {
  incidentId: string;
  proposalId: string;
}): Promise<void>
```

### Analysis Functions

Functions that analyze profile data (primarily for agent use).

```typescript
// Get suggested next steps based on current profile state
getSuggestedNextSteps(params: {
  organizationId: string;
}): Promise<{
  priority: 'high' | 'medium' | 'low';
  section: ProfileSection;
  suggestion: string;
  reason: string;
}[]>

// Identify inconsistencies or gaps in profile
analyzeProfileGaps(params: {
  organizationId: string;
}): Promise<{
  gaps: Array<{
    type: 'missing' | 'inconsistent' | 'outdated';
    section: ProfileSection;
    description: string;
    suggestion: string;
  }>;
}>

// Cross-reference: find assets not covered by threats, etc.
analyzeCoverage(params: {
  organizationId: string;
}): Promise<{
  assetsWithoutThreats: string[];     // Asset IDs
  threatsWithoutRisks: string[];      // Threat IDs
  risksWithoutMitigations: string[];  // Risk IDs
  adversariesWithoutThreats: string[];
}>
```

---

## REST API (Web Application)

HTTP API for the web application frontend.

### Authentication

Uses session-based authentication with OAuth provider (Google, GitHub, etc.) or email magic links.

```
POST /auth/login
POST /auth/logout
GET  /auth/me
```

### Organizations

```
GET    /api/organizations                    # List user's orgs
POST   /api/organizations                    # Create org
GET    /api/organizations/:slug              # Get org
PATCH  /api/organizations/:slug              # Update org
DELETE /api/organizations/:slug              # Delete org

GET    /api/organizations/:slug/members      # List members
POST   /api/organizations/:slug/members      # Add member
PATCH  /api/organizations/:slug/members/:id  # Update role
DELETE /api/organizations/:slug/members/:id  # Remove member
```

### Profile

```
GET    /api/organizations/:slug/profile                    # Full profile
GET    /api/organizations/:slug/profile/completeness       # Completeness metrics
GET    /api/organizations/:slug/profile/history            # Change history

GET    /api/organizations/:slug/profile/:section           # Get section
PUT    /api/organizations/:slug/profile/:section           # Replace section
PATCH  /api/organizations/:slug/profile/:section           # Partial update

POST   /api/organizations/:slug/profile/:section/validate  # Validate without saving

GET    /api/organizations/:slug/proposals                  # List pending proposals
POST   /api/organizations/:slug/proposals/:id/accept       # Accept proposal
POST   /api/organizations/:slug/proposals/:id/reject       # Reject proposal
```

### Incidents

```
GET    /api/organizations/:slug/incidents            # List incidents
POST   /api/organizations/:slug/incidents            # Create incident
GET    /api/organizations/:slug/incidents/:id        # Get incident
PATCH  /api/organizations/:slug/incidents/:id        # Update incident
DELETE /api/organizations/:slug/incidents/:id        # Delete incident

POST   /api/organizations/:slug/incidents/:id/actions          # Add response action
PATCH  /api/organizations/:slug/incidents/:id/actions/:aid     # Update action
POST   /api/organizations/:slug/incidents/:id/link-proposal    # Link proposal
```

### Export/Import

```
GET    /api/organizations/:slug/export          # Export full profile as JSON
POST   /api/organizations/:slug/import          # Import profile from JSON
GET    /api/organizations/:slug/export/pdf      # Export as PDF report
```

---

## MCP Server (Agent Interface)

Model Context Protocol server exposing profile operations as tools for AI agents.

### Connection & Authentication

Agents connect via MCP protocol. Authentication via API key scoped to an organization.

```typescript
// Agent establishes session with org context
interface AgentSession {
  sessionId: string;
  organizationId: string;
  startedAt: DateTime;
  permissions: 'read' | 'read-propose' | 'read-write';
}
```

### MCP Tools

#### Read Tools

```typescript
// Get organization overview and profile summary
tool: armor_get_overview
params: {}
returns: {
  organization: { name, slug },
  profileCompleteness: { overall, sections },
  lastUpdated: DateTime,
  pendingProposals: number
}

// Get specific profile section
tool: armor_get_section
params: {
  section: ProfileSection
}
returns: {
  data: SectionData | null,
  completeness: { complete, missingFields },
  lastUpdated: DateTime
}

// Get full profile (use sparingly - large payload)
tool: armor_get_full_profile
params: {}
returns: Profile

// Get what's missing or needs attention
tool: armor_get_gaps
params: {}
returns: {
  gaps: Gap[],
  suggestedNextSteps: Suggestion[]
}

// Get coverage analysis
tool: armor_analyze_coverage
params: {}
returns: {
  assetsWithoutThreats: Asset[],
  threatsWithoutRisks: Threat[],
  risksWithoutMitigations: Risk[],
  // ... etc
}
```

#### Write Tools

```typescript
// Propose a change to a section (requires user approval)
tool: armor_propose_change
params: {
  section: ProfileSection,
  changes: Partial<SectionData>,
  rationale: string              // Explain why this change
}
returns: {
  proposalId: string,
  preview: SectionData,
  diff: string,                  // Human-readable diff
  validationResult: { valid, errors }
}

// Add a single item to a list within a section
tool: armor_propose_add_item
params: {
  section: ProfileSection,
  listPath: string,              // e.g., "assets" or "adversaries"
  item: object,
  rationale: string
}
returns: {
  proposalId: string,
  item: object,                  // With generated ID
  validationResult: { valid, errors }
}

// Validate proposed data without creating proposal
tool: armor_validate
params: {
  section: ProfileSection,
  data: Partial<SectionData>
}
returns: {
  valid: boolean,
  errors: ValidationError[]
}
```

#### Incident Tools

```typescript
// List incidents
tool: armor_list_incidents
params: {
  status?: IncidentStatus[];     // Filter by status
  severity?: Severity[];         // Filter by severity
  limit?: number;
}
returns: {
  incidents: Incident[];
  openCount: number;
  recentCount: number;           // Last 30 days
}

// Get incident details
tool: armor_get_incident
params: {
  id: string;
}
returns: Incident

// Suggest profile updates based on incident
tool: armor_suggest_incident_updates
params: {
  incidentId: string;
}
returns: {
  suggestions: Array<{
    section: ProfileSection;
    field: string;
    currentValue: any;
    suggestedChange: string;
    rationale: string;
  }>;
}
```

#### Guidance Tools

```typescript
// Get guidance for completing a section
tool: armor_get_section_guidance
params: {
  section: ProfileSection
}
returns: {
  description: string,           // What this section captures
  methodology: string,           // How to approach it (from methodology.md)
  requiredFields: FieldInfo[],
  optionalFields: FieldInfo[],
  examples: Example[],
  tips: string[]
}

// Get questions to ask user to complete a section
tool: armor_get_interview_questions
params: {
  section: ProfileSection,
  currentData: SectionData | null  // To skip already-answered
}
returns: {
  questions: Array<{
    id: string,
    question: string,
    fieldPath: string,           // Where answer maps to in schema
    answerType: 'text' | 'select' | 'multiselect' | 'scale',
    options?: string[],          // For select types
    required: boolean,
    followUp?: string            // Conditional follow-up question
  }>
}

// Map a user's natural language answer to structured data
tool: armor_parse_answer
params: {
  section: ProfileSection,
  fieldPath: string,
  userResponse: string
}
returns: {
  parsed: any,                   // Structured data
  confidence: number,            // 0-1
  clarificationNeeded?: string   // If confidence low
}
```

### MCP Resources

Expose schemas and methodology as readable resources.

```typescript
// Profile schemas
resource: armor://schemas/mission
resource: armor://schemas/assets
resource: armor://schemas/adversaries
// ... etc

// Methodology content
resource: armor://methodology/overview
resource: armor://methodology/mission
resource: armor://methodology/assets
// ... etc

// Templates
resource: armor://templates/adversaries  # Pre-built adversary profiles
```

---

## Web Application (SvelteKit)

### Tech Stack

- **SvelteKit** - Full-stack Svelte framework
- **TypeScript** - Type safety, types generated from JSON schemas
- **TailwindCSS** - Utility-first styling
- **Superforms** - Form handling with validation
- **Zod** - Runtime validation (generated from JSON schemas)

### Routes

```
/                                    # Landing / login
/[org]/                              # Dashboard
/[org]/profile                       # Profile overview + completeness
/[org]/profile/[section]             # Section form (mission, assets, etc.)
/[org]/modules/[module]              # Optional module forms
/[org]/proposals                     # Review agent proposals
/[org]/incidents                     # Incident log
/[org]/incidents/[id]                # Incident detail
/[org]/incidents/new                 # Log new incident
/[org]/conversations                 # Agent conversation history
/[org]/conversations/[id]            # Conversation detail + linked proposals
/[org]/history                       # Profile change log
/[org]/settings                      # Org settings, members, API keys
/[org]/export                        # Export options
```

### Key Views

#### Dashboard (`/[org]/`)

Overview of organization's security posture:

- **Completeness ring** - Visual progress for overall profile
- **Section cards** - Status per section with quick-edit links
- **Recent activity** - Timeline of changes (user and agent)
- **Pending proposals** - Alert badge, quick review
- **Open mitigations** - Count and link to track
- **Upcoming reviews** - Next scheduled threat model review

#### Profile Section Form (`/[org]/profile/[section]`)

Guided editing for each ARMOR section:

- **Progress indicator** - Fields completed vs required
- **Form fields** - Mapped from JSON schema
- **Inline validation** - Real-time schema validation
- **Help tooltips** - Methodology guidance per field
- **Examples drawer** - Expandable examples from templates
- **Auto-save** - Draft saved locally, explicit "Save" to commit
- **History sidebar** - Recent changes to this section

#### Proposals Review (`/[org]/proposals`)

Queue of agent-proposed changes:

- **Proposal cards** - Section, summary, timestamp, agent session
- **Diff viewer** - Side-by-side or inline diff
- **Rationale display** - Agent's explanation for the change
- **Accept/Reject** - With optional comment
- **Batch actions** - Select multiple, accept/reject all
- **Link to conversation** - Jump to agent session context

#### Incidents (`/[org]/incidents`)

Security incident tracking:

- **Incident list** - Date, type, severity, status
- **Log new incident** - Quick form: what happened, when, who involved
- **Incident detail** - Full description, timeline, response actions
- **Link to profile** - "Update threat model" button
- **Suggested updates** - Agent-proposed profile changes based on incident

#### Conversations (`/[org]/conversations`)

History of agent interactions:

- **Session list** - Date, duration, tool calls, proposals created
- **Conversation view** - Full message history (if stored)
- **Linked proposals** - Proposals generated in this session
- **Profile changes** - What was updated as result

#### Settings (`/[org]/settings`)

Organization configuration:

- **General** - Name, slug
- **Members** - Invite, roles, remove
- **API Keys** - Create, revoke, permissions
- **Review schedule** - Set reminder cadence
- **Export/Import** - Backup, restore
- **Danger zone** - Delete organization

### Components

#### Forms (`/lib/components/forms/`)

Per-section form components:

```
MissionForm.svelte
AssetsForm.svelte
  └── AssetItemForm.svelte      # Repeatable item
AdversariesForm.svelte
  └── AdversaryItemForm.svelte
ThreatsForm.svelte
RisksForm.svelte
  └── RiskScenarioForm.svelte
MitigationsForm.svelte
  └── MitigationItemForm.svelte
```

#### UI (`/lib/components/ui/`)

Generic reusable components:

```
Button.svelte
Input.svelte
Select.svelte
Textarea.svelte
Modal.svelte
Dropdown.svelte
Badge.svelte
Card.svelte
ProgressRing.svelte
DiffViewer.svelte
Timeline.svelte
Toast.svelte
```

#### Layout (`/lib/components/layout/`)

```
Nav.svelte               # Top navigation
Sidebar.svelte           # Section navigation within org
OrgSwitcher.svelte       # Switch between organizations
UserMenu.svelte          # Profile, logout
```

### State Management

Svelte stores for global state:

```typescript
// stores/auth.ts
export const user = writable<User | null>(null);
export const isAuthenticated = derived(user, $user => !!$user);

// stores/org.ts
export const currentOrg = writable<Organization | null>(null);
export const orgSlug = derived(currentOrg, $org => $org?.slug);

// stores/profile.ts
export const profile = writable<Profile | null>(null);
export const completeness = derived(profile, $p => calculateCompleteness($p));

// stores/proposals.ts
export const pendingProposals = writable<Proposal[]>([]);
export const proposalCount = derived(pendingProposals, $p => $p.length);
```

### API Client

Typed API client for REST endpoints:

```typescript
// lib/api.ts
class ArmorAPI {
  constructor(private baseUrl: string) {}
  
  // Organizations
  async getOrganizations(): Promise<Organization[]>
  async getOrganization(slug: string): Promise<Organization>
  
  // Profile
  async getProfile(orgSlug: string): Promise<Profile>
  async getProfileSection(orgSlug: string, section: Section): Promise<SectionData>
  async updateProfileSection(orgSlug: string, section: Section, data: SectionData): Promise<Profile>
  async getCompleteness(orgSlug: string): Promise<Completeness>
  
  // Proposals
  async getProposals(orgSlug: string): Promise<Proposal[]>
  async acceptProposal(orgSlug: string, proposalId: string): Promise<void>
  async rejectProposal(orgSlug: string, proposalId: string, reason?: string): Promise<void>
  
  // Incidents
  async getIncidents(orgSlug: string): Promise<Incident[]>
  async createIncident(orgSlug: string, data: CreateIncident): Promise<Incident>
  async updateIncident(orgSlug: string, id: string, data: UpdateIncident): Promise<Incident>
  
  // History
  async getHistory(orgSlug: string, options?: HistoryOptions): Promise<ProfileChange[]>
}

export const api = new ArmorAPI(PUBLIC_API_URL);
```

---

## Agent Interaction Patterns

### Pattern 1: Guided Profile Completion

Agent helps complete a section through conversation.

```
Agent: I see your Assets section is only 30% complete. Would you like 
       help identifying your organization's critical assets?

User:  Yes, let's do that.

Agent: [calls armor_get_section(assets)]
       [calls armor_get_interview_questions(assets)]
       
       I'll ask you some questions to identify what needs protection.
       Let's start with your beneficiaries and clients.
       
       Who does your organization serve? What information do you 
       collect or maintain about them?

User:  We work with refugees applying for asylum. We have their 
       application documents, personal histories, and contact info.

Agent: [calls armor_parse_answer(...)]
       [calls armor_propose_add_item(assets, {...})]
       
       I've drafted an asset entry for "Refugee client data" including
       application documents, personal histories, and contact information.
       I've marked it as Critical value since exposure could endanger 
       their safety. 
       
       [Shows preview]
       
       Does this look right? You can review and approve this in the 
       Proposals tab, or I can adjust it.
```

### Pattern 2: Analysis & Recommendations

Agent analyzes existing profile and suggests improvements.

```
Agent: [calls armor_analyze_coverage]
       
       I've analyzed your profile and found some gaps:
       
       - 3 assets have no associated threats mapped
       - Your "Nation-state actor" adversary has no linked threats
       - 2 high-priority risks have no mitigation plans
       
       Would you like to work on any of these?
```

### Pattern 3: Incident-Triggered Update

User reports something that should update the profile.

```
User:  We just had a staff member fall for a phishing email. They 
       clicked a link but we think we caught it in time.

Agent: [calls armor_get_section(threats)]
       [calls armor_get_section(risks)]
       
       I'm sorry to hear that. Let me help you update your threat 
       model based on this incident.
       
       I see you have "Phishing & Social Engineering" listed as a 
       threat with Medium likelihood. Given this incident, should 
       we update it to High?
       
       Also, I'd recommend:
       1. Adding this to your incident history
       2. Reviewing the vulnerability score for email-related risks
       3. Checking if your mitigations need updating
       
       Want me to draft these updates for your review?
```

---

## Technical Implementation Notes

### Data Store: SQLite with Encryption

SQLite with SQLCipher for encryption at rest. Simple, portable, no separate database server.

**Why SQLite:**
- Single-file database, easy backup and portability
- No server process to manage
- Sufficient for expected scale (single org, handful of users)
- JSON1 extension for JSON column queries
- Battle-tested, reliable

**Encryption: SQLCipher**
- Transparent 256-bit AES encryption of entire database file
- Key derived from user-provided passphrase or environment variable
- Encryption at rest protects against disk/backup theft
- Open source, audited implementation

**Schema approach:**
- Organizations, Users, Memberships as regular tables
- Profile sections stored as JSON text columns
- JSON schema validation in application layer (not database)
- Indexes on JSON fields via SQLite's JSON1 `json_extract()`

**Go Libraries:**
- `github.com/mattn/go-sqlite3` with SQLCipher build tags
- Or `github.com/mutecomm/go-sqlcipher` (SQLCipher bindings)
- `github.com/mark3labs/mcp-go` for MCP server

**Key management:**
```
# Key passed via environment variable
ARMOR_DB_KEY=<256-bit-key-base64>

# Or derived from passphrase at startup
ARMOR_DB_PASSPHRASE=<user-passphrase>
```

**Database file location:**
```
~/.armor/armor.db           # Default local
/var/lib/armor/armor.db     # Server deployment
./data/armor.db             # Development
```

### SQL Schema

```sql
-- Organizations
CREATE TABLE organizations (
  id TEXT PRIMARY KEY,              -- UUID
  name TEXT NOT NULL,
  slug TEXT NOT NULL UNIQUE,
  created_at TEXT NOT NULL,         -- ISO 8601
  updated_at TEXT NOT NULL
);

CREATE INDEX idx_organizations_slug ON organizations(slug);

-- Users
CREATE TABLE users (
  id TEXT PRIMARY KEY,              -- UUID
  email TEXT NOT NULL UNIQUE,
  name TEXT NOT NULL,
  created_at TEXT NOT NULL,
  last_login_at TEXT
);

CREATE INDEX idx_users_email ON users(email);

-- Organization Memberships
CREATE TABLE organization_memberships (
  organization_id TEXT NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
  user_id TEXT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  role TEXT NOT NULL CHECK (role IN ('owner', 'admin', 'editor', 'viewer')),
  created_at TEXT NOT NULL,
  PRIMARY KEY (organization_id, user_id)
);

-- Profiles (one per organization)
CREATE TABLE profiles (
  organization_id TEXT PRIMARY KEY REFERENCES organizations(id) ON DELETE CASCADE,
  version INTEGER NOT NULL DEFAULT 1,
  updated_at TEXT NOT NULL,
  updated_by TEXT,                  -- User ID or 'agent'
  
  -- Core components (JSON text, validated in application)
  mission TEXT,                     -- JSON matching mission.schema.json
  assets TEXT,                      -- JSON matching assets.schema.json
  adversaries TEXT,                 -- JSON matching adversaries.schema.json
  threats TEXT,                     -- JSON matching threats.schema.json
  risks TEXT,                       -- JSON matching risks.schema.json
  mitigations TEXT,                 -- JSON matching mitigations.schema.json
  
  -- Optional modules
  deep_adversary_profiling TEXT,
  information_operations TEXT,
  opsec TEXT,
  response_capability TEXT,
  technical_deep_dive TEXT
);

-- Profile Change History
CREATE TABLE profile_changes (
  id TEXT PRIMARY KEY,              -- UUID
  organization_id TEXT NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
  version INTEGER NOT NULL,         -- Profile version after this change
  timestamp TEXT NOT NULL,
  user_id TEXT REFERENCES users(id),
  agent_session_id TEXT,
  change_type TEXT NOT NULL CHECK (change_type IN ('create', 'update', 'delete')),
  section TEXT NOT NULL,
  path TEXT NOT NULL,               -- JSON path within section
  previous_value TEXT,              -- JSON
  new_value TEXT,                   -- JSON
  source TEXT NOT NULL CHECK (source IN ('web', 'agent', 'api', 'import'))
);

CREATE INDEX idx_profile_changes_org ON profile_changes(organization_id);
CREATE INDEX idx_profile_changes_timestamp ON profile_changes(timestamp);

-- Agent Proposals (pending changes awaiting approval)
CREATE TABLE proposals (
  id TEXT PRIMARY KEY,              -- UUID
  organization_id TEXT NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
  agent_session_id TEXT NOT NULL,
  created_at TEXT NOT NULL,
  status TEXT NOT NULL DEFAULT 'pending' CHECK (status IN ('pending', 'accepted', 'rejected')),
  resolved_at TEXT,
  resolved_by TEXT REFERENCES users(id),
  rejection_reason TEXT,
  
  section TEXT NOT NULL,
  changes TEXT NOT NULL,            -- JSON patch
  rationale TEXT NOT NULL,          -- Agent's explanation
  preview TEXT NOT NULL             -- JSON: what section would look like
);

CREATE INDEX idx_proposals_org_status ON proposals(organization_id, status);

-- API Keys (for programmatic access)
CREATE TABLE api_keys (
  id TEXT PRIMARY KEY,              -- UUID
  organization_id TEXT NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
  name TEXT NOT NULL,               -- Human-readable label
  key_hash TEXT NOT NULL,           -- SHA-256 hash of key
  permissions TEXT NOT NULL CHECK (permissions IN ('read', 'read-propose', 'read-write')),
  created_at TEXT NOT NULL,
  created_by TEXT NOT NULL REFERENCES users(id),
  last_used_at TEXT,
  expires_at TEXT                   -- NULL = never expires
);

CREATE INDEX idx_api_keys_hash ON api_keys(key_hash);

-- Agent Sessions (for audit trail)
CREATE TABLE agent_sessions (
  id TEXT PRIMARY KEY,              -- UUID
  organization_id TEXT NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
  api_key_id TEXT REFERENCES api_keys(id),
  started_at TEXT NOT NULL,
  ended_at TEXT,
  tool_calls INTEGER DEFAULT 0      -- Count of MCP tool invocations
);

CREATE INDEX idx_agent_sessions_org ON agent_sessions(organization_id);

-- Incidents
CREATE TABLE incidents (
  id TEXT PRIMARY KEY,              -- UUID
  organization_id TEXT NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
  title TEXT NOT NULL,
  description TEXT NOT NULL,
  type TEXT NOT NULL CHECK (type IN ('phishing', 'account_compromise', 'data_breach', 
    'malware', 'physical', 'harassment', 'disinformation', 'other')),
  severity TEXT NOT NULL CHECK (severity IN ('critical', 'high', 'medium', 'low')),
  status TEXT NOT NULL DEFAULT 'open' CHECK (status IN ('open', 'investigating', 'resolved', 'closed')),
  occurred_at TEXT NOT NULL,
  discovered_at TEXT NOT NULL,
  resolved_at TEXT,
  created_at TEXT NOT NULL,
  created_by TEXT NOT NULL REFERENCES users(id),
  
  affected_assets TEXT,             -- JSON array of asset IDs
  related_threats TEXT,             -- JSON array of threat IDs
  linked_proposals TEXT,            -- JSON array of proposal IDs
  response_actions TEXT,            -- JSON array of ResponseAction objects
  lessons_learned TEXT
);

CREATE INDEX idx_incidents_org ON incidents(organization_id);
CREATE INDEX idx_incidents_status ON incidents(organization_id, status);
CREATE INDEX idx_incidents_occurred ON incidents(occurred_at);
```

### Schema Validation

- Use AJV (JavaScript) or equivalent for JSON Schema validation
- Validate on every write operation
- Return detailed validation errors with JSON paths
- Schemas loaded from existing `/schemas/*.json` files

### MCP Server Implementation

- Use official MCP SDK
- Stateless tools (session context passed per-call)
- Rate limiting per organization
- Audit log all tool invocations

### Security Considerations

- All profile data encrypted at rest
- API keys scoped to single organization
- Agent sessions have configurable permission levels
- Proposal system prevents agents from directly modifying data
- Audit trail for all changes (user and agent)
- PII handling: flag sensitive fields, enable field-level encryption

---

## Future Considerations

### Multi-tenancy for Consultants

Consultants who work with multiple orgs need:
- Single login, multiple org access
- Role per organization
- Aggregate views across clients (anonymized)

### Collaborative Editing

- Real-time collaboration (multiple users editing)
- Presence indicators
- Conflict resolution

### Integrations

- Import from existing security tools
- Export to ticketing systems (mitigations → Jira)
- Calendar integration (review reminders)
- Slack/Teams notifications

### Offline/Local Mode

- Local-first option for sensitive orgs
- Sync when connected
- Full encryption, user-held keys

---

## Open Questions

1. **Proposal expiration**: Should agent proposals expire after N days?

2. **Batch proposals**: Should agents be able to propose multiple related changes as a single reviewable unit?

3. **Confidence thresholds**: Should high-confidence agent proposals auto-apply (with undo), or always require approval?

4. **Multi-org agents**: Should an agent session be able to access multiple orgs (for consultants)?

5. **Version snapshots**: Should we support named snapshots (e.g., "Q1 2024 Assessment") beyond the change log?

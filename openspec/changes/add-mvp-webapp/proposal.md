# Change: Add MVP Web Application

## Why

We need a proof-of-concept to validate the core ARMOR workflow: creating and editing threat model profiles through a web interface. This MVP proves the end-to-end flow before investing in auth, multi-org, agent integration, and encryption.

## What Changes

- New Go backend server with REST API for profile CRUD
- New SvelteKit frontend with forms for all 6 core profile sections
- SQLite database (unencrypted) for persistence
- JSON schema validation on all profile writes
- Simple runtime password protection (single shared password)
- Field-fill percentage completeness tracking per section

## Scope Boundaries

**In scope:**
- Single profile (no org/user management)
- All 6 core sections: Mission, Assets, Adversaries, Threats, Risks, Mitigations
- Create new profile / load existing profile
- Edit and save each section independently
- Validation against JSON schemas
- Completeness percentage display
- Local deployment only

**Out of scope (future work):**
- User authentication / authorization
- Organization management
- Agent/MCP integration
- Proposal workflow
- Encryption (SQLCipher)
- Optional modules (Deep Adversary, Info Ops, OPSEC, etc.)
- Incidents tracking
- History/audit log
- Cloudflared deployment

## Impact

- Affected specs: New `webapp` capability
- Affected code: New `server/` and `web/` directories
- No breaking changes (greenfield)

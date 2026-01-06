# Tasks: MVP Web Application

## 1. Backend Foundation

- [x] 1.1 Initialize Go module in `server/`
- [x] 1.2 Set up SQLite database connection (no encryption)
- [x] 1.3 Create database schema (profiles table with UUID, name, JSON columns)
- [x] 1.4 Implement database migrations (create table on startup)
- [x] 1.5 Load JSON schemas from `schemas/` directory
- [x] 1.6 Implement JSON schema validator
- [x] 1.7 Implement completeness calculator (field-fill percentage)

## 2. Backend API

- [x] 2.1 Set up HTTP router (net/http or chi)
- [x] 2.2 Implement password middleware (check Authorization header)
- [x] 2.3 `GET /api/profiles` - list all profiles with completeness
- [x] 2.4 `POST /api/profiles` - create new profile (name, description)
- [x] 2.5 `GET /api/profiles/:id` - return full profile
- [x] 2.6 `DELETE /api/profiles/:id` - delete profile
- [x] 2.7 `GET /api/profiles/:id/:section` - return single section
- [x] 2.8 `PUT /api/profiles/:id/:section` - update section with validation
- [x] 2.9 Add CORS headers for local development

## 3. Frontend Foundation

- [x] 3.1 Initialize SvelteKit project in `web/`
- [x] 3.2 Set up TailwindCSS
- [x] 3.3 Create API client with password header injection
- [x] 3.4 Create password gate component (prompt + sessionStorage)
- [x] 3.5 Create base layout with navigation

## 4. Frontend - Profiles List & Management

- [x] 4.1 Create profiles dashboard page (`/profiles`)
- [x] 4.2 Display profile cards with name, description, completeness
- [x] 4.3 Create new profile page (`/profiles/new`)
- [x] 4.4 Add delete profile functionality

## 5. Frontend - Profile Detail

- [x] 5.1 Create profile detail page (`/profiles/:id`)
- [x] 5.2 Display section cards with completeness percentages
- [x] 5.3 Add navigation links to each section form

## 6. Frontend - Section Forms

- [x] 6.1 Create Mission form (`/profiles/:id/mission`)
- [x] 6.2 Create Assets form with list management (`/profiles/:id/assets`)
- [x] 6.3 Create Adversaries form with list management (`/profiles/:id/adversaries`)
- [x] 6.4 Create Threats form (`/profiles/:id/threats`)
- [x] 6.5 Create Risks form with list management (`/profiles/:id/risks`)
- [x] 6.6 Create Mitigations form with list management (`/profiles/:id/mitigations`)

## 7. Validation & Error Handling

- [x] 7.1 Display validation errors from backend in forms
- [x] 7.2 Add loading states for API calls
- [x] 7.3 Add success/error toast notifications

## 8. Integration & Testing

- [x] 8.1 Test full create profile → edit sections → save flow
- [x] 8.2 Verify schema validation rejects invalid data
- [x] 8.3 Verify completeness calculation updates correctly
- [x] 8.4 Test multiple profiles work independently
- [x] 8.5 Write README with setup and run instructions

## Dependencies

- Tasks 2.x depend on 1.x (backend foundation)
- Tasks 4.x, 5.x, 6.x depend on 3.x (frontend foundation)
- Tasks 5.x and 6.x depend on 2.x (API endpoints)
- Task 7.x can be done in parallel with 5.x and 6.x
- Task 8.x requires all previous tasks

## Parallelizable Work

- Backend (1.x, 2.x) and Frontend foundation (3.x) can proceed in parallel
- Different section forms (6.1-6.6) can be built in parallel once 6.1 pattern established

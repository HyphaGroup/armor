# Capability: Web Application

MVP web application for creating and editing ARMOR threat model profiles.

## ADDED Requirements

### Requirement: Password Protection

The system SHALL require a password to access the application.

The password SHALL be configured via `ARMOR_PASSWORD` environment variable on the backend.

The frontend SHALL prompt for password on first access and store it in sessionStorage.

All API requests SHALL include the password in the `Authorization: Bearer <password>` header.

#### Scenario: Valid password grants access
- **GIVEN** the backend is running with `ARMOR_PASSWORD=secret123`
- **WHEN** a user enters `secret123` in the password prompt
- **THEN** the user is granted access to the application
- **AND** subsequent API calls include the password header

#### Scenario: Invalid password is rejected
- **GIVEN** the backend is running with `ARMOR_PASSWORD=secret123`
- **WHEN** an API request is made with `Authorization: Bearer wrongpassword`
- **THEN** the server responds with 401 Unauthorized

#### Scenario: Missing password is rejected
- **GIVEN** the backend is running with `ARMOR_PASSWORD` set
- **WHEN** an API request is made without an Authorization header
- **THEN** the server responds with 401 Unauthorized

---

### Requirement: Profile Storage

The system SHALL store multiple threat model profiles in SQLite.

Each profile SHALL have a unique ID, name, and optional description.

Each profile SHALL contain six sections: mission, assets, adversaries, threats, risks, mitigations.

Each section SHALL be stored as a JSON column validated against its corresponding schema.

#### Scenario: Create new profile
- **GIVEN** a user wants to create a new threat model
- **WHEN** `POST /api/profiles` is called with name "Acme Org"
- **THEN** a new profile is created with a unique ID
- **AND** all sections are initialized to null
- **AND** the response contains the new profile with its ID

#### Scenario: Profile persists across restarts
- **GIVEN** a profile with mission data saved
- **WHEN** the server is restarted
- **THEN** the mission data is still present when retrieved

---

### Requirement: List Profiles

The system SHALL provide an endpoint to list all profiles with their completeness.

#### Scenario: List profiles with completeness
- **GIVEN** two profiles exist: "Acme" at 50% complete and "Beta" at 25% complete
- **WHEN** `GET /api/profiles` is called
- **THEN** the response contains both profiles
- **AND** each profile includes its overall completeness percentage

#### Scenario: List empty profiles
- **GIVEN** no profiles exist
- **WHEN** `GET /api/profiles` is called
- **THEN** the response contains an empty array

---

### Requirement: View Full Profile

The system SHALL provide an endpoint to retrieve a complete profile by ID.

#### Scenario: Get full profile
- **GIVEN** a profile with ID "abc123" and mission and assets sections populated
- **WHEN** `GET /api/profiles/abc123` is called
- **THEN** the response contains the profile with all six sections
- **AND** unpopulated sections are returned as null

#### Scenario: Get non-existent profile
- **WHEN** `GET /api/profiles/nonexistent` is called
- **THEN** the server responds with 404 Not Found

---

### Requirement: Delete Profile

The system SHALL provide an endpoint to delete a profile.

#### Scenario: Delete existing profile
- **GIVEN** a profile with ID "abc123" exists
- **WHEN** `DELETE /api/profiles/abc123` is called
- **THEN** the profile is removed from the database
- **AND** the response confirms deletion

#### Scenario: Delete non-existent profile
- **WHEN** `DELETE /api/profiles/nonexistent` is called
- **THEN** the server responds with 404 Not Found

---

### Requirement: View Single Section

The system SHALL provide an endpoint to retrieve a single profile section.

#### Scenario: Get populated section
- **GIVEN** a profile "abc123" with mission data saved
- **WHEN** `GET /api/profiles/abc123/mission` is called
- **THEN** the response contains the mission section data

#### Scenario: Get empty section
- **GIVEN** a profile "abc123" with no assets data
- **WHEN** `GET /api/profiles/abc123/assets` is called
- **THEN** the response contains null or empty object

#### Scenario: Invalid section name
- **WHEN** `GET /api/profiles/abc123/invalid` is called
- **THEN** the server responds with 404 Not Found

---

### Requirement: Update Section with Validation

The system SHALL provide an endpoint to update a single profile section.

The system SHALL validate the section data against its JSON schema before saving.

The system SHALL reject invalid data with detailed error messages.

#### Scenario: Valid section update
- **GIVEN** valid mission data conforming to mission.schema.json
- **WHEN** `PUT /api/profiles/abc123/mission` is called with the data
- **THEN** the mission section is saved
- **AND** the response contains the saved data
- **AND** `updated_at` timestamp is set

#### Scenario: Invalid section update
- **GIVEN** mission data missing required `mission_statement` field
- **WHEN** `PUT /api/profiles/abc123/mission` is called with the data
- **THEN** the server responds with 400 Bad Request
- **AND** the response contains validation error details including the field path

#### Scenario: Partial section update
- **GIVEN** an existing profile with mission fully populated
- **WHEN** `PUT /api/profiles/abc123/mission` is called with only some fields
- **THEN** the entire section is replaced with the new data (not merged)

---

### Requirement: Profile Completeness

The system SHALL calculate completeness percentage for each section.

Completeness SHALL be calculated as: (non-empty required fields / total required fields) × 100.

Overall profile completeness SHALL be the average of all section completeness percentages.

#### Scenario: Completeness in profile list
- **GIVEN** profiles with varying completeness
- **WHEN** `GET /api/profiles` is called
- **THEN** each profile includes its overall completeness percentage

#### Scenario: Completeness for empty profile
- **GIVEN** a new profile with no sections populated
- **WHEN** the profile is retrieved
- **THEN** overall completeness is 0%

#### Scenario: Completeness for partial profile
- **GIVEN** a profile with mission at 100% and all other sections at 0%
- **WHEN** the profile is retrieved
- **THEN** overall completeness is approximately 17% (1/6 sections complete)

---

### Requirement: Profiles Dashboard

The frontend SHALL display a dashboard listing all profiles.

The dashboard SHALL show name, description, and overall completeness for each profile.

The dashboard SHALL provide a link to create a new profile.

The dashboard SHALL provide navigation to each profile's detail view.

#### Scenario: View dashboard with multiple profiles
- **GIVEN** two profiles exist: "Acme Org" at 50% and "Beta Corp" at 25%
- **WHEN** the user navigates to the profiles dashboard
- **THEN** both profiles are listed with their names and completeness
- **AND** each profile links to its detail view

#### Scenario: View empty dashboard
- **GIVEN** no profiles exist
- **WHEN** the user navigates to the profiles dashboard
- **THEN** a message indicates no profiles exist
- **AND** a button to create a new profile is displayed

---

### Requirement: Create Profile

The frontend SHALL provide a form to create a new profile.

#### Scenario: Create new profile
- **GIVEN** a user on the create profile page
- **WHEN** the user enters name "Acme Org" and submits
- **THEN** a new profile is created
- **AND** the user is navigated to the new profile's detail view

---

### Requirement: Profile Detail View

The frontend SHALL display a detail view for a single profile showing all sections.

The detail view SHALL show completeness percentage for each section.

The detail view SHALL provide navigation to each section's edit form.

#### Scenario: View profile detail with partial data
- **GIVEN** a profile with mission at 50% and assets at 0%
- **WHEN** the user navigates to the profile detail view
- **THEN** the mission card shows 50% complete
- **AND** the assets card shows 0% complete
- **AND** each card links to its edit form

---

### Requirement: Section Edit Forms

The frontend SHALL provide edit forms for each of the six profile sections.

Each form SHALL load existing data on page load.

Each form SHALL save data to the backend on submit.

Each form SHALL display validation errors returned by the backend.

#### Scenario: Edit mission section
- **GIVEN** a user on the mission edit form
- **WHEN** the user enters a mission statement and submits
- **THEN** the data is saved to the backend
- **AND** a success message is displayed
- **AND** completeness is updated

#### Scenario: Display validation error
- **GIVEN** a user on the mission edit form
- **WHEN** the user submits without required fields
- **THEN** validation errors are displayed inline
- **AND** the form is not cleared

#### Scenario: Edit list-based section
- **GIVEN** a user on the assets edit form
- **WHEN** the user adds a new asset item
- **THEN** a new item form appears
- **AND** the user can fill in asset details
- **AND** the user can remove the item

---

### Requirement: Navigation

The frontend SHALL provide navigation between dashboard and section forms.

The frontend SHALL indicate the current location in the navigation.

#### Scenario: Navigate from dashboard to section
- **GIVEN** a user on the dashboard
- **WHEN** the user clicks the mission section card
- **THEN** the user is navigated to the mission edit form

#### Scenario: Navigate back to dashboard
- **GIVEN** a user on a section edit form
- **WHEN** the user clicks the dashboard link in navigation
- **THEN** the user is navigated to the dashboard

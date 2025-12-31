# ARMOR Methodology

Detailed instructions for conducting ARMOR threat modeling. This document describes *what* to do at each step, regardless of delivery format (workshop, self-guided, agent-assisted).

For conceptual background, see `framework.md`. For format-specific guidance, see the implementation guides.

---

## Overview

### Structure

**Core Components** (required, ~3 hours total):
1. Mission & Impact Framework (30 min)
2. Asset Identification (45 min)
3. Adversary Profiling (30 min)
4. Threat Mapping (45 min)
5. Risk Assessment (45 min)
6. Mitigation Planning (included in closing)

**Modules** (as needed, 45-60 min each):
- Deep Adversary Profiling
- Information Operations
- OPSEC Analysis
- Incident Response
- Technical Deep-Dive

### Prerequisites

Before starting, gather:
- Organization's mission statement
- List of key programs/activities
- Overview of IT systems and data storage (diagram if available)
- Any recent security incidents or concerns
- Known adversaries or threat actors

---

## Core Component 1: Mission & Impact Framework

### Purpose
Establish context for all security decisions by defining what matters most to THIS organization.

### Inputs
- Mission statement
- Knowledge of organizational priorities

### Process

**Step 1: Capture Mission Statement**

Document the organization's mission in their own words:
- What are you trying to achieve?
- Who depends on you achieving it?
- What would mission failure look like?

**Step 2: Define Impact Areas**

Assess and rank these impact areas by organizational priority (1=highest, 6=lowest):

| Impact Area | Description |
|-------------|-------------|
| Safety & Security | Physical safety of staff, volunteers, beneficiaries, sources |
| Mission Delivery | Ability to deliver core programs and services |
| Trust & Reputation | Stakeholder confidence, public perception, credibility |
| Financial Resources | Funding, operational costs, fines |
| Legal & Compliance | Regulatory requirements, legal exposure |
| Partner Relations | Relationships with donors, allies, coalitions |

**Step 3: Define Impact Thresholds**

For the top 3 priority impact areas, define what HIGH/MEDIUM/LOW impact means:

| Impact Area | HIGH | MEDIUM | LOW |
|-------------|------|--------|-----|
| [Area 1] | | | |
| [Area 2] | | | |
| [Area 3] | | | |

*Example for Safety:*
- HIGH: Physical harm to staff or beneficiaries
- MEDIUM: Credible threats requiring security changes
- LOW: General increase in hostile attention

**Step 4: Map Support Network**

Identify who you can turn to for help when security issues arise:

| Support Type | Contact/Resource | Relationship Status |
|--------------|------------------|---------------------|
| **Peer Organizations** | Orgs facing similar threats, can share intel | |
| **Legal Support** | Lawyer/firm for security-related legal issues | |
| **Digital Security** | Helpline, consultant, or trained staff | |
| **IT Support** | Technical response capability | |
| **Physical Security** | If relevant to your context | |
| **Communications/PR** | For incident response communications | |
| **Mental Health** | Support for staff dealing with incidents | |
| **Platform Contacts** | Direct contacts at social media platforms | |

**Key questions:**
- Do you have named contacts, or just "we'd find someone"?
- Are relationships established BEFORE you need them?
- Do peer organizations know to alert you about threats in your sector?

### Outputs
- Mission statement captured
- Impact areas ranked 1-6
- Impact thresholds defined for top 3 areas
- Support network mapped with contact status

### Schema Reference
`mission.schema.json`

---

## Core Component 2: Asset Identification

### Purpose
Identify what needs protection—focusing on information assets and where they live.

### Inputs
- Mission context from Component 1
- Knowledge of organizational data and systems

### Process

**Step 1: Brainstorm Information Assets**

Work through each category, identifying specific assets:

| Category | Guiding Questions |
|----------|-------------------|
| Beneficiary/Client Data | Who do you serve? What do you know about them? |
| Source Data | Do you have confidential sources? Whistleblowers? |
| Donor/Supporter Data | Who funds you? What do you know about them? |
| Staff/Volunteer Data | Personal information, HR records, contact details |
| Financial Data | Banking, transactions, budgets, payroll |
| Strategic Data | Plans, unpublished research, internal decisions |
| Communications | Emails, messages, meeting notes, call records |
| Credentials | Passwords, keys, access tokens |
| Operational Data | Schedules, locations, travel plans |

For each asset identified, document:
- What it is
- Why it matters (what harm if compromised?)

**Step 2: Map Containers**

For each critical asset, identify where it lives:

| Container Type | Examples |
|----------------|----------|
| Technical | Cloud services, databases, devices, email systems |
| Physical | Filing cabinets, printed documents, office locations |
| Human | People who know the information, have access |

Note third-party services that hold critical data.

**Step 3: Assign Value and Requirements**

For each asset:

| Field | Options | Notes |
|-------|---------|-------|
| Value | Critical / High / Medium / Low | How damaging if compromised? |
| Risk Score Value | 3 / 2 / 2 / 1 | Numeric mapping for scoring |
| Primary Requirement | Confidentiality / Integrity / Availability | What matters most? |

**Value Mapping:**
| Value | Risk Score | Meaning |
|-------|------------|---------|
| Critical | 3 | Loss would severely impact mission or cause serious harm |
| High | 2 | Loss would significantly disrupt operations |
| Medium | 2 | Loss would cause moderate disruption |
| Low | 1 | Loss would be inconvenient but manageable |

### Outputs
- Asset inventory with descriptions
- Container mapping for critical assets
- Value ratings (used in Risk Assessment)
- Primary security requirements

### Schema Reference
`assets.schema.json`

---

## Core Component 3: Adversary Profiling

### Purpose
Identify who might target you and why.

### Inputs
- Mission context from Component 1
- Asset inventory from Component 2

### Process

**Step 1: Assess Adversary Categories**

Review each category and assess relevance:

| Category | Relevance Indicators | Universal? |
|----------|---------------------|------------|
| Nation-State / Intelligence | Work challenges government authority, operates in/covers authoritarian contexts, supports opposition/dissidents | No |
| Ideological Opposition | Politically contentious issues, high visibility on divisive topics | No |
| Cybercriminal | Has money, data, or systems worth stealing | Yes |
| Insider Threat | Has staff/volunteers with access to sensitive information | Yes |
| Competitor / Opposing Org | Competitive funding environment, institutional opposition | No |
| Opportunistic | Has internet-connected systems | Yes |

For each relevant category, assess: **Confirmed / Likely / Possible / Unlikely**

**Step 2: Customize Adversary Profiles**

For each relevant adversary, document:

| Field | Description |
|-------|-------------|
| Name/Label | Specific or generic identifier |
| Category | From the list above |
| Why You? | Why would they target YOUR organization? |
| What They Want | Information? Disruption? Discrediting? |
| Relevance | Confirmed / Likely / Possible |
| Relevance Rationale | Evidence for this assessment |

**Step 3: Select Primary Adversaries**

Identify 2-4 adversaries to focus on for threat mapping. Selection criteria:
- Confirmed or Likely relevance
- Higher capability
- History of targeting similar organizations

### Outputs
- Adversary categories assessed
- Relevant adversaries profiled
- Primary adversaries selected (2-4)

### Schema Reference
`adversaries.schema.json`, `adversary-templates.json`

---

## Core Component 4: Threat Mapping

### Purpose
Map specific threats based on your adversaries and assets.

### Inputs
- Asset inventory from Component 2
- Adversary profiles from Component 3

### Process

**Step 1: Assess Threat Categories**

For each threat category, assess relevance and provide examples:

**Account & Access Threats:**
| Threat | Description |
|--------|-------------|
| Phishing & Social Engineering | Deceptive messages to steal credentials or information |
| Account Takeover | Unauthorized access to accounts |
| Unauthorized Access | Breaking into systems |
| Insider Misuse | Staff/volunteers abusing their access |

**Data & Information Threats:**
| Threat | Description |
|--------|-------------|
| Data Breach | Unauthorized access to sensitive data |
| Data Tampering | Unauthorized modification of data |
| Data Loss | Accidental or malicious destruction |
| Surveillance | Monitoring of communications or activities |

**Disruption Threats:**
| Threat | Description |
|--------|-------------|
| Service Disruption | DDoS, system outages |
| Infrastructure Attack | Targeting core systems |

**Information & Reputation Threats:**
| Threat | Description |
|--------|-------------|
| Narrative Attacks | False claims, conspiracy theories, distorted facts |
| Impersonation | Fake accounts, personas posing as org/staff |
| Harassment | Coordinated pile-ons, threats, doxing |
| Amplification | Bot networks, coordinated inauthentic behavior |
| Platform Manipulation | Mass reporting, SEO attacks |
| Document Leaks | Theft and release of internal documents |

**Physical Threats:**
| Threat | Description |
|--------|-------------|
| Physical Intrusion | Unauthorized facility access |
| Physical Surveillance | Being watched, followed |
| Intimidation & Violence | Threats, harassment, attacks |

**Operational Threats (non-adversarial):**
| Threat | Description |
|--------|-------------|
| Human Error | Mistakes, accidents |
| Technical Failure | System crashes, data corruption |
| Natural Disaster | Fire, flood, earthquake |
| Third-Party Failure | Vendor/partner failures |

**Step 2: Link Threats to Adversaries and Assets**

For each relevant threat, document:
- Which adversaries might use this threat
- Which assets it could affect
- Concrete example of how it might manifest

**Step 3: Assign Likelihood**

For each relevant threat:

| Likelihood | Risk Score | Meaning |
|------------|------------|---------|
| High (Expected/Active) | 3 | Threat is active or expected within the year |
| Medium (Possible) | 2 | Threat could occur, some indicators exist |
| Low (Unlikely) | 1 | Threat is possible but no current indicators |

### Outputs
- Relevant threats identified with examples
- Threats linked to adversaries and assets
- Likelihood ratings (used in Risk Assessment)

### Schema Reference
`threats.schema.json`

---

## Core Component 5: Risk Assessment

### Purpose
Combine assets, adversaries, and threats into prioritized risk scenarios.

### Inputs
- Assets with values from Component 2
- Threats with likelihood from Component 4

### Process

**Step 1: Generate Risk Scenarios**

Create 8-12 risk scenarios using this format:

> "There is a risk that **[adversary]** could **[threat action]** affecting **[asset]**, which would impact **[impact area]**."

*Examples:*
- "There is a risk that nation-state actors could conduct phishing attacks affecting staff email accounts, which would impact beneficiary safety."
- "There is a risk that ideological opposition could launch coordinated harassment affecting visible staff members, which would impact mission delivery and staff wellbeing."

**Scenario Validation Questions:**
- Is it specific and clear?
- Can we identify the cause?
- Is the impact tangible?
- Is it within our scope?
- Is it actionable?

**Step 2: Assess Vulnerability**

For each scenario, assess current exposure:

| Field | Description |
|-------|-------------|
| Existing Controls | What protections currently exist? |
| Control Gaps | What's missing or insufficient? |
| Vulnerability Score | 1-3 rating |

**Vulnerability Scoring:**
| Score | Level | Indicators |
|-------|-------|------------|
| 1 | Well-protected | Strong controls exist, regularly tested, few gaps |
| 2 | Some gaps | Partial controls, known weaknesses, or untested |
| 3 | Exposed | Minimal controls, significant gaps, or no protection |

**Step 3: Calculate Risk Scores**

For each scenario:

| Factor | Score | Source |
|--------|-------|--------|
| Asset Value | 1-3 | From Component 2 |
| Likelihood | 1-3 | From Component 4 |
| Vulnerability | 1-3 | From Step 2 |
| **Risk Score** | 1-27 | Asset Value × Likelihood × Vulnerability |

**Priority Bands:**
| Score | Priority | Action |
|-------|----------|--------|
| 18-27 | Critical | Immediate action required |
| 10-17 | High | Near-term action needed |
| 4-9 | Moderate | Plan and address |
| 1-3 | Low | Monitor, accept, or defer |

**Step 4: Identify Priority Risks**

Select top 5 risks for immediate focus. Note which factor is driving each:
- High asset value → Protect the asset
- High likelihood → Address the adversary/threat
- High vulnerability → Fix the gaps (most actionable)

**Step 5: Establish Security Indicators**

Identify early warning signs that your threat environment is changing:

| Indicator Type | What to Monitor | Who Monitors | How Often |
|----------------|-----------------|--------------|-----------|
| **Peer Incidents** | Security incidents at similar organizations | | Monthly |
| **Sector Targeting** | Reports of adversary activity in your sector | | Ongoing |
| **Online Hostility** | Increase in negative mentions, harassment | | Weekly |
| **Legal/Regulatory** | New laws, investigations affecting civil society | | Monthly |
| **Adversary Activity** | Known adversaries becoming more active | | Ongoing |
| **Staff Reports** | Staff noticing surveillance, unusual contacts | | Ongoing |
| **Technical Alerts** | Unusual login attempts, phishing increase | | Daily |

**Trigger thresholds:**
Define what level of indicator activity should trigger a threat model review:
- Any confirmed incident at peer organization in same region/sector
- Significant increase in online hostility or harassment
- New legal threats to organizational type
- Staff report of physical surveillance
- Successful phishing attempt or account compromise

### Outputs
- Risk scenarios documented
- Vulnerability assessed for each
- Risk scores calculated
- Top 5 priority risks identified
- Security indicators defined with monitoring owners

### Schema Reference
`risks.schema.json`

---

## Core Component 6: Mitigation Planning

### Purpose
Turn risks into actionable improvements.

### Inputs
- Priority risks from Component 5

### Process

**Step 1: Identify Mitigations**

For each priority risk, identify potential mitigations:

| Mitigation Type | Examples |
|-----------------|----------|
| Technical Control | MFA, encryption, access controls, monitoring |
| Policy/Procedure | Security policies, incident response plans, access reviews |
| Training/Awareness | Phishing training, security awareness, incident reporting |
| Physical Control | Locks, cameras, secure disposal |
| Organizational | Role changes, separation of duties, vendor management |

**Step 2: Assess Feasibility**

For each mitigation:

| Field | Options |
|-------|---------|
| Priority | Critical / High / Medium / Low |
| Effort | Minimal / Low / Medium / High / Major |
| Cost | None / Low / Medium / High |
| Owner | Who is responsible? |

**Step 3: Plan Timeline**

| Timeframe | Description |
|-----------|-------------|
| Immediate (30 days) | Quick wins, critical fixes |
| Short-term (90 days) | Planned improvements |
| Long-term (6-12 months) | Major initiatives |

**Step 4: Define Success Criteria**

For each mitigation:
- How will you know it's implemented?
- How will you know it's working?

### Outputs
- Mitigation plans for priority risks
- Owners and timelines assigned
- Quick wins identified
- Success criteria defined

### Schema Reference
`mitigations.schema.json`

---

## Closing: Next Steps

After completing the core components:

**1. Identify Quick Wins**
What can be done in the next 1-2 weeks?

**2. Recommend Modules**
Based on what emerged, which modules are relevant?

| Module | Recommend If... |
|--------|-----------------|
| Deep Adversary Profiling | Confirmed targeted threats, nation-state actors |
| Information Operations | Narrative attacks, harassment, public-facing work |
| OPSEC Analysis | Sensitive sources/beneficiaries, hostile environment |
| Incident Response | History of incidents, high-risk environment |
| Technical Deep-Dive | Complex infrastructure, specific technical concerns |

**3. Schedule Follow-Up**
- When will modules be completed?
- Who compiles the threat model profile?
- When is the first review?

---

## Module: Deep Adversary Profiling

### Purpose
Develop detailed understanding of specific adversaries to anticipate their behavior.

### When to Use
- Confirmed or strongly suspected targeted threats
- A specific adversary has targeted similar organizations
- Need to anticipate adversary behavior, not just react
- Core identified nation-state or persistent threat actors

### Process

**Step 1: Select Adversaries**
Choose 1-3 high-relevance adversaries from the core session.

**Step 2: Analyze Adversary Identity**

| Element | Questions |
|---------|-----------|
| Specific Identification | Who specifically? (if known) |
| Organizational Structure | How are they organized? |
| Primary Motivation | What drives them? |
| Objectives | What are they trying to achieve? |
| Resources | What level of resources? |
| Persistence | One-time or ongoing threat? |
| History | What's their track record with similar orgs? |

**Step 3: Analyze Capabilities**

| Domain | Level (Advanced/Moderate/Basic/Minimal) | Known Techniques |
|--------|----------------------------------------|------------------|
| Technical | | |
| Social Engineering | | |
| Information Operations | | |
| Physical | | |
| Legal | | |

**Step 4: Analyze Infrastructure**

| Type | Details |
|------|---------|
| Technical Infrastructure | Owned servers? Compromised? Commercial? |
| Social Infrastructure | Fake accounts? Front orgs? Proxy actors? |
| Physical Presence | On the ground capability? |
| Known Indicators | Domains, accounts, patterns to monitor |

**Step 5: Analyze Targeting**

| Question | Analysis |
|----------|----------|
| Why you specifically? | |
| What do they want from you? | |
| Your value as stepping stone? | Access to sources, partners, beneficiaries? |
| Primary targets within org? | Who/what would they target first? |
| Likely attack vectors? | Based on capabilities |
| Targeting indicators? | How would you know you're being targeted? |

**Step 6: Anticipate Scenarios**

For each adversary, document 2-3 anticipated attack scenarios:
- Scenario description
- Likelihood
- Primary technique
- Indicators to watch

### Outputs
- Detailed adversary profiles
- Anticipated attack scenarios
- Targeting indicators to monitor
- Defensive priorities per adversary

### Schema Reference
`deep-adversary-profiling.schema.json`

---

## Module: Information Operations

### Purpose
Assess exposure to and build resilience against narrative attacks, harassment, and reputation threats.

### When to Use
- Public-facing work (advocacy, journalism, research)
- Experienced or anticipate coordinated online attacks
- Staff have been targeted with harassment
- Work is politically contentious
- Reputation is a strategic asset

### Process

**Step 1: Assess Exposure**

| Factor | Assessment |
|--------|------------|
| Public visibility | High / Medium / Low |
| Social media presence | Platforms, followers, engagement |
| Media coverage | Frequency, visibility |
| Contentious positions | What attracts opposition? |
| Known opposition | Who actively opposes your work? |

**Step 2: Assess Staff Exposure**

For public-facing staff:
- Visibility level
- Personal social media presence
- Previous harassment experience
- Identity factors that may increase targeting

**Step 3: Assess Threat Categories**

For each category, assess relevance, current evidence, and vulnerability:

| Category | What to Assess |
|----------|----------------|
| Narrative Attacks | False claims currently circulating? Anticipated narratives? |
| Impersonation | Known fake accounts? Brand monitoring in place? |
| Harassment | Harassment history? At-risk individuals? Support available? |
| Doxing | Staff personal info exposure level? |
| Amplification | Bot activity observed? Coordinated behavior? |
| Platform Manipulation | Mass reporting experienced? Platform contacts? |
| Document Operations | Leak risk? Sensitive document exposure? |

**Step 4: Assess Monitoring Capability**

| Type | In Place? | Tool/Method | Who Monitors? |
|------|-----------|-------------|---------------|
| Social media mentions | | | |
| News/media mentions | | | |
| Staff name monitoring | | | |
| Domain monitoring | | | |
| Hashtag monitoring | | | |

**Step 5: Assess Response Capability**

| Element | In Place? |
|---------|-----------|
| Response decision framework (when to respond vs. ignore) | |
| Pre-drafted responses | |
| Designated spokesperson | |
| Platform reporting knowledge | |
| Documentation process | |
| Legal support | |
| PR/comms support | |
| Staff harassment support protocol | |
| Mental health resources | |

### Outputs
- Info ops exposure assessment
- Threat category vulnerability ratings
- Monitoring gaps
- Response capability gaps
- Priority recommendations

### Schema Reference
`information-operations.schema.json`

---

## Module: OPSEC Analysis

### Purpose
Identify what adversaries can learn about you through observation.

### When to Use
- Protecting sensitive sources or beneficiaries
- Operating in hostile environments
- Investigating powerful actors
- Staff travel to high-risk locations
- Work requires operational secrecy

### Process

**Step 1: Audit Public Exposure**

| Source | What's Public? | Concern Level |
|--------|----------------|---------------|
| Website (staff, addresses, partners, donors) | | |
| Social media accounts | | |
| Regulatory filings (990s, registrations) | | |
| Media coverage | | |

**Step 2: Audit Digital Footprint**

| Element | Assessment | What It Reveals |
|---------|------------|-----------------|
| Domain WHOIS | Private/Public | |
| Email infrastructure | | |
| Third-party services | | |
| Data broker presence | | |

**Step 3: Identify Operational Patterns**

| Pattern | Predictable? | Publicly Known? | Exploitable? |
|---------|--------------|-----------------|--------------|
| Regular meetings | | | |
| Travel patterns | | | |
| Communication patterns | | | |
| Funding cycles | | | |
| Event schedules | | | |

**Step 4: Assess Trust Relationships**

For key relationships:
- Could it be exploited through impersonation?
- How might an adversary exploit this trust?
- What mitigation is possible?

**Step 5: Assess Source/Beneficiary Protection** (if applicable)

| Element | Assessment |
|---------|------------|
| Initial contact methods | |
| Ongoing communication security | |
| Data trails that could identify sources | |
| Data minimization practices | |
| Information protection measures | |

**Step 6: Apply OPSEC Thinking**

For each finding, ask:
- Does an adversary want this information?
- Can they collect it?
- What can they do with it?
- What's the potential harm?
- Can we reduce exposure without harming mission?

**Step 7: Assess Personal Safety Exposure** (for at-risk individuals)

For staff who may be personally targeted:

| Assessment Area | Status | Action Needed |
|-----------------|--------|---------------|
| **Data Broker Presence** | Check if personal info is on data broker sites | Consider removal service |
| **Social Media Privacy** | Review personal account privacy settings | Audit and lock down |
| **Home Address Exposure** | Is home address findable online? | Remove from public records if possible |
| **Family Exposure** | Are family members identifiable/contactable? | Discuss family security |
| **Financial Exposure** | Credit monitoring in place? | Consider credit freeze |
| **Personal Device Security** | Personal devices secured? | Ensure MFA, encryption |

**Personal safety measures to consider:**

| Measure | Purpose | Priority |
|---------|---------|----------|
| Data broker removal service | Remove personal info from aggregator sites | High for public-facing staff |
| Credit freeze (all bureaus) | Prevent identity theft and info gathering | Medium |
| Google Alerts on name | Monitor for new exposure | Low effort, high value |
| "Have I Been Pwned" checks | Monitor for credential exposure | Regular check |
| Family code word | Verify identity in unusual situations | For high-risk contexts |
| Separate work/personal devices | Limit exposure if work device compromised | Medium |

**For high-risk individuals:**
- PimEyes or similar reverse image search to find photo exposure
- Regular review of what's findable via search
- Consider pseudonym for public-facing work
- Physical security assessment of residence

### Outputs
- Public exposure audit
- Digital footprint assessment
- Operational patterns identified
- OPSEC vulnerabilities prioritized
- Personal safety assessment (for at-risk staff)
- Remediation recommendations

### Schema Reference
`opsec.schema.json`

---

## Module: Incident Response

### Purpose
Assess and build capability to detect, respond to, and recover from security incidents.

### When to Use
- History of security incidents
- High-risk environment where incidents are likely
- Limited current response capability
- Want to establish protocols before needed

### Process

**Step 1: Define "Incident"**

Document what constitutes a security incident for this organization:
- Definition
- Examples that would trigger response
- Examples that would not

**Step 2: Assess Preparation**

| Element | In Place? | Notes |
|---------|-----------|-------|
| Incident response plan documented | | |
| Response team defined | | |
| Roles and decision authority clear | | |
| Contact lists (internal) | | |
| Contact lists (legal, IT, security) | | |
| Contact lists (platform, peer orgs) | | |
| Communication channels defined | | |
| Secure backup channel | | |
| Staff trained on reporting | | |
| Tabletop exercises conducted | | |
| Documentation templates ready | | |

**Step 3: Assess Detection**

| Monitoring Type | In Place? | Tool/Method | Who Monitors? |
|-----------------|-----------|-------------|---------------|
| Email security | | | |
| Account login alerts | | | |
| Endpoint protection | | | |
| Social media | | | |
| News/mentions | | | |
| Domain monitoring | | | |
| Data breach monitoring | | | |

- Do staff know what to report?
- Do staff know how to report?

**Step 4: Assess Response Capability**

| Capability | In Place? |
|------------|-----------|
| Can isolate compromised systems | |
| Can revoke/reset credentials quickly | |
| Can block accounts/access | |
| Backups exist | |
| Backups tested | |
| IT support available | |
| Digital security support available | |
| Legal support available | |
| PR/comms support available | |

**Step 5: Assess Post-Incident Process**

| Element | In Place? |
|---------|-----------|
| Post-incident review process | |
| Lessons learned documentation | |
| Threat model update trigger | |

**Step 6: Review Incident History**

Document past incidents:
- Date, type, summary
- How it was handled
- Lessons learned

### Minimum Viable Incident Response
- Incident definition exists
- Someone responsible for coordinating response
- Contact list (internal + key external)
- Secure communication channel
- Basic documentation template
- Post-incident review process

### Outputs
- Incident definition
- Preparation gaps
- Detection gaps
- Response capability gaps
- Contact lists verified/created
- Priority improvements with owners

### Schema Reference
`response-capability.schema.json`

---

## Module: Technical Deep-Dive

### Purpose
Detailed assessment of technical infrastructure and security controls.

### When to Use
- Complex IT infrastructure
- Preparing for security audit
- Specific technical concerns from core session
- Technical staff available to implement findings

### Process

**Step 1: Map Infrastructure**

**Systems Inventory:**
| System/Service | Type | Owner | Criticality | Managed By |
|----------------|------|-------|-------------|------------|
| | Endpoint/Server/Cloud/Mobile | | Critical/High/Med/Low | Internal/Third-party |

**Third-Party Services:**
| Service | Provider | Purpose | Data Shared | Security Reviewed? |
|---------|----------|---------|-------------|-------------------|

**Data Flows:**
| Data Type | Source | Destination | Encrypted? |
|-----------|--------|-------------|------------|

**Step 2: Assess Authentication**

| Question | Response | Gap? |
|----------|----------|------|
| MFA enabled on critical accounts? | | |
| MFA coverage (% of accounts) | | |
| Password manager in use? | | |
| Password policy exists? | | |
| SSO implemented? | | |

**Step 3: Assess Encryption**

| Question | Response | Gap? |
|----------|----------|------|
| Device encryption (laptops)? | | |
| Device encryption (mobile)? | | |
| Data at rest encrypted? | | |
| Data in transit encrypted? | | |
| Email encryption available? | | |

**Step 4: Assess Patching**

| Question | Response | Gap? |
|----------|----------|------|
| Auto-updates enabled (endpoints)? | | |
| Server/infrastructure patching? | | |
| Patch timeline for critical vulns? | | |

**Step 5: Assess Backups**

| Question | Response | Gap? |
|----------|----------|------|
| Critical data backed up? | | |
| Backup frequency | | |
| Backups tested? | | |
| Offsite/offline backup? | | |
| Recovery time estimate | | |

**Step 6: Assess Monitoring**

| Question | Response | Gap? |
|----------|----------|------|
| Logging enabled? | | |
| Alerts configured? | | |
| Who monitors? | | |
| Log retention | | |

**Step 7: Assess Access Control**

| Question | Response | Gap? |
|----------|----------|------|
| Least privilege implemented? | | |
| Access reviews conducted? | | |
| Offboarding process defined? | | |
| Admin access limited? | | |

**Quick Assessment Checklist:**
- [ ] Is MFA enabled on all critical accounts?
- [ ] Are all devices encrypted?
- [ ] Are systems set to auto-update?
- [ ] When were backups last tested?
- [ ] Who has admin access, and is it necessary?
- [ ] What happens to access when someone leaves?

**Step 8: Summarize Vulnerabilities**

| Vulnerability | Area | Severity | Exploitability | Remediation |
|---------------|------|----------|----------------|-------------|

**Step 9: Create Prioritized Roadmap**

| Priority | Improvement | Area | Effort | Impact | Owner | Target |
|----------|-------------|------|--------|--------|-------|--------|

**Step 10: Overall Assessment**

| Area | Score (Strong/Adequate/Weak/None) |
|------|-----------------------------------|
| Authentication | |
| Encryption | |
| Patching | |
| Backups | |
| Monitoring | |
| Access Control | |
| **Overall** | |

### Outputs
- Infrastructure inventory
- Technical assessment by area
- Vulnerabilities identified and prioritized
- Improvement roadmap with owners
- Overall technical posture rating

### Schema Reference
`technical-deep-dive.schema.json`

---

## Documentation Outputs

After completing threat modeling, compile:

| Output | Contents |
|--------|----------|
| **Threat Model Profile** | Structured JSON data conforming to schemas |
| **Executive Summary** | 1-2 page overview for leadership |
| **Risk Register** | Prioritized list of risks with scores |
| **Action Plan** | Mitigations with owners, timelines, success criteria |
| **Quick Wins** | Immediate actions (next 30 days) |

---

## Review Cadence

| Frequency | Activity |
|-----------|----------|
| Weekly (first month) | Quick wins progress check |
| Monthly | High-priority mitigation review |
| Quarterly | Emerging threats, adversary relevance check |
| Annually | Full threat model refresh |

**Triggered Reviews:**
- After significant security incidents
- After major organizational changes
- After changes in operating environment
- After new adversary activity against similar orgs

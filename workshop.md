# Threat Modeling Workshop Guide

A facilitation guide for running threat modeling sessions using the Civil Society Threat Modeling Methodology.

---

## Overview

**Total time:** 3.5-4 hours (core) + 45-60 min per module

**Structure:**
- Core Session 1: Mission & Impact (30 min)
- Core Session 2: Asset Identification (45 min)
- Core Session 3: Adversary Selection (30 min)
- Core Session 4: Threat Mapping (45 min)
- Core Session 5: Risk Prioritization (45 min)
- Closing: Next Steps (15 min)
- Modules: As needed (45-60 min each)

**Participants:** 3-8 people including leadership, program staff, and anyone with security responsibilities.

---

## Pre-Workshop Preparation

### Materials Needed
- [ ] Whiteboard or digital workspace (Miro, Mural, etc.)
- [ ] Sticky notes or digital equivalent (3 colors)
- [ ] Timer
- [ ] This guide
- [ ] Worksheets (printed or digital)

### Information to Gather
- [ ] Organization's mission statement
- [ ] List of key programs/activities
- [ ] Overview of IT systems and data storage
- [ ] Any recent security incidents or concerns
- [ ] Known adversaries or threat actors

### Participant Prep (Send in Advance)
Ask participants to think about:
- What information/systems are most critical to your work?
- Who might want to harm your organization and why?
- Recent security concerns or close calls
- Current security measures in place

---

## Core Session 1: Mission & Impact (30 min)

### Purpose
Establish why security matters for THIS organization. Ground all subsequent work in mission impact.

### Facilitator Notes
- Start here even if participants want to jump to threats
- Get buy-in that security serves mission, not the other way around
- Document in participants' own words

### Process

**1. Mission Review (5 min)**

> "Let's start with why you exist. Can someone read or summarize your mission statement?"

Capture the mission statement. Ask:
- What are you ultimately trying to achieve?
- Who depends on you achieving it?

**2. Impact Area Ranking (15 min)**

> "If something went wrong security-wise, there are different ways it could hurt you. Let's rank which types of harm matter most to YOUR organization."

Present the impact areas and have participants rank 1-6:

| Impact Area | Description | Priority |
|-------------|-------------|----------|
| **Safety & Security** | Physical safety of staff, volunteers, beneficiaries, sources | |
| **Mission Delivery** | Ability to deliver core programs and services | |
| **Trust & Reputation** | Stakeholder confidence, public perception, credibility | |
| **Financial Resources** | Funding, operational costs, fines | |
| **Legal & Compliance** | Regulatory requirements, legal exposure | |
| **Partner Relations** | Relationships with donors, allies, coalitions | |

Discuss any disagreements. Consensus isn't required but understand the reasoning.

**3. Impact Thresholds (10 min)**

For the top 3 impact areas, define what HIGH/MEDIUM/LOW would look like:

> "For [top impact area], what would be a HIGH impact event? What about MEDIUM? LOW?"

Example for Safety:
- **HIGH**: Physical harm to staff or beneficiaries
- **MEDIUM**: Credible threats requiring security changes  
- **LOW**: General increase in hostile attention

### Output
- [ ] Mission statement captured
- [ ] Impact areas ranked
- [ ] Thresholds defined for top 3

---

## Core Session 2: Asset Identification (45 min)

### Purpose
Identify what needs protection—focusing on information assets and where they live.

### Facilitator Notes
- Push beyond "our data" to specific categories
- Containers matter—same data in different places has different risk
- Value is about mission impact, not monetary value

### Process

**1. Brainstorm Information Assets (20 min)**

> "What information do you have that would hurt people or your mission if it got into the wrong hands, was changed, or became unavailable?"

Work through categories:

| Category | Your Assets |
|----------|-------------|
| **Beneficiary/Client Data** | |
| **Source Data** (for journalism/investigations) | |
| **Donor/Supporter Data** | |
| **Staff/Volunteer Data** | |
| **Financial Data** | |
| **Strategic Data** (plans, unpublished research) | |
| **Communications** (emails, messages) | |
| **Credentials** (passwords, keys) | |
| **Operational Data** (schedules, locations) | |

For each asset identified, briefly note:
- What it is
- Why it matters

**2. Map Containers (15 min)**

> "Where does each of these assets actually live?"

For critical assets, identify containers:

| Asset | Technical Containers | Physical Containers | Human Containers |
|-------|---------------------|--------------------|--------------------|
| | (systems, services) | (locations, documents) | (who knows it) |

Note any third-party services that hold critical data.

**3. Prioritize & Assign Requirements (10 min)**

For top 5-10 assets:

| Asset | Value | Primary Requirement |
|-------|-------|---------------------|
| | Critical/High/Medium/Low | Confidentiality/Integrity/Availability |

> "For each asset—is the main concern keeping it secret (confidentiality), keeping it accurate (integrity), or keeping it accessible (availability)?"

### Output
- [ ] Asset list with descriptions
- [ ] Container mapping for critical assets
- [ ] Value ratings
- [ ] Primary security requirements

---

## Core Session 3: Adversary Selection (30 min)

### Purpose
Identify who might target you and why—moving from generic threats to specific adversaries.

### Facilitator Notes
- Some adversaries are universal (cybercriminals, opportunistic)
- Others depend on your work (nation-state, ideological)
- Push for specificity on "why you"

### Process

**1. Review Adversary Categories (10 min)**

> "Let's think about who might want to harm your organization. I'll walk through some categories—tell me which seem relevant and why."

| Category | Relevant? | Notes |
|----------|-----------|-------|
| **Nation-State / Intelligence** | | Working on issues challenging government? Operating in/covering authoritarian contexts? |
| **Ideological Opposition** | | Politically contentious issues? High visibility on divisive topics? |
| **Cybercriminal** | ✓ (universal) | |
| **Insider Threat** | ✓ (universal) | |
| **Competitor / Opposing Org** | | Competitive funding? Institutional opposition? |
| **Opportunistic** | ✓ (universal) | |

**2. Customize Relevant Adversaries (15 min)**

For each relevant category:

> "Let's get specific. For [category], who specifically might this be? Why would they target YOU? What would they want?"

| Adversary | Why You? | What They Want | Relevance |
|-----------|----------|----------------|-----------|
| | | | Confirmed/Likely/Possible |

**3. Select Primary Adversaries (5 min)**

> "Of these, which 2-4 should we focus our attention on?"

Mark primary adversaries. These will drive threat mapping.

### Output
- [ ] Adversary categories assessed
- [ ] Relevant adversaries customized
- [ ] Primary adversaries selected (2-4)

---

## Core Session 4: Threat Mapping (45 min)

### Purpose
Map specific threats based on your adversaries and assets.

### Facilitator Notes
- Don't try to cover everything—focus on relevant categories
- Link threats back to adversaries identified
- Get concrete examples

### Process

**1. Walk Through Threat Categories (35 min)**

> "Now let's think about what your adversaries might actually DO. For each category, tell me if it's relevant and give me an example of what it might look like for you."

**Account & Access Threats**

| Threat | Relevant? | Example for You | Primary Adversary |
|--------|-----------|-----------------|-------------------|
| Phishing & Social Engineering | | | |
| Account Takeover | | | |
| Unauthorized Access | | | |
| Insider Misuse | | | |

**Data & Information Threats**

| Threat | Relevant? | Example for You | Primary Adversary |
|--------|-----------|-----------------|-------------------|
| Data Breach | | | |
| Data Tampering | | | |
| Data Loss | | | |
| Surveillance | | | |

**Disruption Threats**

| Threat | Relevant? | Example for You | Primary Adversary |
|--------|-----------|-----------------|-------------------|
| Service Disruption | | | |
| Infrastructure Attack | | | |

**Information & Reputation Threats**

| Threat | Relevant? | Example for You | Primary Adversary |
|--------|-----------|-----------------|-------------------|
| Narrative Attacks | | | |
| Impersonation | | | |
| Harassment Campaigns | | | |
| Amplified Attacks | | | |
| Platform Manipulation | | | |
| Document Leaks | | | |

**Physical Threats**

| Threat | Relevant? | Example for You | Primary Adversary |
|--------|-----------|-----------------|-------------------|
| Physical Intrusion | | | |
| Physical Surveillance | | | |
| Intimidation & Violence | | | |

**Non-Adversarial Threats**

| Threat | Relevant? | Example for You |
|--------|-----------|-----------------|
| Human Error | | |
| Technical Failure | | |
| Natural Disaster | | |
| Third-Party Failure | | |

**2. Assign Likelihood (10 min)**

For relevant threats, assign likelihood:

> "How likely is each of these? HIGH means you expect it or it's already happening. MEDIUM means it could happen. LOW means it's possible but unlikely."

### Output
- [ ] Relevant threats identified
- [ ] Examples documented
- [ ] Threats linked to adversaries
- [ ] Likelihood ratings

---

## Core Session 5: Risk Prioritization (45 min)

### Purpose
Combine assets, adversaries, and threats into prioritized risk scenarios with three-factor scoring.

### Facilitator Notes
- Use the risk statement format consistently
- Three factors: Asset Value, Likelihood, Vulnerability
- Don't debate scores too long—directionally right is fine
- Vulnerability scoring is key—it's what you can actually change

### Process

**1. Generate Risk Scenarios (15 min)**

> "Let's write out specific risk scenarios combining what we've discussed."

Format:
> "There is a risk that **[adversary]** could **[threat action]** affecting **[asset]**, which would impact **[impact area]**."

Generate 8-12 scenarios covering the highest-concern combinations.

| # | Risk Scenario |
|---|---------------|
| 1 | |
| 2 | |
| 3 | |
| ... | |

**2. Assess Vulnerability for Each Scenario (15 min)**

> "For each risk, how exposed are we? What protections exist? What's missing?"

For each scenario, document:

| # | Existing Controls | Gaps/Weaknesses | Vulnerability (1-3) |
|---|-------------------|-----------------|---------------------|
| 1 | | | |
| 2 | | | |

**Vulnerability scoring guide:**
- **1 (Well-protected)**: Strong controls in place, few gaps
- **2 (Some gaps)**: Partial controls, known weaknesses
- **3 (Exposed)**: Minimal controls, significant gaps, or no protection

**3. Score Risks (10 min)**

For each scenario, combine three factors:

| # | Asset Value (1-3) | Likelihood (1-3) | Vulnerability (1-3) | Score | Priority |
|---|-------------------|------------------|---------------------|-------|----------|
| 1 | | | | | |
| 2 | | | | | |

**Scoring guide:**
- **Asset Value**: 1=Low, 2=Medium, 3=High/Critical (from Session 2)
- **Likelihood**: 1=Unlikely, 2=Possible, 3=Expected/Active (from Session 4)
- **Vulnerability**: 1=Well-protected, 2=Some gaps, 3=Exposed (from step 2 above)
- **Score**: Asset Value × Likelihood × Vulnerability (1-27)
- **Priority**: Critical (18-27), High (10-17), Moderate (4-9), Low (1-3)

**4. Identify Top Risks (5 min)**

> "Which 5 risks should be our priority focus?"

Mark the top 5 risks. Note which are high because of:
- High asset value (protect the asset)
- High likelihood (address the adversary/threat)
- High vulnerability (fix the gaps—most actionable)

### Output
- [ ] Risk scenarios documented
- [ ] Vulnerability assessed for each
- [ ] Three-factor scores assigned
- [ ] Top 5 risks identified with action rationale

---

## Closing: Next Steps (15 min)

### Purpose
Translate findings into action.

### Process

**1. Quick Wins (5 min)**

> "Are there any quick wins—things we could do in the next week or two that would address these risks?"

List quick wins with owners.

**2. Module Recommendations (5 min)**

> "Based on what we discussed, I'd recommend these additional modules..."

| Module | Recommended If... | Recommend? |
|--------|-------------------|------------|
| Deep Adversary Profiling | Confirmed targeted threats | |
| Information Operations | Facing narrative attacks, harassment | |
| OPSEC Analysis | Protecting sensitive sources | |
| Incident Response | History of incidents | |
| Technical Deep-Dive | Complex infrastructure | |

**3. Schedule Follow-Up (5 min)**

- When will we complete recommended modules?
- Who will compile the threat model profile?
- When is our first review?

### Output
- [ ] Quick wins assigned
- [ ] Modules scheduled
- [ ] Follow-up dates set

---

## Modules

### Module: Deep Adversary Profiling (45 min)

**When to use:** Confirmed or highly likely targeted threats, need to anticipate adversary behavior.

**Process:**

For each primary adversary, complete:

**Who They Are**
- Specific identification (if known)
- Organizational structure
- Motivation and objectives
- Resources available
- Historical behavior patterns

**What They Can Do**
- Technical capability (Advanced/Moderate/Basic)
- Social engineering capability
- Information operations capability
- Physical capability
- Legal capability
- Known tools and techniques

**What They Use**
- Types of infrastructure
- Known indicators
- Operational patterns

**Why You**
- Why your organization specifically
- What they're after
- How you fit their objectives
- Your value as stepping stone to others

**Output:** Detailed adversary profiles, anticipated attack patterns, indicators to monitor

---

### Module: Information Operations (45 min)

**When to use:** Currently experiencing or anticipating information attacks, public-facing work on contentious issues, staff experiencing harassment.

**Process:**

**1. Current Exposure Assessment (10 min)**
- Social media presence and engagement
- Media coverage and public profile
- Known opposition narratives

**2. Threat Assessment (20 min)**

Walk through information operations categories:

| Category | Relevant? | Current Evidence | Anticipated? |
|----------|-----------|------------------|--------------|
| Narrative Attacks | | | |
| Impersonation | | | |
| Harassment | | | |
| Amplification | | | |
| Platform Manipulation | | | |
| Document Leaks | | | |

**3. Response Planning (15 min)**
- What to monitor
- When to respond vs. ignore
- Pre-drafted responses
- Platform reporting procedures
- Documentation protocols

**Output:** Info ops threat assessment, monitoring plan, response playbook

---

### Module: OPSEC Analysis (45 min)

**When to use:** Protecting sensitive sources or beneficiaries, operating in hostile environments, staff at elevated personal risk.

**Process:**

**1. Public Exposure Audit (15 min)**

| Exposure Type | What's Public? | Concern Level |
|---------------|----------------|---------------|
| Website (staff, addresses, partners) | | |
| Social media | | |
| Regulatory filings | | |
| Media coverage | | |

**2. Operational Patterns (15 min)**

| Pattern | Predictable? | Exploitable? |
|---------|--------------|--------------|
| Regular meetings | | |
| Travel | | |
| Communication | | |
| Financial cycles | | |

**3. Trust Relationships (10 min)**

| Relationship | Could Be Exploited? | How? |
|--------------|---------------------|------|
| | | |

**4. Recommendations (5 min)**

List OPSEC improvements with priority.

**Output:** OPSEC vulnerability list, prioritized remediation steps

---

### Module: Incident Response Capability (45 min)

**When to use:** High-risk environment, history of incidents, limited current readiness.

**Process:**

**1. Preparation Assessment (15 min)**

| Element | In Place? | Notes |
|---------|-----------|-------|
| Incident response plan | | |
| Incident definition | | |
| Response team & roles | | |
| Contact lists | | |
| Communication channels | | |
| Training | | |

**2. Detection Capability (10 min)**

| Monitoring Type | In Place? | Who Monitors? |
|-----------------|-----------|---------------|
| Email security | | |
| Account alerts | | |
| Social media | | |
| News/mentions | | |

**3. Response Capability (10 min)**

| Capability | In Place? | Notes |
|------------|-----------|-------|
| Can isolate systems | | |
| Can revoke access | | |
| Backups exist & tested | | |
| External support available | | |

**4. Gap Identification (10 min)**

List priority gaps with owners.

**Output:** Response readiness assessment, priority improvements

---

### Module: Technical Deep-Dive (60 min)

**When to use:** Complex IT infrastructure, specific technical concerns, preparing for audit.

**Process:**

**1. Infrastructure Mapping (20 min)**
- Systems inventory
- Data flows
- Access controls
- Third-party services

**2. Technical Assessment (25 min)**

| Area | Current State | Gaps |
|------|---------------|------|
| Authentication | | |
| Encryption | | |
| Patching | | |
| Backups | | |
| Monitoring | | |

**3. Prioritization (15 min)**

List technical improvements by priority and effort.

**Output:** Technical vulnerability list, prioritized roadmap

---

## Facilitation Tips

### Keep it Accessible
- Use plain language, explain any jargon
- Ground everything in mission impact
- Use concrete examples from civil society
- Make space for all participants to contribute

### Manage Scope
- Stay focused on what's actionable
- It's okay to note something and move on
- Defer deep dives to modules
- Document assumptions and limitations

### Common Pitfalls to Avoid
- Getting too technical too quickly
- Focusing only on cyber threats
- Ignoring human factors
- Creating plans without owners
- Skipping documentation
- Letting perfect be enemy of good

### If You Get Stuck
- Return to mission: "What would hurt your ability to do X?"
- Get concrete: "Can you give me an example?"
- Move on: "Let's note that and come back if we have time"

---

## Post-Workshop

### Documentation
1. Compile findings into threat model profile
2. Create executive summary for leadership
3. Develop action plan with timelines and owners

### Follow-Up Schedule
- **Week 1**: Quick wins implemented
- **Month 1**: Module sessions completed
- **Monthly**: Review high-priority mitigations
- **Quarterly**: Progress update, check for emerging threats
- **Annually**: Full threat model refresh

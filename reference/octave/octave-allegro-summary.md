# OCTAVE Allegro Reference

**Source:** Software Engineering Institute (SEI), Carnegie Mellon University  
**Full Report:** "Introducing OCTAVE Allegro: Improving the Information Security Risk Assessment Process" (2007)  
**Training:** https://www.sei.cmu.edu/training/

---

## Overview

OCTAVE Allegro (Operationally Critical Threat, Asset, and Vulnerability Evaluation) is a streamlined risk assessment methodology designed to help organizations conduct information security risk assessments with minimal investment in time and resources. It focuses on the interplay between people, technology, and facilities in relation to information and business processes.

Key differentiator: OCTAVE Allegro is **asset-centric** and **mission-focused**, starting with what matters to the organization rather than starting with technical vulnerabilities.

---

## Core Principles

### Mission Focus
- Security exists to support the mission
- Risk assessment should be driven by organizational objectives
- Impact is measured in terms of mission consequences

### Asset-Centric Approach
- Start with information assets, not technology
- Understand how assets support the mission
- Consider assets in their operational context

### Operationally Focused
- Assess risks where they occur—in operations
- Consider the full environment: people, technology, facilities
- Practical outputs that drive action

---

## OCTAVE Allegro Process

### Step 1: Establish Risk Measurement Criteria

**Purpose:** Define how the organization will evaluate risk impact.

**Activities:**
- Define impact areas (e.g., reputation, financial, safety, legal)
- Establish impact scales (high/medium/low)
- Prioritize impact areas for your organization

**Example Impact Areas for Civil Society:**

| Impact Area | Description |
|-------------|-------------|
| **Safety & Security** | Physical safety of staff, volunteers, beneficiaries |
| **Mission Delivery** | Ability to deliver core programs and services |
| **Reputation & Trust** | Stakeholder confidence, public perception |
| **Financial** | Funding, operational costs, fines |
| **Legal & Compliance** | Regulatory requirements, legal exposure |
| **Partner Relations** | Relationships with donors, allies, collaborators |

**Impact Scale Example:**

| Level | Safety Impact | Mission Impact | Reputation Impact |
|-------|---------------|----------------|-------------------|
| High | Physical harm possible | Programs cannot operate | Major public damage |
| Medium | Elevated risk | Significant disruption | Notable concern |
| Low | Minimal risk | Minor inconvenience | Limited impact |

### Step 2: Develop Information Asset Profiles

**Purpose:** Identify and document critical information assets.

**For each asset, document:**
- Description and rationale (why it matters)
- Owners and custodians
- Security requirements (confidentiality, integrity, availability)
- Most important security requirement

**Information Asset Categories:**

| Category | Examples |
|----------|----------|
| Beneficiary/Client Data | Personal information, case files, identities |
| Donor/Supporter Data | Contact information, giving history |
| Financial Information | Banking details, budgets, transactions |
| Strategic Information | Plans, strategies, unpublished research |
| Communication Records | Emails, messages, meeting notes |
| Credentials | Passwords, access keys, authentication data |

**Asset Profile Questions:**
1. What is this asset?
2. Who owns it? Who is accountable?
3. What are the security requirements?
4. What is the most important requirement and why?

### Step 3: Identify Information Asset Containers

**Purpose:** Determine where information assets live.

**Container Types:**
- **Technical**: Databases, servers, applications, devices, cloud services
- **Physical**: Filing cabinets, printed documents, storage rooms
- **Human**: People who know or have access to information

**For each container:**
- Identify what assets it holds
- Assess how it's protected
- Consider who has access

### Step 4: Identify Areas of Concern

**Purpose:** Brainstorm potential threats and vulnerabilities.

**Consider threats from:**
- Human actors (deliberate): External attackers, malicious insiders
- Human actors (accidental): Errors, mistakes, negligence
- Technical failures: System crashes, data corruption
- Environmental: Natural disasters, physical security incidents

**Area of Concern Template:**
1. The actor/event that creates the threat
2. How they would access the asset container
3. What they would do to the asset
4. The resulting impact

**Example:**
> "A disgruntled former volunteer (actor) who retains credentials (access) could download and leak (action) our beneficiary database, resulting in safety risks to vulnerable clients (impact)."

### Step 5: Identify Threat Scenarios

**Purpose:** Develop specific, realistic threat scenarios.

**Build scenarios by combining:**
- Asset at risk
- Container where it resides
- Actor or event
- Method of attack/access
- Outcome

**Scenario Structure:**
> "There is a risk that [actor/event] could exploit [vulnerability in container] to [action against asset], resulting in [impact to organization]."

### Step 6: Identify Risks

**Purpose:** Determine the potential impacts of each threat scenario.

**For each scenario, assess impact in each area:**
- Safety & Security
- Mission Delivery
- Reputation & Trust
- Financial
- Legal & Compliance
- Partner Relations

### Step 7: Analyze Risks

**Purpose:** Calculate risk scores to enable prioritization.

**Risk Score = Impact × Likelihood**

Or use the impact area scoring:
- Score each impact area (using Step 1 criteria)
- Weight by priority of impact areas
- Calculate total risk score

**Likelihood Assessment:**
| Level | Description |
|-------|-------------|
| High | Expected to occur, has happened before |
| Medium | Possible, some indicators exist |
| Low | Unlikely given current conditions |

### Step 8: Select Mitigation Approach

**Purpose:** Decide how to address each risk.

**Options:**
| Approach | When to Use | Example |
|----------|-------------|---------|
| **Mitigate** | Can reduce likelihood or impact | Implement encryption, access controls |
| **Defer** | Need more information | Research solutions before deciding |
| **Accept** | Cost exceeds benefit | Document acceptance and rationale |
| **Transfer** | Can shift to third party | Insurance, outsource to specialist |

**Mitigation Planning:**
- Identify specific controls
- Assign ownership
- Set timeline
- Estimate resources needed
- Define success criteria

---

## Adapting OCTAVE for Civil Society

### Simplified Impact Areas

| Area | Civil Society Framing |
|------|----------------------|
| **People Safety** | Could anyone be harmed? (staff, beneficiaries, sources) |
| **Mission Impact** | Can we still do our work? |
| **Trust & Reputation** | Will stakeholders still support us? |
| **Resources** | Can we afford the consequences? |

### Key Questions for Civil Society Risk Assessment

1. **What information would hurt people if exposed?**
   - Beneficiary identities, locations
   - Source identities (for journalism/investigations)
   - Staff home addresses, personal information

2. **What information would hurt the mission if exposed?**
   - Strategic plans adversaries could counter
   - Investigation findings before publication
   - Internal communications that could be weaponized

3. **What information would hurt the organization if exposed?**
   - Donor information
   - Financial records
   - Internal disagreements or challenges

4. **Where does this information live?**
   - Which systems, devices, people have it?
   - Which third parties have access?
   - Where does it travel?

5. **Who wants this information or wants to harm us?**
   - Map to adversary profiles
   - Consider opportunistic threats too

---

## OCTAVE Allegro Worksheets

### Worksheet 1: Risk Measurement Criteria

| Impact Area | Priority (1-5) | High | Medium | Low |
|-------------|----------------|------|--------|-----|
| Safety | | | | |
| Mission | | | | |
| Reputation | | | | |
| Financial | | | | |
| Legal | | | | |

### Worksheet 2: Information Asset Profile

| Field | Response |
|-------|----------|
| Asset Name | |
| Description | |
| Owner | |
| Confidentiality Requirement | |
| Integrity Requirement | |
| Availability Requirement | |
| Most Important Requirement | |

### Worksheet 3: Asset Containers

| Container | Type | Assets Held | Access Controls |
|-----------|------|-------------|-----------------|
| | Technical/Physical/Human | | |

### Worksheet 4: Risk Register

| Scenario | Impact (H/M/L) | Likelihood (H/M/L) | Score | Priority | Mitigation |
|----------|----------------|-------------------|-------|----------|------------|
| | | | | | |

---

## Resources

- SEI OCTAVE Page: https://www.sei.cmu.edu/our-work/octave/
- OCTAVE Allegro Training: https://www.sei.cmu.edu/training/
- Original Report: https://resources.sei.cmu.edu/library/asset-view.cfm?assetid=8419

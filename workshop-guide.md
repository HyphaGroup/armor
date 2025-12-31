# Threat Modeling Workshop Guide

A facilitation guide for running threat modeling sessions using the ARMOR methodology.

---

## Overview

**Total time:** ~3.5 hours (210 min) for core + 45-60 min per module

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

| Asset | Value | Risk Score Value | Primary Requirement |
|-------|-------|------------------|---------------------|
| | Critical/High/Medium/Low | 3/2/2/1 | Confidentiality/Integrity/Availability |

> "For each asset—how critical is it? And is the main concern keeping it secret (confidentiality), keeping it accurate (integrity), or keeping it accessible (availability)?"

**Value mapping for risk scoring:**
- **Critical** → 3 (highest risk weight)
- **High** → 2
- **Medium** → 2
- **Low** → 1 (lowest risk weight)

### Output
- [ ] Asset list with descriptions
- [ ] Container mapping for critical assets
- [ ] Value ratings (will be used in Session 5 risk scoring)
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

| Threat | Likelihood | Risk Score Value |
|--------|------------|------------------|
| | High/Medium/Low | 3/2/1 |

**Likelihood mapping for risk scoring:**
- **High** (Expected/Active) → 3
- **Medium** (Possible) → 2
- **Low** (Unlikely) → 1

### Output
- [ ] Relevant threats identified
- [ ] Examples documented
- [ ] Threats linked to adversaries
- [ ] Likelihood ratings (will be used in Session 5 risk scoring)

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
- **Asset Value**: Use values from Session 2 (Critical=3, High/Medium=2, Low=1)
- **Likelihood**: Use ratings from Session 4 (Expected/Active=3, Possible=2, Unlikely=1)
- **Vulnerability**: From step 2 above (Exposed=3, Some gaps=2, Well-protected=1)
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

Each module provides depth in a specific area. Select modules based on what emerged from the core session and organizational context.

---

### Module: Deep Adversary Profiling (45 min)

**When to use:** 
- Confirmed or highly suspected targeted threats
- A specific adversary has targeted similar organizations
- Need to anticipate adversary behavior, not just react
- Core session identified nation-state or persistent threat actors

**Facilitator Notes:**
- This module goes deep on 1-3 adversaries, not all of them
- Focus on adversaries marked "confirmed" or "likely" in core session
- The goal is actionable intelligence, not academic analysis

### Process

**1. Select Adversaries for Deep Dive (5 min)**

> "Based on our core session, which adversaries should we analyze in detail? I'd recommend focusing on [highest relevance ones]."

Select 1-3 adversaries for detailed analysis.

**2. Adversary Identity Analysis (10 min)**

For each selected adversary:

> "Let's build a detailed picture of who this adversary is."

| Element | Analysis |
|---------|----------|
| Specific identification | Who specifically? (if known) |
| Organizational structure | How are they organized? |
| Primary motivation | What drives them? |
| Objectives | What are they trying to achieve? |
| Resources | What level of resources do they have? |
| Persistence | One-time or persistent threat? |
| History | What's their track record? |

> "Have they targeted organizations like yours before? What happened?"

**3. Capability Analysis (10 min)**

> "What can this adversary actually do? Let's assess their capabilities across different domains."

| Capability Area | Level (Advanced/Moderate/Basic/Minimal) | Known Techniques |
|-----------------|----------------------------------------|------------------|
| Technical (hacking, malware) | | |
| Social Engineering (phishing, pretexting) | | |
| Information Operations (narrative attacks, harassment) | | |
| Physical (surveillance, intimidation) | | |
| Legal (lawsuits, regulatory pressure) | | |

> "What specific techniques have they used against similar organizations?"

**4. Infrastructure Analysis (5 min)**

> "What resources do they use to conduct operations?"

| Infrastructure Type | Details |
|---------------------|---------|
| Technical infrastructure | Owned servers? Compromised? Commercial services? |
| Social infrastructure | Fake accounts? Front organizations? Proxy actors? |
| Physical presence | On the ground capability? |
| Known indicators | Domains, accounts, patterns to watch for |

**5. Targeting Analysis (10 min)**

> "Why would they target YOUR organization specifically?"

| Question | Analysis |
|----------|----------|
| Why you? | What about your work attracts their attention? |
| What do they want? | Information? Disruption? Discrediting? |
| Primary targets within org | Who/what would they go after first? |
| Stepping stone value | Could you be used to reach others? (sources, partners, beneficiaries) |
| Likely attack vectors | Based on capabilities, how would they most likely come at you? |
| Targeting indicators | What would signal you're being targeted? |

**6. Anticipated Scenarios (5 min)**

> "Based on this analysis, what attack scenarios should we anticipate?"

| Scenario | Likelihood | Primary Technique | What to Watch For |
|----------|------------|-------------------|-------------------|
| | High/Med/Low | | |

### Output
- [ ] Detailed profile for each selected adversary
- [ ] Anticipated attack scenarios
- [ ] Targeting indicators to monitor
- [ ] Priority defensive measures for each adversary

---

### Module: Information Operations (45 min)

**When to use:**
- Public-facing work (advocacy, journalism, research)
- Experienced or anticipate coordinated online attacks
- Staff have been targeted with harassment
- Work is politically contentious or opposes powerful interests
- Reputation is a strategic asset requiring protection

**Facilitator Notes:**
- This is often highly relevant but overlooked
- Be sensitive—harassment discussion can be difficult
- Focus on organizational resilience, not individual blame

### Process

**1. Exposure Assessment (10 min)**

> "Let's understand your organization's exposure to information operations."

**Public Profile:**

| Factor | Assessment |
|--------|------------|
| Overall visibility | High / Medium / Low |
| Social media presence | Platforms: _____ Followers: _____ |
| Media coverage | Frequent / Occasional / Rare |
| Contentious positions | What positions attract opposition? |
| Known opposition | Who actively opposes your work? |

**Staff Exposure:**

> "Who on your team is most visible? Has anyone experienced online harassment?"

| Role | Public Visibility | Personal Social Media | Harassment History |
|------|-------------------|----------------------|-------------------|
| | High/Med/Low | Y/N | Y/N |

**2. Threat Category Assessment (15 min)**

> "Let's walk through specific types of information operations and assess your exposure to each."

**Narrative Attacks:**
> "Have you seen false claims about your organization—funding, motives, foreign connections, conspiracies?"

| Assessment | Response |
|------------|----------|
| Currently experiencing? | Y/N |
| Examples of current narratives | |
| Anticipated future narratives | |
| Vulnerability level | High / Medium / Low |

**Impersonation:**
> "Have you seen fake accounts using your organization's name or posing as staff?"

| Assessment | Response |
|------------|----------|
| Known impersonation accounts/sites | |
| Brand monitoring in place? | Y/N |
| Vulnerability level | High / Medium / Low |

**Harassment & Doxing:**
> "Have staff experienced coordinated harassment, pile-ons, or threats? Has personal information been exposed?"

| Assessment | Response |
|------------|----------|
| Harassment incidents | |
| Doxing incidents | |
| Staff personal info exposure | High / Medium / Low |
| Support resources available? | Y/N |
| Vulnerability level | High / Medium / Low |

**Amplification & Coordination:**
> "Have you seen bot activity or coordinated inauthentic behavior targeting you?"

| Assessment | Response |
|------------|----------|
| Bot activity observed | Y/N |
| Coordinated pile-ons | Y/N |
| Vulnerability level | High / Medium / Low |

**Platform Manipulation:**
> "Have you experienced mass reporting of your content or accounts? SEO attacks?"

| Assessment | Response |
|------------|----------|
| Mass reporting experienced | Y/N |
| Content/accounts removed | Y/N |
| Platform contacts established? | Y/N |
| Vulnerability level | High / Medium / Low |

**Document Operations:**
> "Are you concerned about internal documents being stolen and leaked, possibly altered?"

| Assessment | Response |
|------------|----------|
| Leak risk | High / Medium / Low |
| Most sensitive documents | |
| Vulnerability level | High / Medium / Low |

**3. Monitoring Capability (5 min)**

> "Can you detect when these attacks are happening?"

| Monitoring Type | In Place? | Tool/Method | Who Monitors? |
|-----------------|-----------|-------------|---------------|
| Social media mentions | | | |
| News/media mentions | | | |
| Staff name monitoring | | | |
| Domain monitoring | | | |
| Hashtag monitoring | | | |

> "What are the gaps in your monitoring?"

**4. Response Capability (10 min)**

> "If an attack happens, are you ready to respond?"

| Element | In Place? | Notes |
|---------|-----------|-------|
| Response decision framework (when to respond vs. ignore) | | |
| Pre-drafted responses | | |
| Designated spokesperson | | |
| Platform reporting knowledge | | |
| Documentation process (screenshots, archives) | | |
| Legal support available | | |
| PR/comms support available | | |
| Staff harassment support protocol | | |
| Mental health resources | | |

**5. Recommendations (5 min)**

> "Based on this assessment, here are the priority improvements..."

| Recommendation | Category | Priority | Effort |
|----------------|----------|----------|--------|
| | Monitoring/Prevention/Response/Staff Support | High/Med/Low | Low/Med/High |

### Output
- [ ] Info ops exposure assessment
- [ ] Threat category vulnerability ratings
- [ ] Monitoring gaps identified
- [ ] Response capability gaps identified
- [ ] Priority recommendations
- [ ] Staff support needs identified

---

### Module: OPSEC Analysis (45 min)

**When to use:**
- Protecting sensitive sources or beneficiaries
- Operating in hostile environments (authoritarian contexts, conflict zones)
- Investigating powerful actors who might counter-investigate
- Staff travel to high-risk locations
- Work requires operational secrecy

**Facilitator Notes:**
- Think like an adversary—what would they find?
- OPSEC failures enable other attacks
- Balance security with mission (some exposure may be necessary)

### Process

**1. Public Exposure Audit (15 min)**

> "Let's see what an adversary could learn about you through open-source research."

**Website Exposure:**

| Element | Public? | Concern Level |
|---------|---------|---------------|
| Staff names listed | Y/N | High/Med/Low |
| Staff contact info | Y/N | |
| Office address | Y/N | |
| Partner organizations | Y/N | |
| Donor list | Y/N | |
| Program/activity details | Y/N | |

**Social Media Exposure:**

| Element | Assessment |
|---------|------------|
| Organization accounts | List: |
| Staff personal accounts linked to org? | Y/N |
| Location sharing visible | Y/N |
| Activity patterns visible (posting times, locations) | Y/N |

**Regulatory/Public Filings:**

| Filing Type | Public? | What It Reveals |
|-------------|---------|-----------------|
| 990 (US nonprofits) | Y/N | |
| Corporate registration | Y/N | |
| Other | | |

**Media Coverage:**

| Element | Assessment |
|---------|------------|
| Staff quoted by name | Y/N |
| Activities/locations covered | Y/N |
| What could adversary learn from media? | |

**2. Digital Footprint (10 min)**

> "What does your technical footprint reveal?"

| Element | Assessment | Reveals |
|---------|------------|---------|
| Domain WHOIS | Private/Public | |
| Email infrastructure (MX records) | | Provider, setup |
| Third-party services | List: | |
| Data broker presence | Checked? | |

**3. Operational Patterns (10 min)**

> "What predictable patterns could an adversary exploit?"

| Pattern Type | Predictable? | Publicly Known? | Exploitable? |
|--------------|--------------|-----------------|--------------|
| Regular meetings (timing, location) | | | |
| Travel patterns | | | |
| Communication patterns | | | |
| Funding cycles | | | |
| Event schedules | | | |

**4. Trust Relationships (5 min)**

> "Which relationships could be exploited through impersonation or infiltration?"

| Relationship | Could Be Exploited? | How? | Mitigation |
|--------------|---------------------|------|------------|
| | Y/N | | |

**5. OPSEC Vulnerabilities Summary (5 min)**

For each finding, apply OPSEC thinking:
- Does an adversary want this information?
- Can they collect it?
- What can they do with it?
- What's the potential harm?
- Can we reduce exposure without harming mission?

| Vulnerability | What Adversary Learns | Severity | Remediation | Priority |
|---------------|----------------------|----------|-------------|----------|
| | | High/Med/Low | | High/Med/Low |

### Output
- [ ] Public exposure audit complete
- [ ] Digital footprint assessed
- [ ] Operational patterns identified
- [ ] Trust relationships assessed
- [ ] OPSEC vulnerabilities prioritized
- [ ] Remediation recommendations

---

### Module: Incident Response Capability (45 min)

**When to use:**
- High-risk environment where incidents are likely
- History of security incidents
- Limited current incident response capability
- Want to establish protocols before they're needed

**Facilitator Notes:**
- Preparation is 80% of incident response
- Don't overwhelm—focus on minimum viable capability first
- Include information operations incidents, not just technical

### Process

**1. Define "Incident" (5 min)**

> "First, let's define what constitutes a security incident for YOUR organization."

| Question | Response |
|----------|----------|
| What's an incident? | |
| Examples that would trigger response | |
| Examples that wouldn't | |

**2. Preparation Assessment (15 min)**

> "Let's assess your preparation—what's in place before an incident happens?"

| Element | In Place? | Current State | Priority Gap? |
|---------|-----------|---------------|---------------|
| Incident response plan documented | Y/N/Partial | | |
| Response team defined | Y/N | | |
| Roles and decision authority clear | Y/N | | |
| Contact lists (internal) | Y/N | | |
| Contact lists (legal counsel) | Y/N | | |
| Contact lists (IT support) | Y/N | | |
| Contact lists (digital security help) | Y/N | | |
| Contact lists (platform contacts) | Y/N | | |
| Contact lists (peer organizations) | Y/N | | |
| Primary communication channel | Y/N | What: | |
| Backup/secure communication channel | Y/N | What: | |
| Staff trained on incident reporting | Y/N | | |
| Tabletop exercises conducted | Y/N | When: | |
| Documentation templates ready | Y/N | | |

**3. Detection Capability (10 min)**

> "Would you know if something was happening?"

| Monitoring Type | In Place? | Tool/Method | Who Monitors? | Gap? |
|-----------------|-----------|-------------|---------------|------|
| Email security alerts | | | | |
| Account login alerts | | | | |
| Endpoint protection | | | | |
| Social media monitoring | | | | |
| News/mention monitoring | | | | |
| Domain monitoring | | | | |
| Data breach monitoring | | | | |

> "Do staff know what to report and how?"

| Question | Response |
|----------|----------|
| Staff know how to report incidents? | Y/N |
| Reporting method | |

**4. Response Capability (10 min)**

> "If something happens, can you respond effectively?"

| Capability | In Place? | Notes |
|------------|-----------|-------|
| Can isolate compromised systems | Y/N | |
| Can revoke/reset credentials quickly | Y/N | |
| Can block accounts/access | Y/N | |
| Backups exist | Y/N | Frequency: |
| Backups tested | Y/N | Last test: |
| Recovery time estimate | | |
| IT support available | Y/N | Contact: |
| Digital security support available | Y/N | Contact: |
| Legal support available | Y/N | Contact: |
| PR/comms support available | Y/N | Contact: |

**5. Post-Incident Learning (5 min)**

> "After an incident, how do you learn and improve?"

| Element | In Place? |
|---------|-----------|
| Post-incident review process | Y/N |
| Lessons learned documented | Y/N |
| Threat model updated after incidents | Y/N |

**6. Incident History Review**

> "Have you experienced incidents before? What happened and what did you learn?"

| Date | Type | Summary | Lessons |
|------|------|---------|---------|
| | | | |

**7. Gap Prioritization**

> "Based on this assessment, here are the priority gaps..."

| Gap | Priority | Owner | Target Date |
|-----|----------|-------|-------------|
| | Critical/High/Med/Low | | |

**Minimum Viable Incident Response Checklist:**
- [ ] Incident definition exists
- [ ] Someone responsible for coordinating response
- [ ] Contact list (internal + key external)
- [ ] Secure communication channel identified
- [ ] Basic documentation template
- [ ] Post-incident review process

### Output
- [ ] Incident definition documented
- [ ] Preparation gaps identified
- [ ] Detection gaps identified
- [ ] Response capability gaps identified
- [ ] Contact lists verified/created
- [ ] Priority improvements with owners

---

### Module: Technical Deep-Dive (60 min)

**When to use:**
- Complex IT infrastructure requiring detailed assessment
- Preparing for a technical security audit
- Specific technical concerns emerged from core session
- Want to assess specific control areas
- Have technical staff who can implement findings

**Facilitator Notes:**
- Need someone technical in the room
- Focus on fundamentals before advanced measures
- Quick wins matter—identify easy improvements

### Process

**1. Infrastructure Mapping (20 min)**

> "Let's map your technical infrastructure—what systems and services do you use?"

**Systems Inventory:**

| System/Service | Type | Owner | Criticality | Managed By |
|----------------|------|-------|-------------|------------|
| | Endpoint/Server/Cloud/Mobile | | Critical/High/Med/Low | Internal/Third-party |

**Third-Party Services:**

| Service | Provider | Purpose | Data Shared | Security Reviewed? |
|---------|----------|---------|-------------|-------------------|
| | | | | Y/N |

**Data Flows:**

> "How does sensitive data move through your systems?"

| Data Type | Source | Destination | Encrypted? |
|-----------|--------|-------------|------------|
| | | | Y/N |

**2. Technical Assessment (25 min)**

> "Let's assess security controls in key areas."

**Authentication:**

| Question | Response | Gap? |
|----------|----------|------|
| MFA enabled on critical accounts? | Y/N/Partial | |
| MFA coverage (% of accounts) | | |
| Password manager in use? | Y/N | |
| Password policy exists? | Y/N | |
| SSO implemented? | Y/N | |

**Encryption:**

| Question | Response | Gap? |
|----------|----------|------|
| Device encryption (laptops)? | Y/N/Partial | |
| Device encryption (mobile)? | Y/N/Partial | |
| Data at rest encrypted? | Y/N/Partial | |
| Data in transit encrypted? | Y/N | |
| Email encryption available? | Y/N | |

**Patching & Updates:**

| Question | Response | Gap? |
|----------|----------|------|
| Auto-updates enabled (endpoints)? | Y/N/Partial | |
| Server/infrastructure patching? | Regular/Irregular/None | |
| Patch timeline for critical vulns? | | |

**Backups:**

| Question | Response | Gap? |
|----------|----------|------|
| Critical data backed up? | Y/N | |
| Backup frequency | | |
| Backups tested? | Y/N | Last: |
| Offsite/offline backup? | Y/N | |
| Recovery time estimate | | |

**Monitoring & Logging:**

| Question | Response | Gap? |
|----------|----------|------|
| Logging enabled? | Y/N/Partial | |
| Alerts configured? | Y/N | |
| Who monitors? | | |
| Log retention | | |

**Access Control:**

| Question | Response | Gap? |
|----------|----------|------|
| Least privilege implemented? | Y/N/Partial | |
| Access reviews conducted? | Y/N | Frequency: |
| Offboarding process defined? | Y/N | |
| Admin access limited? | Y/N | # of admins: |

**Quick Assessment Questions:**
- [ ] Is MFA enabled on all critical accounts?
- [ ] Are all devices encrypted?
- [ ] Are systems set to auto-update?
- [ ] When were backups last tested?
- [ ] Who has admin access, and is it necessary?
- [ ] What happens to access when someone leaves?

**3. Vulnerability Summary (10 min)**

> "Let's summarize the technical vulnerabilities we've identified."

| Vulnerability | Area | Severity | Exploitability | Remediation |
|---------------|------|----------|----------------|-------------|
| | Auth/Encryption/Patching/Backups/Monitoring/Access | Critical/High/Med/Low | Easy/Moderate/Difficult | |

**4. Prioritized Roadmap (5 min)**

> "Let's prioritize improvements by impact and effort."

| Priority | Improvement | Area | Effort | Impact | Owner | Target |
|----------|-------------|------|--------|--------|-------|--------|
| 1 | | | Low/Med/High | Critical/High/Med/Low | | |
| 2 | | | | | | |
| 3 | | | | | | |

**Overall Assessment:**

| Area | Score |
|------|-------|
| Authentication | Strong / Adequate / Weak / None |
| Encryption | Strong / Adequate / Weak / None |
| Patching | Strong / Adequate / Weak / None |
| Backups | Strong / Adequate / Weak / None |
| Monitoring | Strong / Adequate / Weak / None |
| Access Control | Strong / Adequate / Weak / None |
| **Overall Technical Posture** | Strong / Adequate / Weak / None |

### Output
- [ ] Infrastructure inventory
- [ ] Technical assessment by area
- [ ] Vulnerabilities identified and prioritized
- [ ] Improvement roadmap with owners
- [ ] Overall technical posture rating

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

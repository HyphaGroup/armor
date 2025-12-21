# Civil Society Threat Modeling Methodology

A structured approach for civil society organizations to identify, assess, and address security threats across digital, physical, and information domains.

## Purpose

Civil society organizations—human rights groups, journalists, advocacy organizations, legal aid providers—face sophisticated threats from well-resourced adversaries while operating with limited security budgets and capacity. This methodology provides:

- A systematic way to understand your specific threat landscape
- Prioritization based on mission impact, not generic risk scores
- Actionable outputs that match your resource constraints
- A foundation for ongoing security improvement

*This methodology draws on established security frameworks (OCTAVE, Diamond Model, STRIDE, DISARM, NIST 800-61) but presents them as a unified approach without requiring practitioners to learn those frameworks.*

---

## Core Concepts

### Mission-First Security

Security exists to protect your mission. This methodology starts with understanding what you're trying to achieve and what would constitute mission failure, then works backward to identify what needs protection and from whom.

### The Threat Modeling Chain

```
MISSION → ASSETS → ADVERSARIES → THREATS → RISKS → MITIGATIONS
   ↓         ↓          ↓            ↓         ↓          ↓
 Why you   What to    Who might   What they  Which      What to
 exist     protect    target you  could do   matter     do about it
```

Each element informs the next:
- Your **mission** determines which **assets** are critical
- Your **assets** and mission attract specific **adversaries**
- Your **adversaries** have particular **threats** they can deploy
- **Threats** against **assets** create **risks**
- **Risks** are addressed through **mitigations**

### Adversary-Informed Defense

Rather than defending against all possible threats equally, this methodology identifies your likely adversaries first, then focuses defensive resources on the threats those adversaries actually use.

---

## Methodology Components

### 1. Mission & Impact Framework

**Purpose:** Establish the context for all security decisions.

**Key elements:**
- **Mission statement**: What you're trying to achieve
- **Core activities**: Programs and work that fulfill the mission
- **Impact areas**: Categories of harm (safety, mission, reputation, financial, legal, partners)
- **Impact thresholds**: What HIGH/MEDIUM/LOW impact means for your organization

**Why it matters:** A threat that would be minor for one organization could be catastrophic for another. Defining your impact areas and thresholds ensures risk scoring reflects YOUR priorities, not generic benchmarks.

### 2. Asset Identification

**Purpose:** Identify what needs protection.

**Key elements:**
- **Information assets**: Data, documents, communications, credentials
- **Containers**: Where assets live (technical systems, physical locations, people's knowledge)
- **Security requirements**: Confidentiality, integrity, availability needs
- **Value**: How critical each asset is to mission (Critical/High/Medium/Low → maps to 3/2/2/1 in risk scoring)

**Why it matters:** You can't protect everything equally. Asset identification focuses attention on what actually matters—often information about people (beneficiaries, sources, staff) rather than generic "data." Asset values established here carry forward into risk scoring.

### 3. Adversary Profiling

**Purpose:** Understand who might target you and why.

**Key elements:**
- **Identity**: Who they are (category or specific actor)
- **Motivation**: Why they'd target you
- **Capability**: What they can do (technical, social, physical, legal, information operations)
- **Infrastructure**: Resources they use
- **Targeting rationale**: Why YOU specifically

**Adversary categories:**
| Category | Typical Motivation |
|----------|-------------------|
| Nation-State / Intelligence | Political control, surveillance, disruption |
| Ideological Opposition | Discrediting, silencing, harassment |
| Cybercriminal | Financial gain |
| Insider | Revenge, ideology, financial |
| Competitor / Opposing Org | Competitive advantage, policy wins |
| Opportunistic | Whatever value can be extracted |

**Why it matters:** Different adversaries use different techniques. A nation-state and an ideological opposition group both might target you, but they'll do it differently. Understanding your adversaries focuses defense on relevant threats.

### 4. Threat Mapping

**Purpose:** Identify specific threats relevant to your situation.

**Threat categories:**

| Domain | Threats |
|--------|---------|
| **Account & Access** | Phishing, account takeover, unauthorized access, insider misuse |
| **Data & Information** | Data breach, tampering, loss, surveillance |
| **Disruption** | Service disruption, infrastructure attacks |
| **Information & Reputation** | Narrative attacks, impersonation, harassment, amplification, platform manipulation, document leaks |
| **Physical** | Intrusion, surveillance, intimidation |
| **Operational** | Human error, technical failure, natural disaster, third-party failure |

**Why it matters:** Civil society faces threats across multiple domains—not just "cyber." Information operations (narrative attacks, harassment) and physical threats are often as significant as technical attacks. This methodology covers all domains. Threat likelihood assessed here carries forward into risk scoring.

### 5. Risk Assessment

**Purpose:** Prioritize which threats to address.

**Risk formula:**
```
RISK = ASSET VALUE × THREAT LIKELIHOOD × VULNERABILITY
```

Each factor captures something distinct:
- **Asset Value**: How critical is what's at risk? (from asset identification)
- **Threat Likelihood**: How likely is this threat to be attempted? (from adversary capability and intent)
- **Vulnerability**: How exposed are we? (from control gaps and weaknesses)

**Risk scenario format:**
> "There is a risk that [adversary] could [threat action] affecting [asset], which would impact [mission area]."

**Scoring:**
- **Asset Value** (1-3): From asset identification—Low (1), Medium/High (2), Critical (3)
- **Likelihood** (1-3): From threat mapping—Unlikely (1), Possible (2), Expected/Active (3)
- **Vulnerability** (1-3): Assessed per risk—Well-protected (1), Some gaps (2), Exposed (3)
- **Risk Score** (1-27): Asset Value × Likelihood × Vulnerability
- **Priority**: Critical (18-27), High (10-17), Moderate (4-9), Low (1-3)

**Why it matters:** Three-factor scoring distinguishes between risks that are serious because of asset value vs. likelihood vs. exposure. This helps prioritize action—vulnerability is what you can actually change through mitigation.

### 6. Mitigation Planning

**Purpose:** Turn risks into actions.

**Key elements:**
- **Mitigation type**: Technical control, policy, training, etc.
- **Priority**: Based on risk score
- **Effort/Cost**: Resource requirements
- **Owner**: Who's responsible
- **Timeline**: When it will be done
- **Success criteria**: How you know it's working

**Why it matters:** Threat modeling without action is just worry. Mitigation planning ensures findings become improvements.

---

## Modular Structure

The methodology uses a modular structure to balance thoroughness with practicality.

### Core (Required)

Every organization completes the core:

| Session | Purpose | Time |
|---------|---------|------|
| Mission & Impact | Establish context | 30 min |
| Asset Identification | Identify what to protect | 45 min |
| Adversary Selection | Identify who might target you | 30 min |
| Threat Mapping | Map relevant threats | 45 min |
| Risk Prioritization | Assess vulnerability, score, prioritize | 45 min |

**Total core time:** ~3.5 hours

### Modules (As Needed)

Modules provide depth in specific areas based on organizational context:

| Module | When to Use | Time |
|--------|-------------|------|
| **Deep Adversary Profiling** | Confirmed targeted threats, need to anticipate behavior | 45 min |
| **Information Operations** | Facing narrative attacks, harassment, public-facing work | 45 min |
| **OPSEC Analysis** | Protecting sensitive sources/beneficiaries, hostile environment | 45 min |
| **Incident Response** | History of incidents, high-risk environment | 45 min |
| **Technical Deep-Dive** | Complex infrastructure, specific technical concerns | 60 min |

**Module selection:** Based on what emerges from the core session and organizational context.

---

## Module Descriptions

### Module: Deep Adversary Profiling

**Purpose:** Develop detailed understanding of specific adversaries to anticipate their behavior and prioritize defenses.

**When to use:**
- You have confirmed or strongly suspected targeted threats
- A specific adversary has targeted similar organizations
- You need to anticipate adversary behavior, not just react
- Core adversary selection identified nation-state or persistent threat actors

**Key elements:**

The module examines each relevant adversary through four interconnected lenses:

| Element | Key Questions |
|---------|---------------|
| **Adversary Identity** | Who are they? What's their organizational structure? What motivates them? How persistent are they? What's their history? |
| **Capability** | What technical, social, physical, and legal capabilities do they have? What techniques do they typically use? What's their sophistication level? |
| **Infrastructure** | What resources do they use? Own servers or compromised infrastructure? Social media accounts? Front organizations? Physical presence? |
| **Targeting Rationale** | Why would they target YOU specifically? What do they want from you? Are you a primary target or stepping stone to others? |

**Process:**
1. For each high-relevance adversary from the core session, work through all four elements
2. Research known campaigns against similar organizations
3. Identify the specific techniques this adversary uses
4. Document targeting indicators—how would you know if they're targeting you?
5. Map adversary capabilities to your specific assets and vulnerabilities

**Output:** Detailed adversary profiles that inform defensive priorities and help recognize targeting when it occurs.

**Why it matters:** Generic security advice treats all threats equally. Deep adversary profiling lets you focus on the specific techniques YOUR adversaries actually use, making limited security resources more effective.

---

### Module: Information Operations

**Purpose:** Assess exposure to and build resilience against narrative attacks, harassment campaigns, and reputation threats.

**When to use:**
- You do public-facing work (advocacy, journalism, research)
- You've experienced or anticipate coordinated online attacks
- Staff have been targeted with harassment
- Your work is politically contentious or opposes powerful interests
- You need to protect organizational reputation as a strategic asset

**Key elements:**

| Threat Category | What It Involves |
|-----------------|------------------|
| **Narrative Attacks** | False claims about funding, motives, foreign connections; conspiracy theories; distorted facts; fake research countering your work |
| **Impersonation** | Fake accounts using your name/brand; personas posing as supporters, journalists, or researchers; infiltration of your networks |
| **Harassment** | Coordinated pile-ons; identity-based attacks on staff; doxing (publishing private information); threats and intimidation |
| **Amplification** | Bot networks; coordinated inauthentic behavior; flooding your channels; swarming attacks on individuals |
| **Platform Manipulation** | Mass reporting to trigger account suspension; SEO attacks; algorithm gaming to suppress your content |
| **Document Operations** | Theft and selective leaking of internal documents; altered documents presented as leaks |

**Process:**
1. Assess current exposure: How visible is your organization? What's your attack surface?
2. Review past incidents: Have you experienced any of these attacks?
3. Map adversary capability: Which of your adversaries use information operations?
4. Identify vulnerable staff: Who is most exposed or likely to be targeted?
5. Evaluate monitoring capabilities: Can you detect when attacks are happening?
6. Assess response readiness: Do you have processes for responding to narrative attacks?

**Indicators to monitor:**
- Sudden spikes in hostile engagement
- Coordinated messaging with similar language/timing
- New accounts with limited history engaging with your content
- Content appearing on opposition media outlets
- Staff receiving unusual interview requests or partnership offers

**Output:** Information operations risk assessment, monitoring recommendations, and response protocols.

**Why it matters:** For many civil society organizations, reputation attacks and harassment are more common than technical intrusions. Information operations can silence voices, burn out staff, and undermine mission effectiveness without ever touching your computer systems.

---

### Module: OPSEC Analysis

**Purpose:** Identify what adversaries can learn about you through observation, and reduce exploitable information exposure.

**When to use:**
- You protect sensitive sources or beneficiaries whose exposure could cause harm
- You operate in hostile environments (authoritarian contexts, conflict zones)
- You investigate powerful actors who might counter-investigate you
- Staff travel to high-risk locations
- Your work requires operational secrecy

**Key elements:**

| Analysis Area | What to Examine |
|---------------|-----------------|
| **Public Exposure** | What's on your website (staff, addresses, partners, donors)? Social media presence? Regulatory filings? Media coverage? |
| **Operational Patterns** | Predictable meetings, travel, communication patterns? Funding cycles visible? Regular activities that could be anticipated? |
| **Digital Footprint** | Domain registration (WHOIS private?)? Email infrastructure visible? Third-party services revealing information? |
| **Trust Relationships** | Which relationships could be exploited through impersonation? Where does trust create vulnerability? |
| **Human Sources** | How are sources/beneficiaries contacted? What trails exist? How is their information protected? |

**Process:**
1. Map what an adversary could learn through open-source research
2. Identify operational patterns that could be predicted or exploited
3. Assess digital footprint and what it reveals
4. Examine trust relationships for exploitation potential
5. For each finding, assess: What could an adversary do with this information?
6. Prioritize OPSEC vulnerabilities by severity and addressability

**OPSEC Thinking:**
For each piece of information you expose, ask:
- Does an adversary want this information?
- Can they collect it?
- What can they do with it?
- What's the potential harm?
- Can we reduce exposure without harming mission?

**Output:** OPSEC vulnerability assessment with prioritized recommendations for reducing exploitable exposure.

**Why it matters:** Adversaries don't just attack—they first gather information. OPSEC analysis helps you see yourself as an adversary sees you, identifying what you're revealing that could enable targeting, social engineering, or physical threats.

---

### Module: Incident Response

**Purpose:** Assess and build capability to detect, respond to, and recover from security incidents.

**When to use:**
- You've experienced security incidents and want to improve response
- You operate in a high-risk environment where incidents are likely
- You have limited incident response experience and need to build basic capability
- You want to establish response protocols before they're needed

**Key elements:**

The module assesses capability across four phases:

| Phase | Key Questions |
|-------|---------------|
| **Preparation** | Is there an incident response plan? Are roles defined? Are contact lists current? Has staff been trained? Have you done tabletop exercises? |
| **Detection** | What monitoring exists? Would you know if something was happening? Do staff know what to report and how? What are the gaps? |
| **Response** | Can you contain incidents (isolate systems, revoke access)? Do backups exist and work? Is external support available (IT, legal, digital security)? |
| **Recovery & Learning** | Is there a post-incident review process? Do lessons lead to improvements? Does an incident trigger threat model updates? |

**Incident types to consider:**
- Account compromise (email, social media, systems)
- Device compromise (malware, theft, seizure)
- Data breach (unauthorized access, exfiltration)
- Ransomware
- Information operations (narrative attacks, impersonation)
- Physical incidents (office intrusion, surveillance, intimidation)
- Insider incidents

**Process:**
1. Define what constitutes an "incident" for your organization
2. Assess current capability in each phase
3. Review past incidents and how they were handled
4. Identify gaps and priorities for improvement
5. Develop or update incident response plan
6. Identify external resources and support contacts

**Minimum viable incident response:**
- A definition of what's an incident
- Someone responsible for coordinating response
- A contact list (internal, IT support, legal, digital security help)
- A secure communication channel for incident discussion
- Basic documentation template
- Post-incident review process

**Output:** Incident response capability assessment, gap analysis, and improvement roadmap.

**Why it matters:** Incidents will happen. The difference between a manageable incident and a catastrophe often comes down to preparation and response capability. This module builds readiness before it's tested.

---

### Module: Technical Deep-Dive

**Purpose:** Conduct detailed assessment of technical infrastructure, security controls, and technical vulnerabilities.

**When to use:**
- You have complex IT infrastructure requiring detailed assessment
- You're preparing for a technical security audit
- Specific technical concerns emerged from the core session
- You want to assess specific control areas (authentication, encryption, backups, etc.)
- You have technical staff who can implement findings

**Key elements:**

| Assessment Area | What to Examine |
|-----------------|-----------------|
| **Infrastructure Mapping** | Systems inventory, data flows, access controls, third-party services, cloud vs. on-premise |
| **Authentication** | MFA deployment and coverage, password policies, credential management, SSO |
| **Encryption** | Data at rest, data in transit, email encryption, device encryption |
| **Patching** | Update policies, auto-update status, vulnerability management |
| **Backups** | Backup frequency, testing, offsite/offline copies, recovery time |
| **Monitoring** | Logging, alerting, who monitors, detection capabilities |
| **Access Control** | Least privilege, access reviews, offboarding processes |

**Process:**
1. Map technical infrastructure—what systems, services, and data flows exist?
2. For each assessment area, document current state
3. Identify gaps between current state and appropriate security level
4. Prioritize technical vulnerabilities by severity and exploitability
5. Develop prioritized technical improvement roadmap
6. Assign ownership and timelines

**Quick assessment questions:**
- Is MFA enabled on all critical accounts?
- Are devices encrypted?
- Are systems set to auto-update?
- When were backups last tested?
- Who has admin access, and is it necessary?
- What happens to access when someone leaves?

**Output:** Technical vulnerability assessment, prioritized improvement roadmap with effort estimates.

**Why it matters:** Technical controls are the foundation of digital security. This module provides detailed assessment for organizations with the capacity to implement technical improvements, ensuring the basics are covered before advanced measures.

---

## Outputs

### Threat Model Profile

The methodology produces a structured **threat model profile** containing:

```
profile/
├── Organization metadata
├── Mission & impact framework
├── Asset inventory
├── Adversary profiles
├── Threat mapping
├── Risk register
├── OPSEC assessment (if module completed)
├── Response capability (if module completed)
└── Mitigation plan
```

### Structured Data

All outputs conform to JSON schemas enabling:
- **Structured capture**: Consistent data format across engagements
- **Platform integration**: Foundation for interactive tools
- **Agent-assisted processing**: LLM extraction from transcripts/notes
- **Cross-organization analysis**: Anonymized pattern identification

See `/schemas/` for detailed schema definitions.

---

## Lifecycle

Threat modeling is not a one-time event.

### Initial Assessment
Complete core + relevant modules to establish baseline.

### Ongoing Maintenance
- **Monthly**: Review high-priority mitigations
- **Quarterly**: Check for emerging threats, update adversary relevance
- **Annually**: Full threat model refresh

### Triggered Reviews
Update after:
- Significant security incidents
- Major organizational changes (new programs, new regions)
- Changes in operating environment
- New adversary activity against similar organizations

---

## Implementation Paths

### Facilitated Workshop
Run the methodology as a structured workshop with an external facilitator. Best for initial assessment or annual refresh.

### Self-Guided
Organization works through the methodology independently using the workshop guide. Best for organizations with some security capacity.

### Agent-Assisted
Use AI assistance to process transcripts, notes, or conversations into structured threat model profiles. Best for rapid assessments or ongoing updates.

### Platform (Future)
Interactive web platform for building and maintaining threat model profiles. Enables ongoing updates, incident tracking, and peer comparison.

---

## Principles

### Accessibility
- Plain language over jargon
- Mission impact over technical metrics
- Actionable outputs over comprehensive documentation

### Proportionality
- Match defensive investment to actual threat
- Accept some risks rather than trying to eliminate all
- Focus on high-impact, achievable mitigations

### Adaptability
- Modular structure fits different needs
- Scales from individual practitioners to large organizations
- Works across different civil society contexts

### Sustainability
- Outputs that can be maintained over time
- Clear ownership of risks and mitigations
- Integration with existing organizational processes

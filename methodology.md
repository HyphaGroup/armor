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
- **Value**: How critical each asset is to mission

**Why it matters:** You can't protect everything equally. Asset identification focuses attention on what actually matters—often information about people (beneficiaries, sources, staff) rather than generic "data."

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

**Why it matters:** Civil society faces threats across multiple domains—not just "cyber." Information operations (narrative attacks, harassment) and physical threats are often as significant as technical attacks. This methodology covers all domains.

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
- **Asset Value** (1-3): Low (1), Medium (2), High/Critical (3)
- **Likelihood** (1-3): Unlikely (1), Possible (2), Expected/Active (3)
- **Vulnerability** (1-3): Well-protected (1), Some gaps (2), Exposed (3)
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

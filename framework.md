# Civil Society Threat Modeling Framework

A security framework for civil society organizations to systematically identify, assess, and address threats across digital, physical, and information domains.

---

## Purpose

Civil society organizations—human rights groups, journalists, advocacy organizations, legal aid providers—face sophisticated threats from well-resourced adversaries while operating with limited security budgets and capacity.

This framework provides:

- A systematic way to understand your specific threat landscape
- Prioritization based on mission impact, not generic risk scores
- Actionable outputs that match your resource constraints
- A foundation for ongoing security improvement

---

## Core Principles

### Mission-First Security

Security exists to protect your mission, not the other way around. Every security decision should be evaluated against the question: *does this help us achieve our mission more safely?*

Start with what you're trying to achieve and what would constitute mission failure, then work backward to identify what needs protection and from whom.

### Adversary-Informed Defense

Rather than defending against all possible threats equally, identify your likely adversaries first, then focus defensive resources on the threats those adversaries actually use.

Generic security advice wastes resources. Understanding YOUR adversaries lets you prioritize the specific techniques they employ.

### Proportional Response

Match defensive investment to actual threat levels. Not every risk requires mitigation—some should be accepted, transferred, or monitored. Focus limited resources on high-impact, achievable improvements.

### Accessibility Over Jargon

Security frameworks should be usable by the people who need them. Plain language, mission-grounded explanations, and practical outputs matter more than technical comprehensiveness.

### Sustainability

Security is not a one-time project. Outputs should be maintainable over time, with clear ownership, regular review cycles, and integration into existing organizational processes.

### Well-Being as Security Foundation

Security work that burns out staff is not sustainable security. The stress of working under threat, responding to incidents, and maintaining vigilance takes a toll. Effective security strategies acknowledge this by:

- Recognizing that staff well-being is a security asset (exhausted people make mistakes)
- Building support systems before they're needed (mental health resources, peer support)
- Distributing security responsibilities rather than concentrating them
- Celebrating security wins, not just responding to failures

This framework focuses on threat modeling, not well-being programming. For organizations seeking deeper integration of well-being and security, see [Holistic Security](https://holistic-security.tacticaltech.org/) by Tactical Tech.

---

## The Threat Modeling Chain

The framework follows a logical chain where each element informs the next:

```
MISSION → ASSETS → ADVERSARIES → THREATS → RISKS → MITIGATIONS
   ↓         ↓          ↓            ↓         ↓          ↓
 Why you   What to    Who might   What they  Which      What to
 exist     protect    target you  could do   matter     do about it
```

**Flow of logic:**

1. Your **mission** determines which **assets** are critical
2. Your **assets** and mission attract specific **adversaries**
3. Your **adversaries** have particular **threats** they can deploy
4. **Threats** against **assets** create **risks**
5. **Risks** are addressed through **mitigations**

This chain ensures security work remains grounded in organizational reality rather than abstract threat catalogs.

---

## Framework Components

### Core Components (Required)

Every threat model addresses these six areas:

| Component | Purpose |
|-----------|---------|
| **Mission & Impact** | Establish context—why security matters for THIS organization |
| **Asset Identification** | Identify what needs protection and where it lives |
| **Adversary Profiling** | Understand who might target you and why |
| **Threat Mapping** | Identify specific threats relevant to your situation |
| **Risk Assessment** | Prioritize which threats to address |
| **Mitigation Planning** | Turn risks into actionable improvements |

### Modules (As Needed)

Modules provide additional depth based on organizational context:

| Module | When Relevant |
|--------|---------------|
| **Deep Adversary Profiling** | Confirmed targeted threats, need to anticipate specific adversary behavior |
| **Information Operations** | Facing narrative attacks, harassment, or reputation threats |
| **OPSEC Analysis** | Protecting sensitive sources/beneficiaries, operating in hostile environments |
| **Incident Response** | History of incidents, need to build response capability |
| **Technical Deep-Dive** | Complex infrastructure, specific technical security concerns |

---

## Threat Domains

Civil society faces threats across multiple domains—not just "cyber":

| Domain | Examples |
|--------|----------|
| **Account & Access** | Phishing, account takeover, unauthorized access, insider misuse |
| **Data & Information** | Data breach, tampering, loss, surveillance |
| **Disruption** | Service disruption, infrastructure attacks |
| **Information & Reputation** | Narrative attacks, impersonation, harassment, platform manipulation |
| **Physical** | Intrusion, surveillance, intimidation, violence |
| **Operational** | Human error, technical failure, third-party failure |

Information operations and physical threats are often as significant as technical attacks. This framework covers all domains.

---

## Adversary Categories

Different adversaries require different defensive approaches:

| Category | Typical Motivation | Typical Capability |
|----------|-------------------|-------------------|
| **Nation-State / Intelligence** | Political control, surveillance, disruption | Advanced across all domains |
| **Ideological Opposition** | Discrediting, silencing, harassment | Moderate info ops, basic technical |
| **Cybercriminal** | Financial gain | Moderate technical, opportunistic |
| **Insider** | Revenge, ideology, financial | Varies, has existing access |
| **Competitor / Opposing Org** | Competitive advantage, policy wins | Legal, reputational attacks |
| **Opportunistic** | Whatever value can be extracted | Basic, automated |

---

## Risk Model

### Three-Factor Scoring

Risk is assessed using three factors, each scored 1-3:

```
RISK SCORE = ASSET VALUE × THREAT LIKELIHOOD × VULNERABILITY
```

| Factor | What It Measures | Score Range |
|--------|------------------|-------------|
| **Asset Value** | How critical is what's at risk? | 1 (Low) to 3 (Critical) |
| **Likelihood** | How likely is this threat to be attempted? | 1 (Unlikely) to 3 (Expected) |
| **Vulnerability** | How exposed are we to this threat? | 1 (Protected) to 3 (Exposed) |

**Why three factors?**

- Asset Value and Likelihood are largely fixed—you can't change what's valuable or who's targeting you
- Vulnerability is what you can actually change through mitigation
- Distinguishing these helps prioritize action

### Priority Bands

| Score Range | Priority | Meaning |
|-------------|----------|---------|
| 18-27 | Critical | Immediate action required |
| 10-17 | High | Near-term action needed |
| 4-9 | Moderate | Plan and address |
| 1-3 | Low | Monitor, accept, or defer |

---

## Outputs

### Threat Model Profile

A complete threat model produces a structured profile containing:

- Organization metadata and context
- Mission and impact framework
- Asset inventory with values
- Adversary profiles
- Threat mapping with likelihood
- Risk register with scores
- Mitigation plans
- Module outputs (as completed)

### Structured Data

All outputs conform to JSON schemas enabling:

- Consistent data format across engagements
- Platform and tool integration
- AI/agent-assisted processing
- Cross-organization analysis (anonymized)

---

## Lifecycle

Threat modeling is not a one-time event.

### Review Cadence

| Frequency | Activity |
|-----------|----------|
| Monthly | Review high-priority mitigation progress |
| Quarterly | Check for emerging threats, update adversary relevance |
| Annually | Full threat model refresh |

### Triggered Reviews

Update after:
- Significant security incidents
- Major organizational changes (new programs, regions, staff)
- Changes in operating environment
- New adversary activity against similar organizations

---

## Implementation Paths

This framework can be implemented through different delivery methods:

| Path | Best For | Guide |
|------|----------|-------|
| **Facilitated Workshop** | Initial assessment, annual refresh, group engagement | `workshop-guide.md` |
| **Self-Guided** | Organizations with security capacity, ongoing updates | `self-guided.md` |
| **Agent-Assisted** | Rapid assessments, conversational extraction, ongoing updates | `agent-assisted.md` |
| **Platform** | Interactive web interface (future) | `platform-spec.md` |

All paths produce the same structured outputs (schemas).

---

## Framework Lineage

This framework synthesizes established security methodologies, adapted for civil society contexts:

| Framework | Contribution |
|-----------|--------------|
| **OCTAVE** | Asset-centric, mission-focused risk assessment |
| **Diamond Model** | Adversary profiling (identity, capability, infrastructure, targeting) |
| **STRIDE** | Threat categorization foundation |
| **DISARM** | Information operations threat taxonomy |
| **NIST 800-61** | Incident response lifecycle |

These frameworks are integrated into a unified approach—practitioners don't need to learn them separately.

---

## Success Metrics

How to measure threat modeling effectiveness:

| Metric | What It Measures |
|--------|------------------|
| Mitigation completion rate | Are identified improvements being implemented? |
| Time to detect incidents | Is detection capability improving? |
| Incident impact reduction | Are incidents less damaging over time? |
| Staff security awareness | Do people know what to do? |
| Threat model currency | Is the model being maintained? |

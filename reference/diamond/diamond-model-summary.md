# Diamond Model of Intrusion Analysis Reference

**Source:** Caltagirone, Pendergast, and Betz (2013)  
**Original Paper:** "The Diamond Model of Intrusion Analysis"

---

## Overview

The Diamond Model is a framework for analyzing cyber intrusions by examining four interconnected features that form a diamond shape. It provides structure for threat intelligence and helps analysts understand the relationships between attackers, their methods, and their targets.

---

## The Four Core Features

```
                    ADVERSARY
                   (Who are they?)
                        │
                        │
    INFRASTRUCTURE ─────┼───── CAPABILITY
   (What resources do   │    (What can they do?)
    they use?)          │
                        │
                      VICTIM
                   (Why us? What do
                    they want?)
```

### 1. Adversary

**Definition:** The actor or group responsible for the intrusion.

**Key Questions:**
- Who is behind the attack?
- What is their identity (individual, group, nation-state)?
- What are their motivations?
- What are their objectives?
- What resources do they have?
- What is their history and pattern of behavior?

**Adversary Attributes:**
- **Operator**: The person(s) executing the intrusion
- **Customer**: The entity directing/funding the operation
- **Motivation**: Financial, political, ideological, espionage, etc.
- **Sophistication**: Technical capability level
- **Persistence**: One-time opportunistic vs. persistent campaign

**Civil Society Relevance:**
Understanding adversary motivation helps predict:
- What they're after (data, disruption, discrediting)
- How persistent they'll be
- What techniques they're likely to use
- What defenses will be most effective

### 2. Capability

**Definition:** The tools, techniques, and procedures (TTPs) used by the adversary.

**Key Questions:**
- What tools are they using (malware, exploits, social engineering)?
- What techniques do they employ?
- What skills do they possess?
- What are their known attack patterns?

**Capability Categories:**
- **Reconnaissance**: Information gathering methods
- **Weaponization**: Tools and exploits used
- **Delivery**: How attacks are delivered (email, web, physical)
- **Exploitation**: Vulnerabilities targeted
- **Installation**: Persistence mechanisms
- **Command & Control**: Communication methods
- **Actions on Objectives**: What they do once inside

**Civil Society Relevance:**
Capability analysis helps:
- Anticipate attack methods
- Prioritize defensive measures
- Recognize attack patterns
- Share intelligence with peers

### 3. Infrastructure

**Definition:** The physical and digital resources used by the adversary.

**Key Questions:**
- What servers, domains, IP addresses are used?
- What communication channels do they use?
- What supporting resources enable their operations?
- Is infrastructure owned, leased, or compromised?

**Infrastructure Types:**
- **Type 1 (Owned)**: Adversary-controlled resources
- **Type 2 (Compromised)**: Legitimate resources taken over
- **Service Providers**: Hosting, email, VPN services used

**Infrastructure Components:**
- Command and control servers
- Phishing domains and landing pages
- Email accounts for social engineering
- Social media accounts (authentic or fake)
- Hosting for malware delivery
- Payment infrastructure (for financial crimes)

**Civil Society Relevance:**
Infrastructure awareness helps:
- Block known malicious resources
- Identify related campaigns
- Report abuse to providers
- Track adversary evolution

### 4. Victim

**Definition:** The target of the intrusion and why they were targeted.

**Key Questions:**
- Who is being targeted? (person, organization, sector)
- Why were they targeted?
- What assets are they after?
- What access does the victim have that interests the adversary?
- What vulnerabilities made targeting successful?

**Victim Analysis:**
- **Persona**: Role, position, access level
- **Assets**: What the victim has that adversary wants
- **Vulnerabilities**: Exploitable weaknesses
- **Relationships**: Connections that make victim valuable

**Civil Society Relevance:**
Understanding victimology helps:
- Recognize why you're a target
- Identify what adversaries want from you
- Understand your value as a stepping stone to others
- Prioritize protection for high-value assets/people

---

## Using the Diamond Model

### For Incident Analysis

When analyzing an incident, work through each feature:

1. **Start with what you know** (often Victim and some Capability)
2. **Pivot to discover more** - each feature reveals others
3. **Build the picture** - relationships between features reveal patterns
4. **Document and share** - structured analysis supports collaboration

### Pivoting Between Features

| From | To | How |
|------|-----|-----|
| Capability | Infrastructure | Malware phones home to C2 servers |
| Infrastructure | Adversary | Domain registration, hosting patterns |
| Victim | Adversary | Who has motivation to target this victim? |
| Adversary | Capability | Known tools and techniques of this actor |

### For Threat Profiling (Proactive)

Use the model to profile potential adversaries before incidents:

1. **Identify likely adversaries** based on your mission/profile
2. **Research their capabilities** from public reporting
3. **Understand their infrastructure** patterns
4. **Analyze why you'd be a target** (victim analysis)

---

## Diamond Model for Civil Society Adversary Profiling

### Example: Nation-State Intelligence Service

| Feature | Analysis |
|---------|----------|
| **Adversary** | Intelligence agency, well-resourced, highly persistent, motivated by surveillance and disruption of opposition |
| **Capability** | Advanced: custom malware, zero-days, sophisticated social engineering, legal pressure, physical surveillance |
| **Infrastructure** | Extensive: owned and compromised infrastructure, front companies, domestic ISP cooperation |
| **Victim (Why You)** | Your work threatens their interests; your contacts are intelligence targets; your data reveals opposition networks |

### Example: Ideological Opposition Group

| Feature | Analysis |
|---------|----------|
| **Adversary** | Loosely organized online group, low-moderate resources, motivated by ideological opposition, opportunistic |
| **Capability** | Basic-moderate: doxing, harassment campaigns, social media manipulation, DDoS-for-hire |
| **Infrastructure** | Platforms: social media accounts, anonymous forums, encrypted chat groups |
| **Victim (Why You)** | Your mission conflicts with their ideology; you're visible; attacking you generates attention |

### Example: Cybercriminal Group

| Feature | Analysis |
|---------|----------|
| **Adversary** | Criminal organization, moderate resources, financially motivated, opportunistic |
| **Capability** | Moderate: ransomware, phishing kits, credential theft, cryptocurrency laundering |
| **Infrastructure** | Criminal ecosystem: bulletproof hosting, money mules, ransomware-as-a-service |
| **Victim (Why You)** | You have money/data worth stealing; your security is weak; you'd likely pay ransom |

---

## Integrating with Threat Modeling

The Diamond Model complements other frameworks:

| Framework | Purpose | Diamond Model Adds |
|-----------|---------|-------------------|
| STRIDE | Categorize threat types | Who would use each threat type against you |
| Kill Chain | Understand attack phases | Who's at each phase, with what tools |
| MITRE ATT&CK | Catalog techniques | Map techniques to specific adversaries |

---

## Resources

- Original Paper: "The Diamond Model of Intrusion Analysis" (Caltagirone et al., 2013)
- MITRE ATT&CK Framework: https://attack.mitre.org/
- Threat Intelligence Platforms often implement Diamond Model concepts

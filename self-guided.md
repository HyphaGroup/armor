# Self-Guided Threat Modeling

A guide for organizations conducting threat modeling independently, without external facilitation.

For step-by-step methodology, see `methodology.md`. For conceptual background, see `framework.md`.

---

## Overview

Self-guided threat modeling allows organizations to work through the methodology at their own pace using internal resources. This guide adapts the methodology for independent use.

### When Self-Guided Works Best

- Organization has internal security knowledge
- Budget constraints prevent external facilitation
- Prefer to build internal capacity
- Need to work incrementally over time
- Sensitive context makes external facilitation complicated

### Challenges to Address

| Challenge | Mitigation |
|-----------|------------|
| Lack of external perspective | Include diverse internal voices, review with peers |
| Security knowledge gaps | Use provided templates, consult references |
| Time management | Schedule dedicated sessions, assign coordinator |
| Maintaining momentum | Set milestones, assign accountability |
| Groupthink | Explicitly seek dissenting views |

---

## Preparation

### Roles

| Role | Responsibility | Time Commitment |
|------|----------------|-----------------|
| **Coordinator** | Schedule sessions, prepare materials, track progress, compile outputs | 4-6 hours over 2-3 weeks |
| **Participants** (3-6 people) | Contribute knowledge, participate in sessions, review outputs | 3-4 hours total |
| **Reviewer** (optional) | External peer or advisor to review draft outputs | 1-2 hours |

**Participant Selection:**
- Include diverse organizational perspectives
- Must include someone with IT/technical knowledge
- Must include someone with programmatic knowledge
- Should include leadership representation
- Consider including someone external for fresh perspective

### Materials

Before starting, gather:
- [ ] Organization's mission statement
- [ ] List of key programs and activities
- [ ] Overview of IT systems (even informal list)
- [ ] Any previous security assessments
- [ ] Recent incidents or concerns
- [ ] Known adversaries or threats

Prepare:
- [ ] Whiteboard or shared document for collaboration
- [ ] Printed or shared copies of methodology sections
- [ ] Schema files for structured output

### Scheduling

**Option A: Intensive (1 day)**
- Full day dedicated session
- Core components: 3 hours with breaks
- Good for focused assessment

**Option B: Distributed (2-3 weeks)**
- Four 90-minute sessions
- Session 1: Mission + Assets
- Session 2: Adversaries + Threats
- Session 3: Risk Assessment
- Session 4: Mitigation Planning + Next Steps
- Better for busy teams, allows reflection between sessions

---

## Session Guidelines

### Before Each Session

**Coordinator tasks:**
1. Review relevant methodology section
2. Prepare any templates or prompts
3. Send pre-work to participants (if any)
4. Test any shared documents or tools

### During Sessions

**Ground rules:**
- All information stays confidential
- No wrong answers during brainstorming
- Build on each other's contributions
- Document everything, refine later
- Time-box discussions to avoid rabbit holes

**Facilitation tips:**
- Use round-robin for initial brainstorming
- Explicitly ask "What are we missing?"
- Parking lot for off-topic issues
- Take breaks every 45-60 minutes

### After Each Session

**Coordinator tasks:**
1. Clean up and organize notes
2. Identify any gaps requiring follow-up
3. Prepare input for next session
4. Share summary with participants

---

## Working Through Core Components

### Component 1: Mission & Impact Framework
**Estimated time:** 30 minutes

**Self-guided approach:**
1. Start with written mission statement
2. Have each participant independently rank impact areas
3. Compare rankings, discuss differences
4. Reach consensus or note disagreements
5. Define impact thresholds for top 3 areas

**Common pitfalls:**
- Defaulting to "everything is critical"
- Not defining concrete thresholds
- Skipping this because it seems obvious

### Component 2: Asset Identification
**Estimated time:** 45 minutes

**Self-guided approach:**
1. Work through each asset category as a group
2. For each identified asset, immediately ask "what's the harm if compromised?"
3. Assign values as you go rather than at the end
4. Don't aim for completeness—focus on what matters most

**Common pitfalls:**
- Listing every piece of data
- Forgetting human containers (who knows what?)
- Not including third-party services

### Component 3: Adversary Profiling
**Estimated time:** 30 minutes

**Self-guided approach:**
1. Review adversary categories individually first
2. Share assessments, note where you agree/disagree
3. For relevant adversaries, brainstorm "why us specifically?"
4. Use adversary templates as starting points

**Common pitfalls:**
- Either paranoid (everyone is an adversary) or dismissive (we're not important enough)
- Forgetting opportunistic threats and insider risk
- Not considering adversaries of partners/funders

### Component 4: Threat Mapping
**Estimated time:** 45 minutes

**Self-guided approach:**
1. Review threat categories against your adversaries and assets
2. For each relevant threat, ask: "Has this happened to us? To similar orgs?"
3. Generate concrete examples, not abstract categories
4. Assign likelihood based on evidence

**Common pitfalls:**
- Getting lost in theoretical threats
- Not linking threats to specific adversaries and assets
- Spending too much time on low-relevance threats

### Component 5: Risk Assessment
**Estimated time:** 45 minutes

**Self-guided approach:**
1. Generate scenarios using the template format
2. Validate each scenario against the criteria
3. Assess vulnerability honestly (be critical)
4. Calculate scores, but don't game the math
5. Select top 5 for focus

**Common pitfalls:**
- Scenarios too vague to be actionable
- Overestimating existing controls
- Not distinguishing between likelihood and vulnerability

### Component 6: Mitigation Planning
**Estimated time:** 30 minutes

**Self-guided approach:**
1. Focus only on top 5 risks
2. Brainstorm mitigations before assessing feasibility
3. Be realistic about capacity and budget
4. Identify clear owners and timelines
5. Prioritize quick wins

**Common pitfalls:**
- Planning more than can be implemented
- Not assigning clear ownership
- Forgetting to define success criteria

---

## Completing Modules

Modules can be completed:
- As follow-up sessions after core
- Asynchronously by designated individuals
- With external support for specialized topics

**Module prioritization:**
Based on what emerged from core, identify 1-2 modules for immediate follow-up.

| Module | Self-Guided Feasibility |
|--------|------------------------|
| Deep Adversary Profiling | Medium - requires research capability |
| Information Operations | High - can use existing social media knowledge |
| OPSEC Analysis | Medium - benefits from outside perspective |
| Incident Response | High - largely internal assessment |
| Technical Deep-Dive | Medium - requires technical expertise |

---

## Compiling Outputs

### During the Process

Capture structured data as you go using the JSON schemas:
- `mission.schema.json`
- `assets.schema.json`
- `adversaries.schema.json`
- `threats.schema.json`
- `risks.schema.json`
- `mitigations.schema.json`

### After Core Completion

**Coordinator compiles:**

1. **Threat Model Profile** (JSON)
   - Combine all schema outputs
   - Use `meta.schema.json` for wrapper

2. **Executive Summary** (1-2 pages)
   - Mission context (2-3 sentences)
   - Top 3-5 risks
   - Priority mitigations
   - Quick wins

3. **Action Plan** (table)
   - Mitigation, owner, timeline, success criteria

### Quality Check

Before finalizing, verify:
- [ ] Risk scenarios are specific enough to act on
- [ ] All priority risks have at least one mitigation
- [ ] Each mitigation has an owner
- [ ] Quick wins are actually quick
- [ ] No sensitive information that shouldn't be documented

---

## Maintaining the Threat Model

### Schedule

| Activity | Frequency | Owner |
|----------|-----------|-------|
| Quick wins check | Weekly (first month) | Coordinator |
| Mitigation progress | Monthly | Coordinator |
| Emerging threats review | Quarterly | Team |
| Full refresh | Annually | Team |

### Triggers for Ad-Hoc Review

- Security incident
- Major organizational change
- New adversary activity in sector
- Significant operational change
- New program in high-risk area

### Version Control

Track changes to threat model:
- Date of update
- What changed and why
- Who updated

---

## Getting Help

### When to Seek External Support

- First time conducting threat modeling
- Identified risks beyond internal expertise
- Need validation of findings
- Want to conduct tabletop exercises
- Specific technical assessment needed

### Resources

**Peer Review:**
- Ask a peer organization to review your threat model
- Offer reciprocal review

**Community Support:**
- Digital security helplines (Access Now, etc.)
- Sector-specific security communities
- Regional digital rights organizations

**Professional Support:**
- Digital security trainers and consultants
- Security audit firms with civil society experience

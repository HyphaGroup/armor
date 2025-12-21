# Agent-Assisted Threat Modeling

A guide for conducting threat modeling with AI agent support.

For step-by-step methodology, see `methodology.md`. For conceptual background, see `framework.md`.

---

## Overview

Agent-assisted threat modeling uses AI agents to guide organizations through the methodology, provide context-appropriate suggestions, and help structure outputs. The agent serves as facilitator and documentation assistant while the human provides organizational knowledge.

### When Agent-Assisted Works Best

- Limited access to human facilitators
- Need structured guidance through methodology
- Want consistent, thorough documentation
- Comfortable with AI interaction
- Can work asynchronously

### Limitations

| Limitation | Mitigation |
|------------|------------|
| Agent doesn't know your organization | Provide thorough context upfront |
| Agent may miss cultural/political nuance | Human reviews all outputs |
| Agent can't verify claims | Human validates accuracy |
| Agent lacks field experience | Supplement with reference materials |

---

## Preparation

### Human Roles

| Role | Responsibility |
|------|----------------|
| **Primary Informant** | Provides organizational knowledge, answers agent questions |
| **Reviewer** | Validates agent outputs, adds context |
| **Decision Maker** | Approves final threat model, assigns owners |

### Context to Provide

Before starting, prepare to share:

**Organizational Context:**
- Mission statement
- Type of work (advocacy, journalism, humanitarian, etc.)
- Geographic focus
- Size and structure
- Operating environment (open democracy, restricted, conflict zone)

**Security Context:**
- Previous incidents (if any)
- Known threats or adversaries
- Current security measures
- Recent concerns

**Technical Context:**
- Key systems and services used
- Data handling overview
- Technical capacity level

### Session Options

**Option A: Single Session (~2-3 hours)**
- Work through all core components
- Good for focused assessment

**Option B: Multiple Sessions**
- Session 1: Context, Mission, Assets
- Session 2: Adversaries, Threats
- Session 3: Risk Assessment, Mitigations
- Better for thorough exploration

---

## Agent Prompting Guide

### Starting the Session

**Initial prompt example:**

```
I want to conduct a threat model for my organization using the civil society threat modeling methodology.

About us:
- [Organization type]: [description]
- Mission: [mission statement]
- Key activities: [list]
- Geographic scope: [regions/countries]
- Size: [staff count, volunteer count]
- Operating environment: [characterize]

Previous security experience:
- [Any incidents, concerns, or assessments]

Please guide me through the threat modeling process, starting with the Mission & Impact Framework.
```

### Component-Specific Prompts

**Mission & Impact Framework:**
```
Help me define impact areas and thresholds for our organization.
Our mission is: [mission]
Our key stakeholders are: [list]
```

**Asset Identification:**
```
Help me identify our information assets. 
We work on [activities] and handle data about [types].
Our key systems are: [list].
```

**Adversary Profiling:**
```
Help me assess potential adversaries.
Our work involves: [description]
We operate in: [locations/contexts]
We've experienced: [any relevant history]
```

**Threat Mapping:**
```
Based on our assets and adversaries, help me map relevant threats.
Assets of concern: [list top assets]
Relevant adversaries: [list]
```

**Risk Assessment:**
```
Help me create risk scenarios combining our adversaries, threats, and assets.
Here's what we've identified so far: [summary]
```

**Mitigation Planning:**
```
For our top risks, help me identify practical mitigations.
Our top risks are: [list]
Our capacity/constraints: [description]
```

### Module Prompts

**Deep Adversary Profiling:**
```
Help me develop a detailed profile for [adversary name/type].
What we know: [existing information]
Why they're relevant: [reason]
```

**Information Operations:**
```
Assess our exposure to information operations.
Our public presence: [description]
Our social media: [platforms, reach]
Public-facing staff: [who, roles]
Previous incidents: [any history]
```

**OPSEC Analysis:**
```
Help me analyze our operational security.
Sensitive aspects of our work: [description]
Current public exposure: [what's visible]
Adversaries who might monitor us: [list]
```

**Incident Response:**
```
Assess our incident response capability.
Current resources: [IT support, security contacts]
Past incidents: [what happened, how handled]
Current documentation: [what exists]
```

**Technical Deep-Dive:**
```
Help me assess our technical security.
Our infrastructure: [systems, services]
Current security measures: [what's in place]
Technical capacity: [who manages IT]
```

---

## Working with Agent Outputs

### Validation Checklist

After each component, verify:
- [ ] Outputs reflect organizational reality
- [ ] Nothing sensitive inappropriately documented
- [ ] Assessments feel accurate (not over/understated)
- [ ] Examples are relevant to your context

### Adding Human Context

The agent will miss:
- Internal politics and dynamics
- Informal knowledge (who really knows what)
- Historical context not shared
- Nuance about relationships and trust
- Gut feelings based on experience

Add this context as you review outputs.

### Correcting Course

If agent outputs feel off:
```
That doesn't quite fit our situation. In our case, [correction]. 
Please adjust the assessment.
```

```
You're missing important context: [context]. 
How does this change your analysis?
```

```
This feels [too paranoid / too dismissive] for our situation. 
Here's why: [explanation]
```

---

## Generating Structured Outputs

### During the Session

Ask the agent to structure outputs per schema:
```
Please format the assets we've identified according to the assets.schema.json structure.
```

### Compiling Final Outputs

At session end:
```
Please compile our complete threat model in the following formats:
1. Structured JSON per the schemas
2. Executive summary (1-2 pages)
3. Risk register (prioritized list)
4. Action plan with owners and timelines
```

### Meta Information

Include in final output:
- Date of assessment
- Participants (human roles)
- Agent model used
- Session duration
- Modules completed

---

## Security Considerations

### What to Share with Agent

**Generally safe:**
- Organizational structure and mission
- Types of work and activities
- General threat landscape
- Technical infrastructure overview
- Hypothetical scenarios

**Exercise caution:**
- Specific names of sources or beneficiaries
- Detailed operational plans
- Specific vulnerability details
- Internal conflicts or politics
- Unverified suspicions

**Avoid sharing:**
- Credentials or access information
- Specific evidence of surveillance
- Legal matters under privilege
- Information that could identify protected individuals

### Data Handling

Consider:
- Agent conversation may be logged
- Use organizational accounts, not personal
- Review outputs before sharing externally
- Redact sensitive specifics in final documents

---

## After the Session

### Human Review

All agent outputs should be:
1. Reviewed by someone with organizational knowledge
2. Validated against ground truth
3. Supplemented with human context
4. Approved by decision maker before adoption

### Action Items

Convert mitigation plans into:
- Assigned tasks with owners
- Calendar reminders for timelines
- Progress tracking system

### Follow-Up Sessions

Schedule:
- Quick wins review (2 weeks)
- Progress check (monthly)
- Refresh (quarterly or annually)

---

## Example Session Flow

**Session 1: Foundation (45-60 min)**
1. Provide organizational context
2. Complete Mission & Impact Framework
3. Begin Asset Identification
4. *Save state for next session*

**Session 2: Threats (45-60 min)**
1. Complete Asset Identification
2. Adversary Profiling
3. Threat Mapping
4. *Save state*

**Session 3: Assessment & Planning (60-90 min)**
1. Risk Assessment
2. Mitigation Planning
3. Quick wins identification
4. Module recommendations
5. Compile outputs

**Optional Module Sessions (45-60 min each)**
- Schedule based on recommendations
- Can be done asynchronously

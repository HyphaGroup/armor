# NIST SP 800-61: Incident Response Reference

**Source:** NIST Special Publication 800-61 Rev. 3 (April 2025)  
**Full Document:** https://csrc.nist.gov/pubs/sp/800/61/r3/final

---

## Overview

NIST SP 800-61 provides comprehensive guidance on computer security incident handling. Rev. 3 aligns incident response with the NIST Cybersecurity Framework (CSF) 2.0's five core functions: Identify, Protect, Detect, Respond, and Recover.

---

## The Four Phases of Incident Response

### Phase 1: Preparation

**Purpose:** Establish and maintain incident response capability before incidents occur.

**Key Activities:**
- Develop incident response policy and plan
- Define what constitutes an incident for your organization
- Establish roles and responsibilities
- Set up communication channels and contact lists
- Conduct regular training and simulations
- Ensure necessary tools and resources are in place
- Document procedures for common incident types

**Outputs:**
- Incident response plan
- Contact lists (internal, external, legal, PR)
- Communication templates
- Trained response team

### Phase 2: Detection and Analysis

**Purpose:** Identify potential security incidents through monitoring and analysis.

**Key Activities:**
- Implement monitoring and alerting systems
- Establish baselines for normal activity
- Correlate events from multiple sources
- Analyze indicators of compromise
- Determine incident scope and impact
- Document findings and preserve evidence

**Detection Sources:**
- Security information and event management (SIEM)
- Intrusion detection/prevention systems
- Antivirus and endpoint detection
- Network monitoring
- User reports
- Third-party notifications

**Analysis Questions:**
- What systems are affected?
- What data may be exposed?
- Is the incident ongoing?
- What is the potential impact?
- Who needs to be notified?

### Phase 3: Containment, Eradication, and Recovery

**Containment Purpose:** Limit the damage and prevent further spread.

**Containment Strategies:**
- Short-term: Immediate actions to stop bleeding (isolate systems, block IPs)
- Long-term: Temporary fixes while preparing for eradication
- Evidence preservation: Maintain forensic integrity

**Eradication Purpose:** Remove the threat and root cause.

**Eradication Activities:**
- Remove malware and unauthorized access
- Identify and patch vulnerabilities exploited
- Change compromised credentials
- Verify removal was complete

**Recovery Purpose:** Restore systems to normal operations.

**Recovery Activities:**
- Restore from clean backups
- Rebuild compromised systems
- Verify system integrity before return to production
- Monitor for re-compromise
- Gradually restore services

### Phase 4: Post-Incident Activity

**Purpose:** Learn from incidents to improve future response.

**Key Activities:**
- Conduct post-incident review (within 1-2 weeks)
- Document lessons learned
- Update incident response procedures
- Address root cause vulnerabilities
- Share information with relevant parties (as appropriate)
- Update training based on findings

**Post-Incident Questions:**
- What happened and when?
- How well did staff and management perform?
- What information was needed sooner?
- Were any steps or actions that might have inhibited recovery?
- What would staff do differently next time?
- What corrective actions can prevent similar incidents?

---

## Incident Categories (Examples)

| Category | Description |
|----------|-------------|
| Unauthorized Access | Gaining access to systems without permission |
| Denial of Service | Disrupting availability of systems or services |
| Malicious Code | Viruses, worms, trojans, ransomware |
| Improper Usage | Violations of acceptable use policies |
| Multiple Component | Incidents involving multiple categories |

---

## Incident Severity Levels

| Level | Impact | Examples |
|-------|--------|----------|
| High | Significant impact on mission | Data breach, ransomware, critical system down |
| Medium | Moderate impact, recoverable | Single system compromise, limited data exposure |
| Low | Minimal impact | Failed attack attempts, policy violations |

---

## Key Principles

1. **Speed matters** - Faster detection and response limits damage
2. **Preparation is essential** - Most work happens before incidents
3. **Document everything** - Supports investigation, legal, and learning
4. **Communicate appropriately** - Right information to right people at right time
5. **Learn and improve** - Every incident is an opportunity to strengthen defenses

---

## Adapting for Civil Society Organizations

**Preparation considerations:**
- Contact lists should include digital security organizations, platform contacts, legal counsel
- Training should cover information operations (harassment, impersonation) not just technical attacks
- Establish relationships with peer organizations for mutual support

**Detection considerations:**
- Include social media monitoring for impersonation and narrative attacks
- Monitor for credential dumps and data leaks mentioning organization
- Staff should know what to report (unusual account activity, suspicious contacts)

**Response considerations:**
- Information operation incidents require different response (communications-focused)
- Consider reputational impact alongside technical impact
- May need to coordinate with journalists, fact-checkers, platforms

**Post-incident considerations:**
- Share lessons with peer organizations (safely)
- Consider whether incident reveals need for OPSEC improvements
- Update threat model based on adversary behavior observed

---

## Resources

- NIST SP 800-61r3: https://csrc.nist.gov/pubs/sp/800/61/r3/final
- NIST Cybersecurity Framework 2.0: https://www.nist.gov/cyberframework

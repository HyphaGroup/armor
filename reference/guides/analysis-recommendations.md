# Analysis: External Guides vs. Our Methodology

Comparative analysis of civil society security guides and recommendations for refinements to the ARMOR framework.

---

## Summary Assessment

**Our methodology is already well-positioned** relative to the major guides. It incorporates most key elements and adds several innovations (three-factor risk scoring, structured JSON outputs, modular architecture). However, analysis reveals several areas for potential enhancement.

---

## Strengths of Our Current Approach

### What We Do Well (Already Covered)

1. **Comprehensive Threat Modeling**
   - Our 6-question framework aligns well with EFF's "Your Security Plan"
   - Our adversary profiling goes deeper than most guides
   - Threat categories are more comprehensive than typical guides

2. **Structured Risk Assessment**
   - Three-factor scoring (Asset Value × Likelihood × Vulnerability) is more rigorous than most
   - Explicit priority bands and action mapping
   - Risk scenarios with validation criteria

3. **Information Operations Coverage**
   - Our InfoOps module is notably absent from most other guides
   - DISARM framework integration is unique
   - Harassment and narrative attacks are underserved elsewhere

4. **Modular Architecture**
   - Flexible depth based on organizational needs
   - Core + modules structure allows customization
   - Better than one-size-fits-all approaches

5. **Structured Outputs**
   - JSON schemas enable tooling and automation
   - Machine-readable outputs are unique to our approach
   - Supports agent-assisted delivery

---

6. **Self-Assessment Checklist**
   - The "General Best Practices" guide ends with a practical checklist
   - Our methodology could benefit from similar verification tools

---

## Gaps and Potential Refinements

### 1. WELL-BEING AND STRESS INTEGRATION (High Priority)
**Gap:** Holistic Security and Front Line Defenders emphasize well-being as integral to security. Our methodology treats this minimally.

**Recommendation:**
- Add a "Well-being Assessment" component to core methodology
- Include stress/burnout indicators in Incident Response module
- Add "sustainability of security practices" as success metric
- Consider: "Holistic Security's framing: "Well-being in action"

**Specific additions to consider:**
- Staff burnout risk assessment
- Psychosocial support resources
- "Do No Harm" considerations for beneficiaries
- Vicarious trauma for those handling sensitive cases

### 2. ALLIES AND SUPPORT NETWORKS (Medium Priority)
**Gap:** EFF's "Who are my allies?" question is absent from our framework.

**Recommendation:**
- Add "Support Network Assessment" to Mission & Impact component
- Include: peer organizations, legal support, digital security contacts, media allies
- Map which relationships are protective vs. which add exposure

**Implementation:**
```
Support Network:
- Peer organizations for mutual aid
- Legal support (named contacts, retainer status)
- Digital security support (helplines, consultants)
- Media/communications allies
- Physical security contacts
- Mental health resources
```

### 3. EMERGENCY RESPONSE INTEGRATION (Medium Priority)
**Gap:** Our Incident Response module is assessment-focused. Digital First Aid Kit provides triage/diagnostic approach for active incidents.

**Recommendation:**
- Add "Incident Triage" decision tree to methodology
- Include "During Incident" guidance, not just preparedness
- Link to external support resources (Access Now, CiviCERT, etc.)

**Specific additions:**
- "What to do RIGHT NOW if..." quick guides
- Decision tree: escalate vs. handle internally
- External resource directory

### 4. SECURITY INDICATORS / EARLY WARNING (Medium Priority)
**Gap:** Holistic Security emphasizes "security indicators" - early warning signs that context is changing.

**Recommendation:**
- Add "Security Indicators" section to methodology
- Connect to review triggers
- Example indicators: increased online hostility, peer org incidents, legal threats, surveillance signs

**Implementation in methodology:**
Add to Risk Assessment or as separate component:
```
Security Indicators to Monitor:
- Changes in government rhetoric about civil society
- Incidents at peer organizations
- New legislation affecting operations
- Unusual contact from authorities
- Increased online attention/harassment
- Staff reporting being followed/watched
```

### 5. CONTEXT ANALYSIS DEPTH (Medium Priority)
**Gap:** Holistic Security has robust context analysis framework. Our Mission component is lighter.

**Recommendation:**
- Expand Mission & Impact to include:
  - Political environment assessment
  - Civil society space (open/shrinking/closed)
  - Legal framework for organizations
  - Recent trends affecting similar orgs

### 6. PHYSICAL SECURITY INTEGRATION (Lower Priority)
**Gap:** CPJ and Front Line Defenders have strong physical security coverage. Our methodology is digital-focused.

**Assessment:** This may be intentional scope limitation, but consider:
- Adding physical security module for orgs in high-risk environments
- At minimum, cross-reference physical security in OPSEC module

### 7. BORDER CROSSING / TRAVEL (Lower Priority)
**Gap:** CPJ and Security in a Box have dedicated travel/border crossing guidance.

**Recommendation:**
- Add travel security considerations to OPSEC module
- Include device preparation for border crossing
- Address journalist/activist-specific border risks

### 8. PERSONAL SAFETY / ANTI-DOXXING (Medium Priority)
**Gap:** The "General Best Practices" guide has strong personal safety section that goes beyond digital.

**Recommendation:**
- Add personal safety considerations to OPSEC module or as separate section
- Include doxxing prevention (data broker removal, credit freezes)
- Family safety measures (code words, transitioning family to secure comms)
- Monitoring for personal exposure (Google Alerts, PimEyes checks)

**Specific additions:**
```
Personal Safety Measures:
- Data broker removal services
- Credit freeze with all major bureaus
- Family "code word" for identity verification
- Google Alerts for personal identifiers
- Regular "Have I Been Pwned" checks
- Social media privacy audit
```

### 9. TOOL CONFIGURATION SPECIFICITY (Lower Priority)
**Gap:** General Best Practices provides specific tool configuration requirements (Signal settings, Zoom security, etc.). Our methodology is tool-agnostic.

**Assessment:** This is likely intentional - our methodology is meant to be durable while tools change. However:

**Recommendation:**
- Consider a companion "Implementation Checklist" document
- Could be a Level 3 guide: "Technical Implementation Checklist"
- Specific settings for common platforms (Signal, Zoom, Google Workspace)
- Updated more frequently than methodology

### 10. DISCOURAGED TECHNOLOGIES FRAMING (Lower Priority)
**Gap:** General Best Practices explicitly lists what NOT to use, not just what to use.

**Recommendation:**
- In Technical Deep-Dive module, consider adding:
  - "Discouraged" category for tools/practices
  - Clear alternatives for each discouraged item
  - Rationale for each (helps with buy-in)

---

## Elements We Have That Others Lack

Our methodology has innovations that external guides don't:

1. **Information Operations Module** - Most guides ignore narrative attacks, online harassment campaigns, platform manipulation
2. **Structured JSON Outputs** - No other guide provides machine-readable schemas
3. **Agent-Assisted Delivery** - We've designed for AI-guided implementation
4. **Three-Factor Risk Scoring** - More rigorous than typical guides
5. **Deep Adversary Profiling** - More detailed than standard threat actor analysis
6. **Technical Deep-Dive Module** - Comprehensive technical assessment for orgs with capacity

---

## Recommended Refinements (Prioritized)

### Immediate (Add to current structure)

1. **Add "Allies & Support Network" section to Mission & Impact**
   - Who can you call for help?
   - What peer organizations share your threat landscape?
   - Legal, digital security, media, physical security contacts

2. **Add Well-Being considerations to framework.md principles**
   - Security sustainability
   - Staff/volunteer burnout as security risk
   - Mental health resources

3. **Add Security Indicators to Risk Assessment**
   - Early warning signs to monitor
   - Context change triggers

### Near-Term (Module enhancements)

4. **Enhance Incident Response module with triage guidance**
   - Active incident decision tree
   - Escalation criteria
   - External resource directory

5. **Add Travel/Border considerations to OPSEC module**
   - Device preparation
   - Data minimization for travel
   - Border crossing scenarios

6. **Add Personal Safety section to OPSEC module**
   - Anti-doxxing measures (data broker removal, credit freezes)
   - Family safety protocols
   - Personal exposure monitoring

### Future Consideration

7. **Physical Security module** (for high-risk contexts)
8. **Well-Being module** (integrate holistic security approach)
9. **Context Analysis expansion** (political/legal environment assessment)
10. **Technical Implementation Checklist** (Level 3 guide with specific tool configs)

---

## Implementation Notes

### For framework.md
Add to principles:
- Security as sustainability (not just protection)
- Well-being as security enabler
- Community/network security (allies matter)

### For methodology.md
Add to Mission & Impact:
- Support network mapping
- Context assessment questions

Add to Risk Assessment:
- Security indicators section
- Context change triggers

### For Incident Response module
Add:
- Active incident triage
- Decision trees
- External resource directory

### For OPSEC module
Add:
- Travel security section
- Border crossing considerations

---

## External Resources to Reference

Consider adding a resources section pointing to:
- Access Now Digital Security Helpline: help@accessnow.org
- Digital First Aid Kit: digitalfirstaid.org
- Security in a Box: securityinabox.org
- EFF Surveillance Self-Defense: ssd.eff.org
- Front Line Defenders: frontlinedefenders.org/en/emergency-contact

---

## Conclusion

Our methodology is comprehensive and in many ways more structured than existing guides. The main gaps are:
1. **Well-being integration** (biggest gap - from Holistic Security, FLD)
2. **Support network / allies mapping** (from EFF SSD)
3. **Active incident triage** (from Digital First Aid Kit)
4. **Security indicators / early warning** (from Holistic Security)
5. **Personal safety / anti-doxxing** (from General Best Practices)

These can be addressed with targeted additions without restructuring the core methodology. The modular architecture makes this straightforward.

### Note on Tool-Specificity

The "General Best Practices" guide is notably more prescriptive about specific tools and configurations. Our methodology intentionally stays tool-agnostic for durability, but we could consider:
- A companion "Implementation Checklist" at Level 3
- Linking to external tool guides (Security in a Box, EFF SSD) rather than duplicating

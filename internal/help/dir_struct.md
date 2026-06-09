# OperaTree Project Structure Guide

**Repository:** [OperaTree GitHub Repository](https://github.com/hanymamdouh82/operatree.git?utm_source=chatgpt.com)


## Introduction

OperaTree is built around a simple principle:

> Organize project knowledge according to its purpose, not its file type.

Most projects eventually become difficult to navigate because documents are grouped by format:

```plaintext
Documents/
Photos/
PDFs/
Presentations/
```

This approach makes it difficult to answer practical questions such as:

- Where is the research behind this decision?
- Which factory visit produced this finding?
- Which files support this objective?
- Which datasets were used for this report?

OperaTree organizes information by **business meaning and lifecycle** rather than by file extension.

The directory structure acts as a project operating system that helps teams:

- Preserve knowledge
- Improve discoverability
- Maintain traceability
- Support collaboration
- Prepare for AI-assisted knowledge retrieval
- Reduce information loss over long projects


## Core Philosophy

Every file belongs to one of six categories:

| Category           | Question Answered          |
| ------------------ | -------------------------- |
| Events             | What happened?             |
| Project Management | What are we doing?         |
| Research           | What do we know?           |
| Engineering        | What are we building?      |
| Data               | What have we measured?     |
| Deliverables       | What are we communicating? |

When deciding where to place a file, identify the question it answers.

## Top-Level Directory Overview

```plaintext
00_ADMIN
01_EVENTS
02_PROJECT_MANAGEMENT
03_LEGAL
04_RESEARCH
05_ENGINEERING
06_DATA
97_MEDIA_LIBRARY
98_DELIVERABLES
99_ARCHIVE
```

---

## 00_ADMIN

### Purpose

Administrative and governance information for the project.

### Contents

| Directory  | Purpose                                             |
| ---------- | --------------------------------------------------- |
| governance | Project charter, governance model, responsibilities |
| contacts   | Stakeholder and team contact information            |
| templates  | Reusable document templates                         |
| guidelines | Project procedures and standards                    |

### Examples

- Project charter
- Naming conventions
- Meeting templates
- Communication guidelines

---

## 01_EVENTS

### Purpose

Stores all information related to specific project events.

An event is anything that occurs at a specific point in time.

Examples:

- Factory visit
- Workshop
- Seminar
- Stakeholder meeting
- Field survey
- Training session

---

## Event Structure

```plaintext
01_EVENTS/
└── 2026-04-15-factory-visit/
    ├── 01_AGENDA
    ├── 02_MEDIA
    ├── 03_NOTES
    ├── 04_DOCUMENTS
    ├── 05_OUTCOMES
    └── META.yaml
```

---

## Subdirectories

| Directory    | Purpose                             |
| ------------ | ----------------------------------- |
| 01_AGENDA    | Plans, schedules, invitations       |
| 02_MEDIA     | Photos, videos, recordings          |
| 03_NOTES     | Raw notes and observations          |
| 04_DOCUMENTS | Presentations and shared materials  |
| 05_OUTCOMES  | Meeting minutes, actions, decisions |

---

## Rule

Store raw event information here.

Store long-term conclusions in:

```plaintext
04_RESEARCH/03_SUMMARIES
```

---

# 02_PROJECT_MANAGEMENT

## Purpose

Manages project execution.

This area answers:

> What work is currently being performed?

---

## Structure

```plaintext
02_PROJECT_MANAGEMENT/
├── 07_TASKS
├── reports
├── budgets
├── risks
├── communications
└── planning
```

---

## Tasks

Each task contains its own working environment.

Example:

```plaintext
07_TASKS/
└── website-development/
    ├── 01_INPUTS
    ├── 02_WORKING
    ├── 03_REVIEW
    ├── 04_FINAL
    └── META.yaml
```

---

## Examples

- NDA preparation
- Framework creation
- Website development
- Procurement analysis
- Vendor evaluation

---

## Rule

If a file is actively supporting execution of work, it belongs here.

---

# 03_LEGAL

## Purpose

Official legal documentation.

---

## Structure

| Directory  | Purpose                             |
| ---------- | ----------------------------------- |
| contracts  | Signed contracts                    |
| ndas       | Executed NDAs                       |
| compliance | Regulatory and compliance materials |
| approvals  | Formal approvals                    |
| templates  | Legal templates                     |

---

## Rule

Store finalized legal documents here.

Drafts generally remain under the corresponding task until approved.

---

# 04_RESEARCH

## Purpose

The knowledge and intelligence center of the project.

---

# 04_RESEARCH/09_TOPICS

## Purpose

Concept-centric knowledge.

Answers:

> What is this concept?

---

## Examples

- Forecasting methods
- Predictive maintenance
- Digital twins
- Lean manufacturing

---

## Typical Contents

```plaintext
forecasting_methods/
├── overview.md
├── notes.md
├── diagrams/
└── attachments/
```

---

## Store Here

- Explanations
- Theory
- Frameworks
- Educational material
- Topic-specific references

---

## Do Not Store Here

- Project decisions
- Factory-specific findings
- Strategic actions

---

# 04_RESEARCH/10_OBJECTIVES

## Purpose

Goal-centric intelligence.

Answers:

> What are we trying to achieve?

---

## Examples

- Reduce downtime
- Improve forecasting accuracy
- Create industry-agnostic platform
- Improve production visibility

---

## Typical Contents

```plaintext
objective_reduce_downtime/
├── definition.md
├── findings.md
├── strategy.md
├── notes/
├── discussions/
└── attachments/
```

---

## Store Here

- Findings
- Discussions
- Research related to the objective
- Brainstorming
- Supporting evidence
- Strategy proposals

---

## Rule

Objectives are active thinking spaces.

---

# 04_RESEARCH/11_SUMMARIES

## Purpose

Distilled knowledge.

Answers:

> What matters?

---

## Examples

| Type               | Example                      |
| ------------------ | ---------------------------- |
| Event Summary      | Factory visit conclusions    |
| Research Summary   | Key findings from papers     |
| Discussion Summary | DeepSeek discussion outcomes |
| Objective Summary  | Final strategic conclusion   |
| Executive Summary  | Management-level overview    |

---

## Rule

Summaries contain conclusions, not raw information.

---

# 04_RESEARCH/12_REFERENCES

## Purpose

External knowledge library.

Answers:

> What have we read?

---

## Examples

- Books
- Research papers
- Standards
- Whitepapers
- Regulations
- Vendor documentation

---

## Suggested Structure

```plaintext
04_REFERENCES/
├── articles
├── books
├── papers
├── standards
├── regulations
└── vendor_docs
```

---

## Rule

Original external resources belong here.

Your understanding belongs in:

```plaintext
09_TOPICS
```

---

# 04_RESEARCH/13_AUDIO_NOTES

## Purpose

Knowledge captured verbally.

---

## Structure

| Directory      | Purpose                   |
| -------------- | ------------------------- |
| raw            | Original recordings       |
| transcriptions | Speech-to-text output     |
| indexed        | Organized audio knowledge |

---

# 04_RESEARCH/14_ATTACHMENTS

## Purpose

General supporting files that do not fit neatly elsewhere.

Examples:

- Scans
- Miscellaneous PDFs
- Supporting images

---

# 05_ENGINEERING

## Purpose

Technical implementation and system design.

---

## Structure

| Directory      | Purpose                              |
| -------------- | ------------------------------------ |
| architecture   | High-level designs                   |
| diagrams       | Technical diagrams                   |
| specifications | Requirements and specifications      |
| prototypes     | Proofs of concept                    |
| simulations    | Simulations and experiments          |
| decisions      | Architecture Decision Records (ADRs) |

---

## Rule

If it describes how the system works, it belongs here.

---

# 06_DATA

## Purpose

Manage the complete lifecycle of project data.

---

## Data Flow

```plaintext
Sources
   ↓
Raw
   ↓
Staging
   ↓
Processed
   ↓
Analytics
   ↓
Exports
```

---

## 15_SOURCES

Metadata describing data origins.

Examples:

- Kaggle dataset description
- Team-provided dataset metadata
- API source definitions

---

## 01_RAW

Immutable source data.

Never modify files here.

Examples:

- CSV exports
- XLSX files
- Database dumps

---

## 02_STAGING

Intermediate processing area.

Examples:

- Cleaned datasets
- Merged datasets
- Temporary transformations

---

## 03_PROCESSED

Trusted project-ready datasets.

Examples:

- Aggregated results
- Feature sets
- Curated datasets

---

## 04_ANALYTICS

Exploration and analysis.

Examples:

- Notebooks
- Experiments
- Visualizations

---

## 05_MODELS

Machine learning artifacts.

Examples:

- Trained models
- Checkpoints
- Evaluation metadata

---

## 06_EXPORTS

Data prepared for external consumption.

Examples:

- Dashboard feeds
- Reporting datasets
- Client exports

---

# 97_MEDIA_LIBRARY

## Purpose

Reusable media assets shared across the project.

---

## Rule

Ask:

> Can this asset be reused by multiple tasks, reports, presentations, or deliverables?

If yes, place it here.

---

## Suggested Structure

```plaintext
07_MEDIA_LIBRARY/
├── branding
├── people
├── website
├── photos_clean
├── videos
├── diagrams
└── presentation_assets
```

---

## Branding

Examples:

- Logos
- Email signatures
- Brand guidelines

---

## People

Examples:

- Team member photos
- Speaker photos
- Stakeholder photos

---

## Website

Examples:

- Website banners
- Hero images
- Illustrations
- Optimized web assets

---

## Photos Clean

Curated photos selected from events for long-term reuse.

---

# 98_DELIVERABLES

## Purpose

Final outputs provided to stakeholders.

---

## Examples

| Directory        | Contents                |
| ---------------- | ----------------------- |
| reports          | Final reports           |
| presentations    | Presentation decks      |
| submissions      | Formal submissions      |
| client_documents | Client-facing documents |

---

## Rule

Only finalized outputs belong here.

Drafts remain in tasks until approved.

---

# 99_ARCHIVE

## Purpose

Historical preservation.

---

## Examples

- Deprecated materials
- Obsolete datasets
- Closed workstreams
- Legacy versions

---

# File Placement Decision Tree

When adding a file, ask:

### Was it created during a specific event?

→ `01_EVENTS`

### Is it helping execute current work?

→ `02_PROJECT_MANAGEMENT`

### Is it a legal document?

→ `03_LEGAL`

### Is it knowledge or research?

→ `04_RESEARCH`

### Is it technical design or implementation?

→ `05_ENGINEERING`

### Is it data?

→ `06_DATA`

### Is it a reusable asset?

→ `97_MEDIA_LIBRARY`

### Is it a finalized external output?

→ `98_DELIVERABLES`

### Is it obsolete but must be retained?

→ `99_ARCHIVE`

---

# Research Layer Cheat Sheet

This is the most important distinction in the system.

| Directory       | Question                       |
| --------------- | ------------------------------ |
| `01_TOPICS`     | What is this?                  |
| `02_OBJECTIVES` | What are we trying to achieve? |
| `03_SUMMARIES`  | What matters?                  |
| `04_REFERENCES` | What have we read?             |

If users understand these four directories, they will understand the entire knowledge model behind OperaTree.

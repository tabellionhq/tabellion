```
Document source: Tabellion project                          A. Crommer
Document: draft-tabellion-curriculum-00                     October 2026
Category: Beta specification, not an RFC
```

# Tabellion Skills Curriculum: Data Model and Rules

## Abstract

This document specifies a data model and rules for recording the skills a person has practised over time: the occupations held, the practices of a skill on a concept at an occupation over a period, and the views through which the person exposes part of this record to others. Skills are identified by the European Skills, Competences, Qualifications and Occupations classification (ESCO); concepts and contexts by Wikidata items. The document defines what a server keeping a record, a client using it, and a plugin doing part of the server's work must do to interoperate.

## Status of This Memo

This document is a beta specification published by the Tabellion project. It is not an RFC. It has not been submitted to any RFC stream, and has not been reviewed by the IETF, the IAB, the IRTF or the Independent Submissions Editor.

Its name follows the Internet-Draft naming convention; the number at the end increases with each published revision. This document may be updated, replaced or obsoleted by other documents at any time.

## Copyright Notice

Copyright (c) 2026 Alban Crommer and the Tabellion authors. This document is licensed under the Creative Commons Attribution 4.0 International licence (CC BY 4.0), <https://creativecommons.org/licenses/by/4.0/>.

## Table of Contents

1. [Introduction](#1-introduction)
   1. [Requirements Language](#11-requirements-language)
2. [Terminology](#2-terminology)
3. [Data Model](#3-data-model)
4. [Rules](#4-rules)
5. [Vocabulary](#5-vocabulary)
6. [HTTP Binding](#6-http-binding)
7. [Plugin Contracts](#7-plugin-contracts)
8. [Conformance](#8-conformance)
9. [Internationalization Considerations](#9-internationalization-considerations)
10. [Security Considerations](#10-security-considerations)
11. [Privacy Considerations](#11-privacy-considerations)
12. [IANA Considerations](#12-iana-considerations)
13. [References](#13-references)

Author's Address

## 1. Introduction

A person's skills are usually recorded in documents written for one reader at a time: a CV, an application form, a yearly review. Each is rewritten from memory, and none can be reused by another system.

This document specifies a record kept for one person and owned by that person. The record is a set of facts: occupations held, and practices of a skill on a concept, in a context, at an occupation, over a period. Facts are added, and reviewed by the person, who can correct any of them. The person exposes part of the record through views, read by others under a grant.

The record does not assess the person. How a level is evaluated, and whether a reader trusts a fact, are outside this document.

### 1.1. Requirements Language

The key words "MUST", "MUST NOT", "REQUIRED", "SHALL", "SHALL NOT", "SHOULD", "SHOULD NOT", "RECOMMENDED", "NOT RECOMMENDED", "MAY", and "OPTIONAL" in this document are to be interpreted as described in BCP 14 [RFC2119] [RFC8174] when, and only when, they appear in all capitals, as shown here.

## 2. Terminology

Person
: The subject of a record, identified by a UUID.

Owner
: The person, acting on their own record.

Record
: The set of facts about one person.

Fact
: An occupation held or a practice. A fact has an author, a recording time and a status.

Occupation held
: An occupation the person held over a period: work, pleasure or study.

Practice
: The statement that the person practised a skill on a concept, in an optional context, at an optional occupation, over a period, at an optional level.

Skill
: An ESCO skill, identified by its URI.

Concept
: What a skill was practised on, identified by one Wikidata item, or composed of several (Section 4.4).

Context
: An optional Wikidata item framing a practice, such as a programming language.

Level
: An integer from 1 to 5 (Section 4.3).

Status
: The review state of a fact: proposed, accepted or rejected.

Completion
: The share of the fields of a fact, or of a record, that are filled (Section 4.8).

Evidence
: Data supporting a fact, such as an excerpt of a CV. Opaque to this specification.

View
: A stored selection of facts and of the fields shown, written by the owner.

Reader
: A person, an organisation or a system that reads a record through a view.

Grant
: The permission given by the owner to a reader to read one view.

Task
: Slow work run by a server for the owner, such as importing a CV. Not a fact.

Preference
: How the owner likes a server, or a kind of plugin, to work for them, such as the languages labels are shown in. Not a fact, and not part of the record.

Consent
: The owner's dated agreement to an act that sends or receives personal data, such as calling a plugin on another host. Kept with the record.

Server
: A system that keeps records and enforces the rules of this document.

Client
: A system that reads or changes a record through a server, for a person: the owner or a reader. A client application, such as a web interface or a mobile app, is a client.

Plugin
: A system that a server calls to do part of a task, such as reading a CV. A plugin never reads or changes a record: it returns a result to the server that called it (Section 4.10).

## 3. Data Model

Types: `uuid` is a UUID [RFC9562]; `date` is a year, a year and month, or a full date (YYYY, YYYY-MM or YYYY-MM-DD), at the precision its source gives, and stands for the whole year or month it names; `datetime` is a date-time [RFC3339]; `uri` is an absolute URI [RFC3986]. A field marked optional MAY be absent or null.

### 3.1. Person

| Field | Type | Requirement |
|---|---|---|
| id | uuid, version 4 | required |

A person has no other field. Names and contacts are identities (Section 3.2).

### 3.2. Identity

How a person presents themself to a reader. A server MUST support several identities per person.

| Field | Type | Requirement |
|---|---|---|
| id | string | required |
| person | Person id | required |
| label | string | required; a name or a pseudonym |
| contact | string | optional; e.g. an email address |

### 3.3. Common fields of facts

| Field | Type | Requirement |
|---|---|---|
| id | uuid | required |
| person | Person id | required |
| author | string | required; who states the fact: the person, another person, or the task that proposed it |
| recorded | datetime | required |
| status | proposed, accepted, rejected | required |
| evidence | any JSON value [RFC8259] | optional |

### 3.4. Occupation held

| Field | Type | Requirement |
|---|---|---|
| occupation | uri | optional; an ESCO occupation |
| label | string | required; as stated, e.g. "Backend developer, ACME" |
| kind | work, pleasure, study | required |
| from | date | required |
| to | date | optional; absent: ongoing |

### 3.5. Practice

| Field | Type | Requirement |
|---|---|---|
| skill | uri | required; an ESCO skill |
| concept | Concept (Section 3.6) | required |
| context | uri | optional; a Wikidata item |
| occupation | Occupation held id | optional |
| from | date | required, except on a proposed practice (Section 4.1) |
| to | date | optional; absent: a single date, or ongoing |
| level | integer, 1 to 5 | optional |

### 3.6. Concept

| Field | Type | Requirement |
|---|---|---|
| is | uri | required; a Wikidata item, the genus |
| with | list of uri | required, MAY be empty; Wikidata items that set the concept apart |
| label | string | optional; a name for display |

### 3.7. View

| Field | Type | Requirement |
|---|---|---|
| id | string | required |
| person | Person id | required |
| include.occupation_kind | list of kinds | optional; absent: no restriction |
| include.occupation | list of Occupation held ids | optional |
| include.period | from and to, each an optional date | optional |
| include.skill_under | list of uri, ESCO skill groups | optional |
| include.status | list of proposed, accepted | optional; default in Section 4.5 |
| include_facts | list of fact ids | optional; allow list |
| exclude_facts | list of fact ids | optional; deny list |
| show.level | boolean | required |
| show.evidence | boolean | required |

### 3.8. Grant and access log

| Field | Type | Requirement |
|---|---|---|
| id | uuid | required |
| view | View id | required |
| reader | string | required |
| identity | Identity id | optional; how the person presents themself to this reader |
| status | active, revoked | required |
| expires | datetime | optional |

An access log entry holds a grant id and the datetime of a read.

### 3.9. Task

| Field | Type | Requirement |
|---|---|---|
| id | uuid | required |
| person | Person id | required |
| kind | string, e.g. cv-import | required |
| state | ongoing, finished, abandoned, interrupted | required |
| stage | string | required while ongoing; defined by the kind of task |
| stages | list of stages, each with a name and a progress: done, current, skipped, to do | optional; defined by the kind of task |
| started | datetime | required |
| ended | datetime | optional |
| result | object | optional; per kind |
| problem | problem details [RFC9457] | optional; when interrupted |

### 3.10. Preference

| Field | Type | Requirement |
|---|---|---|
| scope | `server`, or a kind of plugin | required |
| key | string | required; unique within its scope |
| value | any JSON value [RFC8259] | required |
| updated | datetime | required |

Preferences of scope `server` defined by this document:

| Key | Value |
|---|---|
| languages | list of language tags [BCP47], in the owner's order |

### 3.11. Consent

| Field | Type | Requirement |
|---|---|---|
| id | uuid | required |
| person | Person id | required |
| kind | remote-plugin, hosted-terms, grant, proposal | required |
| subject | string | optional; what the consent is for: a plugin URL, a URL of terms, a Grant id, an author. Absent: agreement by default to the whole kind |
| given | datetime | required |
| withdrawn | datetime | optional |

## 4. Rules

### 4.1. Facts

- A record grows by adding facts. A change over time, such as a new level or a new period of practice, MUST be recorded as a new practice: the history of a person's practices is their progress.
- A correction of a wrong fact MUST be made by changing that fact. The fact keeps its id.
- A fact created by anyone other than the owner, including a task, MUST be created with status proposed.
- A server MUST let the owner create facts with status accepted.
- A server MUST let the owner change any field of any fact in their record, its status included, whatever its current status and author.
- A server MUST refuse a practice with status accepted and no `from` date.
- A server MUST let the owner delete facts. A deleted fact is no longer in the record.
- A server MUST NOT delete an occupation held that practices refer to, unless the owner, in the same operation, moves those practices to another occupation held or deletes them.
- A fact proposed from a source SHOULD hold only what the source states. A field the source does not state, or states unclearly, is left empty: an uncertain fact lacks definitions, and its completion shows it.
- A server MUST store and return evidence unchanged. It does not interpret evidence. Changing a fact does not change its evidence: evidence signed by the fact's author may no longer match the fact.

### 4.2. Identifiers

- A person MUST be identified by a version 4 UUID that never changes. Facts carry only this identifier of the person.
- Skills and occupations MUST be ESCO URIs. Concepts and contexts MUST be Wikidata item URIs.
- Status, kind and state are exchanged as the names given in Section 3. How a server stores them is not specified.

### 4.3. Levels

| Level | Name | Meaning |
|---|---|---|
| 1 | Basic | requires guidance |
| 2 | Intermediate | autonomous: applies solutions |
| 3 | Experienced | experience: solves problems |
| 4 | Advanced | defines problems |
| 5 | Expert | finds problems |

Who set a level is the author of the fact. How a level is evaluated is outside this document.

### 4.4. Concepts

- A concept that exists as a Wikidata item is given by that item, with an empty `with` list.
- A concept that does not exist in Wikidata is composed: `is` names the closest Wikidata item, and `with` the Wikidata items that set the concept apart. Example: a Kafka pipeline is a pipeline (Q1192422) with Apache Kafka (Q16235208).
- A server MUST NOT accept a concept identified by anything other than Wikidata items.

### 4.4.1. Human languages

A human language SHOULD be recorded as a practice of the ESCO language skill, with the Wikidata item of the language as concept, over the period the person has used it; for a native language, since birth.

### 4.4.2. Skills named without a period

A skill named in a source without a period, such as in the summary of a CV, SHOULD be proposed as a practice without a `from` date. The owner sets the date before accepting it (Section 4.1).

### 4.5. Views

- A view selects the facts that match its `include` rules or its `include_facts` list, and do not appear in its `exclude_facts` list. The deny list wins.
- A fact matches `include.period` if its own period overlaps it. A `from` date is read as the first day it stands for, a `to` date as the last.
- Without `include.status`, a view selects proposed and accepted facts. A view read under a grant selects only accepted facts unless its `include.status` says otherwise.
- Rejected facts are never selected.
- Fields hidden by `show` MUST NOT be returned to a reader.
- A server MUST apply a view by a deterministic filter.
- A server MUST filter facts through the view before any search, count or derived result, and MUST compute derived results only from the facts the view selects. Searching first would reveal hidden facts through scores, counts or truncated lists.
- A view is not a fact. A server MUST let the owner change a view at any time.

### 4.6. Access

- A reader MUST NOT read a record except through a view, under an active grant that has not expired.
- The owner reads the whole record, and any of their views, without a grant.
- A server MUST record each read under a grant in the access log.
- A server MUST let the owner revoke a grant at any time. A revoked grant gives no further access.

### 4.7. Tasks

- A task is ongoing until it ends: finished, abandoned by the owner, or interrupted by a failure.
- While ongoing, a task is at a stage. The stages of a task, and the transitions between them, are defined by its kind, not by this document.
- Finished and abandoned are final. An interrupted task MAY be resumed, if its kind allows it.
- A server MUST keep a task, with its result, until the owner deletes it. What a result holds, such as terms an import did not match, stays available after the task ends.
- A client MUST NOT assume the stages of a kind it does not know. It shows the stages and the actions the task carries.
- A fact proposed by a task has the task as its author.
- A server MUST let the owner reject, in one operation, every fact still proposed by a task.
- A server MUST let the owner leave out a term an import did not match. Leaving a term out removes it from the result and records no fact.
- A server SHOULD NOT keep the source document of an import once the task has ended. The evidence of each fact keeps the excerpt it comes from.

### 4.8. Completion

- The completion of a practice is the share of these fields that are set: `from`, `level`.
- The completion of an occupation held is 1 if its `occupation` is set, 0 otherwise.
- The completion of a record is the average completion of its facts that are not rejected.
- Completion is derived: a server computes it, and does not store it as a fact.

### 4.9. Server authority

- A server MUST check every change it receives against the rules of this document, whatever the client, and MUST refuse a change that breaks them.
- A client MAY keep data, work offline and apply its own rules on the device. Its changes take effect only when the server accepts them.

### 4.10. Plugins

- A server calls a plugin by its URL, for a kind of plugin that the plugin declares (Section 7.1). A server MUST NOT call a plugin for a kind it does not declare.
- A server MUST send a plugin only the data its call needs, never access to a record.
- A server MUST check what a plugin returns as it checks a change from a client (Section 4.9). Facts built from a plugin's result are proposed by the task that called it (Section 4.7).
- A plugin MUST NOT keep the data it receives once it has answered. It MAY keep a cache that holds no personal data and that it can rebuild, such as vectors of ESCO and Wikidata labels.
- The rules of Sections 3 and 4 MUST NOT depend on a plugin. Without a plugin, a server still lets the owner create, change, review and delete facts, and write views.
- A server SHOULD NOT offer an action that needs a plugin it cannot reach. It MAY tell the owner which kind of plugin is missing.
- A server MAY send a plugin the preferences of the kind it is called for. The plugin MAY return changes to them with its result (Section 7.1). A plugin stores preferences only this way.

### 4.11. Preferences

- Preferences are not facts: a server changes them in place and keeps no history.
- A server keeps the preferences of scope `server`, and those of each kind of plugin it declares. A server, not a plugin, creates a scope.
- Only the owner reads and changes preferences. A grant gives no access to them.
- A server MUST let the owner read the preferences of a kind of plugin, and remove them all in one operation.
- A server SHOULD limit the size of a value and of a scope. Preferences do not hold documents.
- When `languages` is set, a server MUST return labels in those languages, in that order, before the languages a client asks for (Section 9).
- A client MAY keep its own preferences instead of the server's. Keeping the two in step is the client's choice.

### 4.12. Consents

- A server MUST record the owner's consent before it:
  - calls a plugin outside its own host or private network (`remote-plugin`);
  - keeps a record under the terms of a hosted service (`hosted-terms`);
  - gives a reader access to a view (`grant`);
  - accepts facts proposed by an author other than the owner and the owner's tasks (`proposal`).
- A consent without a subject agrees by default to every act of its kind. It is a consent like the others, kept with the record.
- Creating a grant records its consent, with the grant as subject. Revoking the grant withdraws it.
- A server MUST let the owner withdraw a consent at any time. A withdrawn consent allows no further act. A server keeps a withdrawn consent, with its date.

## 5. Vocabulary

The terms of Sections 2 and 3 are exchanged as JSON-LD 1.1 [JSON-LD11]. The context is published with the HTTP binding [TABELLION-API]. Actions use schema.org [SCHEMA-ORG]; identities use vCard [VCARD-RDF].

The IRI of the vocabulary is <https://tabellion.app/vocab#>.

## 6. HTTP Binding

The HTTP binding is described by an OpenAPI contract [OPENAPI], published by the Tabellion project [TABELLION-API]. It interprets this document: it follows its requirements and chooses among its options. Where they disagree, this document prevails.

The binding chooses, among others: JSON-LD responses with the actions available on each resource, problem details [RFC9457] for refused changes, JSON merge patch [RFC7396] for changes to facts, and entity tags for freshness and concurrent changes [RFC9110].

## 7. Plugin Contracts

A contract is what a server sends to a plugin of one kind, and what the plugin returns. This document does not define an HTTP binding for these contracts.

### 7.1. Common to all plugins

- A plugin MUST describe itself: the kinds it implements, the revision of this document it follows, and the options of each kind.
- A plugin that cannot do a call MUST return a problem detail [RFC9457]. The server then interrupts the task (Section 4.7).
- Every call MAY carry `preferences`: the preferences of the kind, as an object of keys and values. Every result MAY carry `preferences`: the keys to change, a null value removing the key (Section 4.10).

### 7.2. Term matching

Proposes facts from a text, such as the text of a document read by another plugin.

| Direction | Field | Type | Requirement |
|---|---|---|---|
| Sent | lines | list of strings | required |
| Sent | language | language tag of the text, e.g. fr | optional |
| Returned | occupations | occupations held: the fields of Section 3.4, without id, author, status or recording time | required, may be empty |
| Returned | practices | practices: the fields of Section 3.5, without id, author, status or recording time; the occupation as an index in `occupations` | required, may be empty |
| Returned | unmatched | terms found without a skill or a concept, each with its excerpt and line | required, may be empty |

- The line numbers of the text are the positions in `lines`, from 1. Evidence refers to them.
- Each returned occupation and practice carries its evidence: the excerpt it comes from and its line.
- A plugin SHOULD return only what the text states (Section 4.1).

## 8. Conformance

A conforming server implements Sections 3 and 4.

A conforming client sends only changes allowed by Section 4, and handles the refusal of any change.

A conforming plugin implements Sections 4.10 and 7.1, and the contract of each kind it declares.

Test data for conformance, a record and the document it was written from, is published by the Tabellion project [TABELLION-TESTDATA].

## 9. Internationalization Considerations

- Skills, occupations and concepts are identifiers; their labels exist in many languages in ESCO and Wikidata. A server SHOULD return labels in the language the client asks for, after the owner's `languages` when set (Section 4.11).
- Status, kind and state are names, not text for display; clients translate them.
- Identity labels and occupation labels are free text, in any language.

## 10. Security Considerations

- A grant gives access to one view. A server MUST check the grant on every read, as it can be revoked or expire.
- Evidence is opaque and can come from any author. A client MUST NOT run or render evidence as active content.
- A server reachable from a web browser, including one listening on localhost only, can receive requests from any web page the browser opens. A server SHOULD authenticate every request, also on localhost.
- A plugin may run on another host, or be run by another provider. A server SHOULD authenticate the plugins it calls, and MUST use TLS to call a plugin outside its own host or private network.
- What a plugin returns is untrusted input, like any change from a client (Section 4.10).
- Facts state what their author claims. This document gives no means to verify a claim; a reader decides whether to trust it.

## 11. Privacy Considerations

- A record describes a person's working life: it is personal data.
- A person is identified by a UUID only. Names and contacts are kept apart, in identities, and given to a reader only through the identity chosen for a grant.
- The person decides what each reader sees; the deny list of a view always wins.
- The access log tells the person who read which view, and when.
- Preferences are personal data. They are read by the owner only, and by a plugin only for its own kind (Section 4.11).
- Source documents, such as a CV, are not kept beyond their import (Section 4.7).
- A server MUST NOT send usage data, such as telemetry, to anyone. It sends personal data only to the plugins the owner's tasks call.
- A plugin receives personal data, such as the text of a CV, and does not keep it (Section 4.10). A server SHOULD tell the owner which plugin, and which provider, receives their data before a task sends it, and MUST have the owner's consent for a plugin outside its own host or private network (Section 4.12).

## 12. IANA Considerations

This document has no IANA actions.

## 13. References

### 13.1. Normative References

- [RFC2119] Bradner, S., "Key words for use in RFCs to Indicate Requirement Levels", BCP 14, RFC 2119, March 1997, <https://www.rfc-editor.org/info/rfc2119>.
- [RFC3339] Klyne, G. and C. Newman, "Date and Time on the Internet: Timestamps", RFC 3339, July 2002, <https://www.rfc-editor.org/info/rfc3339>.
- [RFC3986] Berners-Lee, T., Fielding, R., and L. Masinter, "Uniform Resource Identifier (URI): Generic Syntax", STD 66, RFC 3986, January 2005, <https://www.rfc-editor.org/info/rfc3986>.
- [RFC8174] Leiba, B., "Ambiguity of Uppercase vs Lowercase in RFC 2119 Key Words", BCP 14, RFC 8174, May 2017, <https://www.rfc-editor.org/info/rfc8174>.
- [RFC8259] Bray, T., Ed., "The JavaScript Object Notation (JSON) Data Interchange Format", STD 90, RFC 8259, December 2017, <https://www.rfc-editor.org/info/rfc8259>.
- [RFC9562] Davis, K., Peabody, B., and P. Leach, "Universally Unique IDentifiers (UUIDs)", RFC 9562, May 2024, <https://www.rfc-editor.org/info/rfc9562>.
- [BCP47] Phillips, A., Ed., and M. Davis, Ed., "Tags for Identifying Languages", BCP 47, RFC 5646, September 2009, <https://www.rfc-editor.org/info/rfc5646>.
- [ESCO] European Commission, "ESCO: European Skills, Competences, Qualifications and Occupations", <https://esco.ec.europa.eu>.
- [WIKIDATA] Wikimedia Foundation, "Wikidata", <https://www.wikidata.org>.
- [JSON-LD11] Sporny, M., Longley, D., Kellogg, G., Lanthaler, M., Champin, P., and N. Lindström, "JSON-LD 1.1", W3C Recommendation, July 2020, <https://www.w3.org/TR/json-ld11/>.

### 13.2. Informative References

- [RFC7396] Hoffman, P. and J. Snell, "JSON Merge Patch", RFC 7396, October 2014, <https://www.rfc-editor.org/info/rfc7396>.
- [RFC9110] Fielding, R., Ed., Nottingham, M., Ed., and J. Reschke, Ed., "HTTP Semantics", STD 97, RFC 9110, June 2022, <https://www.rfc-editor.org/info/rfc9110>.
- [RFC9457] Nottingham, M., Wilde, E., and S. Dalal, "Problem Details for HTTP APIs", RFC 9457, August 2023, <https://www.rfc-editor.org/info/rfc9457>.
- [SCHEMA-ORG] "Schema.org", <https://schema.org>.
- [TABELLION-API] Tabellion project, "Tabellion HTTP binding: OpenAPI contract and JSON-LD context", <https://github.com/tabellionhq/tabellion/tree/main/api>.
- [TABELLION-TESTDATA] Tabellion project, "Tabellion conformance test data", <https://github.com/tabellionhq/tabellion/tree/main/testdata>.
- [VCARD-RDF] Iannella, R. and J. McKinney, "vCard Ontology - for describing People and Organizations", W3C Interest Group Note, May 2014, <https://www.w3.org/TR/vcard-rdf/>.
- [OPENAPI] OpenAPI Initiative, "OpenAPI Specification v3.1.0", February 2021, <https://spec.openapis.org/oas/v3.1.0>.

## Author's Address

Alban Crommer\
Tabellion project\
URI: <https://github.com/tabellionhq/tabellion>

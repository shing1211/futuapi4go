# Graph Report - docs  (2026-05-07)

## Corpus Check
- Corpus is ~4,030 words - fits in a single context window. You may not need a graph.

## Summary
- 143 nodes · 140 edges · 18 communities
- Extraction: 100% EXTRACTED · 0% INFERRED · 0% AMBIGUOUS
- Token cost: 0 input · 0 output

## Community Hubs (Navigation)
- [[_COMMUNITY_Changelog v0.0-0.5 Entries|Changelog v0.0-0.5 Entries]]
- [[_COMMUNITY_Python API - Account & Functions|Python API - Account & Functions]]
- [[_COMMUNITY_Proto - General Definitions|Proto - General Definitions]]
- [[_COMMUNITY_Proto - Encryption & Protocol ID|Proto - Encryption & Protocol ID]]
- [[_COMMUNITY_Python Quote API - Market Data|Python Quote API - Market Data]]
- [[_COMMUNITY_Proto - Basic Functions|Proto - Basic Functions]]
- [[_COMMUNITY_Proto - Protocol Design|Proto - Protocol Design]]
- [[_COMMUNITY_Changelog - Tests & Features|Changelog - Tests & Features]]
- [[_COMMUNITY_Python - Config & OpenD|Python - Config & OpenD]]
- [[_COMMUNITY_Changelog - Categories|Changelog - Categories]]
- [[_COMMUNITY_Changelog v0.5.1|Changelog v0.5.1]]
- [[_COMMUNITY_Changelog v0.5.0|Changelog v0.5.0]]
- [[_COMMUNITY_Changelog v0.3.0|Changelog v0.3.0]]
- [[_COMMUNITY_Changelog v0.2.5|Changelog v0.2.5]]
- [[_COMMUNITY_Changelog v0.2.4|Changelog v0.2.4]]
- [[_COMMUNITY_Changelog v0.2.2 Phase 3|Changelog v0.2.2 Phase 3]]
- [[_COMMUNITY_Changelog v0.2.1 Phase 2|Changelog v0.2.1 Phase 2]]
- [[_COMMUNITY_Changelog v0.0.5|Changelog v0.0.5]]

## God Nodes (most connected - your core abstractions)
1. `Changelog` - 21 edges
2. `Quote API (Market Data)` - 11 edges
3. `General Definitions` - 9 edges
4. `[0.2.0] - 2026-04-25` - 5 edges
5. `[0.0.5] - 2026-04-23 — Feature Parity Achieved` - 5 edges
6. `Introduction` - 5 edges
7. `[Unreleased]` - 4 edges
8. `[0.5.1] - 2026-04-28` - 4 edges
9. `[0.5.0] - 2026-04-27` - 4 edges
10. `[0.3.0] - 2026-04-25` - 4 edges

## Surprising Connections (you probably didn't know these)
- None detected - all connections are within the same source files.

## Communities (18 total, 0 thin omitted)

### Community 0 - "Changelog v0.0-0.5 Entries"
Cohesion: 0.07
Nodes (27): [0.0.1] - 2026-04-12, [0.0.2] - 2026-04-18, [0.0.3] - 2026-04-18, [0.0.4] - 2026-04-19, [0.0.6] - 2026-04-24, [0.2.0] - 2026-04-25, [0.2.3] - 2026-04-25, [0.2.6] - 2026-04-25 (+19 more)

### Community 1 - "Python API - Account & Functions"
Cohesion: 0.11
Nodes (17): Account Types, API Categories, API Frequency Limits, Common Error Codes, Error Categories, Error Handling, Features, Functionality (+9 more)

### Community 2 - "Proto - General Definitions"
Cohesion: 0.12
Nodes (17): code:protobuf (message ProgramStatus), code:protobuf (enum RetType), code:protobuf (enum ProtoFmt), code:protobuf (enum PacketEncAlgo), code:protobuf (enum ProgramStatusType), code:protobuf (enum GtwEventType), code:protobuf (enum NotifyType), code:protobuf (message PacketID) (+9 more)

### Community 3 - "Proto - Encryption & Protocol ID"
Cohesion: 0.18
Nodes (10): AES Encryption, Basic Protocols, Encrypted Communication Process, Futu OpenAPI Protocol Reference (Proto), Proto ID Reference, Protocol Introduction, Protocol Request Process, Quote Protocols (3xxx) (+2 more)

### Community 4 - "Python Quote API - Market Data"
Cohesion: 0.18
Nodes (11): Capital Flow, K-line Data, Market Snapshot, Market State, Order Book, Plates, Quote API (Market Data), Stock Filter (+3 more)

### Community 5 - "Proto - Basic Functions"
Cohesion: 0.25
Nodes (8): Basic Functions, code:protobuf (message C2S), code:protobuf (syntax = "proto2";), Connection Encryption, InitConnect, KeepAlive, Protocol ID Table, Push Protocol Format

### Community 6 - "Proto - Protocol Design"
Cohesion: 0.25
Nodes (8): code:c (struct APIProtoHeader), code:protobuf (message C2S), code:protobuf (message S2C), Protobuf Request Structure, Protobuf Response Structure, Protocol Body, Protocol Design, Protocol Header

### Community 7 - "Changelog - Tests & Features"
Cohesion: 0.4
Nodes (5): [0.0.5] - 2026-04-23 — Feature Parity Achieved, Added, Changed, Fixed, Tests

### Community 8 - "Python - Config & OpenD"
Cohesion: 0.4
Nodes (5): code:python (from futu import OpenQuoteContext, OpenSecTradeContext), Configuration, Connection, Installation, OpenD

### Community 9 - "Changelog - Categories"
Cohesion: 0.5
Nodes (4): Added, Documentation, Fixed, [Unreleased]

### Community 10 - "Changelog v0.5.1"
Cohesion: 0.5
Nodes (4): [0.5.1] - 2026-04-28, Added, Changed, Fixed

### Community 11 - "Changelog v0.5.0"
Cohesion: 0.5
Nodes (4): [0.5.0] - 2026-04-27, Added, Changed, Documentation

### Community 12 - "Changelog v0.3.0"
Cohesion: 0.5
Nodes (4): [0.3.0] - 2026-04-25, Added, Already Existed, Breaking Changes (v0.3.0)

### Community 13 - "Changelog v0.2.5"
Cohesion: 0.5
Nodes (4): [0.2.5] - 2026-04-25, Added (P4-4), Changed, Completed (Previously Existed)

### Community 14 - "Changelog v0.2.4"
Cohesion: 0.67
Nodes (3): [0.2.4] - 2026-04-25, Added, Fixed

### Community 15 - "Changelog v0.2.2 Phase 3"
Cohesion: 0.67
Nodes (3): [0.2.2] - 2026-04-25, Added (Phase 3 Infrastructure), code:go (if constant.IsTimeout(err) { /* handle timeout */ })

### Community 16 - "Changelog v0.2.1 Phase 2"
Cohesion: 0.67
Nodes (3): [0.2.1] - 2026-04-25, Added (Phase 2 Ease of Use), code:go (trd.NewOrder(accID, market, env).Buy("00700", 100).At(350.5))

### Community 17 - "Changelog v0.0.5"
Cohesion: 0.67
Nodes (3): [0.0.5] - 2026-04-21, Added, Changed

## Knowledge Gaps
- **89 isolated node(s):** `Fixed`, `Added`, `Documentation`, `Added`, `Changed` (+84 more)
  These have ≤1 connection - possible missing edges or undocumented components.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `Changelog` connect `Changelog v0.0-0.5 Entries` to `Changelog - Tests & Features`, `Changelog - Categories`, `Changelog v0.5.1`, `Changelog v0.5.0`, `Changelog v0.3.0`, `Changelog v0.2.5`, `Changelog v0.2.4`, `Changelog v0.2.2 Phase 3`, `Changelog v0.2.1 Phase 2`, `Changelog v0.0.5`?**
  _High betweenness centrality (0.194) - this node is a cross-community bridge._
- **Why does `General Definitions` connect `Proto - General Definitions` to `Proto - Encryption & Protocol ID`?**
  _High betweenness centrality (0.054) - this node is a cross-community bridge._
- **Why does `Protocol Introduction` connect `Proto - Encryption & Protocol ID` to `Proto - Protocol Design`?**
  _High betweenness centrality (0.041) - this node is a cross-community bridge._
- **What connects `Fixed`, `Added`, `Documentation` to the rest of the system?**
  _89 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `Changelog v0.0-0.5 Entries` be split into smaller, more focused modules?**
  _Cohesion score 0.07 - nodes in this community are weakly interconnected._
- **Should `Python API - Account & Functions` be split into smaller, more focused modules?**
  _Cohesion score 0.11 - nodes in this community are weakly interconnected._
- **Should `Proto - General Definitions` be split into smaller, more focused modules?**
  _Cohesion score 0.12 - nodes in this community are weakly interconnected._
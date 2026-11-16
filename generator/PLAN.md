# Fixing the generator — plan

**Goal:** the generator is again the one source of every generated file. Regenerating it
reproduces the library, and every rule it applies is either a clause of TS 29.244 or a decision
recorded here.

**Where we start** (reviewed 2026-11-16):
- **The input:** `pfcp-tlv-etri.py`, an open5gs port, reads TS 29.244 v17.7.1 (`29244-h71`) from
  a hand-edited docx into `cache/`. It emits `tmp/`: `tags.go`, a file per grouped IE and one
  `message.go`, on ndn-dpdk's TLV.
- **Since then:** the library was laid out and fixed by hand. Only 3 of 112 generated files
  still match its output after renames.
- **Three kinds of difference:**
  - mechanical changes the templates never learned;
  - files hand-written by design;
  - hand edits to generated content.
- **The generator's own faults:** repetition taken from open5gs's needs, not the
  specification; a row dropped by name; decoding that breaks TS 29.244 §7.6.9 and §7.6.10.

---

## Phase A — the generator under version control

1. Commit `generator/` to this repository:
   - the script;
   - `cache/`, which is the generator's real input: a run never reads the docx while the cache
     exists;
   - `field-ies.py`, both docx files, and a README.
2. The README states the pipeline: docx → `cache/` → generated files. It names the files that
   are **generated**, the ones that are **hand-written** (and must never be overwritten), and how
   to regenerate.
3. Pin `python-docx`. The cache was built with an unknown version, and its parse of merged cells
   is what the tables depend on.

### Which files are which

| kind | files |
| --- | --- |
| generated | `ie/tags.go`; every grouped IE in `ie/`; every message file in `message/` except those below |
| hand-written, the leaf codecs | the `ie/*IE.go` of the types `field-ies.py` lists: F-TEID, SDF Filter, Application ID, Gate Status, MBR, GBR, Volume Threshold, Apply Action, F-SEID, Node ID, Measurement Method, FQ-CSID, Volume Quota, Outer Header Creation, UE IP Address, Outer Header Removal, Measurement Information |
| hand-written, the rest | `tlv/`; `common.go`; `ie/{cause,mics,VolumeMeasurement,bitrate,DLBufferingSuggestedPacketCountIE}.go`; `message/{header,codec,msgtype,json,Association}.go`; every `_test.go` |

## Phase B — reproduce today's library

Change only the templates, until regenerating writes the library as it is. Nothing changes for
any caller in this phase.

- **Form:**
  - packages `ie` and `message`, and their imports;
  - this repository's `tlv`;
  - decoding through `tlv.Elements` with its error returned;
  - the loop variable `i`;
  - a file per message pair, named as today;
  - `gofmt` on the output.
- **Hand edits that are rules, stated in the generator:**
  - **An IE that is not mandatory is a pointer**, a slice or `[]byte` (4d0c2a1's MBS fix is this
    rule applied).
  - **A mandatory grouped or custom-typed IE is a pointer**, as the first commit made Create PDR's
    PDI and Create QER's Gate Status. This is the rule as it stands; phase C2 decides whether it
    stays.
  - **`NewForwardingParameters`** moves to a hand-written helper file.

**Acceptance:** regenerate, then diff against the library. The only differences left are the
ones phase C makes on purpose, listed here:
- the seven repeated-bytes fields the first commit made single;
- the PDR lists of 2c99995, which C1 produces from the specification.

## Phase C — the rules fixed

One commit per rule. Each is regenerated and its diff reviewed against the expectation written
here.

- **C1 — repetition from the specification's own text.**
  - **List or single** is read from the row's comment, not from open5gs's `tlv_more` list. The
    phrases are:
    - "Several IE(s) with/within the same IE type";
    - "More than one IE with this type";
    - "There may be several instances";
    - "Two IEs with the same IE type";
    - "the list of all the … IDs".
  - **A short override table** covers what the text cannot say, each entry citing its row.
  - **A repeated `[]byte` IE is a list again** (`[][]byte`). One function in the generator maps an
    IE to its Go type, list or single, so giving a repeated IE a named type later is one edit
    there.
  - **Expected diff:** the fields of the 2026-11-16 comparison become lists, and nothing else
    changes. That is about 25 in messages, among them Query URR, the traffic endpoints, MAR, SRR,
    Updated PDR and Session Report. It is about 35 in grouped IEs, among them PDI's QFI and framed
    routes, Duplicating Parameters, Linked URR ID and Downlink Data Report's PDR ID.
- **C2 — presence.** The pointer rule for mandatory IEs, as decided (below). If mandatory IEs are
  pointers, the decoder reports a mandatory IE that is missing.
- **C3 — §7.6.9: an unknown or unexpected IE is discarded and the message processed.** The
  generated decoder skips it, where today it fails the whole message with `tlv.ErrCritical`.
  How it is reported is decided below.
- **C4 — §7.6.10: of a repeated IE where repetition is not specified, the first is kept.** A
  single field takes the first occurrence and ignores the rest; today the last wins.
- **C5 — no row is dropped.** Two rows of one IE type in a grouped IE are both kept: PDI's Local
  F-TEID and Local F-TEID for Redundant Transmission, removed today by name. The representation
  is decided below.
- **C6 — fewer things that break silently on a new release:**
  - message tables found by title, not by hard-coded docx index;
  - leaf sizes checked against the IE definitions;
  - the two docx edits ("Adjust table cell", "Add User Plane IP Resource Information") made a
    script.

## Phase D — tests the generator writes

- **A conformance table from `cache/`:** for every message and grouped IE, each field's list or
  single, and its presence, against its row. A rule that regresses fails here, not on a peer.
- **A round trip for every grouped IE and message:** every field set, encoded, decoded and
  compared; every list given two entries.
- **The decoding rules:** an unknown IE skipped (C3), and the first of a repeated single kept
  (C4).
- **Kept:** the existing wire tests against go-pfcp's captured encodings, and `header_test.go`.

## Phase E — release and adoption

1. Regenerate, build and test. Rerun the go-pfcp comparison: no single-against-list difference
   should remain except FQ-CSID, where ours is the better form, and User Plane IP Resource
   Information, which the specification has removed.
2. Volantis: adapt to the fields that became lists. The SMF's QAURR workaround for Query URR can
   go, since a modification can name several URRs.
3. Release **v1.1.0**: many field types change, and every caller has to follow.

## Phase F — later: Release 19

A separate step, so that its diff holds only what changed between releases:
- the v19.x docx, downloaded by hand, since 3GPP's server refuses scripts;
- the edit script of C6 run on it;
- the cache rebuilt, regenerated, and the new and changed IEs reviewed.

---

## Decided

- **Repeated IEs are lists**, a repeated `[]byte` IE included (`[][]byte`). The Go type of a
  repeated IE stays open for later change, behind one function of the generator (C1).

- **Mandatory IEs are pointers (C2)**, and the decoder reports a mandatory IE that is missing: a
  value cannot tell a missing mandatory IE from a zero one.
- **An unknown or unexpected IE is skipped and its type returned beside the decoded message
  (C3)**, so the caller can log it as §7.6.9 asks.
- **Two rows of one IE type in a grouped IE are two fields, filled in order (C5)**, as the rows
  name two meanings.
- **The correction stays on v17.7.1**; Release 19 is phase F.

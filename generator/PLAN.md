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
| hand-written, the rest | `tlv/`; `common.go`; `errors.go`; `ie/{cause,mics,VolumeMeasurement,bitrate,DLBufferingSuggestedPacketCountIE,constructors}.go`; `message/{header,codec,msgtype,json,Association}.go`; every `_test.go` |

## Phase B — reproduce today's library

*Done 2026-11-16.* 96 of the 107 generated files come out exactly. The other 11 differ only
in the three ways the end of this phase lists.

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

**As built:**
- The form is generated: packages, imports written only when used, qualified marshallers,
  `tlv.Elements`, `i` in messages, a file per message pair.
- **The pointer rule is narrower than it looked.** Only Create PDR's PDI and Create QER's Gate
  Status were made pointers by hand; nine other mandatory custom-typed fields are values. The
  generator names the two (`pointer_fields`), and C2 makes it the rule for all.
- **A presence override** (`presence_overrides`): MBS Session N4mb Control Information is M on
  N4mb alone. Its Sxa, Sxb, Sxc and N4 columns are "-", which the generator does not read, so as
  a value it went out on every N4 request (4d0c2a1).
- **The two hand-written constructors** moved to `ie/constructors.go`.

**What still differs, all of it phase C's:**
- **The repeated-bytes IEs:** User Plane IP Resource Information in the association messages,
  and PDI's and Create Traffic Endpoint's Framed-Route and Framed-IPv6-Route. The generator
  writes lists; the library made them single by hand. They stay lists (decided).
- **Create PDR's QER ID, and Update PDR's QER and URR IDs:** lists in the library (2c99995),
  single in the generator until C1.
- **Reporting Triggers and Usage Report Trigger:** C7.

## Phase C — the rules fixed

*Done 2026-11-16*, C1 to C8. C7 needs no change to the generator, which writes the two
trigger IEs in 3 octets already: the library takes it at regeneration (phase E). The
output builds; the library's own tests set fields that became lists, and follow in phase E.

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
- **C7 — the 3-octet IEs.** Reporting Triggers (§8.2.19, Figure 8.2.19-1) and Usage Report
  Trigger (§8.2.41, Figure 8.2.41-1) are octets 5 to 7. The generator encodes them as 3 octets;
  the library encodes 4 (`MarshalUint32`, since its first commit), one more than the
  specification and readable by no other implementation. Regenerating fixes the encoding. Each
  is a `uint32` of flags, so Volantis must be checked for which bits the octets land in: its flag
  constants, and the UPF's byte reversal of the usage report trigger.
- **C8 — Extendable IEs decoded as TS 29.244 8.1.2 has it.** Nearly every sized IE is
  Extendable: PDR ID, Precedence, Recovery Time Stamp, the flags. Its octet count is the minimum
  it was first defined with, and "the legacy receiving entity shall ignore the unknown octets".
  The decoders demanded an exact length, so a peer of a later release adding an octet had its
  message refused. An Extendable sized IE now decodes through `pfcp.UnmarshalExtendable`: from
  its first-defined length up, missing known octets 0, further octets ignored.
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

- **Mandatory IEs stay values (C2, revised 2026-11-16).** Pointers for all of them would change
  most of the API: PDR ID, FAR ID, Node ID and the rest. The decoder reports a missing one
  instead, as a typed error naming its IE type, from the seen flags C4 keeps. A UPF answers it
  with Cause "Mandatory IE missing" and the Offending IE. Only rows mandatory **on N4** are
  checked: a row is M on the interfaces its columns mark, and MBS Session N4mb Control
  Information is M on N4mb and "-" on N4. So the parse reads the N4 column. The two hand-made
  pointers, PDI and Gate Status, go back to values.
- **An unknown or unexpected IE is skipped and recorded (C3, revised 2026-11-16).** Every message
  and grouped IE has `Unknown []uint32` (`json:"-"`), the types skipped. A grouped IE's are
  merged into its parent's, so a caller reads the message's alone and logs it, as §7.6.9 asks.
  `UnmarshalBinary` has no room to return them beside the message, and an error would read as a
  failure.
- **Two rows of one IE type in a grouped IE (C5, revised 2026-11-16):** the two F-TEID rows of
  Created PDR, Created Traffic Endpoint and Updated PDR. TS 29.244 gives no order between them,
  and Updated PDR lists them the other way round. So it is a fixed rule, stated as an
  assumption: the first F-TEID on the wire is the Local F-TEID, a second the one for redundant
  transmission, and they are encoded in that order whatever the table's.
- **The correction stays on v17.7.1**; Release 19 is phase F.

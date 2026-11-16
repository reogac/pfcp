# The generator

`pfcp-tlv-etri.py` writes the grouped IEs, the messages and the IE type constants of this
library from the tables of TS 29.244. It is a port of Open5GS's `pfcp-tlv.py` (AGPL-3.0, its
header kept) from C to Go. `PLAN.md` is the work to make it, again, the one source of every
file it generates.

## The pipeline

```
29244-h71.docx  --docx-edits.py-->  29244-h71-modified.docx  --parse-->  cache/*.py  --generate-->  Go files
```

- **The specification:** TS 29.244 **v17.7.1** (`29244-h71`, Release 17). It is converted from
  `.doc` to `.docx`, then edited by `docx-edits.py` (below) into `29244-h71-modified.docx`. The docx files are
  3GPP's text and are not committed: download `29244-h71.zip` from the 3GPP archive, convert it,
  and apply the edits.
- **`cache/`** is the generator's real input. While it exists the docx is not read at all, so
  regenerating needs only the cache. It is exactly the parse of the docx: each row's name,
  presence, comment and whether it applies on N4. Delete it to parse the docx again.
- **`field-ies.py`** lists the leaf IEs whose codec is written by hand. The generator writes a
  stub `<Type>IE.go` for each, once, and never overwrites it.

### The edits

`docx-edits.py` turns `29244-h71.docx` into `29244-h71-modified.docx`, and `--check` compares
its result with a copy, table by table. The edits:

| where | edit |
| --- | --- |
| Association Setup and Update, request and response | User Plane IP Resource Information back in, as Release 16 had it: Release 17 removed it, and Open5GS still used it |
| the IE type table | its type 116 back in; three names spelled as the rest of the specification spells them |
| two grouped IE headers | a footnote marker dropped; type 303 named "MBS Session N4mb Information", as the IE type table names it |
| IP Multicast Addressing Info | its second Source IP Address row, M on N4mb alone, dropped: the IE has two rows of one type |
| the generator, not the docx | Update Forwarding Parameters' "Network instance" read as "Network Instance" (`name_overrides`), so the field is `NetworkInstance` like every other one |

They were recovered from a copy edited by hand: `--check` finds every table equal, and the cache
parsed from the script's output equals the committed one.

## Regenerating

```bash
pip install -r requirements.txt        # the versions the cache was checked with
python3 pfcp-tlv-etri.py -f 29244-h71-modified.docx -o <dir>/
gofmt -w <dir>
```

The script needs the docx path even when it reads only the cache. It writes the library's
layout under `<dir>`:
- `ie/tags.go`;
- a file per grouped IE in `ie/`;
- a file per message and its response in `message/`;
- in `ie/`, the stubs of `field-ies.py` not yet there.

**Regenerate into a scratch directory and compare, never over the library**, until `PLAN.md`
phase C is done. Today the output differs from the library in 11 files, each in a way phase C
settles (`PLAN.md`, phase B).

**Checked 2026-11-16, with the versions in `requirements.txt`:**
- From the docx, the script rebuilds `cache/` exactly.
- From `cache/`, it writes 96 of the library's 107 generated files exactly.

## Which files are which

| kind | files |
| --- | --- |
| generated | `ie/tags.go`; every grouped IE in `ie/`; every message file in `message/` except those below; `internal/samples/samples.go`; `ie/` and `message/`'s `roundtrip_test.go` and `shapes_test.go` |
| hand-written, the leaf codecs | the `ie/*IE.go` of the types `field-ies.py` lists |
| hand-written, the samples and the rule tests | `internal/samples/leaves.go`, a value of each hand-written leaf; `message/rules_test.go`, the decoding rules of TS 29.244 7.6 and 8.1.2 |
| hand-written, the rest | `tlv/`; `common.go`; `errors.go`; `ie/{cause,mics,VolumeMeasurement,bitrate,DLBufferingSuggestedPacketCountIE,constructors}.go`; `message/{header,codec,msgtype,json,Association}.go`; every `_test.go` |

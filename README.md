# pfcp

A PFCP (TS 29.244 v17.7.1) library in Go: every message and IE, the header, and the TLV codec.

```go
import (
	"github.com/reogac/pfcp/v2/ie"
	"github.com/reogac/pfcp/v2/message"
)
```

**This repository is generated.** Its source is
[`github.com/reogac/pfcp-gen`](https://github.com/reogac/pfcp-gen): the grouped IEs and the
messages are written from the specification's tables, and the hand-written files are kept in
that repository's `lib/`. A change made here is lost at the next generation, so make it there.

What the decoders do, from TS 29.244:
- a repeated IE is a list wherever the specification lets it repeat;
- a missing IE mandatory on N4 is reported as `*pfcp.MissingIE`, naming its type, but in a
  Heartbeat Request and in a response with a rejection Cause (7.6.6);
- an unknown or unexpected IE is skipped, its type recorded in the message's `Unknown` (7.6.9);
- of a single IE that repeats, the first is kept (7.6.10);
- an Extendable IE is read from its first-defined length up, further octets ignored (8.1.2).

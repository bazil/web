---
title: Roadmap
---

This document serves as a big-picture roadmap of where Bazil is
headed. This is all speculative and may change at any time.

---------------------------------------------------------

## <span id="near"/>Near future

### <span id="sync"/>Synchronize file changes between peers

Work in progress, slowed by lacking RPC infrastructure.

### <span id="kvmulti"/>Concurrent multi-backend KV storage

Current version is sequential and naive.

Introduce concept of cancel/interrupt to CAS and KV APIs: `cancel
<-chan struct{}`, tie into `fuse.Intr` when serving FS request.

Bandwidth/latency estimation to prioritize backend attempts. Staggered
time delays before starting too many concurrent requests, but move
forward faster if too few requests in flight.

Circuit breaker for better behavior under faults.

### <span id="sftp"/>Remote KV store access over SFTP

### <span id="kvshuffle"/>Shuffle data in/out of KV stores interactively

Sketch: `bazil get PATH`, `bazil push PATH KVSTORE`

### <span id="synckv"/>Synchronize KV stores on the background

### <span id="reflink"/>Support reflink for instant file copies

---------------------------------------------------------

## <span id="far"/>Further out

### <span id="pin"/>Pin file content

Persistently remember and respect pins, strive to keep pinned data on
local disk.

Sketch: `bazil pin PATH`, `bazil unpin PATH`

### <span id="arena"/>Arena storage for CAS objects

Once an arena is sealed, make an index with perfect hashing.

Support object deletion by hole punching.

### <span id="s3"/>Remote KV store access over S3

### <span id="review"/>Review crypto usage

Review convergencent encryption against ideas in
[Cryptosphere](https://github.com/cryptosphere/cryptosphere/wiki/Data-Model).

### <span id="gc"/>Garbage collection

Challenging because it's distributed garbage collection in a
weakly-connected system.

### <span id="tree-space"/>Recursive lazy accounting of subtree size etc

Make directory reflect the size, file counts, modification time etc
possible aggregates of the whole subtree.

Update lazily in the background.

Report the non-standard aggregations via xattrs?

See [Ceph](http://ceph.com/) for ideas on the user experience.

---------------------------------------------------------

## <span id="bluesky"/>Blue Sky

These may never happen, but are still worth writing down.

### <span id="remote-client"/>Remote client access through SFTP/CIFS/NFS

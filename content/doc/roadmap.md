---
title: Roadmap
menu:
  doc:
    weight: 7
---

This document serves as a big-picture roadmap of where Bazil is
headed. This is all speculative and may change at any time.

---------------------------------------------------------

##  Near future {#near}

###  Synchronize file changes between peers {#sync}

Work in progress, slowed by lacking RPC infrastructure.

###  Concurrent multi-backend KV storage {#kvmulti}

Current version is sequential and naive.

Introduce concept of cancel/interrupt to CAS and KV APIs: `cancel
<-chan struct{}`, tie into `fuse.Intr` when serving FS request.

Bandwidth/latency estimation to prioritize backend attempts. Staggered
time delays before starting too many concurrent requests, but move
forward faster if too few requests in flight.

Circuit breaker for better behavior under faults.

###  Remote KV store access over SFTP {#sftp}

###  Shuffle data in/out of KV stores interactively {#kvshuffle}

Sketch: `bazil get PATH`, `bazil push PATH KVSTORE`

###  Synchronize KV stores on the background {#synckv}

###  Support reflink for instant file copies {#reflink}

---------------------------------------------------------

##  Further out {#far}

###  Pin file content {#pin}

Persistently remember and respect pins, strive to keep pinned data on
local disk.

Sketch: `bazil pin PATH`, `bazil unpin PATH`

###  Arena storage for CAS objects {#arena}

Once an arena is sealed, make an index with perfect hashing.

Support object deletion by hole punching.

###  Remote KV store access over S3 {#s3}

###  Review crypto usage {#review}

Review convergencent encryption against ideas in
[Cryptosphere](https://github.com/cryptosphere/cryptosphere/wiki/Data-Model).

###  Garbage collection {#gc}

Challenging because it's distributed garbage collection in a
weakly-connected system.

###  Recursive lazy accounting of subtree size etc {#tree-space}

Make directory reflect the size, file counts, modification time etc
possible aggregates of the whole subtree.

Update lazily in the background.

Report the non-standard aggregations via xattrs?

See [Ceph](http://ceph.com/) for ideas on the user experience.

---------------------------------------------------------

##  Blue Sky {#bluesky}

These may never happen, but are still worth writing down.

###  Remote client access through SFTP/CIFS/NFS {#remote-client}

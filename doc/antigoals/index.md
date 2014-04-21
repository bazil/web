---
title: Anti-goals
---

This document collects things that are explicitly not design goals for
Bazil.

Trying to do these things would *probably* force us to trade off
something we consider more desirable. For example, global locks would
prevent weakly connected and offline operation, and thus there is
nothing to be done there.

If you have an idea of how to gain some of these *without trading off*
the unique and desirable aspects of Bazil, please do let us know!


## <span id="posix"/> Full POSIX compatibility

Cannot be done while supporting weakly connected operation.

For example, a `rename` of a file on one peer might not be immediately
visible on another peer. The second peer might even happily write to
the old file name. At some later time, the changes will be
synchronized, in this case resulting in an *update/delete conflict* on
the original file name.

We try to provide useful local operation semantics (for example,
atomic renames), reasonable distributed semantics and *always* detect
conflicts where they've occurred.

## <span id="limits-hardlink"/> Hard links

For simplicity, Bazil does not support hard links. This may never
change, as behavior of hard links when synchronizing remote changes
gets really murky.

``` console
$ ln foo bar    # not supported
```

## <span id="mknod"/> `mknod` support

Bazil is a userspace filesystem, and what would these things even mean
in a distributed context?

``` console
$ mknod foo b 12 34    # not supported
```

## <span id="sync-write"/> Immediate write visibility to other peers

We follow something closer to the AFS "commit-on-close". File state
becomes visible to peers only on `fsync` or `close`.

Current implementation even
[holds dirty file content fully in memory](/doc/status#limits-inmem).


## <span id="global-inodes"/> Global inode numbers across peers

The synchronization operates on filenames, not on inodes. Inodes are
local to one peer.

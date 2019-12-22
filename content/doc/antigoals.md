---
title: Anti-goals
menu:
  doc:
    weight: 6
---

This document collects things that are explicitly not design goals for
Bazil.

Trying to do these things would *probably* force us to trade off
something we consider more desirable. For example, global locks would
prevent weakly connected and offline operation, and thus there is
nothing to be done there.

If you have an idea of how to gain some of these *without trading off*
the unique and desirable aspects of Bazil, please do let us know!


##  Heavy transactional workloads {#tx}

Don't run your SQL database on top of Bazil. It'll never be the best
possible fit for that, both because of the userspace indirection and
because of how Bazil stores its data.

##  Application-specific conflict resolution {#app-conflict}

Writing application and file format specific merging algorithms is an
endless swamp. Bazil won't prevent you from writing your own, and
we'll make it as easy as we can, but as a project we won't spend
effort on it.

##  Full POSIX compatibility {#posix}

Cannot be done while supporting weakly connected operation.

For example, a `rename` of a file on one peer might not be immediately
visible on another peer. The second peer might even happily write to
the old file name. At some later time, the changes will be
synchronized, in this case resulting in an *update/delete conflict* on
the original file name.

We try to provide useful local operation semantics (for example,
atomic renames), reasonable distributed semantics and *always* detect
conflicts where they've occurred.

##  Hard links {#limits-hardlink}

For simplicity, Bazil does not support hard links. This may never
change, as behavior of hard links when synchronizing remote changes
gets really murky.

``` console
$ ln foo bar    # not supported
```

##  `mknod` support {#mknod}

Bazil is a userspace filesystem, and what would these things even mean
in a distributed context?

``` console
$ mknod foo b 12 34    # not supported
```

##  Immediate write visibility to other peers {#sync-write}

We follow something closer to the AFS "commit-on-close". File state
becomes visible to peers only on `fsync` or `close`.

Current implementation even
[holds dirty file content fully in memory](/doc/status#limits-inmem).


##  Global inode numbers across peers {#global-inodes}

The synchronization operates on filenames, not on inodes. Inodes are
local to one peer.

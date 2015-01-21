---
title: Bazil is still in development
---

In early **alpha** testing.

The data formats in Bazil are still undergoing major changes, and
significant functionality is missing or has not been tested well
enough.

While we don't expect it to literally lose data, we will not put extra
effort into data format compatibility until the first formally
released version.

As it is, Bazil is intended primarily for other developers and power
users exploring future options. Do not stare into laser with remaining
eye.

The current version has not been tuned for performance and should not
be taken as an indication of what will be possible later.

Bazil is provided "as is" and with no warranty. See the file
[LICENSE](https://github.com/bazillion/bazil/blob/master/LICENSE) in
the source distribution for more.

With all that said, we are very enthusiastic about the possible uses
of Bazil, and think that the architecture can be very resilient
against data loss.


##  Limitations {#limits}

Bazil is still missing features, and has intentional limits to keep it
simple and focused on specific aspects of the design. These limitation
may be removed later.

See also [anti-goals](/doc/antigoals) which are less likely to be
"fixed".

See also [roadmap](/doc/roadmap) for planned features.

###  No communication between peers yet {#no-sync}

At this point in time, Bazil is a single-node system. It'll store your
data, take snapshots of it, and you could use for example `rsync` to
backup the object store, but Bazil itself is not yet ready to transfer
data between computers, or synchronize file changes between two peers.

The design for peer synchronization exists and is pretty solid, and
current work focuses on a good RPC mechanism; most existing frameworks
assume small messages, can't prioritize data transfers, or are
constrained by to strict client-server request-response messaging that
is not the best match for two-way synchronization traffic.

See [roadmap](/doc/roadmap) for [sync](/doc/roadmap#sync), remote KV
access via [SFTP](/doc/roadmap#sftp) and [S3](/doc/roadmap#s3),
[background synchronization](/doc/roadmap#synckv), and so on.

###  Dirty file data is in memory {#limits-inmem}

Current implementation does not write file contents to disk until it
sees `fsync` or `close`. That means creating a large file consumes a
lot of RAM temporarily.

###  Garbage collection is not implemented yet {#limits-gc}

Right now, objects added to the CAS remain there permanently. A plan
for garbage collection exists (though distributed, weakly connected GC
is tricky!), but implementing it is not yet a priority.

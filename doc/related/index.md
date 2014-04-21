---
title: Related work
---

Bazil stands on the shoulders of some giants, and was inspired by
shoulders of other giants. Or something like that.

## <span id="fuse"/> FUSE

[FUSE](https://en.wikipedia.org/wiki/Filesystem_in_Userspace) is a
protocol between the kernel and a userspace process, letting the
userspace serve file system requests coming from the kernel.

The kernelspace implementation has
[very partial documentation](https://www.kernel.org/doc/Documentation/filesystems/fuse.txt),
but the
[public API](https://git.kernel.org/cgit/linux/kernel/git/torvalds/linux.git/tree/include/uapi/linux/fuse.h)
and the
[implementation](https://git.kernel.org/cgit/linux/kernel/git/torvalds/linux.git/tree/fs/fuse)
are the best guides because they are the only thing that matters in
the end.

Bazil *does not use* the original C library at
[`fuse.sourceforge.net`](http://fuse.sourceforge.net/). Instead, we
have an independent implementation of the protocol, in Go:
[`bazil.org/fuse`](http://bazil.org/fuse/).

[OSXFUSE](http://osxfuse.github.io/) is the FUSE kernel side ported to
Apple's OS X.


## <span id="venti"/> Plan 9's Venti

Bazil's use of a [CAS](#cas) is similar to how
[Plan 9](http://doc.cat-v.org/plan_9/)'s
[Fossil](http://doc.cat-v.org/plan_9/4th_edition/papers/fossil/)
archived nightly snapshots to
[Venti](http://doc.cat-v.org/plan_9/4th_edition/papers/venti/).

Bazil uses the CAS a lot more aggressively, pushing content to it as
soon as reasonable, instead of as part of a nightly dump.

Bazil has multiple peers independently accessing the CAS data, instead
of having a singular service maintaining the state of the filesystem.

Bazil will ([at some point](/doc/status#limits-gc)) perform garbage
collection, instead of relying on any "640kB should be enough for
everyone" mentality.

## <span id="blake2"/> Blake2 hash algorithm

[Blake2](https://blake2.net/) is a fast, cryptographically secure, and
flexible hash. It can be personalized, keyed, and has no fixed output
length.

Every single use of Blake2 in Bazil is *personalized*; you can't build
a "rainbow table" of generic Blake2 results and attack Bazil with
that, the table would have to be specific to Bazil.

Where it makes sense, the hash is *keyed* with a nonce-equivalent, for
example some unique bytes that are going to be always known at the
time the hash is computed. This makes even the above Bazil-specific
rainbow table worthless.

The output lengths are generally chosen to be overly large, for extra
margin of safety and robustness in future -- yet the lengths are just
a matter of tuning. For example, the size of the CAS key is a single
constant, allowing easy experimentation on whether there is a
performance difference (there doesn't seem to be).

And after all this, Blake2 is *fast*.

## <span id="nacl"/>NaCl

[NaCl](http://nacl.cr.yp.to/) is a delightfully simple & fast crypto
library. Bazil uses the
[Go reimplementation](http://godoc.org/code.google.com/p/go.crypto/nacl)
of it.

## <span id="bolt"/>Bolt

[Bolt](https://github.com/boltdb/bolt) is a key/value store, with an
emphasis at very fast reads. It's API is a pleasure to use.

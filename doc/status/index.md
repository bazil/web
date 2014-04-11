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


## <span id="limits"/> Limitations

Bazil is still missing features, and has intentional limits to keep it
simple and focused on specific aspects of the design. These limitation
may be removed later.

### <span id="limits-hardlink"/> No hard links

For simplicity, Bazil does not support hard links. This may never
change, as behavior of hard links when synchronizing remote changes
gets really murky.

``` console
$ ln foo bar    # not supported
```

### <span id="limits-gc"/> Garbage collection is not implemented yet

Right now, objects added to the CAS remain there permanently. A plan
for garbage collection exists (though distributed, weakly connected GC
is tricky!), but implementing it is not yet a priority.

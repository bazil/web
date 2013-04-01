---
title: Go FUSE file system library
go:
  repo: fuse
---

`bazil.org/fuse` is a Go library for writing filesystems. It is a
from-scratch implementation of the kernel-userspace communication
protocol, and does not use the C library from the project called
[FUSE](http://fuse.sourceforge.net/). `bazil.org/fuse` embraces Go
fully for safety and ease of programming.

Here's how to get going:

    go get bazil.org/fuse

Github repository: https://github.com/bazillion/fuse

API docs: http://godoc.org/bazil.org/fuse

Our thanks to [Russ Cox](http://swtch.com/~rsc/) for his [fuse
library](https://code.google.com/p/rsc/source/browse/#hg%2Ffus), which
this project is based on.

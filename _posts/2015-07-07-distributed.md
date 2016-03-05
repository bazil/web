---
title: "Status update: a <em>distributed</em> filesystem"
tags:
- bazil
---

A lot of time has passed, and a lot of code has been written. Bazil is
still in heavy development, but it has reached a good milestone to
blog about: it can synchronize changes from one peer to another.

Warning: at this stage in development, we will put no effort into
compatibility of file formats or protocols. Do not stare into laser
with remaining eye.

What follows is a walkthrough of scenario where we have two computers
sharing files -- find me at [GopherCon](http://www.gophercon.com/) for
a live demo, or follow the steps and run it yourself.


## Installation

First, make sure you have a working [Go](http://golang.org/) (>=1.4)
installation. You are expected to have basic familiarity with Go, at
this point in development.

Unfortunately, to work around a
[missing gRPC feature](https://github.com/grpc/grpc-go/issues/111), we
need a custom branch of it for now. Let's check that out:

```console
$ go get google.golang.org/grpc
$ cd $GOPATH/src/google.golang.org/grpc
$ git checkout 5e5f5df2bbfed81a191eb0484831738cc729f3b9
```

And then install Bazil itself:

```console
$ go get bazil.org/bazil
```


## Initialization

For the rest, we'll assume you have two computers, virtual machines or
containers that will talk to each other.

You can also run the steps on one host, by calling passing the `bazil
-data-dir=PATH` option as appropriate to keep two separate state
directories.

We'll call our two environments `black` and `white`, and differentiate
them with that hostname in the prompt.

```console
white$ bazil create
```

```console
black$ bazil create
```

## Public keys

To introduce the peers to each other, we need to pass their public
keys to each other. As the current code doesn't actually keep track of
any nicknames or aliases for peers, we'll need to refer to these
public keys a lot. Let's set shell variables to remember them.

To see the public key of a node, run

```console
white$ bazil debug pubkey
```

Typically, debug commands access the database directly, and will only
work if the server is not running.

Now set the variable `$BLACK` on the host `white` with the value being
the public key of `black`, and vice versa. If you're running the two
on the same host, the following will work; if not, copy-pasting with
the mouse is needed.

```console
white$ BLACK="$(bazil -data-dir=path/to/datadir/of/black debug pubkey)"
```

```console
black$ WHITE="$(bazil -data-dir=path/to/datadir/of/white debug pubkey)"
```

As is probably obvious from the `debug` in the command name, this is
not the final UX for this.


## Running the server

Bazil has a (per-user) server component that the command-line
utilities communicate with. Let's start the server on `white`.

```console
white$ bazil server run &
bazil: Listening on [::]:34211
```

```console
black$ bazil server run &
bazil: Listening on [::]:nnnnn
```

## Making friends

We believe in the value of encryption. Bazil uses
[convergent encryption](https://bazil.org/doc/architecture/) with
*sharing keys* where the people who know the relevant sharing key can
have access to the data.

The default installation sets up one sharing key, but let's make a new
one for our shared files; it's just 32 bytes of random data. We'll
name our new sharing key `friends`.

```console
white$ dd if=/dev/urandom of=sekrit bs=32 count=1
white$ bazil sharing add friends <sekrit
```

Let's create a volume using the new sharing key, and mount it.

```console
white$ bazil volume create -sharing=friends myfiles
white$ mkdir mnt
white$ bazil volume mount myfiles mnt
```

We now have an encrypted, deduplicating, snapshottable, *local* file
system. Let's share it with `black`, using the public key stored in
`$BLACK` from earlier.

We introduce a new peer, identified by the public key stored in
`$BLACK`. We tell `white` to allow `black` to access its local
[content-addressed storage](https://bazil.org/doc/architecture/), and
the `myfiles` volume we just created.

```console
white$ bazil peer add $BLACK
white$ bazil peer storage allow $BLACK local
white$ bazil peer volume allow $BLACK myfiles
```

Let's tell `black` to use the new volume. First, we introduce the
`white` as a new peer for `black`, and giving the network location
where the server on `white` is listening on. The server prefers the
port 34211 (*bazil*, do you see it?), but will use any free port. We
saw the port output earlier.

```console
black$ bazil peer add $WHITE
black$ bazil peer location set $WHITE 192.0.2.42:34211
```

Later, we'll introduce more rendezvous mechanisms, including multicast
DNS and an internet-wide lookup based on the public key, and
mechanisms for working behind NATs.

`black` needs to know the sharing key from earlier. Copy the `sekrit`
file from `white` to `black` through whatever means are appropriate,
and then run

```console
black$ bazil sharing add friends <sekrit
black$ bazil volume connect -sharing=friends $WHITE myfiles
black$ bazil volume storage add -sharing=friends myfiles peerkey:$WHITE
black$ mkdir mnt
black$ bazil volume mount myfiles mnt
```

We now have the save volume mounted on two machines.


## A *distributed* filesystem

Let's make changes on `white` and observe them on `black`.

```console
white$ echo hello, world >mnt/greeting
```

```console
black$ bazil volume sync myfiles $WHITE
black$ ls mnt
black$ cat mnt/greeting
```

```console
white$ echo hello, again >mnt/greeting
```

```console
black$ bazil volume sync myfiles $WHITE
black$ cat mnt/greeting
```

Hey! It works!


## Limitations

The sync implementation doesn't currently handle deletions or
subdirectories.

There is currently no user interface to resolve conflicts, or to
finish sync merges that were postponed because a file was still open.

At this stage in development, we will put no effort into compatibility
of file formats or protocols.


## Future

After the obvious missing functionality mentioned is done, there's
plenty of work to be done on making the user experience of managing
peers better. The steps above are very manual and discrete right now,
as that is what's easiest to debug.

Once the common usage scenarios have been explored, more convenient
mechanisms can be added on top of these low-level steps, e.g.
bootstrapping a peer connection over ssh, and interacting with friends
over im with humans copy-pasting short messages.

To learn more about the **why** of Bazil, read the
[introductory blog post](https://bazil.org/2014/04/24/introducing-bazil.html).

To understand the architecture of Bazil better, browse the
documentation https://bazil.org/doc/ .

Bazil is still at an early stage in development, but the future looks
really exciting. We'd love to have you participating.

---
title: Usage
---

Bazil uses a per-user daemon, and communicates with it over UNIX
domain sockets.

Our examples use two computers, `squirrel` being a desktop computer
and `pocketgopher` a laptop.

The daemon is not yet automatically started, so the first step is to
start it. This will go away later.

``` console
squirrel$ bazil server run &
```

``` console
pocketgopher$ bazil server run &
```

## Creating a volume

Bazil organizes your files by *volume*. Let's create our first volume:

``` console
squirrel$ bazil volume create mascots
```

As Bazil is a *file system*, we need to `mount` it, to attach it to
the directory tree. We'll use a different name for the volume and
mountpoint for readability.

``` console
squirrel$ mkdir pics
squirrel$ bazil volume mount mascots pics
```

And now you can use the mount:

``` console
squirrel$ cd pics
squirrel$ wget http://blog.golang.org/gopher/gopher.png
```

You can archive a snapshot of the volume by just naming it:

``` console
squirrel$ mkdir .snap/justincase
```

You can later access the snapshot by just browsing `.snap/justincase`.

Once you're done with the volume, unmount it safely with:

``` console
# can't be inside the mount, or it'll stay busy
squirrel$ cd

# Linux
squirrel$ fusermount -u pics

# OS X
squirrel$ umount pics
```

## Adding a peer

 <div class="alert alert-warning">
**Work in progress**: this functionality is not ready yet. It is
provided here to give you an idea of what things might look like.
See the [status](/doc/status#no-sync) page.
</div>

To share volumes with another computer, we need to add it as a *peer*.
The arguments to `bazil peer add` are `NAME` and `DIALER ARGS...`.
`NAME` is just a way to refer to this peer. `DIALER` selects the
communication mechanism.

``` console
pocketgopher$ bazil peer add desktop ssh squirrel.example.com
```

### The `ssh` dialer

The `ssh` dialer will use your SSH authorization (agent, keys and
passphrase, as needed) to connect to the given hostname, and talks to
the Bazil server there. It adds the remote as a peer to your local
instance, and the local instance as a peer to the remote (with no way
to dial back configured, yet).

## Sharing volumes

 <div class="alert alert-warning">
**Work in progress**: this functionality is not ready yet. It is
provided here to give you an idea of what things might look like.
See the [status](/doc/status#no-sync) page.
</div>

Once you have a peer configured, you can create a new local volume
that is synchronized with a volume on the peer. Because our
communication mechanism is `ssh`, we don't need the peer to grant us
permission; we're running commands on the remote computer already,
with access to the relevant files.

The usage is `bazil volume link PEER/VOLUME [NEWVOLUME]`

``` console
pocketgopher$ bazil volume link desktop/mascots mascots
```


## Offline use

 <div class="alert alert-warning">
**Work in progress**: this functionality is not ready yet. It is
provided here to give you an idea of what things might look like.
See the [status](/doc/status#no-sync) page.
</div>

To make file content available locally, without the network
connection, we `pin` the data.

``` console
pocketgopher$ mkdir pics
pocketgopher$ bazil volume mount mascots pics
pocketgopher$ cd pics
pocketgopher$ bazil pin .
```

Pinning just expresses a wish. To make sure the files have actually
copied, we need to wait:

``` console
pocketgopher$ bazil sync wait
```


## Storage locations

 <div class="alert alert-warning">
**Work in progress**: this functionality is not ready yet. It is
provided here to give you an idea of what things might look like.
See the [status](/doc/status#no-sync) page.
</div>

Usage is `bazil volume store VOLUME add NAME DIALER ARGS..`

TODO I'm not 100% happy with this command line; this may end up
happening via editing a file instead.

``` console
squirrel$ bazil volume store mascots add goog google-cloud-storage gs://mahbukkit/
```

TODO talk about data dispersion policies here


## Multiple Bazil instances

By default, configuration and file data are stored in
platform-specific user data directories: `~/.local/share/bazil/` on
many unixes, `~/Library/Application Support/bazil` on OS X.

The bulk file data can be moved around later, but this directory is
used as the starting point. If you want to use multiple Bazil data
stores, with completely separate volumes -- for example, to test a new
build before using it for real -- use

``` console
$ bazil -data-dir=PATH ...
```

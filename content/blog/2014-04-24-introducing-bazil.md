---
date: 2014-04-24T00:00:00Z
tags:
- bazil
- project
- announce
title: Introducing Bazil
slug: introducing-bazil
aliases:
- introducing-bazil.html
---

<img style="float: right; max-height: 10em;" src="/img/gopher-with-disk.png" />

[GopherCon](http://www.gophercon.com/) is here, and it is time to
reveal what Bazil is all about.

Bazil, also known as `bazil.org/bazil`, is a file system that lets
your data reside where it is most convenient for it to reside.

Bazil is still [under heavy development](/doc/status), but welcomes
developers and curious power users. Here's a little teaser of what's
coming.

Imagine you have

  - <i class="fa fa-laptop"></i> A laptop with a a 256GB SSD
  - <i class="fa fa-desktop"></i> A desktop with a 3TB hard disk
  - <i class="fa fa-cloud"></i> A cloud server or storage service
    <br/>(virtual machine with expandable disks, S3, etc)
  - **or** <i class="fa fa-linux"></i> just a cheap computer in a closet with two slow 4TB disks
  - **or** <i class="fa fa-hdd-o"></i> external USB disk you plug in once a week for backups

On the desktop, you naturally want to be able to use the whole 3TB
disk. And you're not always using the desktop, even when you're home
-- the sofa is just so comfortable. You'd like to work with your files
even when you're on the laptop.

## First try: Let's sync to the Cloud<sup style="font-size: 0.6em">[*](https://chrome.google.com/webstore/detail/cloud-to-butt-plus/apmlngnhgbnjpajelfkmabhkfapgnoai)</sup>

So you install the currently fashionable large-corporate-backed
cloud-sync solution.

A file sync based solution will try to copy all of your files from the
desktop to the laptop -- yet the laptop's smaller SSD just *can't hold
that much*! You're forced to play games with *picking-and-choosing*
what folders get synchronized, and just don't have the *convenience*
of grabbing that 8-year-old wedding photo on a whim.

  - <i class="fa fa-desktop"></i> <i class="fa fa-check text-success"></i>
    Desktop use is just ok: you need to keep adjusting what folders are
    synced and what not
  - <i class="fa fa-laptop"></i> <i class="fa fa-frown-o text-danger"></i>
    Laptop use is miserable
  - <i class="fa fa-cloud"></i> <i class="fa fa-meh-o text-danger"></i>
    Cloud storage of 3TB gets expensive: Dropbox or AWS will charge you
    $1000/year for the storage alone. That's about 30TB of hard drives.
    <br/>And most providers don't inspire confidence on the privacy of your
    files.
  - **or** <i class="fa fa-linux"></i> <i class="fa fa-frown-o
    text-danger"></i> The large corporations are not interested in
    supporting your server in a closet.
  - **or** <i class="fa fa-hdd-o"></i> <i class="fa fa-frown-o
    text-danger"></i> These are the wrong corporations to make money off
    of you buying hard drives, so they have no interest in supporting
    that either. Why don't you rent more online space, you're easier
    to monetize that way.

To modernize an aphorism, you can't put ten terabytes of files on a
500 GB SSD. Syncing between very disproportionate systems is
fundamentally a problematic design, and is best for a small
hand-picked set of files, not as an actual storage solution.

Don't take this the wrong way; you really should have some sort of
*remote backups* for important data, in case the building burns down.
S3 RRS/Glacier, Google Cloud Storage DRA seem very promising for
backup cold storage; we'll come back to that later.


## Second try: Use a network file system

Rocking it old school? We're down with that.

A network file system like CIFS or NFS, or something like `sshfs`,
would let you use the files from the desktop on the laptop -- but your
wifi will *never be as fast* as the laptop's local SSD, in either
bandwidth or latency, so now all your file accesses are crawling, and
you end up hunting for an ethernet cable whenever you need to transfer
something bigger.

To speed things up, you end up copying often used files to the SSD.
Now you have several copies of the same files, and *no idea* what was
modified when, or whether you're looking at the last copy, or whether
it's *safe to delete* to free up space on the cramped laptop.

A network file system will also require for you to stay within wifi
range. For travel, you're once again reduced to up *manually copying
files around*, and once again lose track of where's the *latest copy*
of what file.

  - <i class="fa fa-desktop"></i> <i class="fa fa-meh-o
    text-danger"></i> Desktop use is kinda sorta tolerable: you're never
    sure whether the file you are looking at is the latest copy

  - <i class="fa fa-laptop"></i> <i class="fa fa-frown-o
    text-danger"></i> Laptop use is miserable: you're confused about
    which copies of your files are the right ones, the network file
    system is an umbilical tying you to your home network, and
    everything goes over the slow wifi all the time

  - <i class="fa fa-cloud"></i> <i class="fa fa-meh-o text-danger"></i>
    Cloud storage is still expensive, but now you can use it as backup
    only and bypass the synchronization service providers: switching
    between clouds is easier, and cold storage and reduced
    availability is cheaper.

    However, this leaves you installing & configuring cloud backup
    software *in addition to* your network file system woes; not the
    simplest ordeal, and don't expect any kind of file history
    browsing / recovery integration for you network file system
    clients.

  - **or** <i class="fa fa-linux"></i> <i class="fa fa-meh-o
    text-danger"></i> You can choose to back up to your own disks --
	with the same caveats as above

  - **or** <i class="fa fa-hdd-o"></i> <i class="fa fa-frown-o
    text-danger"></i> All of the bad parts of the computer in the
    closet, with the extra of needing to fiddle with the disks
	and remember things.


## What Bazil does

Bazil separates knowledge of a file from the contents of the file,
letting the laptop know *about* all of the files, without having to
store the *contents* of the file.

With Bazil, the laptop SSD contents act as

- a *cache*: file contents accessed recently can be stored
  temporarily, in case they are needed again
- a *buffer*: new content is saved fast on the SSD, and transferred to
  the desktop / cloud / server in a closet in the background
- a *stash*: the user can *pin* files for use when offline or
  just using a slower Internet connection

And because Bazil keeps track of the changes, it can also keep track
of changes and synchronize them between the different peers; no more
confusion about what copy is the latest.

You try to read a file where the contents are not locally stored, the
data will be fetched from desktop or cloud/closet server, whichever
happens to be the fastest way. All the data is accessible even if it
won't fit on the SSD.

You can *pin* files for travel, so you're no longer tied to your home
network, or even any network connectivity.

Bazil *is* the archival solution, with the
[snapshot](/doc/architecture#snapshot) feature. Every Bazil peer can
browse the earlier snapshots, making restoring files easy no matter
what computer you're on. You don't have to manage both a network file
system and a backup solution.

Bazil *is* the redundancy solution, with copies of file contents
stored on multiple computers. The [CAS](/doc/architecture#cas) stores
immutable, write-once objects, so you can even mitigate software bugs
by taking an extra copy of the history with just `rsync`, file system
snapshots, or any other file copy tool. A snapshot is just an object,
and refers to other objects; the objects contain everything needed to
regain access to your files.

All Bazil file storage can be [encrypted](/doc/architecture#crypto) to
guarantee your privacy, whether in the cloud, on your own computers,
or on external hard drives. Encryption is *on by default*.

  - <i class="fa fa-desktop"></i> <i class="fa fa-check
    text-success"></i> Desktop use is good: changes are synced at the
    first opportunity
  - <i class="fa fa-laptop"></i> <i class="fa fa-check
    text-success"></i> Laptop use is good: all the data is accessible,
    often used files are cached, changes are synchronized, files can
    be made available offline
  - <i class="fa fa-cloud"></i> **and** <i class="fa fa-linux"></i> <i
    class="fa fa-check text-success"></i> You can mix-and-match cloud
    storage providers and servers in closet as you please
  - **and** <i class="fa fa-hdd-o"></i> <i class="fa fa-check
    text-success"></i> you can use external disks for extra space, with
    Bazil keeping track of what data is what on disk, even when they
    are unplugged, or even use them to avoid slow Internet transfers


## Current status

Bazil is still [under heavy development](/doc/status), and a lot of
the functionality hinted at above is still not quite there. We welcome
developers and curious power users.

[See the documentation for more](/doc/), and feel free to ask
questions on the
[mailing list](https://groups.google.com/group/bazil-dev)
or Twitter [@BazilFS](https://twitter.com/BazilFS).

{{<small>}}
The original gopher image was made by [Renee French](http://blog.golang.org/gopher).
{{</small>}}

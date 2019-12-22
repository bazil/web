---
title: Introduction to Bazil
menu:
  doc:
    name: "Introduction"
    weight: 1
---

{{<lead>}}
*Each of the concepts below links to more background. <br>
Click them to learn more, or keep reading.*
{{</lead>}}

Bazil is a [file system](/doc/architecture#filesystem), a special
folder on all of your computers where the same contents are available
on every machine -- we'll call each one a **peer** -- at least after
they've had a chance to transfer the data.

Bazil keeps your data private, using
[convergent encryption](/doc/architecture#crypto) that lets you
decide who is able to use your files -- **sharing with friends** when
you want to.

Your files are stored where *you* want to store them. Bazil is not
tied to a proprietary cloud storage provider. You can store your files
fully on **local hard drives**, or your own **private server**, if you
wish.

All of the files you store in a Bazil **volume** are
[deduplicated](/doc/architecture#cas) -- copies of the same file
contents take up no extra space -- and a
[snapshot](/doc/architecture#snapshot) can be stored as a backup,
letting you browse old versions of all the files.

Bazil is *not* just a file synchronizer. Each computer stores only
what it has disk space for, and the rest of the data is fetched over
the network as needed -- preferring the nearby, faster sources -- and
[cached](/doc/architecture#filesystem) locally. This means a laptop
with limited disk space can still access all your files. You can
**pin** files for **offline use**, for example airplane travel. This
also helps when you are **weakly connected** -- that is, your internet
connectivity is slower than usual, for example when tethering to a
cell phone, or at a hotel.

Changes made while **offline** or just **weakly connected** are
synchronized between all **peers**. There is *no* central location,
and *no* cloud service that you would need to pay for, if you don't
want to use one.

If two different **offline** **peers** make changes to the same file
at or nearly the same time, you may get **conflicts**. This means you
will have both versions of the file available, and you need to decide
what to do. This is just like it is with more traditional file
synchronizers, and is fairly rare in practice.

Bazil is **free software**, **open source** and **free to use**. It
builds on a large pool of [previous work](/doc/related).

Bazil is still [in development](/doc/status), and not ready for every
day users.

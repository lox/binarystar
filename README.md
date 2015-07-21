
Binary Star - Bi-directional sync between two hosts
===================================================

Provides a daemon that connects with a daemon on another host (OSX/Linux/Windows) that work together to keep
a local directory synchronized with as little latecy as possible. The daemon monitors the filesystem
for changes and publishes them to it's peer.

Designed for keeping code in-sync between OSX and a VirtualBox VM for Docker Containers.

How it works
------------

The daemon can either act as a server with `-listen :24249` or as a client with `-connect hostname:24249`.

Both sides establish two channels, a event channel on which they receive filesystem events over and a request/response channel for requesting file transfers from the other side.

Authentication is presently not implemented.


Related projects
----------------

 - lsyncd
 - unison
 - csync / csync2

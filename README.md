
# Binary Star - Bi-directional sync between two hosts

Connects two hosts (OSX/Linux/Windows) together to asynchronously replicate files that match a certain pattern when they are changed. Provides a single statically linked binary that will work out-of-the-box in common environments with no dependencies.

Designed for keeping code in-sync between OSX and a virtual machine for Docker containers.

## Usage

The daemon can either act as a server with `-listen :24249` or as a client with `-connect hostname:24249`. One of your hosts must be the server and the other the client. The client can be behind a firewall, but the server must have it's port accessible by the client.

Host A:

```bash
# start one on your host
binarystar --listen 0.0.0.0:8924 $HOME/code

# connect to it from somewhere else, docker, VM's, remote machines, etc
binarystar --connect my.host.blah:89924 ./code

# Files are now replicated on change between these locations
```

## How it works

Each side maintains an internal tree with metadata about the local files. Operating system specific tools are used to listen for filesystem changes and then these are transmitted across the connection as binary diff's.

Merge conflicts are determined at a file level, the most recent change will win. The only exception to this is the initial sync, in which case the listening host will determine the initial state.

## Related projects

 - lsyncd
 - unison
 - csync / csync2

# IDs

## Guild ID

A `uint64` ID. Represents a *Guild*.
This should be unique in a server.

## Channel ID

A `uint64` ID. Represents a *Channel*.
This should be unique in a *Guild*.

## Message ID

A `uint64` ID. Represents a *Message*.
This should be unique in a *Channel*.

## File ID

A `string` ID. Represents a file on a server.
This should be unique in a server.

For clients, this means that they should download
from the server they got the file ID from.

## HarMony Content (HMC)

A URL in the format `hmc://server:port/id`.
Represents a file on a server.

- The `server` part is the domain of the server and must always be provided.
- The `id` part is a file ID and must always be provided.
- The `port` part is the port and must always be provided.
# IDs

## User ID

A `uint64` ID. Represents a user.
This must be unique in a server.

## Guild ID

A `uint64` ID. Represents a guild.
This must be unique in a server.

## Channel ID

A `uint64` ID. Represents a channel.
This must be unique in a *Guild*.

## Message ID

A `uint64` ID. Represents a message.
This must be unique in a channel.

## Emote Pack ID

A `uint64` ID. Represents an emote pack.
This must be unique in a server.

## Local File ID

A `string` ID. Represents a file on a server.
This must be unique in a server.

For clients, this means that they should download
from the server they got the file ID from.

## HarMony Content (HMC)

A URL in the format `hmc://server:port/id`.
Represents a file on a server.

- The `server` part is the domain of the server and must always be provided.
- The `id` part is a file ID and must always be provided.
- The `port` part is the port and must always be provided.

## File ID

This is an ID that represents a file on a server. 
It can be one of the following:

- A local file ID,
- A HMC.
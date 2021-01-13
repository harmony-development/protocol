# High-Level Overview

## Preparing for E2EE

Clients should generate a Ed25519 public and private key.

### Public Key

The public key should be posted to the client's homeserver (not a foreign server) using the PostKey request found in the SecretService.

## Private Key

Clients should prompt for a password from the user. This password should be hashed with HMAC-SHA256 in order to generate a 256-bit key to use with AES.

The generated key should be used to encrypt the private key for storage with the server, and the encrypted private key should be posted to the homeserver using the PostPrivateKey request found in the SecretService.

Whenever the user signs in with a new client, the new client should fetch the encrypted private key and prompt for the password that was used to encrypt it.

## Creating an E2EE chat

A client that wants to create an E2EE chat should call the GenerateChannelPair request on their homeserver's SecretService. This will generate two event streams, one for messages and one for room state.
The client *must* generate two 256-bit AES keys, and assign one to the message stream and the other to the state stream.
These will become the initial keys used to encrypt and decrypt messages.

## Inviting users to the chat

The client should look up the public key of the user it wants to invite, and post a TrustKey event to the state stream.
Afterwards, it should post a new key event to both the message and the state stream.

Afterwards, it should encrypt the following information using the user's public key and send it to the new user:
- Event IDs of the required new key events for the message and state streams
- Current state stream key
- Current message stream key
- An initial room state that should include:
  - Room name
  - Room photo
  - Trusted users

## Posting messages to the chat

Anyone with the room key can post events to the message and state stream, encrypting it using the current key of the stream.

## Receiving messages from the chat

Clients should decrypt incoming events using their current key for the stream. If the incoming message is a valid new_key event or contains a fanout, it should update its current key for the stream and use it for events coming afterwards.

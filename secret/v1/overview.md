# High-Level Overview

## Preparing for E2EE

Clients should generate a 4098-bit RSA public and private key.

### Public Key

The public key should be posted to the client's homeserver (not a foreign server) using the PostKey request found in the SecretService.

## Private Key

Clients should prompt for a password from the user. This password should be hashed with HMAC-SHA256 in order to generate a 256-bit key to use with AES.

The generated key should be used to encrypt the private key for storage with the server, and the encrypted private key should be posted to the homeserver using the PostPrivateKey request found in the SecretService.

Whenever the user signs in with a new client, the new client should fetch the encrypted private key and prompt for the password that was used to encrypt it.

## Creating an E2EE chat

A client that wants to create an E2EE chat should call the GenerateChannelPair request on their homeserver's SecretService. This will generate two event streams, one for messages and one for room state.
The client *must* generate two cryptographically secure random 128-bit numbers. These will be used to seed two olm ratchets, and assigned to the two streams by the client.

## Inviting users to the chat

The client should look up the public key of the user it wants to invite, and post a TrustKey event to the state stream.

Afterwards, it should encrypt the following information using the user's public key and send it to the new user:
- Current keys & counts of the olm ratchets for state & message stream, from where the client would like the other client to see history from.
  Typically this will either be from the start of when the inviting client was able to see history or from the point which the TrustKey event was sent.
- An initial room state that should include:
  - Room name
  - Room photo
  - Trusted users

## Posting messages to a stream

When a client wishes to post a message to a stream, it should advance its olm ratchet by one.
The new current key of the olm ratchet will be used to encrypt the message being sent.
The EncryptedMessage should be equivalent to the client's current olm ratchet counter, in order to allow other clients to increment theirs as necessary.

## Receiving messages to a stream

Clients should decrypt incoming events using their olm ratchet; incrementing it according to the incrementing ID.

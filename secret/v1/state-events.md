# State Events

For most state events, the last state event of a type determines the current value, e.g. room name.
Some work differently.

## TrustKey

TrustKey is essentially an agglutinative event; every event adds a key to the list of keys that are trusted in a room. A TrustKey event is permanent; it cannot be revoked for a room. Clients should only trust events from users in the TrustedKey list, as well as the user that granted them the key to the room.

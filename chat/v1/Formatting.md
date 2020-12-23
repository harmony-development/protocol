# Formatting

Clients are expected to support the Commonmark specification for all message text, with the following changes:

- Images (`![]()`) should be explicitly ignored by the client
- `<@123456789>` should be interpreted as a user mention, and render as the specified user's display name, with other decoration optional
- `<@|123456789>` should be interpreted as a role mention, and render as the specified role's name and colour, with other decorations optional
- `<:123456789:>` should be interpreted as a local emoji, and render as the emoji's content as provided by the homeserver of the guild
- `<:123456789|homeserver.url:>` should be interpreted as a foreign emoji, and render as the emoji's content as provided by the homeserver specified in the emoji

# Harmony Protocol

This is the repository containing the Harmony protocol specification and documentation.

Read [here](FAQ.md) for frequently asked questions.

To read about protocol conventions, please read [GUIDELINES](GUIDELINES.md).

If you are looking for:

- a client to communicate with, check out [Tempest](https://github.com/harmony-development/tempest), [Challah](https://github.com/harmony-development/Challah) or [Crust](https://github.com/harmony-development/Crust).
- a server to host, check out [Scherzo](https://github.com/harmony-development/scherzo).
- an SDK to develop bots / clients / servers with, check out [Rust SDK](https://github.com/harmony-development/harmony_rust_sdk), [Web SDK](https://github.com/harmony-development/harmony-web-sdk) and [C++ SDK](https://github.com/harmony-development/Chometz).

# Stable v. Staging

Harmony has two types of protocol components: stable and staging.

Stable protocols are long-lived protocols that are not expected to be
replaced in the near future, and that have wide server and client support.

Staging protocols are protocols that may be replaced in the near future, and they
may not have wide server or client support.

Both stable and staging protocols follow semantic versioning: breaking changes
are only permitted with a corresponding increase in the major version, while
non-breaking changes result in an increase in the minor version.

Heavily in-dev protocols may be found on work branches in the Harmony repository,
and are subject to no compatibility guarantees. They must have at least 1 server and
2 client implementations before becoming staging protocols.

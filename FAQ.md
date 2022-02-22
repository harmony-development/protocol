---
title: Harmony FAQ
---

## What is Harmony?

A chat protocol which:
- is split into communities which contain channels which contain messages
- is defined entirely within the [Protobuf](https://developers.google.com/protocol-buffers) format, using a simple [RPC mechanism called hRPC](https://github.com/harmony-development/hrpc).

## What are the design goals of Harmony?

Harmony's protocol is designed to be as straightforward and pragmatic as possible. We do not make attempts at creating a "universal" design philosophy which the entire protocol is forced to follow, instead implementing things that make sense as a single cohesive system. Time has proven over time over that design idealism is often a limiting factor in services. 

## Why not Matrix?

- We believe that Matrix's design philosophy works in many cases, but fails in terms of performance and handling large communities.
- We believe that requiring all data to be transferred between servers for federation is inefficient.
- Harmony allows clients to connect directly to foreignservers if they wish for reduced latency and stress on their homeservers.
- Harmony wishes to represent many things that don't work well in Matrix's model of "everything is a JSON message appended to a stream called a room."
- Matrix is lacking in many functions we want from a chat service
- Having an independent protocol allows us to move faster to achieve our goals compared to working on Matrix.

## Why not Revolt?

Revolt is an interesting project, and we actually hadn't known about it while Harmony was under development. However, there are currently a few issues, namely:

- No federation. The developers have stated that federation was too "messy" for a big chat app. We disagree.
- It wants to be "Discord but open source", we want to be more than that.

## Where is E2EE?

E2EE is currently under development as part of our "Secret Manager"
It is currently being worked on in a branch of our [Protocol](https://github.com/harmony-development/protocol/tree/work/e2ee) repository and a library called [Lockdown](https://github.com/harmony-development/lockdown), which allows clients to implement E2EE trivially.

E2EE rooms adopt a different style from the rest of Harmony, allowing us to have all three by having a design specifically tailored to E2EE conversation: good performance, good security, good user experience.

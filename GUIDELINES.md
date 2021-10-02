---
title: Protocol Guidelines
---

# Guidelines

Protocol endpoints adhere to the following standard naming conventions:

- All actions being applied must be **prefixed**.

  **❌ What not to do:**

  - `rpc ProfileUpdate(ProfileUpdateRequest) returns (ProfileUpdateResponse) {}`

  **✅ What to do:**

  - `rpc UpdateProfile(UpdateProfileRequest) returns (UpdateProfileResponse) {}`

- Objects (types that aren't request or response types) should go at the top of the file.
- Objects should not be nested inside request or response types.
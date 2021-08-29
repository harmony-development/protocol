# Guidelines

Protocol endpoints adhere to the following standard naming conventions:

- All actions being applied must be **prefixed**.

  **❌ What not to do:**

  - `rpc ProfileUpdate(ProfileUpdateRequest) returns (ProfileUpdateResponse) {}`

  **✅ What to do:**

  - `rpc UpdateProfile(UpdateProfileRequest) returns (UpdateProfileResponse) {}`

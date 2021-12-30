---
title: RESTful Endpoints
---

## Authentication

If an endpoint expects authentication, the user should set the `Authorization`
header to a valid token from the hRPC API.

## POST `/_harmony/media/upload`

Expects authentication: yes

The body should be POST-ed as a multipart form (`multipart/form-data`), with a
single part named `file` which contains the body of the file being uploaded,
the name of the file and the MIME type of the file.

### Responses

#### 200 OK

The body will contain a JSON object in the following format:

| Field | Description                                |
| ----- | ------------------------------------------ |
| `id`  | The file ID of the file that was uploaded. |

#### 400 Bad Request

The body will contain a JSON object in the following format:

| Field     | Description        |
| --------- | ------------------ |
| `message` | The error message. |

Possible error messages and their meanings:

| Error            | Description                                       |
| ---------------- | ------------------------------------------------- |
| `missing-files`  | You forgot to upload a file.                      |
| `too-many-files` | You have uploaded more than one `file` form part. |

## GET `/_harmony/media/download/:file_id`

Expects authentication: no

The `file_id` should be one of the following:

- An attachment ID returned from POST `/_harmony/media/upload`
- A URI-encoded URL of an image
- A URI-encoded HMC URL.

### Responses

#### 200 OK

| Header                | Value                                                               |
| --------------------- | ------------------------------------------------------------------- |
| `Content-Type`        | The type of the file being downloaded.                              |
| `Content-Disposition` | `attachment; filename=<filename>` or `inline; filename=<filename>`. |

##### Body

The body will contain the content of the requested file.

## GET `/_harmony/about`

Expects authentication: no

### Responses

#### 200 OK

##### Body

The body will contain a JSON object in the following format:

| Field             | Description                                                       |
| ----------------- | ----------------------------------------------------------------- |
| `serverName`      | the name of the Harmony server software being hosted              |
| `version`         | the version of said Harmony server software                       |
| `aboutServer`     | A description of why / who this server is hosted for.             |
| `messageOfTheDay` | "message of the day", can be used to put maintenance information. |

Example response:

```json
{
    "serverName": "Scherzo",
    "version": "git-0c062f6",
    "aboutServer": "The main Harmony server.",
    "messageOfTheDay": "A maintenance will be done between 18:00 - 20:00."
}
```

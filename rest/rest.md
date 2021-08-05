# RESTful Endpoints

## Authentication

If an endpoint expects authentication, the user should set the `Authorization` header to a valid token from the gRPC API.

## POST `/_harmony/media/upload`

Expects authentication: yes

The body should be POST-ed as a multipart form (`multipart/form-data`), with a single part named `file` which contains the body of the file being uploaded.

### URL Parameters

- `filename`: The name of the file being uploaded.
- `contentType`: The MIME type of the file.

### Responses

#### 200 OK

Returns a JSON object with a single key, `id`.

#### 400 Bad Request

##### Body: `missing-files`

You forgot to upload a file.

##### Body: `too-many-files`

You have uploaded more than one `file` form part.

## GET `/_harmony/media/download/:file_id`

Expects authentication: no

The file ID should be one of the following:

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

The body will contain an UTF-8 string in the following format:
```
serverName version

aboutServer

messageOfTheDay
```
`serverName`: the Harmony server software being hosted
`version`: the version of said Harmony server software
`aboutServer`: A description of why / who this server is hosted for.
`messageOfTheDay`: "message of the day", can be used to put maintance information.
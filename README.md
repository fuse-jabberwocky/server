# Jabberwocky API Adapter Layer

This code generated server exposes the [REST API contract](https://github.com/fuse-jabberwocky/jabberwocky-api/blob/main/openapi/openapi.yaml) that is expected to be used by a UI for managing `Connector`s, `Channel`s and `Event Source/Sink`s.

The business logic will be provided by `Camel-K`.

## How to run

Install local libraries (first time you execute):

```
go get github.com/gorilla/mux
```

Launch the local server:

```
go get github.com/gorilla/mux
```

Test the application browsing the API:

```
$ http http://localhost:8080/v0/connector

[
    {
        "configuration": "some yaml configuration here!",
        "name": "my-connector-1",
        "type": "source"
    },
    {
        "configuration": "some yaml configuration here!",
        "name": "my-connector-2",
        "type": "sink"
    }
]

```

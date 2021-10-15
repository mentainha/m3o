# M3O CLI

The codebase for the M3O CLI

## Install

Building from source

```
go get github.com/m3o/m3o
```

Otherwise release binaries coming soon

## Usage

Export your API token as found in the UI

```
export M3O_API_TOKEN=xxxxxx
```

Call a service by name and endpoint

```
$ m3o call helloworld Call '{"name": "Alice"}'
{
 "message": "Hello Alice"
}
```

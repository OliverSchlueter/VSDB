# VSDB-Server

## How to build and run
1. Run build.bat
2. Execute bin/vsdb-server.exe

## Endpoints

### /get

#### Parameters:

``key``: the key to search with

#### Response:

``Status``: 'found' if the entry was found, 'not found' if the entry was not found

``Result``: the value of the entry if found, empty if not found

Example: ``http://localhost/get?key=hello``

```json
{
  "Status": "found",
  "Result": "world"
}
```

### /insert

#### Parameters:

``key``: the key of the entry

``value``: the value of the entry

#### Response:

``Status``: 'inserted' if successfully inserted

``Result``: the key of the inserted entry

Example: ``http://localhost/insert?key=user1&value=oliver``

```json
{
  "Status": "inserted",
  "Result": "user1"
}
```

### /delete

#### Parameters:

``key``: the key of the entry

#### Response:

``Status``: 'not found' if the entry was not found, 'deleted' if the entry was deleted

``Result``: the key of the inserted entry

Example: ``http://localhost/delete?key=hello``

```json
{
  "Status": "deleted",
  "Result": "hello"
}
```
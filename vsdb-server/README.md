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

### /getAllKeys

#### Response:

``Status``: 'success'

``Result``: all saved keys seperated by a semicolon

Example: ``http://localhost/getAllKeys``

```json
{
  "Status": "success",
  "Result": "hello;hello2;hello3"
}
```

### /getAllEntries

#### Response:

``Status``: 'success'

``Result``: all saved entries seperated by a semicolon, key and value are seperated by a colon.

Example: ``http://localhost/getAllEntries``

```json
{
  "Status": "success",
  "Result": "hello:world;hello2:world2;hello3:world3"
}
```
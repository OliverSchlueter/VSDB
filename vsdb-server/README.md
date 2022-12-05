# VSDB-Server

## How to build and run
1. Run build.bat
2. Execute ``bin/vsdb-server.exe``

## Command-line arguments

### -port

Runs the server on a certain port

Syntax: ``$ vsdb-server.exe -port [port number]``

Example: ``$ vsdb-server.exe -port 1337``

### -savePersistent

Enable to save the data persistent in filesystem

Syntax: ``$ vsdb-server.exe -savePersistent``

### -path

Sets the path where to save data if "savePersistent" is enabled

Syntax: ``$ vsdb-server.exe -path [path to a dictionary]``

Example: 

``$ vsdb-server.exe -savePersistent -path "vsdb_data/"``

The data file will be ``vsdb_data/data.json``

## Endpoints

### /get

#### Parameters:

``key``: the key to search with

#### Response:

``Status``: 'found' if the entry was found, 'not found' if the entry was not found

``Result``: the whole entry if found, empty if not found

Example: ``http://localhost/get?key=hello``

```json
{
  "Status": "found",
  "Result": {
    "hello": "world"
  }
}
```

### /insert

#### Parameters:

``key``: the key of the entry

``value``: the value of the entry

#### Response:

``Status``: 'inserted' if successfully inserted

``Result``: the inserted entry

Example: ``http://localhost/insert?key=user1&value=oliver``

```json
{
  "Status": "inserted",
  "Result": {
    "user1": "oliver"
  }
}
```

### /delete

#### Parameters:

``key``: the key of the entry

#### Response:

``Status``: 'not found' if the entry was not found, 'deleted' if the entry was deleted

``Result``: the key of the deleted entry

Example: ``http://localhost/delete?key=hello``

```json
{
  "Status": "deleted",
  "Result": {
    "hello": ""
  }
}
```

### /getAllKeys

#### Response:

``Status``: 'success'

``Result``: all saved keys

Example: ``http://localhost/getAllKeys``

```json
{
  "Status": "success",
  "Result": {
    "hello": "",
    "hello2": ""
  }
}
```

### /getAllEntries

#### Response:

``Status``: 'success'

``Result``: all saved entries

Example: ``http://localhost/getAllEntries``

```json
{
  "Status": "success",
  "Result": {
    "hello": "world",
    "hello2": "world2"
  }
}
```
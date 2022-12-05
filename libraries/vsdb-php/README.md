# VSDB PHP library

## Usage

Include the ``vsdb.php`` file and use its functions

### Get a value from key

Returns the value

```php
//vsdb_get($url, $key): string

$value = vsdb_get("http://localhost", "hello");
```

### Insert a new entry

Returns true if inserted, false if not

```php
//vsdb_insert($url, $key, $value): bool

$inserted = vsdb_get("http://localhost", "hello", "world");
```

### Delete an entry by key

Returns true if deleted, false if not

```php
//vsdb_delete($url, $key): bool

$deleted = vsdb_delete("http://localhost", "hello");
```

### Get all keys

Returns array of strings or false if something went wrong

```php
//vsdb_get_all_keys($url): array|bool

$keys = vsdb_get_all_keys("http://localhost");
```

### Get all entries

Returns associative array or false if something went wrong

```php
//vsdb_get_all_entries($url): array|bool

$entries = vsdb_get_all_entries("http://localhost");
```
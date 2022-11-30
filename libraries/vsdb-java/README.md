# VSDB Java library

## Usage

### Connect to database

```java
Database db = new Database("localhost", 80);
```

### Get a value from key

```java
String value = db.get("hello");
```

### Insert a new value

```java
boolean inserted = db.insert("hello", "world");
```

### Get all keys

```java
String[] keys = db.getAllKeys();
```

### Get all entries

```java
Map<String, String> entries = db.getAllEntries();
```
# VSDB Java library

## Usage

### Connect to database

```java
Database db = new Database("localhost", 80);
```

### Get a value from key

```java
String value = db.get("hello");

int intVal = db.getAsInt("hello");
long longVal = db.getAsLong("hello");
float floatVal = db.getAsFloat("hello");
double doubleVal = db.getAsDouble("hello");
```

### Insert a new value

```java
boolean inserted = db.insert("hello", "world");

boolean insertedInt = db.insert("hello", 123);
boolean insertedLong = db.insert("hello", 123l);
boolean insertedFloat = db.insert("hello", 123.45d);
boolean insertedDouble = db.insert("hello", 123.45d);
```

### Get all keys

```java
String[] keys = db.getAllKeys();
```

### Get all entries

```java
Map<String, String> entries = db.getAllEntries();
```
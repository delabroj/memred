memred is an in-memory, cli key-value store

# Installation

```
go build .
```

# Usage

## Basic Commands

```
> ./memred
SET test-var-name 100
GET test-var-name
100
UNSET test-var-name
GET test-var-name
Nil
SET test-var-name-1 50
SET test-var-name-2 50
NUMEQUALTO 50
2
SET test-var-name-2 10
NUMEQUALTO 50
1
END
```

## Basic Transaction

```
GET test-var-name
Nil
BEGIN
SET test-var-name 100
GET test-var-name
100
COMMIT
GET test-var-name
100
```

## Nested Transaction

```
GET test-var-name
Nil
BEGIN
SET test-var-name 100
GET test-var-name
100
BEGIN
SET test-var-name 120
GET test-var-name
120
BEGIN
SET test-var-name 150
GET test-var-name
150
ROLLBACK
COMMIT
GET test-var-name
120
```

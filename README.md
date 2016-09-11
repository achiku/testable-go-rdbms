## Testable RDBMS backed Go application

### Good Tests?

- Good tests are
    * right
    * independent from eash other
    * fast to complete
    * easy to write

##### Things I talk in 10 min

- independent/repeatable tests for RDBMS backed Go app
- fast tests for RDBMS backed Go app
- easy to write tests for RDBMS backed Go app

##### Things I don't talk

- how to write 'right' tests
- 'tests are necessary/unnecessary' kind of rant


### Problems

- 1. No standard way to setup/teardown database and transaction in Go test
    * Unlike other languages with WAF, it is rather unclear.
    * Django ORM/SQLAlchemy/ActiveRecord have standard way to do this.
    * How to handle transaction in tests is also unclear.
- 2. No standard way to easily create test data
    * Python has factory-boy, Ruby has factory-girl, but...


### Solutions

The following is just one way to tackle the problems previously stated. I really, really want to hear how other Gophers are doing.

- 1. No standard way to setup/teardown database and transaction in Go test
    - 1-1. Use `TestMain` to setup/teardown database(or schema)
    - 1-2. Use single transaction database driver to make test independent and repeatable
    - 1-3. Or mock database access?
- 2. No standard way to easily create test data
    - 2-1. Use patched version of `mergo` to create data
    - 2-2. Or something else?


### 1-1. TestMain to setup/teardown database(or schema)

- Inside `TestMain`
    - Create schema for test (we are using PostgreSQL, so using schema instead of database)
    - Create tables in that schema
    - Run tests
    - Drop cascade test schema


### 1-1. TestMain to setup/teardown database(or schema)

```
code...
```



### 1-1. TestMain to setup/teardown database(or schema)

- We are using SQLAlchemy + alembic to create database objects like tables, index, and constraints
- This combination can be also used for table migration management in production


### 1-2. Single transaction database driver to make test independent and repeatable

- Use single transaction database driver to make test independent and repeatable
    - What's that?
    - https://github.com/DATA-DOG/go-txdb (for MySQL)
    - https://github.com/achiku/pgtxdb (for PostgreSQL)

> Package txdb is a single transaction based database sql driver. When the connection is opened, it starts a transaction and all operations performed on this sql.DB will be within that transaction. If concurrent actions are performed, the lock is acquired and connection is always released the statements and rows are not holding the connection.


### 1-2. Single transaction database driver to make test independent and repeatable

```
code...
```


### 1-2. Single transaction database driver to make test independent and repeatable

- We are extensibly using PostgreSQL
    * I forked `go-txdb` and add some functionalities https://github.com/achiku/pgtxdb
    * PostgreSQL can execute `savepoint`
    * Using `pgtxdb`, developers can test target code with multiple transactions including rollback
- When `conn.Bigin()` is called, this library executes `SAVEPOINT pgtxdb_xxx;` instead of actually begins transaction.
- tx.Commit() does nothing.
- `ROLLBACK TO SAVEPOINT pgtxdb_xxx;` will be executed upon `tx.Rollback()` call so that it can emulate transaction rollback.
- Above features enable us to emulate multiple transactions in one test case.


### 2-1. Patched version of `mergo` to create data

- Use patched version of `mergo` to create data
    - https://github.com/achiku/mergo


### 2-1. Patched version of `mergo` to create data

```
code...
```

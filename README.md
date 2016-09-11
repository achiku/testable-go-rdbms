# testable-go-rdbms


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


### Problem

- 1. No standard way to setup/teardown database and transaction in Go test
    * Unlike other languages with WAF, it is rather unclear.
    * Django ORM/SQLAlchemy/ActiveRecord have standard way to do this.
    * How to handle transaction in tests is unclear.
- 2. No standard way to easily create test data
    * Python has factory-boy, Ruby has factory-girl, Go...


### Solution

The following is just one way to tackle the problems previously stated. I really, really want to hear how other Gophers are doing.

- 1. No standard way to setup/teardown database and transaction in Go test
    - 1-1. Use `TestMain` to setup/teardown database(or schema)
    - 1-2. Use single transaction database driver to make test independent and repeatable
    - 1-3. Or mock?
- 2. No standard way to easily create test data
    - 2-1. Use patched version of `mergo` to create data
    - 2-2. Or something else?


### Solution (1-1)

- Inside `TestMain`
    - Create schema for test
    - Create tables in that schema
    - Run tests
    - Drop cascade test schema


### Solution (1-1) details

```
code...
```



### Solution (1-1) advanced

- We are using SQLAlchemy + alembic to create database objects like tables, index, and constraints
- This combination can be also used for migration management in production


### Solution (1-2)

- Use single transaction database driver to make test independent and repeatable
    - what's that?
    - https://github.com/DATA-DOG/go-txdb (for MySQL)
    - https://github.com/achiku/pgtxdb (for PostgreSQL)

> Package txdb is a single transaction based database sql driver. When the connection is opened, it starts a transaction and all operations performed on this sql.DB will be within that transaction. If concurrent actions are performed, the lock is acquired and connection is always released the statements and rows are not holding the connection.


### Solution (1-2) details

```
code...
```


### Solution (1-2) advanced

- We are extensibly using PostgreSQL
    * PostgreSQL can execute `savepoint`
    * Using `pgtxdb`, developers can test target code with multiple transactions


### Solution (2-1)

- Use patched version of `mergo` to create data


### Solution (2-1) details

```
code...
```

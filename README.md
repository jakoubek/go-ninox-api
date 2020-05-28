# go-ninox-api

*go-ninox-api* is a Go package for interacting with the [Ninox API](https://ninox.com/).

## Ninox API documentation

https://ninox.com/en/manual/api/api-introduction

## Concepts

To access data in your Ninox account (reading and writing) you need to traverse this simple hierarchy:

- team
- database
- table

You can have several teams in your account (to group your databases in different topics). In every team there can be many databases. A database consists of one or more tables.

### API base URL

https://api.ninoxdb.de/v1

### List the teams in your account

```
GET https://api.ninoxdb.de/v1/teams
```

### List the databases in a team

```
GET https://api.ninoxdb.de/v1/teams/$TEAM_ID/databases/
```

### List all tables in a database

```
GET https://api.ninoxdb.de/v1/teams/$TEAM_ID/databases/$DATABASE_ID/tables
```

### Get all records in a table

```
GET https://api.ninoxdb.de/v1/teams/$TEAM_ID/databases/$DATABASE_ID/tables/$TABLE_ID/records
```

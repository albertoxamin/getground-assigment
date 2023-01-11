# GetGround End Year Party Service

## Development Setup

To get you started, first of all start your development db instance.

You can do it with `docker-compose up -d`, the go container is still commented out for now.

Once your db is up, you can start the app with `go run cmd/main.go`

## Project Overview

```
cmd
├── common
│   ├── db
│   └── models
└── routes
    ├── guest_list
    ├── guests
    ├── seats_empty
    └── tables
```

## Testing

> I implemented some tests to give an idea how testing could be expanded

If it's your first time running tests on this repo, you will need to create a new testing db on your container.
You can do that by running this command in your terminal
```sh
mysql -u root --password=password --protocol=TCP -e "create database test"
```

once your database is set up you will be able to run the tests from your ide.

![vscode example](assets/ide-test.png)

If you don't fancy ide help you can also execute the tests from the terminal, for example you can test the AddGuest logic here

```sh
go test -timeout 30s -run ^TestAddGuest$ github.com/getground/tech-tasks/backend/cmd/routes/guest_list
```
# GetGround End Year Party Service

## Development Setup

To get you started, first of all start your development db instance.

You can do it with `docker-compose up -d`, the go container is still commented out for now.

Once your db is up, you can start the app with `go run cmd/main.go`

## Project Overview

In the current implementation, the logic is contained the cmd directory, in which the `main.go` file resides.
The responsibilities of that file are to create the http server and register the different controllers that are implemented in the `routes` directory.

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

### Adding new endpoints

1. Create a new folder under routes and create a `controller.go` file
2. Create the implementation of the business logic in separate files in that folder
3. register your endpoint inside `main.go`

### Adding new Models

1. Create a new file under the models directory
2. Register your newly created model in `common/db/db.go`, in a similar format `db.AutoMigrate(&models.MyNewModel{})`

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
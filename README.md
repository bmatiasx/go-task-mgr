# The Space-X Challenge

Small RESTful API to support communications between management and engineering teams.

## Purpose
The API provided will take certain fields in JSON format and will create tasks in a
Trello board that was created for the sake of interacting between teams faster and 
clearer.

## Instructions
To run the application first open a CLI and then move to the `cmd` directory. There
type in the CLI the following command to run the application:
```
go run main.go
```
There are three task types that we can create with this API. Those are: 
- issue
- bug
- task

The following request and response examples will be shown in cURL format describing
what each use case will return.

### Create an issue
Request:
```
curl --location --request POST 'http://localhost:3000/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "type": "issue",
    "title": "No pilot mode",
    "description": "Enable no pilot mode before departure"
}'
```

Response:
```
{
    "board_id": "63bdd2e8fdf46c026cf9aff2",
    "description": "Enable no pilot mode before departure",
    "id": "63bf7f6c3ab717030125b62c",
    "list_id": "63bdd2e8fdf46c026cf9aff9",
    "message": "card created",
    "title": "No pilot mode",
    "type": "issue",
    "url": "https://trello.com/c/VMiZv94B/25-no-pilot-mode"
}
```

### Create a bug
Request:
```
{
    "type": "bug",
    "description": "Replace old buttons in dashboard"
}
```

Response:
```
{
    "board_id": "63bdd2e8fdf46c026cf9aff2",
    "description": "Replace old buttons in dashboard",
    "id": "63bf7eff993c6e02af87f0fb",
    "list_id": "63bdd2e8fdf46c026cf9affa",
    "message": "card created",
    "type": "bug",
    "url": "https://trello.com/c/VAGKkXnj/23-bug-critical-878"
}
```
### Create a task
For the case of task creation there are three types of categories that are valid.
Those are `Maintenance`,`Research` and `Test`.
If the category is not any of the mentioned it will return an error.

Request:
```
{
    "type": "task",
    "title": "Refill oil in engine to reduce friction",
    "category": "Maintenance"
}
```

Response:
```
{
    "board_id": "63bdd2e8fdf46c026cf9aff2",
    "category": "Maintenance",
    "id": "63bf7f3488350801c9608425",
    "list_id": "63bdd2e8fdf46c026cf9aff9",
    "message": "card created",
    "title": "Refill oil in engine to reduce friction",
    "type": "task",
    "url": "https://trello.com/c/K5blDMHF/24-refill-oil-in-engine-to-reduce-friction"
}
```

## Technical debt

- Unit tests
- Validate string fields to accept alphabet only characters

## Trello board access
Trello board can be checked in the link below after joining as a member:

`https://trello.com/invite/b/kK0RimXn/ATTI7adc4eb61e28a9b4e7dfc0a6c91568f1368B7D0A/sprint-001`



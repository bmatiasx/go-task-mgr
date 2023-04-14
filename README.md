# The Space-X Challenge

Small RESTful API to support communications between management and engineering teams.

## Purpose
The API provided will take certain fields in JSON format and will create tasks in a
Trello board that was created for the sake of interacting between teams faster and 
clearer.

## Instructions
Since this application is containerized Docker is needed to run it. If you don't have it installed you can check the 
following [guide](https://docs.docker.com/engine/install/) to install Docker in your computer.

Once installed go to the root directory and execute the command:

> docker-compose up

Now the application is running on port 9090.

You can check it by sending a request to the welcome API which is at `<host>:9090/api/v1/welcome`
The response should look like:

```
{
    "message": "Welcome to the Card Service" 
}
```

## Cards
There are three card types that we can create with this API. Those are: 
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
curl --location --request POST 'http://localhost:3000/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "type": "bug",
    "description": "Replace old buttons in dashboard"
}'
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
curl --location --request POST 'http://localhost:3000/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "type": "task",
    "title": "Refill oil in engine to reduce friction",
    "category": "Maintenance"
}'
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



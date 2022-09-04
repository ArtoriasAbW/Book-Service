# Book Service

## Data example

| Id  | Book Title                    | Book Author      | Unread |
| ---- | ---- | ----------- | -----------------------------------|
| 1   | History of Western Philosophy | Bertrand Russell | true   |
| 2   | Martin Eden                   | Jack London      | false  |

## Telegram Bot
### Command list

- **/help** - list commands.
- **/list** - list all books.
- **/add "title"|"author"|"status"** - add new book.
- **/delete "id"** - remove book with this id.
- **/update "id"|"title"|"author"|"status"** - edit book with this id.

## REST
[Swagger documentation](swagger)

## GRPC
[Proto files](api)

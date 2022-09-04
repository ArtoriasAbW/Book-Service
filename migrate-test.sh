#!/bin/sh

export MIGRATION_DIR="./migrations"
export DB_DSN="host=localhost port=5433 user=test_user password=test_password dbname=book_service sslmode=disable"

if [ "$1" = "--dryrun" ]; then
    $(LOCAL_BIN)/goose -v -dir ${MIGRATION_DIR} postgres "${DB_DSN}" status
else
    $(LOCAL_BIN)/goose -v -dir ${MIGRATION_DIR} postgres "${DB_DSN}" up
fi
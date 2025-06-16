#!/bin/sh
ls schemas
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
migrate -database "$POSTGRES_URL" -path schemas up
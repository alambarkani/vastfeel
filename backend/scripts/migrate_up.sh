#!/bin/bash

migrate -database "postgres://postgres:XsJTRf431@localhost:5432/vastfeel?sslmode=disable" -path F:/www/vastfeel/backend/internal/database/migrations up
sleep 3
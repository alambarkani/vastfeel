#!/bin/bash

migrate -database "postgres://postgres:XsJTRf431@localhost:5432/vastfeel?sslmode=disable" -path internal/database/migrations force 20240923155801
sleep 3
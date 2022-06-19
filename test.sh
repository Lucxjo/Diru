#!/bin/sh
GOOGLE_APPLICATION_CREDENTIALS=../config/gcloud.json go test ./google ./deepl -race -covermode=atomic -coverprofile=cover.out
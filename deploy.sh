#!/usr/bin/env bash

set -e

go run main.go

if [ ! -d "./dist" ]; then
    echo "Dist directory doesn't exist, exiting..."
    exit 1
fi

gsutil cp -r dist/web/* gs://www.midday.coffee
gsutil iam ch allUsers:objectViewer gs://www.midday.coffee
gsutil -m setmeta -h "Content-Type: image/webp" gs://www.midday.coffee/webp-demo/*.webp
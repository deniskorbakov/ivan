#!/bin/sh
set -e
rm -rf manpages
mkdir manpages
go run ./cmd/ivan man | gzip -c -9 >manpages/ivan.1.gz
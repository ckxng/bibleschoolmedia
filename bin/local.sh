#!/bin/bash

source config/google-local.txt.default

ls config/google-local.txt >/dev/null 2>&1 && \
source config/google-local.txt

exec goapp serve --host=$HOST --port=$PORT

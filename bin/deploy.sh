#!/bin/bash
source config/google-deploy.txt

if [ "$APPID" == "sample" ]; then
  echo "Please configure APPID in config/google-deploy.txt"
  exit 1
fi

exec goapp deploy -application $APPID .

#!/bin/bash

getdep () {
  UNDEFINED="(undefined)"
  PDIR=${1-$UNDEFINED}
  REPO=${2-$UNDEFINED}

  if [ "$PDIR" == "$UNDEFINED" ]; then
    echo "Cannot create repo dir $PDIR"
    return 1
  fi

  if [ "$REPO" == "$UNDEFINED" ]; then
    echo "Cannot fetch from repo $REPO"
    return 1
  fi

  REPO="$PDIR/$REPO"

  ls -d $REPO || ( \
  mkdir -p $PDIR && \
  pushd $PDIR && \
  git clone https://$REPO
  popd )
}

getdep "github.com/gorilla" "context"
getdep "github.com/gorilla" "mux"
getdep "github.com/codegangsta" "negroni"


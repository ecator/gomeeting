#!/bin/bash

wk=`dirname "$0"`
. "$wk/check.sh"

case $1 in
add)
curl -X POST --cookie "auth=$GOMEETING_TOKEN" -d "$2" $GOMEETING_HOST/api/org
;;
del)
curl -X DELETE --cookie "auth=$GOMEETING_TOKEN"  $GOMEETING_HOST/api/org/$2
;;
mod)
curl -X PUT --cookie "auth=$GOMEETING_TOKEN" -d "$3" $GOMEETING_HOST/api/org/$2
;;
get)
curl -X GET --cookie "auth=$GOMEETING_TOKEN" $GOMEETING_HOST/api/org/$2
;;
*)
  echo "The operation can be only one of add del mod get!"
  exit 2
;;
esac
#!/usr/bin/env bash

json=`curl -k -s https://Administrator:2FederateM0re@localhost:9000/pa-admin-api/v3/api-docs/pa/api-docs.json`
for row in $(echo "${json}" | jq '.apis[]| .path'); do
  row="${row%\"}"
  row="${row#\"}"
  row="${row#\/}" 
  # echo "https://Administrator:2FederateM0re@localhost:9000/pa-admin-api/v3/api-docs/pa/${row} models/${row}.json"
  curl -k -s "https://Administrator:2FederateM0re@localhost:9000/pa-admin-api/v3/api-docs/pa/${row}" | jq '.' > "models/${row}.json"
done
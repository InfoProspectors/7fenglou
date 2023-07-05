#!/bin/bash

url="$1"
if [[ -z "$url" ]]; then
  echo "URL is missing. Usage: sh check_data.sh <url>"
  exit 1
fi

file_count=$(ls -l data | grep '^-' | wc -l)
echo "File count: $file_count"

response=$(curl -s "$url")
if [ $? -eq 0 ]; then
  total_page=$(echo "$response" | jq -r '.Data.totalPage')
  echo "API request successful"
  echo "Total pages: $total_page"

  if [ "$file_count" -ne "$total_page" ]; then
    echo "Mismatch: File count does not match total pages"
    echo "Removing log and data directories..."
    rm -rf log
    rm -rf data
  else
    echo "File count matches total pages"
  fi
else
  echo "API request failed, skipping further steps."
fi

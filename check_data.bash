#!/bin/bash

url="$1"
if [[ -z "$url" ]]; then
  echo "URL is missing. Usage: sh check_data.sh <url>"
  exit 1
fi

# Count the number of files in the 'data' directory
file_count=$(ls -l data | grep '^-' | wc -l)
echo "File count: $file_count"

# Make API request to the provided URL and save the response to 'response.json'
echo "Making API request..."
response=$(curl -s "$url" | jq -Rs '.')
echo "$response" > response.json

# Extract the 'totalPage' field from the response and store it in 'total_page' variable
total_page=$(jq -r '.Data.totalPage // empty' response.json)
if [[ -z "$total_page" ]]; then
  echo "Total pages field not found or is null"
  exit 1
fi
echo "Total pages: $total_page"

# Compare the file count with the total pages
if [ "$file_count" -ne "$total_page" ]; then
  echo "Mismatch: File count does not match total pages"
  echo "Removing log and data directories..."
   rm -rf log
   rm -rf data
else
  echo "File count matches total pages"
fi

# Read the 'totalPage' field again from the file
total_page=$(jq -r '.Data.totalPage' response.json | select(. != null))
if [[ -z "$total_page" ]]; then
  echo "Total pages field not found or is null"
  exit 1
fi
echo "Total pages: $total_page"

# Clean up the response file
rm response.json
echo "Response file removed"

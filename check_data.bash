#!/bin/bash

url="$1"
if [[ -z "$url" ]]; then
  echo "URL is missing. Usage: bash check_data.sh <url>"
  exit 1
fi

# Count the number of files in the 'data' directory
file_count=$(ls -l data | grep '^-' | wc -l)
echo "File count: $file_count"

# Make API request to the provided URL and save the response to 'response.json'
echo "Making API request..."
response=$(curl -s "$url" | jq -Rs '.')
echo "$response" > response.json

# Check if the response is a valid JSON object
jq -e . >/dev/null response.json
if [ $? -ne 0 ]; then
  echo "Invalid JSON response"

  # Download the JSON file from the specified URL and save it as response.json
  echo "Downloading response.json from URL..."
  curl -s -o response.json "https://raw.githubusercontent.com/InfoProspectors/7fenglou/main/data/7fenglou1.json"

  # Check if the download was successful
  if [ $? -ne 0 ]; then
    echo "Failed to download the JSON file"
    exit 1
  fi

  echo "response.json downloaded successfully"
fi


# Extract the 'totalPage' field from the response and store it in 'total_page' variable
# Read the 'totalPage' field again from the file
total_page=$(jq -r '.Data.totalPage' response.json | select(. != null))
if [[ -z "$total_page" ]]; then
  echo "Total pages field not found or is null"
  exit 1
fi
echo "Total pages: $total_page"

# Compare the file count with the total pages
if [ "$file_count" -ne "$total_page" ]; then
  echo "Mismatch: File count does not match total pages"
  echo "Removing log and data directories..."
#   rm -rf log
#   rm -rf data
else
  echo "File count matches total pages"
fi



# Clean up the response file
#rm response.json
#echo "Response file removed"

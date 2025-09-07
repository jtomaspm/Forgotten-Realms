#!/bin/bash

if [ -z "$1" ]; then
  echo "Usage: $0 [folder] <filename>"
  exit 1
fi

# Set folder and filename based on the number of arguments
if [ -n "$2" ]; then
  folder="./documentation/$1"
  filename="$2"
else
  folder="./documentation"
  filename="$1"
fi

capitalized_filename="$(echo "${filename:0:1}" | tr '[:lower:]' '[:upper:]')${filename:1}"

mkdir -p "$folder"

cat <<EOF > "../$folder/$filename.md"
# $capitalized_filename

## Full Documentation

* [README](https://github.com/jtomaspm/SimplifiedCrafter/blob/main/README.md)
EOF

echo "Docs for '$capitalized_filename' have been created successfully in '$folder'."
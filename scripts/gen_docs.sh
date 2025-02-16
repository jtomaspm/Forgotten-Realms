#!/bin/bash

if [ -z "$1" ]; then
  echo "Usage: $0 <filename>"
  exit 1
fi

capitalized_argument="$(echo "${1:0:1}" | tr '[:lower:]' '[:upper:]')${1:1}"

cat <<EOF > "../documentation/$1.md"
# $capitalized_argument

## Full Documentation

* [README](https://github.com/jtomaspm/SimplifiedCrafter/blob/main/README.md)
EOF

echo "Docs for '$capitalized_argument' has been created successfully."
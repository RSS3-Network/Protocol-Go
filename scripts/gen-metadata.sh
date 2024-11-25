#!/bin/bash

# create Metadata.yaml if not exists
if [ ! -f "openapi/Metadata.yaml" ]; then
  touch "openapi/Metadata.yaml"
fi
echo "type: object" > "openapi/Metadata.yaml"
echo "oneOf:" >> "openapi/Metadata.yaml"


find ./schema/metadata -name '*.go' | xargs grep -E 'var _ Metadata = \(\*(.*)\)\(nil\)' | sed -E 's/.*\(\*(.*)\)\(nil\)/\1/' | sort -u | while read -r type; do


  # if type contains action
  if [[ $type == *"Action"* ]]; then
    echo "  - \$ref: \"./enum/$type.yaml\"" >> "openapi/Metadata.yaml"
    continue
  else
    # create file if not exists
    # init file with type: object
    if [ ! -f "openapi/metadata/$type.yaml" ]; then
      echo "create openapi/metadata/$type.yaml"
      touch "openapi/metadata/$type.yaml"
      echo "type: object" > "openapi/metadata/$type.yaml"
    fi
    echo "  - \$ref: \"./metadata/$type.yaml\"" >> "openapi/Metadata.yaml"
  fi
done
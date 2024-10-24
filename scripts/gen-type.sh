#!/bin/bash

file=openapi/Type.yaml

# create Metadata.yaml if not exists
if [ ! -f "$file" ]; then
  touch "$file"
fi
echo "type: string" > "$file"
echo "oneOf:" >> "$file"

find ./openapi/type -name '*.yaml' | cut -d '/' -f 4 | while read -r type; do
  echo "  - \$ref: \"./type/$type\"" >> "openapi/Type.yaml"
done
#!/bin/bash
file="schema/type.yaml"
enums=$(yq ea '.enum[]' schema/typex/*.yaml | sort -u)

echo "type: string" > $file

echo "enum:" >> $file

for enum in $enums; do
  if [ "$enum" != "---" ]; then
    echo " - $enum" >> $file
  fi
done
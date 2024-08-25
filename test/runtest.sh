#!/bin/bash

cd ..
go build ./cmd/dec
cd test

for file in test.*; do
  if [[ "$file" != test.*.done && $file != test.*.out ]]; then
    ../dec "$file" > "$file.done"
    diff "$file.done" "$file.out" > /dev/null
    if [[ $? -ne 0 ]]; then
      echo "Error: $file.done does not match $file.out"
    else
        rm "$file.done"
    fi
  fi
done



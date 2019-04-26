#!/bin/bash

for md in 2019*.md
do
    d=$(echo ${md} | cut -f1 -d'-')
    ts="$(date -d${d} +%F) 00:00:00"
    hash=$(md5sum ${md} | cut -f1 -d' ')
    title=$(cat ${md} | grep -E '^# ' | sed -E 's/# //')
    body=$(cat ${md} | sed -E 's/"/\"/g')
    echo "insert into post (filename, timestamp, hash, title, body) values (\"${md}\", \"${ts}\", \"${hash}\", \"${title}\", \"${body}\");"
done

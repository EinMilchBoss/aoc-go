#!/bin/bash

if [ "$#" -ne 1 ]; then
    echo "No relative path provided."
    exit
fi

path=$1

go build -C $path -o app.exe .
$path/app.exe
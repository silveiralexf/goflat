#!/bin/bash

if command -v task > /dev/null; then
    exec echo "task is installed"
    exit 0
else
    echo 'task not found. brew install go-task to have it installed' 1>&2
    exit 1
fi

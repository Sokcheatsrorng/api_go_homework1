#!/bin/bash

message="$1"
curr_branch=$(git branch --show-current)

# Check if the commit message is empty
if [ -z "$message" ]; then
    echo "Please provide a commit message"
    exit 1
fi

# Stage all changes
git add .

# Commit with the provided message
git commit -m "$message"

# Push to the current branch
git push origin "$curr_branch"

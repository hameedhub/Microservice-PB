#!/bin/zsh

directories=("default" "high" "medium" "low" )
for dir in "${directories[@]}"; do
    if [ -d "$dir" ]; then
        cd "$dir"
       if [ -f "main.go" ]; then
            go run main.go &
        else
            echo "main.go not found in $dir"
        fi
        cd ..
    else
        echo "$dir not found"
    fi
done

wait

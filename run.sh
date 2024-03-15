#!/bin/zsh

directories=("default" "high" "medium" "low" )
run() {
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
}

for i in {1..100} ; do
    run "$i"
    sleep 20
done

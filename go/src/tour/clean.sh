ls | grep -v "$(ls *.go *.sh)" | xargs rm 2> /dev/null

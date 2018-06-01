find ~/workspace/project -type f -name "*.rb" -not -iwholename '*.git*' -print0 | xargs -0 grep -oh -E 'https?:\/\/github.com\/.*(pull|issues)\/\d+' | go run main.go


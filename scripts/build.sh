SCRIPT_CONTENT=$(curl -fsSL https://gist.github.com/xkrishguptaa/3c633f4c75f2b319e18880f3b68b1d23/raw/7b82e4f0209b39055565593735bda4ea89e43c61/go-application-builder.bash)

export APP_NAME="go-todo-api"
export GO_BUILD_TAGS="jsoniter"

echo "$SCRIPT_CONTENT" | sh

cwd=$(pwd)

dir=$(dirname "${BASH_SOURCE[0]}")

# Change to script directory

cd $dir

rm -rf ../bin

# Set default values

cross=false
v=""
app="app"
go_build_tags="jsoniter"

if [ -n "$VERSION" ]; then
    v="_$VERSION"
fi

# Check if environment variable is set
if [[ -n "$CROSS_BUILD" ]]; then
    cross=true
fi

if [[ -n "$APP_NAME" ]]; then
    app=$APP_NAME
fi

# Build

if [[ "$cross" = true ]]; then
    mkdir -p ./bin

    platforms=("linux/amd64" "linux/arm64" "darwin/amd64" "darwin/arm64" "windows/amd64" "windows/arm64")

    for platform in "${platforms[@]}"; do
      IFS='/' read -ra p <<< "$platform"

      os=${p[0]}
      arch=${p[1]}

      echo "\033[1mBuilding $app$v for $os $arch\033[0m"
      GOOS=$os GOARCH=$arch go build -tags "$go_build_tags" -o ../bin/$app-$os-$arch$v ../
    done

else
    mkdir -p ./bin
    go build -tags "$go_build_tags" -o ../bin/$app$v ../
fi

cd $cwd

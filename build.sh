package="${PROJECT_NAME:-app}"
outdir="${OUT_DIR:-bin}"
go_tags="${GO_TAGS:-jsoniter}"

echo "info: package: $package"
echo "info: outdir: $outdir"
echo "info: go tags: $go_tags"
echo "info: multi platform: $MULTI_PLATFORM"
echo "............................................"
echo "info: cleaning up $outdir"

rm -rf $outdir

echo "success: cleaned up $outdir"
echo "............................................"
echo "info: creating $outdir"

mkdir -p $outdir

echo "success: created $outdir"
echo "............................................"

platforms=(
"android/arm"
"darwin/386"
"darwin/amd64"
"darwin/arm"
"darwin/arm64"
"dragonfly/amd64"
"freebsd/386"
"freebsd/amd64"
"freebsd/arm"
"linux/386"
"linux/amd64"
"linux/arm"
"linux/arm64"
"linux/ppc64"
"linux/ppc64le"
"linux/mips"
"linux/mipsle"
"linux/mips64"
"linux/mips64le"
"netbsd/386"
"netbsd/amd64"
"netbsd/arm"
"openbsd/386"
"openbsd/amd64"
"openbsd/arm"
"plan9/386"
"plan9/amd64"
"solaris/amd64"
"windows/386"
"windows/amd64"
)

echo "info: fetching dependencies"
go get -d -v ./...
if [[ $? -ne 0 ]]
then
    echo "error: failed to fetch dependencies"
    echo "............................................"
    exit 1
fi
echo "success: fetched dependencies"
echo "............................................"

if [[ "$MULTI_PLATFORM" = "true" ]]
then
    echo "info: platforms: ${platforms[@]}"
    echo "............................................"
    echo "info: enabling cgo"
    export CGO_ENABLED=1
    echo "success: enabled cgo"
    echo "............................................"

    for platform in "${platforms[@]}"
    do
        echo "info: building for $platform"
        platform_split=(${platform//\// })
        GOOS=${platform_split[0]}
        GOARCH=${platform_split[1]}

        echo "info: GOOS: $GOOS"
        echo "info: GOARCH: $GOARCH"

        output_name=$package'-'$GOOS'-'$GOARCH

        if [[ $GOOS = "windows" ]]
        then
            output_name+='.exe'
        fi

        echo "info: output name: $outdir/$output_name"
        echo "info: building"
        echo "info: STDOUT----------------------------------"
        go build -tags=$tags -o bin/$output_name ./...
        echo "info: STDOUT----------------------------------"
        echo "............................................"
    done

    exit 0
else
    echo "info: building for current platform"
    echo "info: GOOS: $GOOS"
    echo "info: GOARCH: $GOARCH"

    output_name=$package

    if [[ $GOOS = "windows" ]]
    then
        output_name+='.exe'
    fi

    echo "info: output name: $outdir/$output_name"

    echo "info: building"
    echo "info: STDOUT----------------------------------"
    go build -tags=$tags -o bin/$output_name ./...
    if [ $? -ne 0 ]
    then
        echo "error: failed to build"
        echo "............................................"
        exit 1
    fi
    echo "info: STDOUT----------------------------------"

    echo "success: built binary: $outdir/$output_name"
    echo "............................................"
fi

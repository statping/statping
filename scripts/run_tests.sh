#!/bin/sh

# set -eux

test_env=$2
if [[ "${test_env}" = "drone" ]]; then
    echo "Setting up code"
    ORIG_DIR=/github/workspace/
    SRC_DIR=/go/src/github.com/razorpay/statping
    mkdir -p ${SRC_DIR}
    cp -Rp ${ORIG_DIR} ${SRC_DIR}
    cd ${SRC_DIR}
    cp -r workspace/* .
fi

if [[ "$1" = "fmt" ]]; then
    echo "Running go fmt"
    files=$(gofmt -l $(find . -type f -name '*.go' -not -path "./vendor/*") 2>&1)
    if [[ "$files" ]]; then
        echo "These files did not pass the gofmt check:"
        echo ${files}
        exit 1
    fi
fi

if [[ "$1" = "test" ]]; then
    echo "Installing dependencies"
    apk add --no-cache git gcc musl-dev
    export GO111MODULE="on"
    echo 'exec echo ${GIT_TOKEN}' > /tmp/askpass.sh
    chmod +x /tmp/askpass.sh
    export GIT_ASKPASS=/tmp/askpass.sh
    go mod vendor
    go version
    echo "Running mysql migrations"
    go run cmd/migration/mysql/main.go -env=drone up
    echo "Running postgres migrations"
    go run cmd/migration/postgres/main.go -env=drone up
    cp configs/drone.toml configs/test.toml

    echo "Running tests ${DRONE_BRANCH}"

    echo "Running Unit tests with coverage Test"
    go generate ./...
        #Interate all the go packages after listing it
    list=$(go list ./...)
    i=1
    #Run the go test for each package and generate a cov with the package name
    for pkg in $list
       do
            go test -coverprofile=pkg-$i.cover.out -coverpkg=./... -covermode=atomic $pkg
            x=$?
            i=$((i+1))
            if [[ $x -ne 0 ]]; then
                echo "Unit tests failed"
                exit $x
            fi
       done

    echo "mode: set" > sonarqube.cov && cat *.cover.out | grep -v mode: | sort -r | \
        # Merge all the cov file and generate sonaqube.cov files
    awk '{if($1 != last) {print $0;last=$1}}' >> sonarqube.cov
        #Renaming the file with the drone_build_number to identify uniquely in drone
    cp sonarqube.cov /github/workspace/sonarqube.cov
    exit $?
fi

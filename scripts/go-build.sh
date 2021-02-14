#!/bin/bash
# https://github.com/complexorganizations/go-build-script

# Detect Operating System
function dist-check() {
    if [ -e /etc/os-release ]; then
        # shellcheck disable=SC1091
        source /etc/os-release
        DISTRO=$ID
    fi
}

# Check Operating System
dist-check

# Pre-Checks system requirements
function installing-system-requirements() {
    if { [ "$DISTRO" == "ubuntu" ] || [ "$DISTRO" == "debian" ] || [ "$DISTRO" == "raspbian" ] || [ "$DISTRO" == "pop" ] || [ "$DISTRO" == "kali" ] || [ "$DISTRO" == "linuxmint" ] || [ "$DISTRO" == "fedora" ] || [ "$DISTRO" == "centos" ] || [ "$DISTRO" == "rhel" ] || [ "$DISTRO" == "arch" ] || [ "$DISTRO" == "manjaro" ] || [ "$DISTRO" == "alpine" ] || [ "$DISTRO" == "freebsd" ]; }; then
        if [ ! -x "$(command -v sha1sum)" ]; then
            if { [ "$DISTRO" == "ubuntu" ] || [ "$DISTRO" == "debian" ] || [ "$DISTRO" == "raspbian" ] || [ "$DISTRO" == "pop" ] || [ "$DISTRO" == "kali" ] || [ "$DISTRO" == "linuxmint" ]; }; then
                sudo apt-get update && apt-get install coreutils -y
            elif { [ "$DISTRO" == "fedora" ] || [ "$DISTRO" == "centos" ] || [ "$DISTRO" == "rhel" ]; }; then
                sudo yum update -y && yum install coreutils -y
            elif { [ "$DISTRO" == "arch" ] || [ "$DISTRO" == "manjaro" ]; }; then
                sudo pacman -Syu --noconfirm iptables coreutils
            elif [ "$DISTRO" == "alpine" ]; then
                sudo apk update && apk add coreutils
            elif [ "$DISTRO" == "freebsd" ]; then
                sudo pkg update && pkg install coreutils
            fi
        fi
    else
        echo "Error: $DISTRO not supported."
        exit
    fi
}

# Run the function and check for requirements
installing-system-requirements

# Build for all the OS
function build-golang-app() {
    APPLICATION="flutter-manager"
    VERSION="1.0.1"
    if [ -x "$(command -v go)" ]; then
        # Darwin
        GOOS=darwin GOARCH=amd64 go build -o build/$APPLICATION-$VERSION-darwin-amd64 .
        GOOS=darwin GOARCH=arm64 go build -o build/$APPLICATION-$VERSION-darwin-arm64 .
        # Linux
        GOOS=linux GOARCH=386 go build -o build/$APPLICATION-$VERSION-linux-386 .
        GOOS=linux GOARCH=amd64 go build -o build/$APPLICATION-$VERSION-linux-amd64 .
        GOOS=linux GOARCH=arm go build -o build/$APPLICATION-$VERSION-linux-arm .
        GOOS=linux GOARCH=arm64 go build -o build/$APPLICATION-$VERSION-linux-arm64 .
        GOOS=linux GOARCH=mips go build -o build/$APPLICATION-$VERSION-linux-mips .
        GOOS=linux GOARCH=mips64 go build -o build/$APPLICATION-$VERSION-linux-mips64 .
        GOOS=linux GOARCH=mips64le go build -o build/$APPLICATION-$VERSION-linux-mips64le .
        GOOS=linux GOARCH=mipsle go build -o build/$APPLICATION-$VERSION-linux-mipsle .
        GOOS=linux GOARCH=ppc64 go build -o build/$APPLICATION-$VERSION-linux-ppc64 .
        GOOS=linux GOARCH=ppc64le go build -o build/$APPLICATION-$VERSION-linux-ppc64le .
        GOOS=linux GOARCH=riscv64 go build -o build/$APPLICATION-$VERSION-linux-riscv64 .
        GOOS=linux GOARCH=s390x go build -o build/$APPLICATION-$VERSION-linux-s390x .
        # Windows
        GOOS=windows GOARCH=386 go build -o build/$APPLICATION-$VERSION-windows-386.exe .
        GOOS=windows GOARCH=amd64 go build -o build/$APPLICATION-$VERSION-windows-amd64.exe .
        GOOS=windows GOARCH=arm go build -o build/$APPLICATION-$VERSION-windows-arm.exe .
        # Get SHA-1 and put everything in a register.
        find build/ -type f -print0 | xargs -0 sha1sum
    else
        echo "Error: In your system, Go wasn't found."
        exit
    fi
}

# Start the build
build-golang-app

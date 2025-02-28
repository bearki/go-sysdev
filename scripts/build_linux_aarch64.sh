#!/bin/bash

# 遇到错误立即停止
set -e

# 获取项目根路径
ProjectRoot=$(realpath "$(dirname "$0")/..")
# 打印工作路径
echo "ProjectRoot: ${ProjectRoot}"

# 配置环境
Toolchain="$HOME/build-tools/gcc-arm-8.3-2019.03-x86_64-aarch64-linux-gnu"
export GOOS="linux"
export GOARCH="arm64"
export CGO_ENABLED="1"
export PKG_CONFIG_PATH="${ProjectRoot}/lib/pkgconfig/linux_aarch64"
export AR="${Toolchain}/bin/aarch64-linux-gnu-ar"
export CC="${Toolchain}/bin/aarch64-linux-gnu-gcc"
export CXX="${Toolchain}/bin/aarch64-linux-gnu-g++"

# pkg-config检查
pkg-config --cflags --libs sysdev

# 执行
go clean -cache
go build -v -ldflags="-extldflags=-static -s -w" "${ProjectRoot}/cmd/sysdev"

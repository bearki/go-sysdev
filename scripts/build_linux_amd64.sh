#!/bin/bash

# 遇到错误立即停止
set -e

# 获取项目根路径
ProjectRoot=$(realpath "$(dirname "$0")/..")
# 打印工作路径
echo "ProjectRoot: ${ProjectRoot}"

# 配置环境
Toolchain="$HOME/build-tools/x86-64--glibc--stable-2021.11-5"
export GOOS="linux"
export GOARCH="amd64"
export CGO_ENABLED="1"
export PKG_CONFIG_PATH="${ProjectRoot}/lib/pkgconfig/linux_x86_64"
# export AR="${Toolchain}/bin/x86_64-buildroot-linux-gnu-ar"
# export CC="${Toolchain}/bin/x86_64-buildroot-linux-gnu-gcc"
# export CXX="${Toolchain}/bin/x86_64-buildroot-linux-gnu-g++"

# pkg-config检查
pkg-config --cflags --libs sysdev

# 执行
go clean -cache
go run "${ProjectRoot}/cmd/sysdev"

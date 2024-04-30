# 注意文件格式，编码必须为UTF8-BOM
# 定义输出编码（对[Console]::WriteLine生效）
$OutputEncoding = [console]::InputEncoding = [console]::OutputEncoding = [Text.UTF8Encoding]::UTF8
# 遇到错误立即停止
$ErrorActionPreference = 'Stop'
# 获取项目根路径
$ProjectRoot = Resolve-Path -Path "${PSScriptRoot}\.."
# 打印工作路径
[Console]::WriteLine("ProjectRoot: ${ProjectRoot}")
# 配置环境
$Env:GOOS = "windows"
$Env:GOARCH = "amd64"
$Env:CGO_ENABLED = "1"
$Env:PATH = "${Env:MinGW64}\bin;${Env:PATH}"
$Env:PKG_CONFIG_PATH = "${ProjectRoot}\lib\pkgconfig\windwos_x86_64"

# 执行
go build "${ProjectRoot}\cmd\sysdev"
#!/bin/bash
# set -x
set -e

echo -e "==> start xcopy ..."

# cp_dst 要复制的目标
cp_dst="$1"

# 如果没有目标参数，打印提示并退出
if [ -z "$cp_dst" ]; then
	echo -e "❗ 错误，缺少目标参数 \n格式: bash_cmd copy_dst"
	exit
else
	echo "复制目标 $cp_dst"
fi

# 当前目录路径
pwdx=$(
	cd "$(dirname "$0")"
	pwd
)

# 当前项目路径 pro
pro=$(
	cd "$pwdx/../.."
	pwd
)

# 工具目录 toolx
toolx=$pro/tools/hooks

# 用 tox 替换 fromx
tox=vuca
fromx=fuwensun

# dst 目标，src 源头
dst=$pro/$cp_dst
src=${dst/$tox/$fromx}

# 执行 copy
rm -rf "$dst"
cp -r "$src" "$dst"

"$toolx"/xcheck.sh

echo -e "==< end xcopy"

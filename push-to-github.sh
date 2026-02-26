#!/bin/bash
# 将 gangguanbao 提交到 GitHub: https://github.com/devuser/gangguanbao
# 在项目根目录执行：bash push-to-github.sh

set -e
cd "$(dirname "$0")"

echo "当前目录: $(pwd)"

# 若 .git 不存在则初始化
if [ ! -d ".git" ]; then
  git init
  echo "已初始化 git 仓库"
fi

# 设置远程
git remote remove origin 2>/dev/null || true
git remote add origin https://github.com/devuser/gangguanbao.git

# 添加并提交
git add .
if git diff --cached --quiet 2>/dev/null; then
  echo "无文件变更"
else
  git commit -m "feat: 规格及厂家查询系统 - 订单/材质/规格/厂家管理、统计分析、Excel导出"
fi

# 推送
git branch -M main
git push -u origin main

echo "完成！已推送到 https://github.com/devuser/gangguanbao"

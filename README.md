# typo

Like thefuck, but he uses Go to implement it more intelligently.

## Thefuck 的几个缺点

Github Repo：https://github.com/nvbn/thefuck

1. 不继续维护了，Python 3.12+ 安装报错；
2. 不支持自定义的 alias 配置；
3. 多级命令时引号会丢失：https://github.com/nvbn/thefuck/issues/1543；
4. 依赖 python 的运行库。

## Typo

设想中的 typo 命令足够专注，它只修正上一次 terminal 输入的错误 command；

不根据上下文推荐此时应该执行什么 command 和执行 command 等其他功能。

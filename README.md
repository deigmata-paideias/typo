# typo

Like thefuck, but he uses Go to implement it more intelligently.

## Thefuck 的几个缺点

Github Repo：https://github.com/nvbn/thefuck

1. 不继续维护了，Python 3.12+ 安装报错；
2. 不支持自定义的 alias 配置；
3. 多级命令时引号会丢失：https://github.com/nvbn/thefuck/issues/1543；
4. 依赖 python 的运行库。

## Typo

设想中的 typo 命令足够专注，它只修正上一次 terminal 输入的错误 command 并且尝试运行；

不根据上下文推荐此时应该执行什么 command 其他功能。

## 实现原理

1. 命令修正：用莱文斯坦距离（Levenshtein distance）等算法计算相似度；
2. 命令存储：
   1. 自动扫描 man 里面的命令列表；
   2. 用户自定义的 alias 配置；
3. 用户选择以执行。

## 二进制运行

![D7AEAA94-35B5-4D95-8ED5-E8AF6AAD0A55](https://github.com/user-attachments/assets/05edefff-544c-417b-8772-5e78640c109e)

## oh-my-zsh 集成

```
mkdir -p ~/.oh-my-zsh/custom/plugins/typo
cp ~/project/indi/typo/zsh/typo.plugin.zsh ~/.oh-my-zsh/custom/plugins/typo/typo.plugin.zsh
```

然后在 ~/.zshrc 中添加 typo 插件：

```
plugins=(... typo)
```

最后执行 source ~/.zshrc，按两下 esc 执行 typo 命令。

### zsh 集成效果演示

![FA5AB79E-A12C-4019-89C6-E9D02758E0E5](https://github.com/user-attachments/assets/73c3d354-8bb0-4ac9-bbe8-03f3635dba1a)



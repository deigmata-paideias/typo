# typo

Like thefuck, but he uses Go to implement it more intelligently.

## Thefuck 的几个缺点

Github Repo：https://github.com/nvbn/thefuck

1. 不继续维护了，Python 3.12+ 安装报错；
2. 不支持自定义的 alias 配置；
3. 多级命令时引号会丢失：https://github.com/nvbn/thefuck/issues/1543；
4. 依赖 python 的运行库;
5. thefuck 有时候可能返回错误的修复命令。

## Typo

设想中的 typo 命令足够专注，它只修正上一次 terminal 输入的错误 command 并且尝试运行；

不根据上下文推荐此时应该执行什么 command 其他功能。

- 使用 typo 可以，修复上一条 typo 的错误命令；
- 在忘记具体的命令时，可以模糊搜索并执行。

> 和 thefuck 不同的是，thefuck 使用规则配置，而 typo 依赖系统命令。

## 实现原理

1. 命令修正：用莱文斯坦距离（Levenshtein distance）等算法计算相似度；
2. 命令存储：
   1. 自动扫描 man 里面的命令列表；
   2. 用户自定义的 alias 配置；
3. 用户选择以执行。

## 二进制运行

1. clone 仓库，git clone https://github.com/deigmata-paideias/typo；
2. 编译，cd typo && make build
3. 扫描 man 并存入数据库：./bin/darwin/arm64/typo scanner --type man；
4. （无配置可跳过）扫描自定义 alias 并存入数据库：./bin/darwin/arm64/typo scanner --type man；
5. 运行  ./bin/darwin/arm64/typo run。

![D7AEAA94-35B5-4D95-8ED5-E8AF6AAD0A55](https://github.com/user-attachments/assets/05edefff-544c-417b-8772-5e78640c109e)

## oh-my-zsh 集成

> 目前处在 dev 阶段，需要手动修改下 zsh 插件的 bin 目录配置！

```
mkdir -p ~/.oh-my-zsh/custom/plugins/typo
cp zsh/typo.plugin.zsh ~/.oh-my-zsh/custom/plugins/typo/typo.plugin.zsh
```

然后在 ~/.zshrc 中添加 typo 插件：

```
plugins=(... typo)
```

最后执行 source ~/.zshrc，按两下 esc 执行 typo 命令。

![AB5E9F0B-9BED-4E30-A3A7-57BD35E05C4E](https://github.com/user-attachments/assets/70ea54c7-5f4c-43a0-a138-8b9cb16dcc6b)

## LLM 集成

参考 etc 目录下的配置文件格式，将此配置复制到 ~/.config/typo 目录下即可。


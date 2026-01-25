# typo

Like thefuck, but he uses Go to implement it more intelligently.

## Thefuck's Shortcomings

Github Repo：https://github.com/nvbn/thefuck

1. No longer maintained, installation errors on Python 3.12+;
2. Does not support custom alias configurations;
3. Quotes are lost with multi-level commands: https://github.com/nvbn/thefuck/issues/1543;
4. Depends on Python runtime libraries;
5. thefuck sometimes returns incorrect fix commands.

## Typo

The envisioned typo command is focused enough - it only corrects the erroneous command from the last terminal input and attempts to run it;

It does not recommend what command should be executed based on context or other functions.

- Using typo can fix the typo error from the previous command;
- When you forget the specific command, you can fuzzy search and execute.

> Unlike thefuck which uses rule-based configuration, typo relies on system commands.

## Implementation Principles

1. Command correction: Use algorithms such as Levenshtein distance to calculate similarity;
2. Command storage:
   1. Automatically scan the command list from man pages;
   2. User-defined alias configurations;
3. User selection to execute.

## Binary Execution

1. Clone the repository: git clone https://github.com/deigmata-paideias/typo;
2. Build: cd typo && make build
3. Scan man pages and save to database: ./bin/darwin/arm64/typo scanner --type man;
4. (Skip if no configuration) Scan custom aliases and save to database: ./bin/darwin/arm64/typo scanner --type alias;
5. Run: ./bin/darwin/arm64/typo run.

![D7AEAA94-35B5-4D95-8ED5-E8AF6AAD0A55](https://github.com/user-attachments/assets/05edefff-544c-417b-8772-5e78640c109e)

## oh-my-zsh Integration

```bash
mkdir -p ~/.oh-my-zsh/custom/plugins/typo
cp zsh/typo.plugin.zsh ~/.oh-my-zsh/custom/plugins/typo/typo.plugin.zsh
```

Then add the typo plugin in ~/.zshrc:

```bash
plugins=(... typo)
```

Finally execute source ~/.zshrc, press ESC twice to run the typo command.

![FA5AB79E-A12C-4019-89C6-E9D02758E0E5](https://github.com/user-attachments/assets/73c3d354-8bb0-4ac9-bbe8-03f3635dba1a)

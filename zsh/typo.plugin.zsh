#!/usr/bin/env zsh

# Typo zsh plugin

__typo_fix() {

    # dev run
    local typo_bin="/Users/shown/workspace/golang/playground/typo/bin/darwin/arm64/typo"

    zle -I

    "$typo_bin" run

    zle -R
    zle reset-prompt
}

zle -N __typo_fix
bindkey '\e\e' __typo_fix

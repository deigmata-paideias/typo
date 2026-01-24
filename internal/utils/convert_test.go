package utils

import (
	"fmt"
	"testing"
)

func Test_Convert_Alias(t *testing.T) {
	var aliasOutput = `
-='cd -'
.='cd ..'
..='cd ../..'
...='cd ../../..'
....=../../..
.....=../../../..
......=../../../../..
1='cd -1'
2='cd -2'
3='cd -3'
4='cd -4'
5='cd -5'
6='cd -6'
7='cd -7'
8='cd -8'
9='cd -9'
_='sudo '
archive='web_search archive'
ask='web_search ask'
baidu='web_search baidu'
bing='web_search bing'
brs='web_search brave'
c=onefetch
chatgpt='web_search chatgpt'
claudeai='web_search claudeai'
ddg='web_search duckduckgo'
deepl='web_search deepl'
dockerhub='web_search dockerhub'
ducky='web_search duckduckgo \!'
ecosia='web_search ecosia'
edit-in-kitty='kitten edit-in-kitty'
egrep='grep -E'
f=fzf
ff='fzf --preview "bat --color=always {}"'
fgrep='grep -F'
gems='web_search gems'
gg='git clone'
github='web_search github'
givero='web_search givero'
globurl='noglob urlglobber '
goodreads='web_search goodreads'
google='web_search google'
gopkg='web_search gopkg'
grep='grep --color=auto'
grok='web_search grok'
history=omz_history
image='web_search duckduckgo \!i'
k=kubectl
ka='kubectl apply'
kcf='kubectl create -f'
kd='kubectl describe'
kdiff='kitty +kitten diff'
kg='kubectl get'
kicat='kitty +kitten icat'
klf='kubectl logs -f'
l='ls -lah'
lD='eza -glD'
lDD='eza -glDa'
lS='eza -gl -ssize'
lT='eza -gl -snewest'
la='eza -gla'
ldot='eza -gld .*'
ll='eza -gl'
ls=eza
lsa='ls -lah'
lsd='eza -gd'
lsdl='eza -gdl'
map='web_search duckduckgo \!m'
md='mkdir -p'
mk=make
news='web_search duckduckgo \!n'
npmpkg='web_search npmpkg'
packagist='web_search packagist'
pi='pip3 install'
ppai='web_search ppai'
py=python3
qwant='web_search qwant'
rd=rmdir
reddit='web_search reddit'
rscrate='web_search rscrate'
rsdoc='web_search rsdoc'
run-help=man
s=neofetch
scholar='web_search scholar'
sp='web_search startpage'
stackoverflow='web_search stackoverflow'
sv='source venv/bin/active'
which-command=whence
wiki='web_search duckduckgo \!w'
wolframalpha='web_search wolframalpha'
x=extract
y=yazi
yahoo='web_search yahoo'
yandex='web_search yandex'
youtube='web_search youtube'
z='zshz 2>&1'
`

	expectContent := `[{Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:- Type:alias Source:alias Description:Alias for: cd - Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:. Type:alias Source:alias Description:Alias for: cd .. Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:.. Type:alias Source:alias Description:Alias for: cd ../.. Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:... Type:alias Source:alias Description:Alias for: cd ../../.. Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:.... Type:alias Source:alias Description:Alias for: ../../.. Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:..... Type:alias Source:alias Description:Alias for: ../../../.. Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:...... Type:alias Source:alias Description:Alias for: ../../../../.. Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:1 Type:alias Source:alias Description:Alias for: cd -1 Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:2 Type:alias Source:alias Description:Alias for: cd -2 Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:3 Type:alias Source:alias Description:Alias for: cd -3 Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:4 Type:alias Source:alias Description:Alias for: cd -4 Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:5 Type:alias Source:alias Description:Alias for: cd -5 Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:6 Type:alias Source:alias Description:Alias for: cd -6 Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:7 Type:alias Source:alias Description:Alias for: cd -7 Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:8 Type:alias Source:alias Description:Alias for: cd -8 Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:9 Type:alias Source:alias Description:Alias for: cd -9 Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:_ Type:alias Source:alias Description:Alias for: sudo Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:archive Type:alias Source:alias Description:Alias for: web_search archive Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:ask Type:alias Source:alias Description:Alias for: web_search ask Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:baidu Type:alias Source:alias Description:Alias for: web_search baidu Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:bing Type:alias Source:alias Description:Alias for: web_search bing Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:brs Type:alias Source:alias Description:Alias for: web_search brave Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:c Type:alias Source:alias Description:Alias for: onefetch Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:chatgpt Type:alias Source:alias Description:Alias for: web_search chatgpt Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:claudeai Type:alias Source:alias Description:Alias for: web_search claudeai Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:ddg Type:alias Source:alias Description:Alias for: web_search duckduckgo Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:deepl Type:alias Source:alias Description:Alias for: web_search deepl Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:dockerhub Type:alias Source:alias Description:Alias for: web_search dockerhub Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:ducky Type:alias Source:alias Description:Alias for: web_search duckduckgo \! Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:ecosia Type:alias Source:alias Description:Alias for: web_search ecosia Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:edit-in-kitty Type:alias Source:alias Description:Alias for: kitten edit-in-kitty Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:egrep Type:alias Source:alias Description:Alias for: grep -E Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:f Type:alias Source:alias Description:Alias for: fzf Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:ff Type:alias Source:alias Description:Alias for: fzf --preview "bat --color=always {} Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:fgrep Type:alias Source:alias Description:Alias for: grep -F Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:gems Type:alias Source:alias Description:Alias for: web_search gems Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:gg Type:alias Source:alias Description:Alias for: git clone Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:github Type:alias Source:alias Description:Alias for: web_search github Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:givero Type:alias Source:alias Description:Alias for: web_search givero Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:globurl Type:alias Source:alias Description:Alias for: noglob urlglobber Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:goodreads Type:alias Source:alias Description:Alias for: web_search goodreads Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:google Type:alias Source:alias Description:Alias for: web_search google Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:gopkg Type:alias Source:alias Description:Alias for: web_search gopkg Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:grep Type:alias Source:alias Description:Alias for: grep --color=auto Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:grok Type:alias Source:alias Description:Alias for: web_search grok Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:history Type:alias Source:alias Description:Alias for: omz_history Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:image Type:alias Source:alias Description:Alias for: web_search duckduckgo \!i Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:k Type:alias Source:alias Description:Alias for: kubectl Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:ka Type:alias Source:alias Description:Alias for: kubectl apply Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:kcf Type:alias Source:alias Description:Alias for: kubectl create -f Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:kd Type:alias Source:alias Description:Alias for: kubectl describe Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:kdiff Type:alias Source:alias Description:Alias for: kitty +kitten diff Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:kg Type:alias Source:alias Description:Alias for: kubectl get Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:kicat Type:alias Source:alias Description:Alias for: kitty +kitten icat Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:klf Type:alias Source:alias Description:Alias for: kubectl logs -f Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:l Type:alias Source:alias Description:Alias for: ls -lah Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:lD Type:alias Source:alias Description:Alias for: eza -glD Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:lDD Type:alias Source:alias Description:Alias for: eza -glDa Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:lS Type:alias Source:alias Description:Alias for: eza -gl -ssize Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:lT Type:alias Source:alias Description:Alias for: eza -gl -snewest Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:la Type:alias Source:alias Description:Alias for: eza -gla Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:ldot Type:alias Source:alias Description:Alias for: eza -gld .* Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:ll Type:alias Source:alias Description:Alias for: eza -gl Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:ls Type:alias Source:alias Description:Alias for: eza Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:lsa Type:alias Source:alias Description:Alias for: ls -lah Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:lsd Type:alias Source:alias Description:Alias for: eza -gd Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:lsdl Type:alias Source:alias Description:Alias for: eza -gdl Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:map Type:alias Source:alias Description:Alias for: web_search duckduckgo \!m Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:md Type:alias Source:alias Description:Alias for: mkdir -p Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:mk Type:alias Source:alias Description:Alias for: make Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:news Type:alias Source:alias Description:Alias for: web_search duckduckgo \!n Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:npmpkg Type:alias Source:alias Description:Alias for: web_search npmpkg Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:packagist Type:alias Source:alias Description:Alias for: web_search packagist Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:pi Type:alias Source:alias Description:Alias for: pip3 install Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:ppai Type:alias Source:alias Description:Alias for: web_search ppai Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:py Type:alias Source:alias Description:Alias for: python3 Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:qwant Type:alias Source:alias Description:Alias for: web_search qwant Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:rd Type:alias Source:alias Description:Alias for: rmdir Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:reddit Type:alias Source:alias Description:Alias for: web_search reddit Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:rscrate Type:alias Source:alias Description:Alias for: web_search rscrate Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:rsdoc Type:alias Source:alias Description:Alias for: web_search rsdoc Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:run-help Type:alias Source:alias Description:Alias for: man Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:s Type:alias Source:alias Description:Alias for: neofetch Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:scholar Type:alias Source:alias Description:Alias for: web_search scholar Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:sp Type:alias Source:alias Description:Alias for: web_search startpage Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:stackoverflow Type:alias Source:alias Description:Alias for: web_search stackoverflow Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:sv Type:alias Source:alias Description:Alias for: source venv/bin/active Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:which-command Type:alias Source:alias Description:Alias for: whence Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:wiki Type:alias Source:alias Description:Alias for: web_search duckduckgo \!w Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:wolframalpha Type:alias Source:alias Description:Alias for: web_search wolframalpha Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:x Type:alias Source:alias Description:Alias for: extract Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:y Type:alias Source:alias Description:Alias for: yazi Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:yahoo Type:alias Source:alias Description:Alias for: web_search yahoo Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:yandex Type:alias Source:alias Description:Alias for: web_search yandex Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:youtube Type:alias Source:alias Description:Alias for: web_search youtube Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:z Type:alias Source:alias Description:Alias for: zshz 2>&1 Aliases:[] Options:[]}]`
	expectLen := 96

	convert, err := Convert(aliasOutput, "alias")
	if err != nil {
		t.Error(err)
	}

	for _, c := range convert {
		fmt.Printf("name: %v \t source: %v \n", c.Name, c.Source)
	}

	if len(convert) != expectLen {
		t.Errorf("expect len: %v, but got: %v", expectLen, len(convert))
	}
	if expectContent != fmt.Sprintf("%+v", convert) {
		t.Errorf("expect content: %v, but got: %v", expectContent, fmt.Sprintf("%+v", convert))
	}

}

func Test_Convert_Git(t *testing.T) {

	var gitOutput = `
alias.br=branch
alias.c=commit -s -m
alias.co=checkout
alias.ch=cherry-pick
alias.dump=cat-file -p
alias.hist=log --pretty=format:'%C(yellow)[%ad]%C(reset) %C(green)[%h]%C(reset) | %C(red)%s %C(bold red){{%an}}%C(reset) %C(blue)%d%C(reset)' --graph --date=short
alias.st=status
alias.type=cat-file -t
`
	expectContent := `[{Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:br Type:alias Source:git Description:Git alias: branch Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:c Type:alias Source:git Description:Git alias: commit -s -m Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:co Type:alias Source:git Description:Git alias: checkout Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:ch Type:alias Source:git Description:Git alias: cherry-pick Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:dump Type:alias Source:git Description:Git alias: cat-file -p Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:hist Type:alias Source:git Description:Git alias: log --pretty=format:'%C(yellow)[%ad]%C(reset) %C(green)[%h]%C(reset) | %C(red)%s %C(bold red){{%an}}%C(reset) %C(blue)%d%C(reset)' --graph --date=short Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:st Type:alias Source:git Description:Git alias: status Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:type Type:alias Source:git Description:Git alias: cat-file -t Aliases:[] Options:[]}]`
	expectLen := 8

	convert, err := Convert(gitOutput, "git")
	if err != nil {
		t.Error(err)
	}

	if len(convert) != expectLen {
		t.Errorf("expect len: %v, but got: %v", expectLen, len(convert))
	}
	if expectContent != fmt.Sprintf("%+v", convert) {
		t.Errorf("expect content: %v, but got: %v", expectContent, fmt.Sprintf("%+v", convert))
	}

}

func Test_Convert_Man(t *testing.T) {

	var manOutput = `
iconvctl(3)              - controlling and diagnostical facility for iconv(3)
iconvlist(3)             - retrieving a list of character encodings supported by iconv(3)
icqctl(1), Other_name_for_same_program()(1), Yet another name for the same program.(1) - This line parsed for whatis database
id(1)                    - return user identity
idebug(n)                - Interactive debugging environment in TkCon
ident(n)                 - Ident protocol client
if(3pm)                  - "use" a Perl module if a condition holds
if(ntcl)                 - Execute scripts conditionally
`

	expectContent := `[{Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:iconvctl(3) Type:man Source:man Description:- controlling and diagnostical facility for iconv(3) Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:iconvlist(3) Type:man Source:man Description:- retrieving a list of character encodings supported by iconv(3) Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:icqctl(1), Type:man Source:man Description:Other_name_for_same_program()(1), Yet another name for the same program.(1) - This line parsed for whatis database Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:id Type:man Source:man Description:- return user identity Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:idebug(n) Type:man Source:man Description:- Interactive debugging environment in TkCon Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:ident(n) Type:man Source:man Description:- Ident protocol client Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:if(3pm) Type:man Source:man Description:- "use" a Perl module if a condition holds Aliases:[] Options:[]} {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:if(ntcl) Type:man Source:man Description:- Execute scripts conditionally Aliases:[] Options:[]}]`
	expectLen := 8

	convert, err := Convert(manOutput, "man")
	if err != nil {
		t.Error(err)
	}

	if len(convert) != expectLen {
		t.Errorf("expect len: %v, but got: %v", expectLen, len(convert))
	}
	if expectContent != fmt.Sprintf("%+v", convert) {
		t.Errorf("expect content: %v, but got: %v", expectContent, fmt.Sprintf("%+v", convert))
	}

}

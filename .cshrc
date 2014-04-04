# $FreeBSD: src/share/skel/dot.cshrc,v 1.14.10.1.2.1 2009/10/25 01:10:29 kensmith Exp $
#
# .cshrc - csh resource script, read at beginning of execution by each shell
#
# see also csh(1), environ(7).
#

alias h		history 25
alias j		jobs -l
alias la	ls -a
alias vi	vim
alias lf	ls -FA
alias ll	ls -lA

# A righteous umask
umask 22

set path = (/sbin /bin /usr/sbin /usr/bin /usr/games /usr/local/sbin /usr/local/bin $HOME/bin $HOME/scripts)
setenv	EDITOR	/usr/local/bin/vim
setenv	PAGER	more
setenv	BLOCKSIZE	K
setenv LC_ALL ru_RU.UTF-8
setenv LANG ru_RU.UTF-8
bindkey "\e[1~" beginning-of-line  # Home
bindkey "\e[7~" beginning-of-line  # Home rxvt
bindkey "\e[2~" overwrite-mode     # Ins
bindkey "\e[3~" delete-char        # Delete
bindkey "\e[4~" end-of-line        # End
bindkey "\e[8~" end-of-line        # End rxvt

if ($?prompt) then
	# An interactive shell -- set some stuff up
	set filec
	set history = 10000
	set savehist = (10000 merge)
	set autolist
	set autoexpand
	set color
	set nobeep
	set promptchars = \$\#
	set prompt = '%B%n@%m%b:%~%# '
	set mail = (/var/mail/$USER)
	setenv CLICOLOR
	setenv LSCOLORS DxHxdxfxgxfgedgbgaacad
	if ( $?tcsh ) then
		bindkey "^W" backward-delete-word
		bindkey "^R" i-search-back
		bindkey -k up history-search-backward
		bindkey -k down history-search-forward
	endif

endif

# fix set -o
DISABLE_MAGIC_FUNCTIONS=true 
EDITOR=vim

PATH=$PATH:$HOME/development/bin

# to support mulpile nodejs versions
export NVM_DIR="$HOME/.nvm"
[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"  # This loads nvm
[ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion"  # This loads nvm bash_completion

source $HOME/development/aliases.env

# useful in termnal to move between words
bindkey -e
# alt+h to move word back
bindkey "^[h" backward-word
# alt+l to move word forward
bindkey "^[l" forward-word

# do first git clone https://github.com/VundleVim/Vundle.vim.git ~/.vim/bundle/Vundle.vim

set nocompatible              " be iMproved, required
filetype off                  " required

" set the runtime path to include Vundle and initialize
set rtp+=~/.vim/bundle/Vundle.vim
call vundle#begin()

" let Vundle manage Vundle, required
Plugin 'VundleVim/Vundle.vim'

Plugin 'https://github.com/posva/vim-vue'
Plugin 'https://github.com/tomasr/molokai'
Plugin 'https://github.com/pangloss/vim-javascript'
Plugin 'https://github.com/w0rp/ale'
" Plugin 'https://github.com/fatih/vim-go'

" All of your Plugins must be added before the following line
call vundle#end()            " required
syntax enable

filetype plugin indent on    " required
" To ignore plugin indent changes, instead use:
" filetype plugin on

"
" Brief help
" :PluginList       - lists configured plugins
" :PluginInstall    - installs plugins; append `!` to update or just :PluginUpdate
" :PluginSearch foo - searches for foo; append `!` to refresh local cache
" :PluginClean      - confirms removal of unused plugins; append `!` to auto-approve removal
"
" see :h vundle for more details or wiki for FAQ
" Put your non-Plugin stuff after this line

" https://medium.com/@alexlafroscia/writing-js-in-vim-4c971a95fd49
" let g:ale_fixers = {
"   \ `javascript`: [`eslint`]
"   \ }

" let g:LanguageClient_serverCommands = {
"    \ 'vue': ['vls']
"    \ }
let g:ale_linters = {
  \ 'go': ['gopls'],
  \}

let g:go_version_warning = 0
let g:rehash256 = 1

" ==== ALE SETTINGS ====
" Set this setting in vimrc if you want to fix files automatically on save.
" This is off by default.
" let g:ale_fix_on_save = 1

" Enable completion where available.
" let g:ale_completion_enabled = 1

" colorscheme molokai

" autocmd FileType vue syntax sync fromstart
" set tabstop=2 shiftwidth=2 softtabstop=2 noexpandtab shiftwidth=2 smarttab
set tabstop=2 shiftwidth=2 softtabstop=2 expandtab shiftwidth=2 smarttab
set expandtab

set undodir=~/.vim/.undo//
set backupdir=~/.vim/.backup//
set directory=~/.vim/.swp//
set list
" set listchars=eol:¬,trail:·,tab:▸\
set listchars=trail:·,tab:-▸

set hlsearch    " highlight all search results
set ignorecase  " do case insensitive search
set incsearch   " show incremental search results as you type
set number      " display line number
set listchars=eol:$,tab:>-,trail:~,extends:>,precedes:<
set list


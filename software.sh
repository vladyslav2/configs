apt-get install -y ca-certificates curl gnupg vim
apt-get install -y awesome awesome-extra redshift
apt-get install -y zerotier-one
# interesting torrent tracker with browser ui
apt-get install -y qbittorrent-nox

# snyk
curl --compressed https://static.snyk.io/cli/latest/snyk-linux -o snyk
chmod +x ./snyk
mv ./snyk $HOME/bin/

# pipewire
apt-get install -y pipewire pamixer

# bc
apt-get install -y bc


# ethereum
curl https://github.com/ethereum/solidity/releases/download/v0.8.21/solc-static-linux -o solc; chmod +x solc; mv solc $HOME/bin
# foundry
curl -L https://foundry.paradigm.xyz | bash
echo 'PATH=$PATH:$HOME/.foundry/bin' >> ~/.zshrc
$HOME/.foundry/foundryup

# calibri to convert pdf to epub
sudo -v && wget -nv -O- https://download.calibre-ebook.com/linux-installer.sh | sudo sh /dev/stdin

# install compfy, composer for awesomewm
apt-get install -y ninja-build meson libev-dev libx11-xcb-dev libxcb-render-util0-dev libxcb-image0-dev libxcb-damage0-dev libxcb-randr0-dev libxcb-sync-dev libxcb-composite0-dev libxcb-xinerama0-dev libxcb-present-dev libxcb-glx0-dev uthash-dev libjson-c-dev libcurlpp-dev libconfig-dev
git clone https://github.com/allusive-dev/compfy
echo 'follow instruction'

# screenrecording
apt-get install -y slop ffmpeg
echo 'put video.sh to the bin folder'

# neovim
apt-get install neovim vim-ale

# obsidian
wget -O ~/bin/obsidian https://github.com/obsidianmd/obsidian-releases/releases/download/v1.6.7/Obsidian-1.6.7.AppImage && chmod +x ~/bin/obsidian

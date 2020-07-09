#!/usr/bin/env bash

coretabs_install_dir() {
  printf %s "${HOME}/.coretabs"
}

coretabs_latest_version() {
  echo "v0.0.2"
}

download_url(){
    echo "https://github.com/TChi91/coretabs/releases/download/$(coretabs_latest_version)/coretabs"
}

# create $HOME/.coretabs directory
mkdir -p $(coretabs_install_dir)

# cd into $HOME/.coretabs directory
cd $(coretabs_install_dir)

#download the last version of coretabs cli
wget -qO coretabs $(download_url)

# export coretabs path to $PATH 
echo "export PATH=\$PATH:/home/tchi/.coretabs" >> ~/.profile

# Give coretabs permission to execute
chmod +x coretabs

# Reload 
source ~/.profile

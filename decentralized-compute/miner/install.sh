#!/bin/bash
current_dir=$(pwd)
eai_file=miner.sh
file="$current_dir/$eai_file"
alias=/usr/local/bin/miner
chmod +x $file
if [ -e "$alias" ]; then
    rm -f $alias
fi
ln -s $file $alias

echo "miner has been installed!!!!"
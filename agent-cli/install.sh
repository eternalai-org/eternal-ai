#!/bin/bash
yarn build

current_dir=$(pwd)
file="$current_dir/bin/eai.js"
alias=/usr/local/bin/eai
chmod +x $file
if [ -e "$alias" ]; then
    rm -f $alias
fi
ln -s $file $alias


# Step 1: Build eai-chat bin
cd ..
make build_decentralize_server_linux
cp ./eai-chat ./agent-cli/src/eternalai/eai-chat-linux

echo "Copy file eai-chat-linux"

if [[ "$OSTYPE" == "linux-gnu"* ]]; then
    echo "Running on Linux"
    # make build_decentralize_server_linux
elif [[ "$OSTYPE" == "darwin"* ]]; then
    echo "Running on macOS"
    make build_decentralize_server_osx
else
    echo "Unknown operating system: $OSTYPE"
fi

# Step 2: Build docker image eternalai-agent
cd $current_dir/src/eternalai
docker build -t eternalai-agent .
echo "Docker image eternalai-agent is built successfully"

echo "eai has been installed!!!"
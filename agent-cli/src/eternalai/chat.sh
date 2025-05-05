#!/bin/bash

agent_id=$1 
config_path=$2 

current_dir=$(pwd)
echo $current_dir

cd ..

export CONFIG_PATH=$config_path
./eai-chat chat $agent_id
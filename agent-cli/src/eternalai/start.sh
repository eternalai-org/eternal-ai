#!/bin/bash

echo "Starting new agent with EternalAI framework"

agent_uid=$1 
ETERNALAI_URL=$2
ETERNALAI_API_KEY=$3
ETERNALAI_CHAIN_ID=$4
ETERNALAI_RPC_URL=$5
ETERNALAI_AGENT_CONTRACT_ADDRESS=$6
ETERNALAI_AGENT_ID=$7
ETERNALAI_MODEL=$8
TWITTER_USERNAME=$9
TWITTER_PASSWORD=${10}
TWITTER_EMAIL=${11}
TWITTER_TARGET_USERS=${12}
AGENT_NAME=${13}
SERVICE_PORT=${14}

echo "Agent UID: $agent_uid"

current_dir=$(pwd)

### make specific folder for new agent 
cd agents
if [ -d $agent_uid ]; then
  echo "Folder $agent_uid already exists."
else
  mkdir $agent_uid
fi

cd $agent_uid

### create .env
env_file_name=".env_$agent_uid"

# Check if the file already exists
# if [ -f "$env_file_name" ]; then
#     echo "$env_file_name already exists. Don't start existing agent again."
#     # exit 0
# else
#     echo "$env_file_name does not exist. Creating it."
# fi

# write .env file
cat <<EOL > "$env_file_name"
ETERNALAI_API_KEY=$ETERNALAI_API_KEY
ETERNALAI_URL=$ETERNALAI_URL
ETERNALAI_CHAIN_ID=$ETERNALAI_CHAIN_ID
ETERNALAI_RPC_URL=$ETERNALAI_RPC_URL
ETERNALAI_AGENT_CONTRACT_ADDRESS=$ETERNALAI_AGENT_CONTRACT_ADDRESS

AGENT_SYSTEM_PROMPT_PATH=config.json

ETERNALAI_AGENT_ID=$ETERNALAI_AGENT_ID
ETERNALAI_MODEL=$ETERNALAI_MODEL


TWITTER_USERNAME=$TWITTER_USERNAME
TWITTER_PASSWORD=$TWITTER_PASSWORD
TWITTER_EMAIL=$TWITTER_EMAIL

ACTION_INTERVAL=1
ENABLE_ACTION_PROCESSING=true
MAX_ACTIONS_PROCESSING=2
ACTION_TIMELINE_TYPE=following
TWITTER_TARGET_USERS=$TWITTER_TARGET_USERS
EOL


# file config for api chat
api_config_file="local_contracts.json"

cat <<EOL > "$api_config_file"
{
  "server_base_url": "http://localhost:$SERVICE_PORT",
  "rpc": "$ETERNALAI_RPC_URL",
  "model_name": "$ETERNALAI_MODEL", 
  "chain_id": "$ETERNALAI_CHAIN_ID",
  "agent_contract_address": "$ETERNALAI_AGENT_CONTRACT_ADDRESS"
}
EOL


# file config for small service
# cd $root_dir/decentralized-inference
s_service_config_file="config.json"

cat <<EOL > "$s_service_config_file"
{
  "server": {
    "port": $SERVICE_PORT
  },
  "mongodb": {
    "uri": "mongodb://localhost:27017",
    "db": "decentralized-inference"
  },
  "file_path_infer": "/tmp/eternal-data",
  "submit_file_path": false,
  "chat_completion_url": "$ETERNALAI_URL/chat/completions",
  "api_key_chat_completion": "$ETERNALAI_API_KEY"
}
EOL

# copy file .eai-chat
cd $current_dir && cd ..
root_dir=$(pwd)
cp $root_dir/eai-chat $current_dir/agents/$agent_uid/eai-chat

### docker run to start agent
cd $current_dir/agents/$agent_uid
docker run -d --name $agent_uid -v ./config.json:/app/eternal-ai/decentralized-inference/config.json -v ./local_contracts.json:/app/eternal-ai/decentralized-compute/worker-hub/env/local_contracts.json -p $SERVICE_PORT:$SERVICE_PORT eternalai-agent 
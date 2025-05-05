#!/bin/bash

echo "Starting new agent with Eliza framework"

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

echo "Agent UID: $agent_uid"

current_dir=$(pwd)

### make specific folder for new agent
cd agents
mkdir $agent_uid
cd $agent_uid

### create .env
env_file_name=".env_$agent_uid"
echo env_file_name

# Check if the file already exists
if [ -f "$env_file_name" ]; then
    echo "$env_file_name already exists. Don't start existing agent again."
    exit 0
else
    echo "$env_file_name does not exist. Creating it."
fi

# custom env
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


### create config.json and custom agent name

cp $current_dir/src/eliza/config.json .
jq --arg new_name "$AGENT_NAME" '.name = $new_name' "config.json" > tmp.json && mv tmp.json "config.json"


### docker run to start agent

docker run -d --env-file $env_file_name --name $agent_uid  -v ./config.json:/app/eliza/agents/config.json eliza











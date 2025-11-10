#!/bin/bash
current_dir=$(pwd)
worker_hub_folder="$current_dir/decentralized-compute/worker-hub" #step 1: eai miner setup
aaas_folder="$current_dir/agent-as-a-service/agent-orchestration/backend" #step 2: aaas start
api_folder="$current_dir/decentralized-inference" #step 4: start APIs
ai721_folder="$current_dir/developer-guides/run-an-end-to-end-decentralized-for-ai-agents/4.how-to-deploy-and-mint-agent" #step 5: mint agent

# Check if at least one argument is provided
if [ $# -lt 1 ]; then
    echo "Usage:"
    echo "- eai miner setup"
    echo "- eai aaas deploy-contract"
    echo "- eai aaas start"
    echo "- eai apis start"
    echo "- eai agent create <prompt-file>"
    echo "- eai agent list"
    echo "- eai agent info <agent_id>"
    echo "- eai agent chat <agent_id>"
    exit 1
fi

check_command() {
    command -v "$1" >/dev/null 2>&1
}

# Function to check for docker-compose or docker compose
check_docker() {
    if check_command "docker-compose"; then
        echo "docker-compose"
    elif check_command "docker"; then
        # Check if 'docker compose' is available (as a subcommand)
        if docker compose version >/dev/null 2>&1; then
            echo "docker compose"
        else
            #docker is installed, but 'docker compose' is not available.
            echo "-1"
        fi
    else
        #Neither docker-compose nor docker is installed.
        echo "-2"
    fi
}

# Function to handle miner commands
handle_miner_commands() {
    case "$1" in
        "setup")
            echo "Setting up local chain and miners..."
            cd "$worker_hub_folder" && make start_cli
            ;;
        *)
            echo "Invalid option: $1 for miner"
            exit 1
            ;;
    esac
}


# Function to handle miner commands
handle_miner_fast_commands() {
    case "$1" in
        "setup")
            echo "Setting up local chain and miners one step..."
            cd "$worker_hub_folder" && make start_cli_fast
            ;;
        *)
            echo "Invalid option: $1 for miner"
            exit 1
            ;;
    esac
}

handle_aaas_commands() {
    case "$1" in
        "deploy-contract")
           echo "aaas deploy-contract"
           cd $ai721_folder && ./deploy-ai721.sh
        ;;
        "start")
            echo "eai aaas start"
            docker_compose=$(check_docker)
            cd $aaas_folder && $docker_compose build && $docker_compose up -d
            ;;
        *)
        echo "Invalid option: $1 for miner"
        exit 1
        ;;
    esac
}

handle_agent_commands() {
    case "$1" in
        "chat")
          cd "$current_dir"  && \
          ./eai-chat server & && \
          ./eai-chat chat $2
        ;;
       "create")
          echo "creating agent with promt-file: $2"
          cd ${ai721_folder} && ./mint-agent.sh $2
          ;;
      "list")
          cd ${ai721_folder} && ./ls-agents.sh
          ;;
      "info")
          cd ${ai721_folder} && ./get-system-prompt.sh $2
          ;;
        *)
        echo "Invalid option: $1 for miner"
        exit 1
        ;;
    esac
}

handle_api_commands() {
    case "$1" in
        "start")
          echo $api_folder
          for pid in $(lsof -t -i:8484); do kill -9 $pid; done && \
          rm -rf "./eai-chat" && \
          cd $api_folder  &&  \
          go build -o eai-chat main.go  && \
          rm -rf config.json && \
          cp -i  config.json.example config.json && \
          cp -i "$api_folder/eai-chat"  "$current_dir/eai-chat" && \
          cd "$current_dir"  && \
          ./eai-chat server &
        ;;
        *)
        echo "Invalid option: $1 for miner"
        exit 1
        ;;
    esac
}

# Handle the main command using a case statement
case "$1" in
    "miner")
        handle_miner_commands "$2"
        ;;
    "miner-fast")
        handle_miner_fast_commands "$2"
        ;;
    "aaas")
        handle_aaas_commands "$2"
        ;;
    "apis")
        handle_api_commands "$2"
        ;;
    "agent")
        handle_agent_commands "$2" "$3"
        ;;
    *)
        echo "Invalid option: $1"
        exit 1
        ;;
esac

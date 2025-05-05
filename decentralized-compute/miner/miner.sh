#!/bin/bash
current_dir=$(pwd)
api_dir="$current_dir/cmd/api"
miner_dir="$current_dir/cmd/miner"
config_file="$current_dir/env/config.env"
build_dir="$current_dir/build"
api_port=8005

create_file() {
    file="$current_dir/build/$1"
    go build  -o="$file"  "$2/main.go"
    while [ ! -e "$file" ]; do
      echo "Waiting for the file to be created..."
      sleep 5  # Wait for 2 seconds before checking again
    done
    chmod +x "$file"
}

#miner
handle_start_node_commands() {
  # Check if the file exists
  file="$build_dir/miner"
  if [ -e "$file" ]; then
      echo "Run miner"
  else
      echo "build miner"
      handle_build_node_commands
      echo "built miner"
  fi

  echo "$build_dir/miner --config-file="$current_dir/env/config_$1.env" &"
  ./build/miner --config-file="$current_dir/env/config_$1.env"
}

handle_stop_node_commands() {
  echo "stop miner"
  for pid in $(ps aux | grep "./build/miner" | awk '{print $2}'); do
      echo "kill $pid"
      kill -9 "$pid"
  done
  echo "stopped miner"
}

handle_build_node_commands() {
    check_golang
    echo "building miner..."
    create_file "miner" "$miner_dir"
}

#api
handle_build_api_commands() {
  check_golang
  echo "building api..."
  create_file "api" "$api_dir"
}

handle_start_api_commands() {
  echo "$api_dir"

  # Check if the file exists
  file="$build_dir/api"
  if [ -e "$file" ]; then
    echo "Run api"
  else
    echo "build api"
    handle_build_api_commands
    echo "built api"


  fi

  echo "start API"
  ./build/api --port=$api_port 
}

handle_stop_api_commands() {
  echo "$api_dir"
  cd "$api_dir"

  for pid in $(lsof -t -i:"$api_port"); do  echo "$pid"; kill -9 $pid; done
}

handle_api_commands() {
  case "$1" in
        "start")
          handle_start_api_commands
        ;;
        "build")
          handle_build_api_commands
        ;;
        "stop")
          handle_stop_api_commands
        ;;
        esac
}

handle_node_commands() {
  case "$1" in
          "start")
          handle_start_node_commands  "$2"
          ;;
          "build")
            handle_build_node_commands
          ;;
         "stop")
            handle_stop_node_commands
          ;;
        esac
}

check_golang() {
  # Check if Go is installed
  if command -v go &> /dev/null; then
      echo "Go is installed."
  else
      # Get the OS type
      os_type=$(uname)
      if [[ "$os_type" == "Darwin" ]]; then
          echo "You are using macOS."
          brew install go
          go version
      elif [[ "$os_type" == "Linux" ]]; then
          echo "You are using Linux."
          sudo apt install golang
          go version
      else
          echo "Unknown operating system: $os_type"
      fi
  fi
}

if [ $# -lt 1 ]; then
    echo "Usage:"
    echo "- miner node build"
    echo "- miner node start"
    echo "- miner node stop"
    echo "- miner api build"
    echo "- miner api start"
    echo "- miner api stop"
    exit 1
fi

#commands
case "$1" in
    "node")
      handle_node_commands "$2" "$3"
      ;;
    "api")
      handle_api_commands "$2"
    ;;
esac

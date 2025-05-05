FROM ubuntu:22.04

# Install dependencies
RUN apt-get update && apt-get install -y \
  build-essential \
  wget \
  curl \
  ca-certificates \
  && apt-get clean

COPY ./eai-chat-linux /app/eternal-ai/eai-chat
COPY default_config.json /app/eternal-ai/decentralized-inference/config.json
COPY default_local_contracts.json /app/eternal-ai/decentralized-compute/worker-hub/env/local_contracts.json

# Change the working directory to the desired folder
WORKDIR /app/eternal-ai

# Ensure that the eai-chat executable is present and is marked as executable
RUN chmod +x ./eai-chat

# Run the command
CMD ["./eai-chat", "server"]

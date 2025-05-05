# Use a base image with the ability to install specific Node.js version
FROM node:18

# Set the HOME environment variable explicitly
ENV HOME=/root

# Install Node Version Manager (NVM)
RUN curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.3/install.sh | bash

# Set NVM_DIR directly to the location where NVM installs Node.js
ENV NVM_DIR=/root/.nvm

# Install Node.js version 23.3.0 using NVM
RUN . $NVM_DIR/nvm.sh && \
    nvm install 23.3.0 && \
    nvm use 23.3.0 && \
    nvm alias default 23.3.0

# Make sure the correct node version is available in the PATH
# ENV PATH=$PATH:$HOME/.nvm/versions/node/v23.3.0/bin
ENV PATH=$NVM_DIR/versions/node/v23.3.0/bin:$PATH

# Set TERM environment variable to avoid warning
ENV TERM=xterm-256color

# RUN node -v

# Set the working directory inside the container
WORKDIR /app

COPY . /app

# # Clone Eternal AI code
# RUN git clone https://github.com/eternalai-org/eternal-ai.git

# COPY mintAgent.ts  /app/eternal-ai/developer-guides/examples/how-to-mint-agent/mintAgent.ts

# # Change the working directory to the desired folder
# WORKDIR /app/eternal-ai/developer-guides/examples/how-to-mint-agent

# # Install dependencies
# RUN npm install

# # Install ts-node and TypeScript globally
# RUN npm install -g ts-node typescript

# # Run the TypeScript script to mint agent
# # CMD ["ts-node", "./mintAgent.ts"]
# # RUN ts-node ./mintAgent.ts

# # Capture the output of ts-node and set it as environment variables
# RUN result=$(ts-node ./mintAgent.ts) && \
#     echo "Result: $result" && \
#     export MY_RESULT=$result && \
#     echo "MY_RESULT=$MY_RESULT" >> ~/.bashrc
   

# WORKDIR /app


# Get Eliza code 
RUN git clone https://github.com/ai16z/eliza.git

COPY config.json  /app/eliza/agents/config.json

# Change the working directory to the desired folder
WORKDIR /app/eliza

# Install dependencies using pnpm and build the project
RUN npm install -g pnpm
RUN pnpm clear 
RUN pnpm install --no-frozen-lockfile
RUN pnpm build


CMD ["sh", "-c", "pnpm start --characters=/app/eliza/agents/config.json"]

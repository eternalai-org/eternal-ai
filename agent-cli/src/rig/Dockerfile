# Use the official Rust image
FROM rust:1.75 AS builder

# Ensure Cargo is up-to-date
RUN rustup update stable && rustup default stable

WORKDIR /app
   
# Step 1: Get Rig code 
RUN git clone https://github.com/0xPlaygrounds/rig.git

# Step 2: Update model, chain ID in file agent_with_eternalai.rs

COPY . .

COPY agent_with_eternalai.rs  /app/rig/rig-eternalai/examples/agent_with_eternalai.rs

# Change the working directory to the desired folder
WORKDIR /app/rig

# Build your Rig agent
RUN cargo update


CMD ["sh", "-c", "cargo run --package rig-eternalai --example agent_with_eternalai"]

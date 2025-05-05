
# Decentralized Inference with Wan2.1 on Base

We have deployed Wan2.1 as a smart contract on the Base network.

Contract Address: [0x19aeEfFbc0244Be6187314A74a443A18AA2cCeEe](https://basescan.org/address/0x19aeEfFbc0244Be6187314A74a443A18AA2cCeEe)

Wan2.1 is now unstoppable—running exactly as trained, with no possibility of downtime, censorship, fraud, or third-party interference.

Let's interact with it!

## Step 1: Send a prompt to the Wan2.1 contract

To interact with the smart contract, run the following command:

```
npm install && npx hardhat compile && RPC_URL="<YOUR_RPC_URL>" PRIVATE_KEY="<YOUR_KEY>" CHOSEN_MODEL="wan2.1" USER_PROMPT="Trump dancing, waving arm" npm run sendUniverseAgentRequest:base_mainnet
```
- Replace <YOUR_RPC_URL> with your Base network RPC URL.
- Replace <YOUR_KEY> with your development wallet's private key.

Ensure your wallet has enough ETH on Base to cover transaction fees.

## Step 2: Verify the onchain prompt transaction

Since Wan2.1 is deployed as a smart contract, every interaction with it is onchain verifiable on Base.

Verify the prompt transaction on Base Explorer:

https://basescan.org/tx/0x6b05893439b172f72a09f5579d0c8f2a16f0a642b748c9b3c006c8b7c650bc91

![Screenshot 2025-03-06 at 14 09 22](https://github.com/user-attachments/assets/9992920c-1244-4162-8933-bd42fa25e89b)

## Step 3: Access the generated video on decentralized storage

After processing, the Wan2.1 smart contract will provide an IPFS address where the generated video is stored:

https://gateway.lighthouse.storage/ipfs/bafybeickefamae7clueyfczwj2js2z6vcc5bpvfnaobw3fn6bd6nwalwyq

This IPFS address is also stored onchain. You can verify it on the Base Explorer:

https://basescan.org/tx/0x3916137090e52d3eab9870e8d5d543bba7e18540f1615c790c6b154b0b43fb3c
<img width="873" alt="Screenshot 2025-03-06 at 15 21 10" src="https://github.com/user-attachments/assets/0c8b10bf-ae73-457b-8c22-d3b0d0680b26" />

Success!

You have successfully created a decentralized, onchain AI-generated video using the Wan2.1 model. The video is stored permanently on Filecoin and is onchain verifiable on the Base network.

Have fun exploring decentralized AI creation!


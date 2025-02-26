import { ethers } from 'ethers';
import { default as Interact } from '../src/interact';
import { default as InteractWithExternalWallet } from '../src/interactWithExternalWallet';
import { InferPayloadWithMessages, InferPayloadWithPrompt } from '../src/types';
import { ChainId } from '../src/constants';

export const AGENT_CONTRACT_ADDRESSES: Record<ChainId, string> = {
  [ChainId.BSC]: '0x3B9710bA5578C2eeD075D8A23D8c596925fa4625',
  [ChainId.BASE]: '0x1E65FCa9b6640bC87AE41f1a897762c334821D1C',
};

// const wallet = new ethers.Wallet("Your private key here");
const wallet = ethers.Wallet.createRandom();

const externalWallet = {
  address: wallet.address,
  getAddress: wallet.getAddress,
  requestSignature: async (
    transaction: ethers.providers.TransactionRequest
  ) => {
    return await wallet.signTransaction(transaction);
  },
};

async function testInferV1() {
  const inferPayload = {
    chainId: ChainId.BSC,
    agentAddress: AGENT_CONTRACT_ADDRESSES[ChainId.BSC],
    prompt: 'Can you tell me about BTC',
  } satisfies InferPayloadWithPrompt;
  {
    const interact = new Interact(wallet);
    await interact.infer(inferPayload);
  }
}

async function testInferV2() {
  const inferPayload = {
    chainId: ChainId.BSC,
    agentAddress: AGENT_CONTRACT_ADDRESSES[ChainId.BSC],
    messages: [
      {
        role: 'system',
        content: 'You are a BTC master',
      },
      {
        role: 'user',
        content: 'Can you tell me about BTC',
      },
    ],
    isLightHouse: true,
  } satisfies InferPayloadWithMessages;
  {
    const interact = new Interact(wallet);
    await interact.infer(inferPayload);
  }
}

async function testInferExternalWalletV1() {
  const inferPayload = {
    chainId: ChainId.BSC,
    agentAddress: AGENT_CONTRACT_ADDRESSES[ChainId.BSC],
    prompt: 'Can you tell me about BTC',
  } satisfies InferPayloadWithPrompt;
  {
    const interact = new InteractWithExternalWallet(externalWallet);
    await interact.infer(inferPayload);
  }
}

async function testInferExternalWalletV2() {
  const inferPayload = {
    chainId: ChainId.BSC,
    agentAddress: AGENT_CONTRACT_ADDRESSES[ChainId.BSC],
    messages: [
      {
        role: 'system',
        content: 'You are a BTC master',
      },
      {
        role: 'user',
        content: 'Can you tell me about BTC',
      },
    ],
    isLightHouse: true,
  } satisfies InferPayloadWithMessages;
  {
    const interact = new InteractWithExternalWallet(externalWallet);
    await interact.infer(inferPayload);
  }
}

testInferV1();

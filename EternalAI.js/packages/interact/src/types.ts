import * as ethers from 'ethers';
import { ChainId } from './constants';
import { Message } from './methods/infer/types';
import { InteractWallet } from './methods/types';

type InferPayloadBase = {
  agentAddress: string;
  isLightHouse?: boolean;
  rpcUrl?: string;
};

export type InferPayloadWithPrompt = InferPayloadBase & {
  chainId: ChainId;
  prompt: string;
};

export type InferPayloadWithMessages = InferPayloadBase & {
  chainId: ChainId;
  messages: Message[];
};

export type ExternalWallet = Pick<InteractWallet, 'getAddress' | 'address'> & {
  requestSignature: (
    transaction: ethers.ethers.providers.TransactionRequest
  ) => Promise<string>;
};

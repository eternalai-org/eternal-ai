import * as ethers from 'ethers';

import { CHAIN_MAPPING, ChainId } from './constants';
import { InferPayloadWithMessages, InferPayloadWithPrompt } from './types';

class BaseInteract {
  protected getProvider(chainId: ChainId, rpcUrl?: string) {
    // create provider from user optional
    if (!!rpcUrl) {
      return new ethers.providers.JsonRpcProvider(rpcUrl);
    }

    if (!CHAIN_MAPPING[chainId]) {
      throw new Error(`Unsupported chainId: ${chainId}`);
    }

    // create provider from default supported chainId
    return new ethers.providers.JsonRpcProvider(CHAIN_MAPPING[chainId]);
  }

  protected normalizePayload(
    payload: InferPayloadWithPrompt | InferPayloadWithMessages
  ) {
    return {
      ...payload,
      isLightHouse: payload.isLightHouse ?? false,
    };
  }
}

export interface IInteract {
  // getNetworkCredential(
  //   chainId: ChainId,
  //   rpcUrl?: string
  // ): {
  //   provider: ethers.providers.JsonRpcProvider;
  //   signer: ethers.ethers.Wallet | InteractWallet;
  // };
  infer(
    payload: InferPayloadWithPrompt | InferPayloadWithMessages
  ): Promise<string | null>;
  // Overload signatures
  // // @ts-ignore
  // public infer(payload: InferPayloadWithPrompt): Promise<string | null>;
  // // @ts-ignore
  // public infer(payload: InferPayloadWithMessages): Promise<string | null>;
}

export default BaseInteract;

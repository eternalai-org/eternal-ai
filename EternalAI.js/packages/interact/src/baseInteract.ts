import { CHAIN_MAPPING, ChainId } from './constants';
import { InferPayloadWithMessages, InferPayloadWithPrompt } from './types';
import * as methods from './methods';
import { InteractWallet } from './methods/types';

import injectDependency from '@/inject';
// this is inject supported packages
const packages = {
  ethers: injectDependency<InjectedTypes.ethers>('ethers'),
};

class BaseInteract {
  protected getProvider(chainId: ChainId, rpcUrl?: string) {
    // create provider from user optional
    if (!!rpcUrl) {
      return new packages.ethers.providers.JsonRpcProvider(rpcUrl);
    }

    if (!CHAIN_MAPPING[chainId]) {
      throw new Error(`Unsupported chainId: ${chainId}`);
    }

    // create provider from default supported chainId
    return new packages.ethers.providers.JsonRpcProvider(
      CHAIN_MAPPING[chainId]
    );
  }

  protected normalizePayload(
    payload: InferPayloadWithPrompt | InferPayloadWithMessages
  ) {
    return {
      ...payload,
      isLightHouse: payload.isLightHouse ?? false,
    };
  }

  protected async sendSignedTransactionAndListenResult(
    signer: InteractWallet,
    signedTx: string,
    agentAddress: string,
    chainId: ChainId
  ) {
    const sendPromptTxHash = await methods.Infer.sendPrompt(signer, signedTx);

    const workerHubAddress = await methods.Infer.getWorkerHubAddress(
      agentAddress,
      signer
    );

    return await methods.Infer.listenPromptResponse(
      chainId,
      signer,
      workerHubAddress,
      sendPromptTxHash
    );
  }
}

export interface IInteract {
  infer(
    payload: InferPayloadWithPrompt | InferPayloadWithMessages
  ): Promise<string | null>;
}

export default BaseInteract;

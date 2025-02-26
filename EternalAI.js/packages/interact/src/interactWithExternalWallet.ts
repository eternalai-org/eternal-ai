import * as ethers from 'ethers';
import {
  ExternalWallet,
  InferPayloadWithMessages,
  InferPayloadWithPrompt,
} from './types';
import * as methods from './methods';
import { CHAIN_MAPPING, ChainId } from './constants';
import { InteractWallet } from './methods/types';

class InteractWithExternalWallet {
  private _wallet: ExternalWallet;

  constructor(wallet: ExternalWallet) {
    this._wallet = wallet;
  }

  private getProvider(chainId: ChainId, rpcUrl?: string) {
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

  private getNetworkCredential(chainId: ChainId, rpcUrl?: string) {
    const provider = this.getProvider(chainId, rpcUrl);
    const signer = {
      ...this._wallet,
      provider,
    } satisfies InteractWallet;
    return {
      provider,
      signer,
    };
  }

  private normalizePayload(
    payload: InferPayloadWithPrompt | InferPayloadWithMessages
  ) {
    return {
      ...payload,
      isLightHouse: payload.isLightHouse ?? false,
    };
  }

  // Overload signatures
  // @ts-ignore
  public async infer(payload: InferPayloadWithPrompt): Promise<string | null>;
  // @ts-ignore
  public async infer(payload: InferPayloadWithMessages): Promise<string | null>;

  // Implementation
  // @ts-ignore
  public async infer(
    payload: InferPayloadWithPrompt | InferPayloadWithMessages
  ): Promise<string | null> {
    try {
      const normalizedPayload = this.normalizePayload(payload);
      console.log('inferWithExternalWallet - start', {
        payload: normalizedPayload,
      });
      if (
        typeof (normalizedPayload as InferPayloadWithPrompt).prompt === 'string'
      ) {
        const result = await this.inferWithPrompt(
          normalizedPayload as InferPayloadWithPrompt
        );
        console.log('inferWithExternalWallet - succeed', result);
        return result;
      } else {
        const result = await this.inferWithMessages(
          normalizedPayload as InferPayloadWithMessages
        );
        console.log('inferWithExternalWallet - succeed', result);
        return result;
      }
    } catch (e) {
      console.log('inferWithExternalWallet - failed', e);
      throw e;
    } finally {
      console.log('inferWithExternalWallet - end');
    }
  }

  private async inferWithPrompt(
    payload: InferPayloadWithPrompt
  ): Promise<string | null> {
    console.log('inferWithEternalWalletWithPrompt - start');
    const { signer } = this.getNetworkCredential(
      payload.chainId,
      payload.rpcUrl
    );

    const params = await methods.Infer.createPayloadWithPrompt(signer, payload);

    const signedTx = await signer.requestSignature(params);

    const sendPromptTxHash = await methods.Infer.sendPrompt(signer, signedTx);

    const workerHubAddress = await methods.Infer.getWorkerHubAddress(
      payload.agentAddress,
      signer
    );

    return await methods.Infer.listenPromptResponse(
      payload.chainId,
      signer,
      workerHubAddress,
      sendPromptTxHash
    );
  }

  private async inferWithMessages(
    payload: InferPayloadWithMessages
  ): Promise<string | null> {
    console.log('inferWithEternalWalletWithMessages - start');

    const { signer } = this.getNetworkCredential(
      payload.chainId,
      payload.rpcUrl
    );

    const params = await methods.Infer.createPayloadWithMessages(
      signer,
      payload
    );

    const signedTx = await signer.requestSignature(params);

    const sendPromptTxHash = await methods.Infer.sendPrompt(signer, signedTx);

    const workerHubAddress = await methods.Infer.getWorkerHubAddress(
      payload.agentAddress,
      signer
    );

    return await methods.Infer.listenPromptResponse(
      payload.chainId,
      signer,
      workerHubAddress,
      sendPromptTxHash
    );
  }
}

export default InteractWithExternalWallet;

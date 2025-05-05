import * as ethers from 'ethers';
import { InferPayloadWithMessages, InferPayloadWithPrompt } from './types';
import * as methods from './methods';
import { ChainId } from './constants';
import BaseInteract, { IInteract } from './baseInteract';

class Interact extends BaseInteract implements IInteract {
  private _wallet: ethers.Wallet;

  constructor(wallet: ethers.Wallet) {
    super();
    if (!ethers.Wallet.isSigner(wallet)) {
      throw new Error('Provided wallet is not a signer');
    }

    this._wallet = wallet;
  }

  private getNetworkCredential(chainId: ChainId, rpcUrl?: string) {
    const provider = this.getProvider(chainId, rpcUrl);
    const signer = this._wallet.connect(provider);
    return {
      provider,
      signer,
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
      console.log('infer - start', {
        payload: normalizedPayload,
      });
      if (
        typeof (normalizedPayload as InferPayloadWithPrompt).prompt === 'string'
      ) {
        const result = await this.inferWithPrompt(
          normalizedPayload as InferPayloadWithPrompt
        );
        console.log('infer - succeed', result);
        return result;
      } else {
        const result = await this.inferWithMessages(
          normalizedPayload as InferPayloadWithMessages
        );
        console.log('infer - succeed', result);
        return result;
      }
    } catch (e) {
      console.log('infer - failed', e);
      throw e;
    } finally {
      console.log('infer - end');
    }
  }

  private async inferWithPrompt(
    payload: InferPayloadWithPrompt
  ): Promise<string | null> {
    console.log('inferWithPrompt - start');
    const { signer } = this.getNetworkCredential(
      payload.chainId,
      payload.rpcUrl
    );

    const params = await methods.Infer.createPayloadWithPrompt(signer, payload);

    const signedTx = await signer.signTransaction(params);

    return await this.sendSignedTransactionAndListenResult(
      signer,
      signedTx,
      payload.agentAddress,
      payload.chainId
    );
  }

  private async inferWithMessages(
    payload: InferPayloadWithMessages
  ): Promise<string | null> {
    console.log('inferWithMessages - start');
    const { signer } = this.getNetworkCredential(
      payload.chainId,
      payload.rpcUrl
    );

    const params = await methods.Infer.createPayloadWithMessages(
      signer,
      payload
    );

    const signedTx = await signer.signTransaction(params);

    return await this.sendSignedTransactionAndListenResult(
      signer,
      signedTx,
      payload.agentAddress,
      payload.chainId
    );
  }
}

export default Interact;

import {
  ExternalWallet,
  InferPayloadWithMessages,
  InferPayloadWithPrompt,
} from './types';
import * as methods from './methods';
import { ChainId } from './constants';
import { InteractWallet } from './methods/types';
import BaseInteract, { IInteract } from './baseInteract';

class InteractWithExternalWallet extends BaseInteract implements IInteract {
  private _wallet: ExternalWallet;

  constructor(wallet: ExternalWallet) {
    super();
    this._wallet = wallet;
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

    return await this.sendSignedTransactionAndListenResult(
      signer,
      signedTx,
      payload.agentAddress,
      payload.chainId
    );
  }
}

export default InteractWithExternalWallet;

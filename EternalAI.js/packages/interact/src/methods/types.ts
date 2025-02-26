import * as ethers from 'ethers';

export interface InteractWallet {
  provider: ethers.providers.Provider;
  getAddress: () => Promise<string>;
  address: string;
}

export type InteractMethod = {
  createPayload: <P, R>(payload: P) => R;
  execute: <R>(signedTx: string) => Promise<R>;
};

import * as ethers from 'ethers';

if (typeof globalThis === 'undefined') {
  (globalThis as any) = global;
}

(globalThis as any).ethers = ethers;

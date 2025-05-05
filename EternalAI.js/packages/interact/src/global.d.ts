declare global {
  namespace InjectedTypes {
    type ethers = typeof import('ethers');
  }
}

export {};

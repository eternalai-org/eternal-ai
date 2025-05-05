// Ensure globalThis is defined
if (typeof globalThis === 'undefined') {
  (globalThis as any) = global;
}

const injectDependency = <T = any>(packageName: 'ethers'): T => {
  if ((globalThis as any)[packageName]) {
    return (globalThis as any)[packageName];
  }

  if (packageName === 'ethers') {
    (globalThis as any)[packageName] = require('ethers');
    return (globalThis as any)[packageName];
  }

  throw new Error(`Package ${packageName} not found`);
};

export default injectDependency;

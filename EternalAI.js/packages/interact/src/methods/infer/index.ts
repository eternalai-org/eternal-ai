import * as ethers from 'ethers';

import { InteractWallet } from '../types';
import { Message, SendInferResponse } from './types';
import {
  AGENT_ABI,
  IPFS,
  LIGHTHOUSE_IPFS,
  PROMPT_SCHEDULER_ABI,
  WORKER_HUB_ABI,
} from './constants';
import { ChainId } from '../../constants';
import { InferPayloadWithMessages, InferPayloadWithPrompt } from '../../types';
import { sleep } from '../../utils/time';
import { Fragment, LogDescription } from 'ethers/lib/utils';
import LightHouse from '@/services/light_house';

import injectDependency from '@/inject';
// this is inject supported packages
const packages = {
  ethers: injectDependency<InjectedTypes.ethers>('ethers'),
};

const contracts: Record<string, ethers.Contract> = {};

const getAgentContract = (contractAddress: string, wallet: InteractWallet) => {
  if (!contracts[contractAddress]) {
    contracts[contractAddress] = new packages.ethers.Contract(
      contractAddress,
      AGENT_ABI,
      wallet.provider
    );
  }
  return contracts[contractAddress];
};

const getWorkerHubContract = (
  contractAddress: string,
  wallet: InteractWallet
) => {
  if (!contracts[contractAddress]) {
    contracts[contractAddress] = new packages.ethers.Contract(
      contractAddress,
      PROMPT_SCHEDULER_ABI,
      wallet.provider
    );
  }
  return contracts[contractAddress];
};

export class InferenceResponse {
  result_uri: string;
  storage: string;
  data: string;

  constructor(result_uri: string, storage: string, data: string) {
    this.result_uri = result_uri;
    this.storage = storage;
    this.data = data;
  }

  static fromJSON(json: string): InferenceResponse {
    const parsed = JSON.parse(json);
    return Object.assign(new InferenceResponse('', '', ''), parsed);
  }
}

const Infer = {
  convertMessagesToBytes: async (
    messages: Message[],
    isLightHouse: boolean
  ) => {
    if (isLightHouse) {
      const uploadedUrl = await LightHouse.upload(JSON.stringify(messages));
      return packages.ethers.utils.toUtf8Bytes(uploadedUrl);
    }

    return packages.ethers.utils.toUtf8Bytes(
      JSON.stringify({
        messages,
      })
    );
  },
  getSystemPrompt: async (contractAddress: string, wallet: InteractWallet) => {
    try {
      console.log('infer getSystemPrompt - start');
      const contract = getAgentContract(contractAddress, wallet);
      const systemPrompt = await contract.getSystemPrompt();
      console.log('infer getSystemPrompt - succeed', systemPrompt);
      return systemPrompt;
    } catch (e) {
      console.log('infer getSystemPrompt - failed');
      throw e;
    } finally {
      console.log('infer getSystemPrompt - end');
    }
  },
  createPayloadWithPrompt: async (
    wallet: InteractWallet,
    payload: InferPayloadWithPrompt
  ) => {
    try {
      console.log('infer createPayloadWithPrompt - start', payload);
      const contractAddress = payload.agentAddress;
      const contract = getAgentContract(contractAddress, wallet);

      const systemPrompt = await Infer.getSystemPrompt(contractAddress, wallet);

      const { chainId, prompt, isLightHouse } = payload;

      const promptPayload = await Infer.convertMessagesToBytes(
        [
          {
            role: 'system',
            content: systemPrompt,
          },
          {
            role: 'user',
            content: prompt,
          },
        ],
        !!isLightHouse
      );
      const callData = contract.interface.encodeFunctionData('prompt(bytes)', [
        promptPayload,
      ]);

      const from = wallet.address || (await wallet.getAddress());
      const [gasLimit, gasPrice, nonce] = await Promise.all([
        contract.estimateGas.prompt(promptPayload),
        wallet.provider.getGasPrice(),
        wallet.provider.getTransactionCount(from),
      ]);

      const params = {
        to: contractAddress, // smart contract address
        from: from, // sender address
        data: callData, // data
        chainId: packages.ethers.BigNumber.from(chainId).toNumber(),
        gasLimit: gasLimit,
        gasPrice: gasPrice,
        nonce: nonce,
      } satisfies ethers.ethers.providers.TransactionRequest;
      console.log('infer createPayloadWithPrompt - succeed', params);
      return params;
    } catch (e) {
      console.log('infer createPayloadWithPrompt - failed');
      throw e;
    } finally {
      console.log('infer createPayloadWithPrompt - end');
    }
  },
  createPayloadWithMessages: async (
    wallet: InteractWallet,
    payload: InferPayloadWithMessages
  ) => {
    try {
      console.log('infer createPayloadWithMessages - start', payload);
      const contractAddress = payload.agentAddress;
      const contract = getAgentContract(contractAddress, wallet);

      const { chainId, messages, isLightHouse } = payload;

      const promptPayload = await Infer.convertMessagesToBytes(
        messages,
        !!isLightHouse
      );
      const callData = contract.interface.encodeFunctionData('prompt(bytes)', [
        promptPayload,
      ]);

      const from = wallet.address || (await wallet.getAddress());
      const [gasLimit, gasPrice, nonce] = await Promise.all([
        contract.estimateGas.prompt(promptPayload),
        wallet.provider.getGasPrice(),
        wallet.provider.getTransactionCount(from),
      ]);

      const params = {
        to: contractAddress,
        from: from,
        data: callData,
        chainId: packages.ethers.BigNumber.from(chainId).toNumber(),
        gasLimit: gasLimit,
        gasPrice: gasPrice,
        nonce: nonce,
      } satisfies ethers.ethers.providers.TransactionRequest;
      console.log('infer createPayloadWithMessages - succeed', params);
      return params;
    } catch (e) {
      console.log('infer createPayloadWithMessages - failed');
      throw e;
    } finally {
      console.log('infer createPayloadWithMessages - end');
    }
  },
  sendPrompt: async (
    wallet: InteractWallet,
    signedTx: string
  ): Promise<SendInferResponse> => {
    try {
      console.log('infer execute - start');

      console.log('infer execute - send transaction', signedTx);
      const txResponse = await wallet.provider.sendTransaction(signedTx);

      console.log('infer execute - waiting', txResponse);
      const receipt = await txResponse.wait();

      return receipt.transactionHash;
    } catch (e) {
      console.log('infer execute - failed');
      throw e;
    } finally {
      console.log('infer execute - end');
    }
  },
  getWorkerHubAddress: async (agentAddress: string, wallet: InteractWallet) => {
    try {
      console.log('infer getWorkerHubAddress - start', {
        agentAddress,
      });
      const contractAddress = agentAddress;
      const contract = getAgentContract(contractAddress, wallet);
      const schedule = await contract.getPromptSchedulerAddress();
      console.log('infer getWorkerHubAddress - succeed', schedule);
      return schedule;
    } catch (e) {
      console.log('infer getWorkerHubAddress - failed');
      throw e;
    } finally {
      console.log('infer getWorkerHubAddress - end');
    }
  },
  getInferId: async (wallet: InteractWallet, promptedTxHash: string) => {
    const txReceipt = await wallet.provider.getTransactionReceipt(
      promptedTxHash
    );

    if (!txReceipt || txReceipt.status != 1) {
      throw new Error('Transaction receipt not found.');
    } else {
      try {
        const iface = new packages.ethers.utils.Interface(
          WORKER_HUB_ABI as ReadonlyArray<Fragment>
        );

        const events = txReceipt.logs
          .map((log) => {
            try {
              return iface.parseLog(log);
            } catch (error) {
              return null;
            }
          })
          .filter((event) => event !== null);

        const newInference = events?.find(
          ((event: LogDescription) => event.name === 'NewInference') as any
        );
        return newInference?.args?.inferenceId;
      } catch (e) {
        throw new Error('No Infer Id');
      }
    }
  },

  processOutput: (out: any) => {
    const str: string = packages.ethers.utils.toUtf8String(out);
    try {
      const result = InferenceResponse.fromJSON(str);
      return result;
    } catch (e) {
      return null;
    }
  },

  processOutputToInferResponse: async (output: ethers.Bytes) => {
    const inferResponse = Infer.processOutput(output);
    if (!inferResponse) {
      return null;
    } else {
      if (
        inferResponse.storage == 'lighthouse-filecoint' ||
        inferResponse.result_uri.includes('ipfs://')
      ) {
        const light_house = inferResponse.result_uri.replace(
          IPFS,
          LIGHTHOUSE_IPFS
        );
        const light_house_reponse = await fetch(light_house);
        if (light_house_reponse.ok) {
          const result = await light_house_reponse.text();
          return result;
        }
        return null;
      } else {
        if (inferResponse.data != '') {
          const decodedString = atob(inferResponse.data);
          return decodedString;
        }
        return null;
      }
    }
  },
  getInferenceById: async (
    wallet: InteractWallet,
    workerHubAddress: string,
    inferId: string,
    chainId: ChainId
  ) => {
    if (chainId === ChainId.BSC) {
      const contract = getWorkerHubContract(workerHubAddress, wallet);
      const inferenceInfo = await contract.getInferenceInfo(inferId);
      const output = inferenceInfo[10];
      const bytesData = packages.ethers.utils.arrayify(output);
      if (bytesData.length != 0) {
        const result = await Infer.processOutputToInferResponse(bytesData);
        if (result) {
          return result;
        } else {
          return null;
        }
      } else {
        throw new Error(`waiting process inference ${inferId}`);
      }
    } else if (chainId === ChainId.BASE) {
      const contract = getWorkerHubContract(workerHubAddress, wallet);
      const assignIds = await contract.getInferenceInfo(inferId);
      if (assignIds.length == 0) {
        throw new Error('No assignment found');
      }
      const assignId = assignIds[0];
      const assignInfo = await contract.getAssignmentInfo(assignId);
      if (assignInfo.length == 0) {
        throw new Error('Inference result not ready');
      }
      const output = assignInfo[7];
      const bytesData = packages.ethers.utils.arrayify(output);
      if (bytesData.length != 0) {
        const result = await Infer.processOutputToInferResponse(bytesData);
        if (result) {
          return result;
        } else {
          return null;
        }
      } else {
        throw new Error(`waiting process inference ${inferId}`);
      }
    } else {
      throw Error('Unsupported chainId');
    }
  },
  listenPromptResponse: async (
    chainId: ChainId,
    wallet: InteractWallet,
    workerHubAddress: string,
    promptedTxHash: string
  ) => {
    try {
      console.log('infer listenPromptResponse - start', {
        chainId,
        workerHubAddress,
        promptedTxHash,
      });

      let result: string | null = null;
      const inferId = await Infer.getInferId(wallet, promptedTxHash);

      while (true) {
        try {
          result = await Infer.getInferenceById(
            wallet,
            workerHubAddress,
            inferId,
            chainId
          );
          break;
        } catch (e) {
          console.log('Retry to get inference by reference id');
          await sleep(30);
        }
      }

      console.log('infer listenPromptResponse - succeed', result);
      return result;
    } catch (e) {
      console.log('infer listenPromptResponse - failed');
      throw e;
    } finally {
      console.log('infer listenPromptResponse - end');
    }
  },
};

export default Infer;

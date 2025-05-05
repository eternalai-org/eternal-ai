import { ethers, network } from "hardhat";
import {
  IHybridModel,
  IPromptScheduler2TX,
  IPromptScheduler3TX,
} from "../typechain-types";
import { assert } from "chai";
import { BaseWallet, JsonRpcProvider, SigningKey } from "ethers";
import * as configJson from "./config.json";
import { Address } from "hardhat-deploy/dist/types";

interface Model {
  modelName: string;
  modelAddress: string;
  description: string;
}

interface NetworkConfig {
  promptSchedulerAddress: string;
  models: Model[];
}

type NetworkModels = {
  [key: string]: NetworkConfig;
};

const config = network.config as any;

async function sendUniverseAgentRequest() {
  const networkName = network.name.toUpperCase();
  let privateKey = config.senderKey;
  let rpcUrl = config.url;
  let chosenNetwork = network.name;
  let chosenModel = process.env.CHOSEN_MODEL;
  let userPrompt = process.env.USER_PROMPT;

  assert(
    privateKey,
    `Missing ${networkName}_PRIVATE_KEY from environment variables!`
  );
  assert(rpcUrl, `Missing ${networkName}_RPC_URL from environment variables!`);
  assert(chosenModel, `Missing CHOSEN_MODEL from environment variables!`);
  assert(chosenNetwork, `Missing CHOSEN_NETWORK from environment variables!`);
  assert(userPrompt, `Missing USER_PROMPT from environment variables!`);

  const sender = await createWallet(privateKey, rpcUrl);

  const inf = getModelInfo(chosenNetwork, chosenModel);
  const modelInstance = (await getContractInstance(
    inf.modelInfo.modelAddress,
    "IHybridModel"
  )) as unknown as IHybridModel;

  let request = "";
  if (chosenModel === "wan2.1") {
    request = userPrompt;
  } else {
    request = buildRequest(inf.modelInfo.modelName, userPrompt);
  }

  // Send inference request
  let feeData;
  try {
    feeData = await ethers.provider.getFeeData();
  } catch (error) {
    console.error("Error getting fee data:", error);
  }

  console.log("Sending inference request...");
  const txRequest = await modelInstance.connect(sender)["infer(bytes)"](
    ethers.toUtf8Bytes(JSON.stringify(request)),
    feeData
      ? {
          maxFeePerGas: feeData.maxFeePerGas, // Adjust this value as needed
          maxPriorityFeePerGas: feeData.maxPriorityFeePerGas, // Adjust this value as needed
        }
      : {}
  );
  const receipt = await txRequest.wait();
  console.log("Tx hash: ", receipt?.hash);
  console.log("Tx status: ", receipt?.status == 1 ? "Success" : "Failed");

  // Get inference result
  let inferResult;
  let inferId = 0;

  if (receipt?.status == 1) {
    // Get inference ID
    inferId = getInferId(receipt, inf.promptSchedulerAddress)[0];
    console.log("Inference ID: ", inferId);
    console.log("Wait for inference result...");

    // Wait for inference result
    while (true) {
      await sleep(20000);
      try {
        inferResult = await tryToGetInferenceResult(
          chosenNetwork,
          inf.promptSchedulerAddress,
          inferId,
          sender
        );

        break;
      } catch (e: any) {
        console.log(e.message);
        continue;
      }
    }
  }

  console.log("Inference result: ", ethers.toUtf8String(inferResult || ""));
}

export async function getContractInstance(
  proxyAddress: string,
  contractName: string
) {
  const contractIns = await ethers.getContractAt(contractName, proxyAddress);
  return contractIns;
}

async function getContractFromArtifact(addreess: string, artifact: any) {
  const contract = await ethers.getContractAt(artifact.abi, addreess);
  return contract;
}

function getModelInfo(network: string, modelName: string) {
  // Get array of models for the chosen network
  // const networkModels = (configJson as NetworkModels)[network];
  const networkInf = (configJson as unknown as NetworkModels)[network];

  if (!networkInf) {
    throw new Error(`Network ${network} not found in config`);
  }

  // Find matching model by name
  const modelInfo = networkInf.models.find(
    (model) => model.modelName === modelName
  );
  if (!modelInfo) {
    throw new Error(`Model ${modelName} not found for network ${network}`);
  }

  return {
    modelInfo,
    promptSchedulerAddress: networkInf.promptSchedulerAddress,
  };
}

function buildRequest(modelName: string, userPrompt: string): any {
  return {
    messages: [
      {
        role: "system",
        content: "You are a helpful assistant",
      },
      {
        role: "user",
        content: userPrompt,
      },
    ],
    max_tokens: 1024,
    model: modelName,
  };
}

async function createWallet(
  privateKey: string,
  rpcUrl: string
): Promise<BaseWallet> {
  const wallet = new BaseWallet(
    new SigningKey(privateKey),
    new JsonRpcProvider(rpcUrl)
  );

  return wallet;
}

export function getInferId(
  receipt: ethers.TransactionReceipt,
  promptSchedulerAddress: string
): number[] {
  return receipt.logs
    .filter(
      (log: any) =>
        log.topics[0] ===
          ethers.id("NewInference(uint256,address,address,uint256,uint256)") &&
        isAddressEq(log.address, promptSchedulerAddress)
    )
    .map((log: any) => {
      return parseInt(log.topics[1], 16);
    });
}

async function tryToGetInferenceResult(
  networkName: string,
  promptSchedulerAddress: string,
  inferId: number,
  sender: ethers.Signer
) {
  let inferResult;

  if (
    networkName == "base_mainnet" ||
    networkName == "symbiosis_mainnet" ||
    networkName == "ethereum_mainnet"
  ) {
    const ins = (await getContractInstance(
      promptSchedulerAddress,
      "IPromptScheduler3TX"
    )) as unknown as IPromptScheduler3TX;

    let assignIds = (
      await ins.connect(sender).getInferenceInfo(inferId)
    )[0] as bigint[];
    if (assignIds.length == 0) {
      throw new Error("No assignment found");
    }

    let assignId = assignIds[0];
    let assignInfo;

    if (networkName == "ethereum_mainnet") {
      const artifact = require("./artifacts/WorkerHub.json");
      const ethIns = await getContractFromArtifact(
        promptSchedulerAddress,
        artifact
      );
      assignInfo = await ethIns.connect(sender).assignments(assignId);
    } else {
      assignInfo = await ins.connect(sender).getAssignmentInfo(assignId);
    }
    let result = assignInfo[7];

    if (result.length == 0) {
      throw new Error("Inference result not ready");
    }

    return result;
  } else {
    const ins = (await getContractInstance(
      promptSchedulerAddress,
      "IPromptScheduler2TX"
    )) as unknown as IPromptScheduler2TX;

    let result = (await ins.getInferenceInfo(inferId))[10];
    if (result === undefined || result == "0x") {
      throw new Error("Inference result not ready");
    }

    return result;
  }
}

function isAddressEq(a: Address, b: Address): boolean {
  return a.toLowerCase() === b.toLowerCase();
}

async function sleep(ms: number): Promise<void> {
  return new Promise((resolve) => setTimeout(resolve, ms));
}

sendUniverseAgentRequest();

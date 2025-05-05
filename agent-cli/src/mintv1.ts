import { assert, ethers } from "ethers";
import { readFile } from "fs/promises";
import * as fs from "fs";
import * as path from "path";
import { logInfo } from "./log";

// interface NetworkConfig {
//     promptSchedulerAddress: string;
//     gpuManagerAddress: string;
// }

// interface Config {
//     [network: string]: NetworkConfig;
// }

// ABI for the mint function
const agentAbi = [
    // "function mint(address to,string calldata uri,bytes calldata data,uint256 fee,string calldata promptKey,address promptScheduler,uint32 modelId) external returns (uint256)",
    // "function mint(address to,string calldata uri,bytes calldata data,uint256 fee) external returns (uint256)",
    "function mint(address _to, string calldata _uri, bytes calldata _data, uint256 _fee) external payable returns (uint256)",
    "function mintPrice() external view returns (uint256)",
];

const gpuManagerAbi = [
    "function getModelIds() external view returns (uint256[] memory)",
];

const mintAgent = async ({
    rpcURL,
    privKey,
    agentSystemPrompPath,
    agentContractAddress,
    modelID,
    // promptSchedulerAddress,
    // gpuManagerAddress,
    agentURI = "EternalAIAgent",
    agentFee = "0",
    // agentPromptKey = "EternalAIAgent",

}: {
    rpcURL: string,
    privKey: string,
    agentSystemPrompPath: string,
    agentContractAddress: string,
    modelID: string,

    // promptSchedulerAddress: string,
    // gpuManagerAddress: string,
    agentURI?: string,
    agentFee?: string,
    // agentPromptKey?: string,
    networkName?: string,

}) => {



    // TODO: 2525 hard code to test
    // return 1925;


    // Get default config
    // const networkName = process.env.NETWORK_NAME;
    // assert(
    //     networkName,
    //     "Missing NETWORK_NAME environment variable",
    //     "INVALID_ARGUMENT"
    // );

    // const [df_promptSchedulerAddress, df_gpuManagerAddress] =
    //     getNetworkAddresses(networkName);
    // Configuration
    // const rpcURL = process.env.RPC_URL;
    // const privKey = process.env.PRIVATE_KEY;
    // const agentContractAddress = process.env.AGENT_CONTRACT_ADDRESS;
    // const agentSystemPrompPath = process.env.AGENT_SYSTEM_PROMPT_PATH;
    // const agentURI = process.env.AGENT_URI || "EternalAIAgent";
    // const agentFee = process.env.AGENT_FEE || "0";
    // const agentPromptKey = process.env.AGENT_PROMPT_KEY || "EternalAIAgent";
    // promptSchedulerAddress = promptSchedulerAddress || df_promptSchedulerAddress;
    // gpuManagerAddress = gpuManagerAddress || df_gpuManagerAddress;

    // const gpuManagerAddress = process.env.GPU_MANAGER_ADDRESS
    //     ? process.env.GPU_MANAGER_ADDRESS
    //     : df_gpuManagerAddress;
    // let modelID = process.env.MODEL_ID;

    assert(rpcURL, "Missing RPC_URL environment variable", "INVALID_ARGUMENT");
    assert(
        privKey,
        "Missing PRIVATE_KEY environment variable",
        "INVALID_ARGUMENT"
    );
    assert(
        agentContractAddress,
        "Missing AGENT_CONTRACT_ADDRESS environment variable",
        "INVALID_ARGUMENT"
    );
    assert(
        agentSystemPrompPath,
        "Missing AGENT_SYSTEM_PROMPT_PATH environment variable",
        "INVALID_ARGUMENT"
    );
    assert(
        agentURI,
        "Missing AGENT_URI environment variable",
        "INVALID_ARGUMENT"
    );
    assert(
        agentFee,
        "Missing AGENT_FEE environment variable",
        "INVALID_ARGUMENT"
    );
    // assert(
    //     agentPromptKey,
    //     "Missing AGENT_PROMPT_KEY environment variable",
    //     "INVALID_ARGUMENT"
    // );
    // assert(
    //     promptSchedulerAddress,
    //     "Missing PROMPT_SCHEDULER_ADDRESS environment variable",
    //     "INVALID_ARGUMENT"
    // );
    assert(
        modelID,
        "Missing modelID variable",
        "INVALID_ARGUMENT"
    );

    // Setup provider and signer
    const provider = new ethers.JsonRpcProvider(rpcURL);
    const wallet = new ethers.Wallet(privKey, provider);

    // if (!modelID) {
    //     assert(
    //         gpuManagerAddress,
    //         "Missing GPU_MANAGER_ADDRESS environment variable",
    //         "INVALID_ARGUMENT"
    //     );

    //     const gpuManager = new ethers.Contract(
    //         gpuManagerAddress,
    //         gpuManagerAbi,
    //         wallet
    //     );

    //     const modelIds = await gpuManager.getModelIds();
    //     console.log("Available models:", modelIds);
    //     if (modelIds.length === 0) {
    //         throw new Error("No models available");
    //     }
    //     modelID = modelIds[0];
    // }

    // Read system prompt from file
    let systemPrompt;
    try {
        systemPrompt = await readFile(agentSystemPrompPath, "utf8");
    } catch (error) {
        throw new Error(`Failed to read system prompt file: ${error}`);
    }

    // console.log("systemPrompt: ", systemPrompt);

    // logInfo(`Mint agent params:  ${{
    //     rpcURL,
    //     privKey,
    //     agentSystemPrompPath,
    //     agentContractAddress,
    //     modelID,
    //     // promptSchedulerAddress,
    //     // gpuManagerAddress,
    //     agentURI,
    //     agentFee,
    //     // agentPromptKey = "EternalAIAgent",

    // }}`)
    // logInfo(`rpcURL: ${rpcURL}`);
    // logInfo(`agentSystemPrompPath: ${agentSystemPrompPath}`);

    // logInfo(`agentContractAddress: ${agentContractAddress}`);
    // logInfo(`modelID: ${modelID}`);
    // logInfo(`agentURI: ${agentURI}`);
    // logInfo(`agentFee: ${agentFee}`);



    try {
        // Create contract instance
        const contract = new ethers.Contract(
            agentContractAddress,
            agentAbi,
            wallet
        );


        // Mint parameters
        const to = wallet.address; // mint to self
        const uri = agentURI.toString();
        const data = ethers.toUtf8Bytes(systemPrompt.toString()); // Convert string to bytes
        const fee = ethers.parseEther(agentFee); // Set agent usage fee
        // const mintPrice = await contract.mintPrice(); // Get mint price
        // const promptKey = agentPromptKey.toString();
        // const promptSchedulerAddress = promptSchedulerAddress.toString();
        const modelId = modelID;

        // console.log("Mint price: ", mintPrice);

        // logInfo(`data: ${data}`);


        // Call mint function
        const tx = await contract.mint(
            to,
            uri,
            data,
            0,
        );
        const receipt = await tx.wait();

        if (receipt?.status === 1) {
            console.log("Minting transaction sent:", receipt.hash);
            console.log("Transaction confirmed in block:", receipt.blockNumber);

            // Get minted token ID from events
            const event = receipt.logs?.find((e: ethers.Log) => {
                return (
                    e.topics[0] ===
                    ethers.id("NewToken(uint256,string,bytes,uint256,address)")
                );
            });
            if (event) {
                console.log("Minted Agent ID:", Number(event.topics[1]));
                return Number(event.topics[1]);
            }
        } else {
            console.error("Minting transaction failed:", receipt);
        }
    } catch (error) {
        console.error("Error minting agent:", error);
    }
}

// Execute
// mintAgent().catch(console.error);

export {
    mintAgent,
}

# How to use onchain LLM

In this code example, we’ll demonstrate how to interact with a LLM model deployed on multiple blockchains, in a few easy steps using TypeScript.

### Step 0

Obtain chain (id) and model (id) from [Onchain Models page](https://docs.eternalai.org/eternal-ai/decentralized-inference-api/onchain-models). (you will need these two information to run the following code in Step 3 below).

### **Step 1**

Call the `infer` function of the `HybridModel` contract with an inference prompt.

```typescript
const inf = getModelInfo(chosenNetwork, chosenModel);
const modelInstance = (await getContractInstance(
    inf.modelInfo.modelAddress,
    "IHybridModel"
)) as unknown as IHybridModel;

const request = buildRequest(inf.modelInfo.modelName, userPrompt || 'Hello, how are you?');
const txRequest = await modelInstance
    .connect(sender)
    ["infer(bytes)"](ethers.toUtf8Bytes(JSON.stringify(request)));
const receipt = await txRequest.wait();
console.log("Tx status: ", receipt?.status == 1 ? "Success" : "Failed");
```

### **Step 2**

Periodically check with the `PromptScheduler` contract to retrieve the response returned by the LLM model’s miners using the `inferId`.

```solidity
inferId = getInferId(receipt, inf.promptSchedulerAddress)[0];
while (true) {
    await sleep(30000);
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
console.log("Inference result: ", inferResult);
```

### Step 3

Complete example code can be found at: [https://github.com/kogisin/eternal-ai/blob/master/developer-guides/examples/universe-dagents/scripts/sendUniverseAgentRequest.ts](../../developer-guides/examples/universe-dagents/scripts/sendUniverseAgentRequest.ts)

You can run the code with the following command:

```
npx hardhat compile && RPC_URL=https://base-mainnet.infura.io/v3/eb492201628143a094aa7afaeb9f32d2  PRIVATE_KEY=<YOUR_KEY>  CHOSEN_MODEL="unsloth/Llama-3.3-70B-Instruct-bnb-4bit"  USER_PROMPT="Hello, how are you?"  npm run sendUniverseAgentRequest:base_mainnet
```

**Note**: replace \<YOUR\_KEY> with your actual wallet private key and the wallet should have some ETH on Base network for paying network fee.

Examples of the prompt request and the miner’s response transaction hashes can be found here:&#x20;

* Prompt tx: https://basescan.org/tx/0x641c26eff85f9486dace2d4ac0558b3c8da576b9e3e79773bb94a817fb8db45c
* Response tx: https://basescan.org/tx/0x133b94f00d31908c4b88b8d1b2602a1241bac7a7ee744c04a248ea3e6f47fc75






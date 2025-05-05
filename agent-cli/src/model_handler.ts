import { ChainIDMap, Network } from "./const";
import { logInfo } from "./log";
import { getSupportedModels } from "./utils";


const getModels = async (chain: Network) => {
    const chainID = ChainIDMap[chain];
    if (!chainID) {

    }
    const supportedModels = await getSupportedModels(chainID);

    logInfo(`List of supported chains: ${Object.keys(supportedModels)}`);
}

export {
    getModels
}
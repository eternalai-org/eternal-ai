import { exec, spawn } from 'child_process';
import { ChainIDMap, Framework, Network, NetworkConfig, ETERNALAI_URL } from "./const";
import * as readline from 'readline';  // Optional: For interactive user input

// for dev
import dotenv from 'dotenv';
import { execCmd, getSupportedModels, padWithPrefix } from "./utils";
import { mintAgent } from "./mintv1";
import { Agent, AgentStatus, getAgentByName, getAgents, insertAgent, updateAgentByName } from "./manager";
import { logError, logInfo, logSuccess, logTable } from "./log";
import { dataLength } from 'ethers';
dotenv.config();

const validateParams = async ({
    chain,
    framework,
    model,
    path,
    name,

}: {
    chain: Network,
    framework: string,
    model: string,
    path: string,
    name: string,
}): Promise<{
    model: string
}> => {

    // validate chain
    const chainID = ChainIDMap[chain];
    const supportedModels = await getSupportedModels(chainID);
    if (model) {
        if (!supportedModels[model]) {
            logError(`model ${model} is not supported in chain ${chain}`);
            // return { model };
        }
    } else {
        const models = Object.keys(supportedModels);
        if (models.length == 0) {
            logError(`no models supported in chain ${chain}`);

            // return { model };
        } else {
            model = models[0];
        }
    }

    // validate framework
    if (!Object.values(Framework).includes(framework as Framework)) {
        logError(`framework must be in supported list: ${Object.values(Framework)}`);
        // return { model };
    }

    return { model };
}

const createAgent = async ({
    chain,
    framework,
    model,
    path,
    name,

}: {
    chain: Network,
    framework: string,
    model: string,
    path: string,
    name: string,
}) => {
    // console.log("options: ", options);

    const agentInfo = await getAgentByName(name);
    if (agentInfo && agentInfo.AgentID) {
        logInfo(`Agent ${name} was created. Use 'eai agent ls' to see your agents.`);
        return;
    }

    const newParam = await validateParams({
        chain,
        framework,
        model,
        path,
        name,
    });
    model = newParam.model;

    // console.log("newParam: ", newParam);

    // get info from chain network
    // const chainRPC = "";
    const chainID = ChainIDMap[chain];
    const networkInfo = NetworkConfig[chainID];
    if (!networkInfo || !networkInfo.agentContractAddress || !networkInfo.url) {
        throw new Error("invalid chain argument")
    }

    const ETERNALAI_RPC_URL = networkInfo.url;
    const ETERNALAI_AGENT_CONTRACT_ADDRESS = networkInfo.agentContractAddress;


    // call contract to mint agent id
    const agentID = await mintAgent({
        rpcURL: ETERNALAI_RPC_URL,
        privKey: process.env.PRIVATE_KEY || "",
        agentSystemPrompPath: path,
        agentContractAddress: ETERNALAI_AGENT_CONTRACT_ADDRESS,
        modelID: model,
        // promptSchedulerAddress: networkInfo.promptSchedulerAddress,
        // gpuManagerAddress: networkInfo.gpuManagerAddress
    });

    if (!agentID) {
        console.error("Mint agent failed! Please retry.");
        return;
    }

    logSuccess(`Create agent successfully: Your Agent ID: ${agentID}`);

    const ETERNALAI_API_KEY = process.env.ETERNALAI_API_KEY;
    const TWITTER_USERNAME = process.env.TWITTER_USERNAME;
    const TWITTER_PASSWORD = process.env.TWITTER_PASSWORD;
    const TWITTER_EMAIL = process.env.TWITTER_EMAIL;
    const TWITTER_TARGET_USERS = process.env.TWITTER_TARGET_USERS;


    // const randomId = randomID();
    const AGENT_UID = getAgentUID(chain, agentID.toString());
    const agentName = name || AGENT_UID;

    // let scriptPath = "";
    // switch (framework) {
    //     case Framework.Eliza: {
    //         // Path to your Bash script
    //         scriptPath = `sh src/eliza/start.sh ${AGENT_UID} ${ETERNALAI_URL} ${ETERNALAI_API_KEY} ${chainID} ${ETERNALAI_RPC_URL} ${ETERNALAI_AGENT_CONTRACT_ADDRESS} ${agentID} ${model} ${TWITTER_USERNAME} ${TWITTER_PASSWORD} ${TWITTER_EMAIL} ${TWITTER_TARGET_USERS} ${agentName}`;

    //         // Run the Bash script

    //         // exec(scriptPath, (error: any, stdout: any, stderr: any) => {
    //         //     if (error) {
    //         //         console.error(`Error executing script: ${error.message}`);
    //         //         // throw error;
    //         //     }
    //         //     if (stderr) {
    //         //         console.error(`stderr: ${stderr}`);
    //         //     }
    //         //     console.log(`stdout: ${stdout}`);
    //         // });
    //         break;

    //     }
    //     case Framework.Rig: {
    //         // Path to your Bash script
    //         scriptPath = `sh src/rig/start.sh ${AGENT_UID} ${ETERNALAI_URL} ${ETERNALAI_API_KEY} ${chainID} ${ETERNALAI_RPC_URL} ${ETERNALAI_AGENT_CONTRACT_ADDRESS} ${agentID} ${model} ${TWITTER_USERNAME} ${TWITTER_PASSWORD} ${TWITTER_EMAIL} ${TWITTER_TARGET_USERS} ${agentName}`;

    //         // Run the Bash script
    //         // exec(scriptPath, (error: any, stdout: any, stderr: any) => {
    //         //     if (error) {
    //         //         console.error(`Error executing script: ${error.message}`);
    //         //         // throw error;
    //         //     }
    //         //     if (stderr) {
    //         //         console.error(`stderr: ${stderr}`);
    //         //     }
    //         //     console.log(`stdout: ${stdout}`);
    //         // });
    //         break;
    //     }
    //     case Framework.EternalAI: {
    //         // Path to your Bash script
    //         scriptPath = `sh src/eternalai/start.sh ${AGENT_UID} ${ETERNALAI_URL} ${ETERNALAI_API_KEY} ${chainID} ${ETERNALAI_RPC_URL} ${ETERNALAI_AGENT_CONTRACT_ADDRESS} ${agentID} ${model} ${TWITTER_USERNAME} ${TWITTER_PASSWORD} ${TWITTER_EMAIL} ${TWITTER_TARGET_USERS} ${agentName}`;

    //         // Run the Bash script
    //         // exec(scriptPath, (error: any, stdout: any, stderr: any) => {
    //         //     if (error) {
    //         //         console.error(`Error executing script: ${error.message}`);
    //         //         // throw error;
    //         //     }
    //         //     if (stderr) {
    //         //         console.error(`stderr: ${stderr}`);
    //         //     }
    //         //     console.log(`stdout: ${stdout}`);
    //         // });
    //         break;
    //     }
    // }

    // try {
    //     const stdout = await execCmd(scriptPath);
    // } catch (e) {
    //     logError(`Start agent error ${e}`);
    // }


    const newAgent: Agent = {
        AgentID: agentID.toString(),
        Name: agentName,
        Framework: framework,
        Network: chain,
        ChainID: chainID,
        Model: model,
        Status: AgentStatus.CREATED,
        CreateAt: "",
        ContainerID: "",
    }

    // insert into the file to manage agents
    await insertAgent(newAgent);
}


const startAgent = async ({
    name,
}: {
    name: string,
}) => {
    // console.log("options: ", options);

    const agentInfo = await getAgentByName(name);
    if (!agentInfo) {
        logError(`Agent ${name} was not created. Use 'eai agent create' to create your agent first.`);
        return;
    }

    if (agentInfo.Status == AgentStatus.RUNNING) {
        logInfo(`Agent ${name} is running currently.`);
        return;
    }

    // const framework = agentInfo.Framework;
    // const chain = agentInfo.Network;
    // const agentID = agentInfo.AgentID;
    // const model = agentInfo.AgentID

    const { Framework: framework, Network: chain, AgentID: agentID, Model: model, ContainerID: containerID } = agentInfo;
    console.log(`Start agent ${containerID}`);


    // get info from chain network
    const chainID = agentInfo.ChainID;
    const networkInfo = NetworkConfig[chainID];
    if (!networkInfo || !networkInfo.agentContractAddress || !networkInfo.url) {
        throw new Error("invalid chain argument")
    }

    const ETERNALAI_RPC_URL = networkInfo.url;
    const ETERNALAI_AGENT_CONTRACT_ADDRESS = networkInfo.agentContractAddress;

    const ETERNALAI_API_KEY = process.env.ETERNALAI_API_KEY;
    const TWITTER_USERNAME = process.env.TWITTER_USERNAME;
    const TWITTER_PASSWORD = process.env.TWITTER_PASSWORD;
    const TWITTER_EMAIL = process.env.TWITTER_EMAIL;
    const TWITTER_TARGET_USERS = process.env.TWITTER_TARGET_USERS;


    // const randomId = randomID();
    const AGENT_UID = getAgentUID(chain, agentID.toString());
    const agentName = name || AGENT_UID;

    const servicePort = padWithPrefix(agentID.toString(), "5", 5);

    let scriptPath = "";
    if (containerID) {
        scriptPath = `docker start ${containerID}`;
    } else {
        switch (framework) {
            case Framework.Eliza: {
                // Path to your Bash script
                scriptPath = `sh src/eliza/start.sh ${AGENT_UID} ${ETERNALAI_URL} ${ETERNALAI_API_KEY} ${chainID} ${ETERNALAI_RPC_URL} ${ETERNALAI_AGENT_CONTRACT_ADDRESS} ${agentID} ${model} ${TWITTER_USERNAME} ${TWITTER_PASSWORD} ${TWITTER_EMAIL} ${TWITTER_TARGET_USERS} ${agentName}`;
                break;
            }
            case Framework.Rig: {
                // Path to your Bash script
                scriptPath = `sh src/rig/start.sh ${AGENT_UID} ${ETERNALAI_URL} ${ETERNALAI_API_KEY} ${chainID} ${ETERNALAI_RPC_URL} ${ETERNALAI_AGENT_CONTRACT_ADDRESS} ${agentID} ${model} ${TWITTER_USERNAME} ${TWITTER_PASSWORD} ${TWITTER_EMAIL} ${TWITTER_TARGET_USERS} ${agentName}`;
                break;
            }
            case Framework.EternalAI: {
                // Path to your Bash script
                scriptPath = `sh src/eternalai/start.sh ${AGENT_UID} ${ETERNALAI_URL} ${ETERNALAI_API_KEY} ${chainID} ${ETERNALAI_RPC_URL} ${ETERNALAI_AGENT_CONTRACT_ADDRESS} ${agentID} ${model} ${TWITTER_USERNAME} ${TWITTER_PASSWORD} ${TWITTER_EMAIL} ${TWITTER_TARGET_USERS} ${agentName} ${servicePort}`;
                break;
            }
        }
    }

    try {
        await execCmd({ cmd: scriptPath });
        // console.log(stdout);
    } catch (e) {
        logError(`Start agent error ${e}`);
    }

    // update status agent
    // agentInfo.Status = AgentStatus.RUNNING;
    // await updateAgentByName(agentName, agentInfo);
}


const chatAgent = async ({
    name,
}: {
    name: string,
}) => {
    // console.log("options: ", options);

    const agentInfo = await getAgentByName(name);
    if (!agentInfo) {
        logError(`Agent ${name} was not created. Use 'eai agent create' to create your agent first.`);
        return;
    }

    if (agentInfo.Status != AgentStatus.RUNNING) {
        logInfo(`Agent ${name} is not running. Use 'eai agent start' to start your agent first`);
        return;
    }

    if (agentInfo.Framework != Framework.EternalAI) {
        logInfo(`Agent ${name} framework is ${agentInfo.Framework} which does not support chat usecase.`);
        return;
    }

    const { Framework: framework, Network: chain, AgentID: agentID, Model: model } = agentInfo;

    // get info from chain network
    const chainID = agentInfo.ChainID;
    const networkInfo = NetworkConfig[chainID];
    if (!networkInfo || !networkInfo.agentContractAddress || !networkInfo.url) {
        throw new Error("invalid chain argument")
    }

    const ETERNALAI_RPC_URL = networkInfo.url;
    const ETERNALAI_AGENT_CONTRACT_ADDRESS = networkInfo.agentContractAddress;

    const AGENT_UID = getAgentUID(chain, agentID.toString());
    const agentName = name || AGENT_UID;


    // console.log('Current working directory:', process.cwd());

    switch (framework) {
        case Framework.EternalAI: {
            // Define the path to the command and arguments
            const configPath = `./agent-cli/agents/${AGENT_UID}/local_contracts.json`
            // const chatConfigPath = `./agent-cli/agents/${AGENT_UID}/config.json`
            const command = 'sh';
            const args = ['./src/eternalai/chat.sh', agentID, configPath];

            // Spawn the child process
            const child = spawn(command, args);

            // Create a readline interface to get user input (optional)

            if (process.stdin.isTTY)
                process.stdin.setRawMode(true);

            const rl = readline.createInterface({
                input: process.stdin,
                output: process.stdout
            });
            // Disable input echoing
            // readline.input.setRawMode(true);  // This prevents the terminal from showing extra characters
            // readline.emitKeypressEvents(process.stdin);

            // Listen for standard output (stdout)
            child.stdout.on('data', (data) => {
                process.stdout.write(data);
            });

            // Listen for standard error (stderr)
            child.stderr.on('data', (data) => {
                console.error(`${data.toString()}`);
            });

            // Handle user input and send it to stdin of the child process
            rl.on('line', (input) => {
                child.stdin.write(input + '\n');  // Write input to the child process
            });

            // Listen for the child process exiting
            child.on('exit', (code) => {
                console.log(`Child process exited with code ${code}`);
                rl.close();  // Close readline when the process exits
            });

            // Optionally handle if there's no input (end of stdin)
            child.on('close', () => {
                console.log('Child process closed');
            });
            break;
        }
    }
}

const listAgents = async () => {
    const agents = await getAgents();
    logTable(agents);
}


const stopAgent = async ({
    name,
}: {
    name: string,
}) => {
    const agents = await getAgents();
    const agentMap = Object.fromEntries(agents.map((a: Agent) => [a.Name, a]));
    const agentInfo = agentMap[name];
    if (!agentInfo) {
        logError(`Agent name ${name} is not found.`);
        return;
    }
    if (agentInfo.Status == AgentStatus.STOPPED) {
        logInfo(`Agent name ${name} was stopped.`);
        return;
    }
    if (agentInfo.Status == AgentStatus.CREATED) {
        logInfo(`Agent name ${name} is not running.`);
        return;
    }

    const agentUID = getAgentUID(agentInfo.Network, agentInfo.AgentID);

    try {
        const stdout = await exec(`docker stop ${agentUID}`);
        // console.log(stdout);
        logSuccess(`Agent ${name} is stopped successfully.`);
    } catch (e) {
        logError(`Stop agent ${name} error ${e}`);
    }
}

const startAgentV2 = async ({
    name,
}: {
    name: string,
}) => {
    const agents = await getAgents();
    const agentMap = Object.fromEntries(agents.map((a: Agent) => [a.Name, a]));
    const agentInfo = agentMap[name];
    if (!agentInfo) {
        logError(`Agent ${name} is not found.`);
        return;
    }
    if (agentInfo.Status == AgentStatus.RUNNING) {
        logInfo(`Agent ${name} is running.`);
        return;
    }
    if (!agentInfo.ContainerID) {
        logInfo(`No docker container for agent ${name}.`);
        // TODO: create and start docker container for agent.
    }

    // console.log("ContainerID: ", agentInfo.ContainerID);

    try {
        const stdout = await exec(`docker start ${agentInfo.ContainerID}`);
        // console.log(stdout);
        logSuccess(`Agent ${name} is started successfully.`);
    } catch (e) {
        logError(`Start agent ${name} error ${e}`);
    }
}





// agentUID is a unique key to identify the agent
// it is also used as the agent's docker container name
const getAgentUID = (chain: string, agentID: string): string => {
    return `${chain}_${agentID}`;

}
export {
    createAgent,
    listAgents,
    stopAgent,
    startAgent,
    chatAgent,
}
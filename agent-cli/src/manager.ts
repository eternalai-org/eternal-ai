import * as fs from 'fs';
import path from 'path';
import { exec } from 'child_process';
import { logError, logInfo } from './log';
import { execCmd } from './utils';

enum AgentStatus {
    CREATED = "Created",  // minted agent id on the contract, but haven't started agent container yet
    RUNNING = "Running",  // agent container is running
    STOPPED = "Stopped",  // agent container is exited
}

interface Agent {
    AgentID: string
    Name: string
    Framework: string
    Network: string
    ChainID: string
    Model: string
    Status: string
    CreateAt: string
    ContainerID: string,
}

const FilePath = "./agents/list.json";


const insertAgent = async (agent: Agent) => {
    try {
        // logInfo(`FilePath: ${FilePath}`);

        if (!fs.existsSync(FilePath)) {
            fs.writeFileSync(FilePath, JSON.stringify([agent], null, 2))

        } else {
            const data = fs.readFileSync(FilePath, 'utf8');
            const currentAgents: Agent[] = JSON.parse(data);
            currentAgents.push(agent);
            fs.writeFileSync(FilePath, JSON.stringify(currentAgents, null, 2))
        }
    } catch (e) {
        logError(`Save agent info ${e}`);
    }
}

const getAgents = async (): Promise<Agent[]> => {
    if (!fs.existsSync(FilePath)) {
        return [];
    }

    const data = fs.readFileSync(FilePath, 'utf8');
    const currentAgents: Agent[] = JSON.parse(data);
    // console.log("currentAgents: ", currentAgents);

    const containers = await listRunningContainers();
    // console.log("containers: ", containers);

    for (let i = 0; i < currentAgents.length; i++) {
        let agent = currentAgents[i];
        const containerName = agent.Network + "_" + agent.AgentID;
        const containerInfo = containers[containerName];
        if (containerInfo) {
            if (containerInfo.State === "running") {
                agent.Status = AgentStatus.RUNNING;
            } else if (containerInfo.State === "exited") {
                agent.Status = AgentStatus.STOPPED;
            }

            // TODO: check more state

            agent.CreateAt = containerInfo.CreatedAt;
            agent.ContainerID = containerInfo.ID;
        }
        currentAgents[i] = agent;
    }

    return currentAgents;
}

const getAgentByName = async (name: string): Promise<Agent> => {
    const agents = await getAgents();
    const agentMap = Object.fromEntries(agents.map((a: Agent) => [a.Name, a]));
    const agentInfo = agentMap[name];
    return agentInfo;
}

const updateAgentByName = async (name: string, updatedAgent: Agent) => {
    const agents = await getAgents();
    agents.map(agent => {
        agent.Name == name ? updatedAgent : agent
    });
    try {
        fs.writeFileSync(FilePath, JSON.stringify(agents, null, 2))
    } catch (e) {
        logError(`Update agent ${name} error ${e}`);
    }
}

interface DockerContainer {
    Command: string
    CreatedAt: string
    ID: string
    Image: string
    Labels: string
    LocalVolumes: string
    Mounts: string
    Names: string
    Networks: string
    Ports: string
    RunningFor: string
    Size: string
    State: string   // running
    Status: string  // time
}


// Running Containers: [
//     {
//         Command: '"docker-entrypoint.s…"',
//         CreatedAt: '2025-02-13 16:18:33 +0700 +07',
//         ID: '2a04d8ee6843',
//         Image: 'eliza',
//         Labels: 'desktop.docker.io/binds/0/Source=/Users/macbookpro/go/src/eternalai-org/eternal-ai/agent-cli/agents/base_4/config.json,desktop.docker.io/binds/0/SourceKind=hostFile,desktop.docker.io/binds/0/Target=/app/eliza/agents/config.json',
//         LocalVolumes: '0',
//         Mounts: '/host_mnt/User…',
//         Names: 'base_4',
//         Networks: 'bridge',
//         Ports: '',
//         RunningFor: '31 seconds ago',
//         Size: '2.54MB (virtual 11.7GB)',
//         State: 'running',
//         Status: 'Up 30 seconds'
//     }
// ]


const listRunningContainers = async (): Promise<Record<string, DockerContainer>> => {
    const stdout = await execCmd({ cmd: 'docker ps -a --format "{{json .}}"', isLog: false });
    if (stdout.length == 0) {
        return {};
    }

    // Split the output into JSON objects (one per line)
    const containers: DockerContainer[] = stdout
        .trim()
        .split('\n')
        .map((line: any) => JSON.parse(line));


    const m = Object.fromEntries(containers.map((c: DockerContainer) => [c.Names, c]));
    return m;
}

export {
    Agent,
    insertAgent,
    getAgents,
    listRunningContainers,
    AgentStatus,
    getAgentByName,
    updateAgentByName,
}
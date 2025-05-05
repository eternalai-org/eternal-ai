import { Command } from 'commander';
import { Framework, Model, Network } from './const';
import { chatAgent, createAgent, listAgents, startAgent, stopAgent } from './handler';
import { getModels } from './model_handler';

// Initialize the commander object
const program = new Command();

// Define the CLI command and its options
program
    .name('eai')
    .description('You can deploy and manage your decentralized agents flexibly and simply - on any chain, any framework.')
    .version('1.0.0');


// Define a parent command: agent
const agentCmd = program
    .command('agent')
    .description('Deploy and manage agents');


// cmd: agent create
agentCmd
    .command('create')
    .description('Create a new agent')
    .option('-c, --chain <chain>', 'The blockchain which the new agent will be deployed on.', Network.Base)
    .option('-f, --framework <framework>', 'The framework is used for new agent.', Framework.EternalAI)
    .option('-m, --model <model>', 'The agent\'s model.')
    .requiredOption('-p, --path <path>', 'The path of the agent\'s character file.')
    .option('-n, --name <name>', 'The agent\'s name.')
    .action((options) => {
        createAgent(options);
    });


// cmd: agent ls
agentCmd
    .command('ls')
    .description('See list of agents')
    .action(() => {
        listAgents();
    });

// cmd: agent stop
agentCmd
    .command('stop')
    .description('Stop the agent')
    .requiredOption('-n, --name <name>', 'The agent\'s name.')
    .action((options) => {
        stopAgent(options);
    });


// cmd: agent start
agentCmd
    .command('start')
    .description('Start the agent')
    .requiredOption('-n, --name <name>', 'The agent\'s name.')
    .action((options) => {
        startAgent(options);
    });

// cmd: agent chat
agentCmd
    .command('chat')
    .description('Chat to the agent')
    .requiredOption('-n, --name <name>', 'The agent\'s name.')
    .action((options) => {
        chatAgent(options);
    });


// Define a parent command: agent
const modelCmd = program
    .command('model')
    .description('Get models\' info.');

modelCmd
    .command('ls')
    .description('List of supported models.')
    .requiredOption('-c, --chain <chain>', '')
    .action((options) => {
        getModels(options);
    });


// Parse the command-line arguments
program.parse(process.argv);

enum Framework {
    EternalAI = "eternalai",
    Eliza = "eliza",
    Rig = "rig",
}

enum Network {
    // EternalAI = "eternalai",
    Symbiosis = "symbiosis",
    BitAI = "bitai",
    Base = "base",
    Ethereum = "ethereum",
    Zksync = "zksync",
    Arbitrum = "arbitrum",
    Polygon = "polygon",
    Avax = "avax",
    BSC = "bsc",
    Tron = "tron",
    Ape = "ape",
    Bittensor = "bittensor",
    Duck = "duck",
    Mode = "mode",
}

enum Model {
    DeepSeek = "DeepSeek-R1-Distill-Llama-70B"
}

const ChainIDMap = {
    // [Network.EternalAI]: "45762",  // todo:

    [Network.Symbiosis]: "45762",
    [Network.BitAI]: "222671",
    [Network.Base]: "8453",
    [Network.Ethereum]: "1",
    [Network.Zksync]: "324",
    [Network.Arbitrum]: "42161",
    [Network.Polygon]: "137",
    [Network.Avax]: "43114",
    [Network.BSC]: "56",
    [Network.Tron]: "728126428",
    [Network.Ape]: "33139",
    [Network.Bittensor]: "964",
    [Network.Duck]: "5545",
    [Network.Mode]: "34443",

}


const ETERNALAI_URL = "https://api.eternalai.org/v1";

interface Config {
    agentContractAddress: string
    // promptSchedulerAddress: string
    // gpuManagerAddress: string
    url: string
}

const NetworkConfig: Record<string, Config> = {
    "45762": {
        "agentContractAddress": "0x5799F6349D7E9DAeD0d5c7f90F5467eC929cc89e",
        "url": "https://rpc.hermeschain.eternalai.org"
    },
    "8453": {  // base mainnet
        // "agentContractAddress": "0x458bE45957F8f29bBf597d5a953097c4095D9231",
        // "promptSchedulerAddress": "0x963691C0b25a8d0866EA17CefC1bfBDb6Ec27894",
        // "gpuManagerAddress": "0x14A008005cfa25621dD48E958EA33d14dd519d0d",
        "agentContractAddress": "0xAed016e060e2fFE3092916b1650Fc558D62e1CCC",  // old version
        "url": "https://base-mainnet.infura.io/v3/eb492201628143a094aa7afaeb9f32d2"
    },
    "bitAI_mainnet": {
        "agentContractAddress": "0x7734c3cd8B3239eA03A8A660095d94183FE63fCD",
        "url": "https://rpc.shard-ai.l2aas.com"
    },
    "1": {
        "agentContractAddress": "0xDdf1720c9689e4e0bf0B383E57b621f12886516C",
        "url": "https://mainnet.infura.io/v3/eb492201628143a094aa7afaeb9f32d2"
    },
    "324": {
        "agentContractAddress": "0xF721bEd9afFc1E584FdD9d8e6d3A5D6540E5D11f",
        "url": "https://mainnet.era.zksync.io"
    },
    "42161": {
        "agentContractAddress": "0x0244f98CFeb64DF810a894726FAaE3e6Fb959c3a",
        "url": "https://arbitrum-mainnet.infura.io/v3/eb492201628143a094aa7afaeb9f32d2"
    },
    "137": {
        "agentContractAddress": "0x472C6f40853C6E83795C50f724e657059f35Ddb1",
        "url": "https://polygon-mainnet.infura.io/v3/eb492201628143a094aa7afaeb9f32d2"
    },
    "43114": {
        "agentContractAddress": "0x4AE5Db75b261108f4a586053805094ED54a14436",
        "url": "https://avalanche-mainnet.infura.io/v3/eb492201628143a094aa7afaeb9f32d2"
    },
    "56": {
        "agentContractAddress": "0x08E3Da99728979024B4973112Ef3F56CBa4D5172",
        "url": "https://bsc-dataseed1.defibit.io"
    },
    "728126428": { // tron_mainnet
        "agentContractAddress": "0xf5336a5785bae1a2a674f1881d81e64e6c01f534",
        "url": "https://api.trongrid.io"
    },
    "33139": {
        "agentContractAddress": "0x3482FA153D1c92c3a686BdbF4F2081728908E838",
        "url": "https://apechain.drpc.org"
    },
    // "abstract_testnet": {
    //     "promptSchedulerAddress": "0x19D5Fb8BFbcBBE0f4E47fC3F7cE37035C16F33EC",
    //     "url": "https://api.testnet.abs.xyz"
    // },
    "964": {
        "agentContractAddress": "0x88feb137Ab8f971df6b8C098A2C05A22d8F480e7",
        "url": "https://evm-subtensor.eternalai.org"
    },
    "5545": {
        "agentContractAddress": "0xD730878c7C96eF99866A690b56D51059e3220805",
        "url": "https://rpc.duckchain.io"
    },
    "34443": {
        "agentContractAddress": "0xF442B5c96C76fF7F5EAaBd5708504CFB17AE5E57",
        "url": "https://mainnet.mode.network",
    }
}


export {
    Framework,
    Network,
    Model,
    ETERNALAI_URL,
    NetworkConfig,
    Config,
    ChainIDMap,

}
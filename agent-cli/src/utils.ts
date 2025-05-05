import { ETERNALAI_URL } from "./const";
import axios from 'axios';
import { exec } from 'child_process';



const getSupportedModels = async (chainID: string) => {
    const url = "https://api.eternalai.org/api/chain-config/get";
    const response = await axios.get(url, { params: { chain_id: chainID } });
    // console.log('Filtered API Data:', response.data);

    if (response.data.status != 1) {
        throw new Error("get supported models status invalid");
    }

    return response.data.data?.support_model_names;
}

const execCmd = async ({
    cmd,
    isLog = true,

}: {
    cmd: string,
    isLog?: boolean,

}): Promise<any> => {
    return new Promise((resolve, reject) => {
        exec(cmd, (error, stdout, stderr) => {
            if (error) {
                console.error(`Exec script ${cmd} error: ${error.message}`);
                reject(error);
            }
            if (stderr) {
                console.error(`Exec script ${cmd} stderr: ${stderr}`);
                reject(stderr);
            }
            if (isLog) {
                console.log(stdout);
            }
            resolve(stdout);
        });
    })
}


const padWithPrefix = (s: string, prefix: string, length: number): string => {
    // If the string is already longer than or equal to the desired length, return it as is
    if (s.length >= length) {
        return s;
    }

    // Calculate how many times to repeat the prefix
    const paddingLength = length - s.length;
    const repeatedPrefix = prefix.repeat(Math.ceil(paddingLength / prefix.length));

    // Pad the string with the repeated prefix and trim to the fixed length
    return (repeatedPrefix + s).slice(0, length);
}



export {
    getSupportedModels,
    execCmd,
    padWithPrefix,
}

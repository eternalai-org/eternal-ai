import * as ethers from 'ethers';
import { PromptPayload } from 'agent-server-definition';

export const prompt = async (payload: PromptPayload): Promise<string> => {
  // your code here
  console.log('payload', payload);

  try {
    if (payload.privateKey) {
      const wallet = new ethers.Wallet(payload.privateKey);

      const address = wallet.address;
      const timestamp = Math.floor(Date.now() / 1000).toString();

      const signature = await wallet.signMessage(timestamp);

      const content = payload.messages?.[0]?.content;

      const params = {
        'XXX-Address': address,
        'XXX-Message': timestamp,
        'XXX-Signature': signature,
      };

      const myHeaders = new Headers();
      myHeaders.append('XXX-Address', address);
      myHeaders.append('XXX-Message', timestamp);
      myHeaders.append('XXX-Signature', signature);
      myHeaders.append('Content-Type', 'application/json');

      console.log('params', params);

      const raw = JSON.stringify({
        content,
      });

      const requestOptions: any = {
        method: 'POST',
        headers: myHeaders,
        body: raw,
        redirect: 'follow',
      };

      const response = await fetch(
        'https://agent.api.eternalai.org/api/utility/twitter/post',
        requestOptions
      );

      const rs = await response.text();

      console.log('rs', rs);

      const parseResult = JSON.parse(rs)?.result;

      console.log('parseResult', parseResult);

      const message = parseResult?.auth_url || parseResult?.message;
      return message;
    }
  } catch (error) {
    console.log('error', JSON.stringify(error));
    return JSON.stringify(error);
  }
  return 'DONE';
};

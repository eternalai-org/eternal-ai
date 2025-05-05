import { ChatRequest, ChatResponse, DirectMessageObj, DirectUnreadMessagesResponse } from "@/types/chat";
import { PromptPayload } from "agent-server-definition";
import { ethers } from "ethers";
import { logger } from '../utils/logger';
import axios, { AxiosResponse } from "axios";

const sendMessage = async (request: ChatRequest): Promise<ChatResponse> => {
    try {
        const response: AxiosResponse<ChatResponse> = await axios.post(
            process.env.NODE_ENV === 'production' ? 'http://localmodel:65534/v1/chat/completions' : 'http://localhost:65534/v1/chat/completions',
            request,
            {
                headers: {
                    'Accept': 'application/json',
                    'Content-Type': 'application/json',
                }
            }
        );
        return response.data;
    } catch (error) {
        logger.error('Error sending chat message:', error);
        throw new Error('Failed to send chat message');
    }
}

export const getAuthHeaders = async (request: PromptPayload) => {
    if (request.privateKey != null && request.privateKey != '') {
        const wallet = new ethers.Wallet(request.privateKey || '');
        const address = wallet.address;
        const timestamp = Math.floor(Date.now() / 1000).toString();
        const signature = await wallet.signMessage(timestamp);
        return {
            'Content-Type': 'application/json',
            'XXX-Address': address,
            'XXX-Message': timestamp,
            'XXX-Signature': signature,
        }
    }
    return {};
}

export const getNotificationMessage = async (request: PromptPayload): Promise<string> => {
    try {
        var authUsername = ''
        if (request.privateKey != null && request.privateKey != '') {
            const headers = await getAuthHeaders(request);
            const response = await axios.get(
                process.env.NODE_ENV === 'production' ? 'https://agent.api.eternalai.org/api/utility/twitter/auth' : 'http://localhost:8480/api/utility/twitter/auth',
                {
                    headers: headers,
                }
            );
            if (response.status !== 200) {
                throw new Error('Failed to authenticate');
            }
            const authUrl = response.data.result.auth_url as string;
            if (authUrl != null && authUrl != '') {
                return authUrl;
            }
            authUsername = response.data.result.username as string;
        }
        const unreadMessages = await getDirectUnreadMessages(request);
        if (unreadMessages == null || unreadMessages.length === 0) {
            return 'No unread messages'
        }
        const prompt = `
I am ${authUsername}, a twitter user.

This is the conversation history with json content:

${JSON.stringify(request.messages)}

This is the direct messages from X with json content:

${JSON.stringify(unreadMessages)}

Based on the conversation history and the direct messages from X, summarize the direct messages from X in the same language, grouping them by dm_conversation_id.

If the direct message is summarized in the conversation history, it will be ignored.

If there is nothing to summarize, reply "No unread messages".

If there is something to summarize, ask the user if they want to auto-reply to these users with suggested responses.
            `.trim();
        console.log('prompt', prompt);
        const response = await sendMessage({
            messages: [{
                role: 'user',
                content: prompt
            }],
            stream: false
        });
        const content = response.choices[0].message.content;
        return content;
    } catch (error) {
        logger.error('Error sending prompt:', error);
        throw new Error('Failed to send prompt');
    }
}

export const sendPrompt = async (request: PromptPayload): Promise<string> => {
    try {
        var authUsername = ''
        if (request.privateKey != null && request.privateKey != '') {
            const headers = await getAuthHeaders(request);
            const response = await axios.get(
                process.env.NODE_ENV === 'production' ? 'https://agent.api.eternalai.org/api/utility/twitter/auth' : 'http://localhost:8480/api/utility/twitter/auth',
                {
                    headers: headers,
                }
            );
            if (response.status !== 200) {
                throw new Error('Failed to authenticate');
            }
            const authUrl = response.data.result.auth_url as string;
            if (authUrl != null && authUrl != '') {
                return authUrl;
            }
            authUsername = response.data.result.username as string;
        }
        const lastMessage = request.messages[request.messages.length - 1];
        if (lastMessage.role === 'user' && lastMessage.content == '.') {
            const unreadMessages = await getDirectUnreadMessages(request);
            if (unreadMessages == null || unreadMessages.length === 0) {
                return 'No unread messages'
            }
            const prompt = `
I am ${authUsername}, a twitter user.

This is the conversation history with json content:

${JSON.stringify(request.messages)}

This is the direct messages from X with json content:

${JSON.stringify(unreadMessages)}

Based on the conversation history and the direct messages from X, summarize the direct messages from X in the same language, grouping them by dm_conversation_id.

If the direct message is summarized in the conversation history, it will be ignored.

If there is nothing to summarize, reply "No unread messages".

If there is something to summarize, ask the user if they want to auto-reply to these users with suggested responses.
            `.trim();
            console.log('prompt', prompt);
            const response = await sendMessage({
                messages: [{
                    role: 'user',
                    content: prompt
                }],
                stream: false
            });
            const content = response.choices[0].message.content;
            return content;
        }
        // handle the case that the last message
        const prompt = `
The conversation history is a conversation between you and the user. You are assistant of the user.

You will help the user to get unread messages from X and summarize them (summarize in the conversation history). Then user will reply to you to auto-reply to some users or not. Sometime user reply somthing for purpose of other things.

This is the conversation history with json content:

${JSON.stringify(request.messages)}

User's X username is ${authUsername || 'anonymous'}.

Based on the conversation history, analyze the user's intent in the last message:
Do they want to auto-reply to some users? If yes, set json field "auto_reply" to true and provide suggested replies in "replies".
If auto-reply is true, suggest one reply per recipient in "replies", reply content should be in the same language as the last message from you that summarized the direct messages, and provide a summary in "reply_summary" that describes what replies will be sent.
Reply content for direct messages should be in the same tone and personality as the conversation.
If auto-reply is false, provide a relevant response about the conversation in "related_content".

Response in json format:
{
    "auto_reply": true,
    "replies": [
        { 
            "dm_conversation_id": "dm_conversation_id",
            "username": "username", 
            "reply": "reply_content" 
        }
    ],
    "reply_summary": "reply_summary_content",
    "related_content": "related_content"
}
      `.trim();
        console.log('prompt', prompt);
        const response = await sendMessage({
            messages: [{
                role: 'user',
                content: prompt
            }],
            stream: false
        });
        var reponse = '';
        let content = response.choices[0].message.content;
        if (content.startsWith('```json')) {
            content = content.slice(7, -3);
        }
        if (content.endsWith('```')) {
            content = content.slice(0, -1);
        }
        console.log('content', content);
        const replyObject = JSON.parse(content);
        if (replyObject.auto_reply) {
            for (const reply of replyObject.replies) {
                try {
                    await sendDirectMessage(request, reply.dm_conversation_id, reply.username, reply.reply);
                } catch (error) {
                    logger.error('Error sending direct message:', error);
                }
            }
            reponse = replyObject.reply_summary;
        } else {
            reponse = replyObject.related_content;
        }
        return reponse;
    } catch (error) {
        logger.error('Error sending prompt:', error);
        throw new Error('Failed to send prompt');
    }
}

const getDirectUnreadMessages = async (request: PromptPayload): Promise<DirectMessageObj[]> => {
    const headers = await getAuthHeaders(request);
    const response = await axios.get(
        process.env.NODE_ENV === 'production' ? 'https://agent.api.eternalai.org/api/utility/twitter/dm' : 'http://localhost:8480/api/utility/twitter/dm',
        {
            headers: headers,
        }
    );
    if (response.status !== 200) {
        throw new Error('Failed to authenticate');
    }
    console.log('getDirectUnreadMessages', response.data.result.data);
    return response.data.result.data;
}

const sendDirectMessage = async (request: PromptPayload, dm_conversation_id: string, username: string, content: string): Promise<void> => {
    const headers = await getAuthHeaders(request);
    console.log('sendDirectMessage', {
        dm_conversation_id: dm_conversation_id,
        recipient_username: username,
        content: content,
    });
    //get conversation chat history
    const resp = await axios.get(
        process.env.NODE_ENV === 'production' ? `https://agent.api.eternalai.org/api/utility/twitter/dm/byid/${dm_conversation_id}` : `http://localhost:8480/api/utility/twitter/dm/byid/${dm_conversation_id}`,
        {
            headers: headers,
        }
    );

    const prompt = `
    Based on the chat conversation history, please analyze and respond to the user while maintaining:

    1. The same tone, style and personality shown in the previous messages
    2. The same language used in the conversation
    3. Consistency with the context and flow of the discussion

    Please provide a response that:
    - Matches the formality level of previous exchanges
    - Uses similar vocabulary and expressions
    - Maintains the established rapport and conversation style
    - Stays relevant to the ongoing discussion topic

    If the conversation history shows specific patterns (casual, formal, technical, friendly, etc.), please reflect those in your response.

    ${JSON.stringify(resp.data.result.data)}
    
    `;


    console.log('prompt', prompt);
    const response = await sendMessage({
        messages: [{
            role: 'user',
            content: prompt
        }],
        stream: false
    });
    var reponse = '';
    content = response.choices[0].message.content;
    if (content.startsWith('```json')) {
        content = content.slice(7, -3);
    }
    if (content.endsWith('```')) {
        content = content.slice(0, -1);
    }
    
    console.log('sendDirectMessage', {
        recipient_username: username,
        content: content,
        dm_conversation_id: dm_conversation_id,
        dm_conversation_data: resp.data.result.data,
    });

    // await axios.post(
    //     process.env.NODE_ENV === 'production' ? 'https://agent.api.eternalai.org/api/utility/twitter/dm/reply' : 'http://localhost:8480/api/utility/twitter/dm/reply',
    //     {
    //         recipient_username: username,
    //         content: content,
    //     },
    //     {
    //         headers: headers,
    //     }
    // );
}
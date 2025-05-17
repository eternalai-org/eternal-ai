export interface ChatMessage {
  role: 'user' | 'assistant';
  content: string;
}

export interface ChatRequest {
  messages: ChatMessage[];
  stream: boolean;
  seed?: number;
  name?: string;
  description?: string;
}

export interface PromptRequest {
  id: string;
  messages: ChatMessage[];
  stream: boolean;
  seed?: number;
  chainId?: number;
  name?: string;
  description?: string;
  privateKey?: string;
}

export interface ChatResponse {
  id: string;
  object: string;
  created: number;
  model: string;
  service_tier: string;
  system_fingerprint: string;
  choices: {
    index: number;
    message: {
      role: string;
      content: string;
      refusal: null | string;
    };
    logprobs: null;
    finish_reason: null | string;
  }[];
}

export interface DirectUnreadMessagesResponse {
  username: string;
  message: string;
}


export interface DirectMessageObj {
  id: string;
  text: string;
  event_type: string;
  created_at: Date;
  sender_id: string;
  sender_username: string;
  recipient_id: string;
  conversation_id: string;
  referenced_tweet?: {
    id: string;
  };
  message?: string;
}

export interface DirectMessageResponse {
  data: DirectMessageObj[];
}



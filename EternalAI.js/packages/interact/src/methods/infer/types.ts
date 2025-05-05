export type Message = {
  role: 'system' | 'user' | 'assistant';
  content: string;
};

export type SendInferResponse = string;

export type ListenInferPayload = SendInferResponse;

// Define the message structure for the 'messages' array
type Message = {
  role: 'system' | 'user' | 'assistant' | 'tool'; // Common roles; 'tool' for tool responses
  content: string; // The text content of the message
  name?: string; // Optional name for multi-agent scenarios
};

// The main PromptPayload interface with all parameters
export type PromptPayload = {
  privateKey?: string;
  messages: Message[]; // Required: Array of message objects
  chainId?: string;
};

# Twitter Infrastructure Agent

A TypeScript-based AI agent that integrates with Twitter's infrastructure through the Eternal AI platform. This agent handles Twitter operations with secure authentication and message processing.

## 📋 Features

- Secure authentication using Ethereum wallet signatures
- Twitter post integration through Eternal AI API
- TypeScript-based implementation
- Error handling and logging
- Environment-based configuration

## 🛠️ Setup

1. Install dependencies:
   ```bash
   npm install
   ```

2. Configure environment variables:
   ```bash
   PORT=4000 # Development port
   NODE_ENV=development
   ```

3. Start the development server:
   ```bash
   npm run dev
   ```

4. For production:
   ```bash
   npm run start
   ```

## 📦 Project Structure

```
twitter-infra/
├── src/
│   ├── index.ts           # Main server entry point
│   └── prompt/
│       ├── index.ts       # Prompt handling implementation
│       └── types.ts       # Type definitions
├── package.json
└── tsconfig.json
```

## 🔐 Authentication

The agent uses Ethereum wallet-based authentication:
1. Signs a timestamp message with the provided private key
2. Sends the signature along with the wallet address
3. Authenticates with the Eternal AI platform

## 🚀 Usage

The agent accepts messages through the Eternal AI platform and processes them with the following flow:

1. Receives a message payload with a private key
2. Generates a wallet signature
3. Sends the request to the Eternal AI Twitter API
4. Returns the authentication URL or response message

## 📚 API Endpoints

- Main endpoint: `https://agent.api.eternalai.org/api/utility/twitter/post`
- Required headers:
  - `XXX-Address`: Ethereum wallet address
  - `XXX-Message`: Current timestamp
  - `XXX-Signature`: Signed message
  - `Content-Type`: application/json

## 🛠️ Development

```bash
# Run tests
npm test

# Lint code
npm run lint

# Build for production
npm run build:eternal

# Watch mode for development
npm run build:watch
```

## 🚀 Deployment to Eternal AI

To deploy this agent to Eternal AI platform:

1. Build and prepare your source code:
   ```bash
   # This will build the project and create a zip file
   npm run build:eternal
   ```

2. Visit [Eternal AI Developer Portal](https://eternalai.org/for-developers/create)

3. Upload the zip file:
   - Log in to your Eternal AI account
   - Navigate to the "Create Agent" section
   - Upload the generated `twitter-infra.zip` file
   - Follow the on-screen instructions to complete deployment

> Note: The `build:eternal` script will:
> - Build the project
> - Create a zip file excluding:
>   - `node_modules` directory
>   - `.git` directory
>   - Any existing zip files
>   - `dist` directory

## 📝 License

MIT License - See LICENSE file for details 
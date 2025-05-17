# TypeScript Node.js API with Express

A modern, well-structured TypeScript Node.js API using Express with best practices.

## Features

- TypeScript support
- Express.js framework
- Environment configuration
- Error handling middleware
- Logging with Winston
- Security with Helmet
- CORS support
- Health check endpoint
- Development and production configurations

## Project Structure

```
src/
├── middleware/     # Express middleware
├── routes/         # API routes
├── utils/          # Utility functions
└── server.ts       # Main application file
```

## Getting Started

1. Install dependencies:
```bash
npm install
```

2. Create a `.env` file in the root directory with the following variables:
```
PORT=3000
NODE_ENV=development
LOG_LEVEL=debug
```

3. Start the development server:
```bash
npm run dev
```

4. Build for production:
```bash
npm run build
```

5. Start production server:
```bash
npm start
```

## Available Scripts

- `npm run dev` - Start development server with hot reload
- `npm run build` - Build TypeScript to JavaScript
- `npm start` - Start production server
- `npm test` - Run tests
- `npm run lint` - Run ESLint
- `npm run format` - Format code with Prettier

## API Endpoints

- `GET /health` - Health check endpoint

## Development

The project uses:
- TypeScript for type safety
- ESLint for code linting
- Prettier for code formatting
- Jest for testing
- Winston for logging
- Helmet for security
- CORS for cross-origin requests 
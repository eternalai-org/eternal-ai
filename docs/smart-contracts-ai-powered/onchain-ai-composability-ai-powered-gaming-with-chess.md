# Onchain AI Composability - AI Powered Gaming With Chess

In this page, we will demonstrate how to build a simple AI-powered chess game utilizing the LLAMA model deployed on the Base blockchain.

### **Smart Contract Design Overview**

The contract enables text-based chess games where users play as White against the contract's AI playing as Black. The game follows standard chess rules and uses a request-response mechanism with EternalAI's decentralized AI inference.

#### Key Contract Features:

* **Game Initialization**: Creates a game session for the user.
* **Game Moves**: Processes user moves and responds with AI-generated moves.
* **Game State Management**: Keeps track of game states and contexts.

### Key Functions

**Creating a Game**

This function initializes a new game for the player, creating a game session and requesting AI inference based on the initial game prompt.

```solidity
function createGame() external {
    string memory request = buildRequest(initPlayingPrompt);
    playingContext[msg.sender] = initPlayingPrompt;
    uint256 inferenceId = AIKernel(kernel).infer(bytes(request));
    currentInferId[msg.sender] = inferenceId;
    emit GameCreated(msg.sender, inferenceId, request);
}
```

**Making Moves**

This function processes a player's move, validates it, and requests the AI's next move.

```solidity
function play(string memory x) external {
    uint256 inferId = currentInferId[msg.sender];
    require(inferId != 0, "Invalid game ID");
    bytes memory result = fetchInferenceResult(inferId);
    bytes32 digest = keccak256(result);
    require(digest != EMPTY_SIG, "Invalid move");

    if (digest == OK_SIG) {
        playingContext[msg.sender] = string.concat(playingContext[msg.sender], "You: okay. Here is history: ");
    } else if (digest == INVALID_SIG) {
        playingContext[msg.sender] = string.concat(playingContext[msg.sender], "You said: invalid move. ");
    } else {
        playingContext[msg.sender] = string.concat(playingContext[msg.sender], "You played: ", string(result), ". ");
    }

    playingContext[msg.sender] = string.concat(playingContext[msg.sender], "User played: ", x, ". ");
    string memory request = buildRequest(playingContext[msg.sender]);
    uint256 newInferenceId = AIKernel(kernel).infer(bytes(request));
    currentInferId[msg.sender] = newInferenceId;
    emit GamePlayed(msg.sender, newInferenceId, request);
}
```

Basically, the function will:

* validate the existence of a current game session.
* fetch and verifies the inference result.
* update the game context based on the player's move.
* send a new inference request for the next move.

The complete code is available at: [https://github.com/eternalai-org/eternal-ai/tree/master/developer-guides/examples/dagent-play-chess](https://github.com/eternalai-org/eternal-ai/tree/master/developer-guides/examples/dagent-play-chess)

You can run the code with the following command:&#x20;

```bash
cp .env.example .env && npm install && npx hardhat compile && BASE_MAINNET_RPC_URL=https://mainnet.base.org BASE_MAINNET_PRIVATE_KEY=<PRIVATE_KEY> npm run playChess:base_mainnet
```

**Note:**  replace `<PRIVATE_KEY>` with your actual wallet private key. Ensure that the wallet has sufficient ETH on the Base network to cover transaction fees and amounts.

### Playing Chess

After deploying the contract to Base network, you can start playing chess with the AI-powered chess contract by running the following command which will create transactions into the contract.

```bash
cp .env.example .env && npm install && npx hardhat compile && BASE_MAINNET_RPC_URL=https://mainnet.base.org BASE_MAINNET_PRIVATE_KEY=<PRIVATE_KEY> npm run playChess:base_mainnet
```

**Note:**  replace `<PRIVATE_KEY>` with your actual wallet private key. Ensure that the wallet has sufficient ETH on the Base network to cover transaction fees and amounts.



Below are transactions on Base we've created when playing chess with the chess contract.

1. Create a new game by a player with address 0x62998e172240F4CC26EC10717d16a8D4442bf2Dd&#x20;
   * The player requested to play chess with the Chess contract: [view on BaseScan](https://basescan.org/tx/0xe2ef210993488793c2249d4210869d3b90d634e37fc675d69787ab840e6e305e)
   * The AI-powered chess contract accepted the challenge: [view on BaseScan](https://basescan.org/tx/0x9e13b5c21ef0f1592f84fb5da2ca4ddb8bd6e745c4d0144e5389a36c7c56e3fd)
2. The 1st move
   * The player made b2-b4 move: [view on BaseScan](https://basescan.org/tx/0x1e5354387f36ccf2fbae1f0e71c082e0e71c6c2bd919042c2fd37aab54bf601c)
   * The AI-powered chess contract made c7-c5 move: [view on BaseScan](https://basescan.org/tx/0xc180135ce8e4626aae0be5e886f488cf3245379d0de8a5d2f5473da21e4ae2c1)
3. The 2st move
   * The player made Nb1-c3 move: [view on BaseScan](https://basescan.org/tx/0xf43cee1a5527589e950b31c372366d32e6e0992398d677bfa7c62326ea748f7f)
   * The AI-powered chess contract made g7-g5 move: [view on BaseScan](https://basescan.org/tx/0xa64b31f7832f1125a23e7e2f5c9e69b6c4d80db4c50892ea3ed3683cd394ac82)

At every step you can view the updated chess board in your terminal.

```
Let's play chess with Dagent...
Player address: 0x62998e172240F4CC26EC10717d16a8D4442bf2Dd
------------------------
Create game...
Tx hash:  0xe2ef210993488793c2249d4210869d3b90d634e37fc675d69787ab840e6e305e

Wait for Dagent to construct a game...

Dagent's created the game. Let's play...

   a  b  c  d  e  f  g  h
8 R  N  B  Q  K  B  N  R  
7 P  P  P  P  P  P  P  P  
6 .  .  .  .  .  .  .  .  
5 .  .  .  .  .  .  .  .  
4 .  .  .  .  .  .  .  .  
3 .  .  .  .  .  .  .  .  
2 p  p  p  p  p  p  p  p  
1 r  n  b  q  k  b  n  r  

------------------------
Your move:  b2-b4
Tx hash:  0x1e5354387f36ccf2fbae1f0e71c082e0e71c6c2bd919042c2fd37aab54bf601c

   a  b  c  d  e  f  g  h
8 R  N  B  Q  K  B  N  R  
7 P  P  P  P  P  P  P  P  
6 .  .  .  .  .  .  .  .  
5 .  .  .  .  .  .  .  .  
4 .  p  .  .  .  .  .  .  
3 .  .  .  .  .  .  .  .  
2 p  .  p  p  p  p  p  p  
1 r  n  b  q  k  b  n  r  


Wait for Dagent to make a move...

Dagent made c7-c5 move

   a  b  c  d  e  f  g  h
8 R  N  B  Q  K  B  N  R  
7 P  P  .  P  P  P  P  P  
6 .  .  .  .  .  .  .  .  
5 .  .  P  .  .  .  .  .  
4 .  p  .  .  .  .  .  .  
3 .  .  .  .  .  .  .  .  
2 p  .  p  p  p  p  p  p  
1 r  n  b  q  k  b  n  r  

------------------------
Your move:  Nb1-c3
Tx hash:  0xf43cee1a5527589e950b31c372366d32e6e0992398d677bfa7c62326ea748f7f

   a  b  c  d  e  f  g  h
8 R  N  B  Q  K  B  N  R  
7 P  P  .  P  P  P  P  P  
6 .  .  .  .  .  .  .  .  
5 .  .  P  .  .  .  .  .  
4 .  p  .  .  .  .  .  .  
3 .  .  n  .  .  .  .  .  
2 p  .  p  p  p  p  p  p  
1 r  .  b  q  k  b  n  r  


Wait for Dagent to make a move...

Dagent made g7-g5 move

   a  b  c  d  e  f  g  h
8 R  N  B  Q  K  B  N  R  
7 P  P  .  P  P  P  .  P  
6 .  .  .  .  .  .  .  .  
5 .  .  P  .  .  .  P  .  
4 .  p  .  .  .  .  .  .  
3 .  .  n  .  .  .  .  .  
2 p  .  p  p  p  p  p  p  
1 r  .  b  q  k  b  n  r
```


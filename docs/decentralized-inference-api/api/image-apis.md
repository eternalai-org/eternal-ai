# Imagine and Reimagine API Documentation

## Overview

This documentation provides a comprehensive guide for integrating with the EternalAI Imagine (image generation) and Reimagine (image editing) APIs.

## Prerequisites

- An active API key from [EternalAI website](https://eternalai.org/apis).
- Basic understanding of HTTP requests and streaming responses
- Knowledge of handling Base64 encoded images (for image editing features)

⚠️ **Important:** All API requests require a valid API key in the `x-api-key` header. Requests without this header will return a `401 Unauthorized` error.

## Imagine API

Use standard message payload with text content, similar to other LLM/AI models:

```bash
curl 'https://agent-api.eternalai.org/agent-ai/9467-uncensored-imagine/prompt' \
  -H 'accept: text/event-stream' \
  -H 'x-api-key: $ETERNALAI_API_KEY' \
  -H 'content-type: application/json' \
  --data-raw '{
    "messages": [
      {
        "role": "user",
        "content": "Generate a red cat image"
      }
    ],
    "stream": true
  }'
```

## Reimagine API

For image editing, add an `image_url` item with base64 data to the content array:

```bash
curl 'https://agent-api.eternalai.org/agent-ai/9467-uncensored-imagine/prompt' \
  -H 'accept: text/event-stream' \
  -H 'x-api-key: YOUR_API_KEY_HERE' \
  -H 'content-type: application/json' \
  --data-raw '{
    "messages": [
      {
        "role": "user",
        "content": [
          {
            "type": "text",
            "text": "Make this cat blue instead of red"
          },
          {
            "type": "image_url",
            "image_url": {
              "url": "data:image/jpeg;base64,/9j/...../2gAMAwEAAhEDEQA//Z",
              "filename": "red_cat.jpg"
            }
          }
        ]
      }
    ],
    "stream": true
  }'
```

## Response Format

The APIs return streaming data (SSE) with different content types:

### 1. Thinking Process (Optional)
```
<think>
Analyzing the request for a red cat image...
I'll use FLUX model to generate a high-quality image...
</think>
```

### 2. Action Logs (Optional)  
```
<action>
Loading FLUX model...
Generating image with prompt: "a red cat"...
Processing complete.
</action>
```

### 3. Generated Image (Main Result)
```
<img src="data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAAAQABAAD/2wBDAAYEBQYFBAYGBQYHBwYIChAKCgkJChQODwwQFxQYGBcUFhYaHSUfGhsjHBYWICwgIyYnKSopGR8tMC0oMCUoKSj/2wBDAQcHBwoIChMKChMoGhYaKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCj/wAARCADEAO..." alt="Generated red cat image"/>
```

### Parsing the Response
- **Thinking content**: `<think>...</think>` - AI reasoning process (optional display)
- **Action content**: `<action>...</action>` - Processing logs (optional display)  
- **Image content**: `<img src="data:image/jpeg;base64,..." alt="..."/>` - **Main result with base64 image data**

Extract the base64 data from the `src` attribute of the `<img>` tag to display or save the generated image.

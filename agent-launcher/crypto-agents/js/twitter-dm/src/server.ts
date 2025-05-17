import express from 'express';
import { getNotificationMessage, sendPrompt } from './prompt';

const app = express();
const port = process.env.PORT || 3000;

// Middleware to parse JSON bodies
app.use(express.json());

// Route for handling prompts
app.post('/prompt', async (req, res) => {
    try {
        const result = await sendPrompt(req.body);
        res.status(200).json(result);
    } catch (error) {
        res.status(500).json({ error: 'Failed to process prompt' });
    }
});

// Route for handling notifications
app.post('/notification', async (req, res) => {
    try {
        // Handle notification logic here
        const result = await getNotificationMessage(req.body);
        res.status(200).json(result);
    } catch (error) {
        res.status(500).json({ error: 'Failed to process notification' });
    }
});

// Start the server
app.listen(port, () => {
    console.log(`Server is running on port ${port}`);
});
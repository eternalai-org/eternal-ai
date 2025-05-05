import { startServer } from "agent-server-definition";
import { prompt } from "./prompt";

const port = process.env.PORT;
startServer(Number(port), prompt);

import cors from "cors"
import mqtt from 'mqtt'
import dotenv from "dotenv";
import express from "express";
import { setup_admin_router, setup_api_router } from "./router";

dotenv.config();

const app = express();

const broker = mqtt.connect(process.env.BROKER_URL || "")

app.use(cors({
  options: {
    origin: '*'
  }
}))

app.use(express.json())

app.use("/api/v1", setup_api_router(express.Router(), broker))
app.use("/api/admin", setup_admin_router(express.Router()))

const port = process.env.PORT || 5174;
const host = process.env.HOST || 'localhost';
app.listen(port, () => {
  console.log(`[server]: Server is running at http://${host}:${port}`);
});

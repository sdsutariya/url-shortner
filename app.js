import dotenv from "dotenv";
dotenv.config();
import express from "express";
const app = express();
import connectdb from "./db/connectdb.js";
import userroute from "./routes/users.js";
import authroute from "./routes/auth.js";
import postroute from "./routes/posts.js";
import helmet from "helmet";
import morgan from "morgan";
const port = process.env.PORT || 3002;
const DATABASE_URL = process.env.DATABASE_URL;

//database connection
connectdb(DATABASE_URL);

//middleware
app.use(express.json());
app.use(helmet());
app.use(morgan("common"));

//routes
app.use("/api/user", userroute);
app.use("/api/auth", authroute);
app.use("/api/post",postroute);

app.listen(port, () => {
  console.log(`server listening on http://localhost:${port}`);
});

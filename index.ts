import http from "http";
import express from "express";

const app = express();
app.use(express.json())

app.get("/", (req, res) => {
	console.log("GET:", JSON.stringify(req.body));
	res.statusCode = 200;
	res.send(req.body);
	res.end();
})
app.post("/", (req, res) => {
	console.log("POST:", JSON.stringify(req.body));
	res.statusCode = 200;
	res.send(req.body);
	res.end();
})

const server = http.createServer(app);
server.listen("3000")

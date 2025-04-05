const express = require("express");
const http = require("http");
const { Server } = require("socket.io");

const app = express();
const server = http.createServer(app);
const io = new Server(server, {
    cors: {
        origin: "*",
    },
});

io.on("connection", (socket) => {
    console.log("Client connected to server2:", socket.id);

    socket.on("message", (data) => {
        console.log("Server2 received from client:", data);

        // Echo the message back to client
        socket.emit("message", `Server received: ${data}`);
    });

    socket.on("disconnect", () => {
        console.log("Client disconnected from server2:", socket.id);
    });
});

server.listen(3002, () => {
    console.log("Socket.IO server listening on port 3002");
});

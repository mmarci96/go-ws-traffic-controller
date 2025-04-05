const { io } = require("socket.io-client");

// Change this to point to your Go proxy later
const socket = io("http://localhost:3000");

socket.on("connect", () => {
    console.log("Connected to server");

    // Send a test message
    socket.emit("message", "Hello from client!");
});

socket.on("message", (data) => {
    console.log("Received from server:", data);
});

socket.on("disconnect", () => {
    console.log("Disconnected from server");
});

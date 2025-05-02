const WebSocket = require('ws');
const readline = require('readline');

// Create an interface for reading from the console
const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout
});

// Connect to the WebSocket server
const ws = new WebSocket('ws://localhost:8080/ws');

// When the connection is open, prompt the user to send a message
ws.on('open', function open() {
  console.log('Connected to server. Type a message and press Enter to send:');
  
  rl.on('line', (input) => {
    ws.send(input); // Send the input message to the server
  });
});

// When a message is received, log it to the console
ws.on('message', function incoming(data) {
  console.log('Received from server:', data);
});

// Handle errors
ws.on('error', function error(err) {
  console.log('Error occurred:', err);
});

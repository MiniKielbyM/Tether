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
    try {
      // Parse the input as JSON if it's a valid JSON string
      const jsonObject = JSON.parse(input);
      ws.send(JSON.stringify(jsonObject)); // Send the JSON object as a string
    } catch (e) {
      console.log('Invalid JSON. Sending as plain text.');
      ws.send(input); // Send as plain text if not valid JSON
    }
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

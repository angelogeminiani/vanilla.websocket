# Vanilla.Websocket client

## How to Use

### HTML

Just download binary version form build folder and add a reference 
to Vanilla.Websocket. 

```html 
<script type="text/javascript" src="../vendor/vwf/vmf.min.js"></script>
```

### Javascript

Below is a simple example on Vanilla.Websocket usage:

```javascript
try {      
         // get Vanilla.Websocket constructor
         const vws = window["__vws"];
         if (!!vws) {
             console.info("Vanilla.Websocket", "version: " + vws.version);
             const client = vws.create("ws://localhost:8080/websocket");
             console.info("Vanilla.Websocket client", "host: " + client.host);
 
             // handle connection error
             client.on("on_error", (err) => {
                 console.error("Vanilla.Websocket client#on_error", err);
             }); 
            client.on("on_message", (message) => {
                 let source = "broadcast server message"
                 if (message["uid"] === "test_message") {
                     source = "response to client request"
                 }
                 console.info("Vanilla.Websocket client#on_message", source, message);
             });
            client.on("on_close", (err) => {
                 console.info("Vanilla.Websocket client#on_close");
             });
 
             // creates a message for the server
             const message = {
                 "uid": "test_message",
                 "payload": {
                     "app_token": "iuhdiu87w23ruh897dfyc2w3r",
                     "command": "echo",
                     "params": ["Hello Vanilla.Websocket"]
                 }
             };
             // send message to server
             client.send(message, (full_response) => {
                 const response = full_response["response"];
                 const error = response["error"];
                 if (!!error) {
                     console.error("Vanilla.Websocket client", error);
                 } else {
                     const data = response["data"];
                     console.info("Response", full_response, data);
                 }
             });
 
         } else {
             console.error("Vanilla.Websocket not loaded!", "");
         }
     } catch (e) {
         console.error(e);
     }
```   

First of all create a Vanilla.Websocket builder reference:
```javascript 
const vws = window["__vws"];
```  
Then with the builder instance ("vws") can invoke "create" passing websocket endpoint:
```javascript 
const client = vws.create("ws://localhost:8080/websocket");
``` 
Now you are ready to handle events or to send a message.

## Events

Vanilla.Websocket emits some events:
* on_error: When an error occurred in connection. Usually because endpoint is not responding.
* on_message: When message was sent to client. Usually this event is handled to receive server messages. 
* on_close: Connection closed


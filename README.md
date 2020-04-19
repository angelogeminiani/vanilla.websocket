# vanilla.websocket
Vanilla Websocket is a pure javascript no dependencies websocket client with callbacks.

## Project Directories

### client
Contains javascript client sources and builds.
Here is what you are looking for if you are a javascript/typescript developer.

### server
This is just a test server written in Go. Not interesting for you if you are looking just for a javascript websocket client.

## How to Use

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
 
             // send message to server
             client.send({message:"hello"}, (full_response) => {
                 const response = full_response["response"];
                 const error = response["error"];
                 if (!!error) {
                     console.error("Vanilla.Websocket client", error);
                 } else {
                     // .... response contains the server data
                     // TODO: do something with this data
                 }
             });
 
         } else {
             console.error("Vanilla.Websocket not loaded!", "");
         }
     } catch (e) {
         console.error(e);
     }
```




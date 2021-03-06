(function () {
    try {
        let vws = window["__vws"];
        if (!!vws) {
            const URL = "ws://localhost:80/websocket"

            log("info", "Vanilla.Websocket", "version: " + vws.version);
            const client = vws.create(URL);
            log("info", "Vanilla.Websocket client", "host: " + client.host);

            // handle connection error
            client.on("on_error", (err) => {
                log("error", "Vanilla.Websocket client#on_error", err);
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

            // send a message
            send(client);

            $("#btn_send").on("click", (e) => {
                e.preventDefault;
                send(client);
            });

        } else {
            log("error", "Vanilla.Websocket not loaded!", "");
        }
    } catch (e) {
        console.error(e);
    }

    function send(client) {
        // creates a message for the server
        let message = {
            "uid": "test_message",
            "payload": {
                "app_token": "iuhdiu87w23ruh897dfyc2w3r",
                "command": "echo",
                "params": ["Hello Vanilla.Websocket"]
            }
        };

        // ANOTHER KIND OF MESSAGE
        // MESSAGES DEPENDS ON SERVER SIDE IMPLEMENTATION
        /*
        const message = {
            "uid": "2b423c92-11ef-33e6-5537-970bff1e4177",
            "lang": "it",
            "payload": {
                "app_token": "iuhdiu87w23ruh897dfyc2w3r",
                "namespace": "system",
                "function": "ping",
                "params": null,
            },
        };*/
         log("info", "Sending message", message)
        // send message to server
        client.send(message, (full_response) => {
            const response = full_response["response"]||{};
            const error = response["error"];
            if (!!error) {
                log("error", "Vanilla.Websocket client", error);
            } else {
                const data = response["data"];
                log("info", "Response", full_response);
            }
        });
    }

    function log(level, context, message) {
        try {
            let style = "alert-primary";
            switch (level) {
                case "error":
                    console.error(context, message);
                    style = "alert-danger"
                    break;
                default:
                    console.info(context, message);
            }
            message = JSON.stringify(message)
            const html = "<div class='row'><div class=\"alert " + style + "\" role=\"alert\">" + context + message + " </div></div>";
            $(html).appendTo($("#logs"));
        } catch (e) {
            console.error(e);
        }
    }

})();
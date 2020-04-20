(function () {
    try {
        let vws = window["__vws"];
        if (!!vws) {
            log("info", "Vanilla.Websocket", "version: " + vws.version);
            const client = vws.create("ws://localhost:80/");
            log("info", "Vanilla.Websocket client", "host: " + client.host);

            // handle connection error
            client.on("on_error", (err) => {
                log("error", "Vanilla.Websocket client#on_error", err);
            });

            // send message to server
            client.send({}, (full_response) => {
                const response = full_response["response"];
                const error = response["error"];
                if (!!error) {
                    log("error", "Vanilla.Websocket client", error);
                } else {

                }
            });

        } else {
            log("error", "Vanilla.Websocket not loaded!", "");
        }
    } catch (e) {
        console.error(e);
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
            const html = "<div class='row'><div class=\"alert " + style + "\" role=\"alert\">" + context + message + " </div></div>";
            $(html).appendTo($("#logs"));
        } catch (e) {
            console.error(e);
        }
    }

})();
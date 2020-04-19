import BaseObject from "./commons/BaseObject";
import console, {LogLevel} from "./commons/console";
import {SocketService} from "./socket/SocketService";
import lang from "./commons/lang";

// ------------------------------------------------------------------------
//                      c o n s t a n t s
// ------------------------------------------------------------------------

const DEBUG: boolean = false;
const VERSION: string = "1.0.0";
const UID: string = "__vws"
const ROOT: any = window

/**
 * Launcher class
 */
class launcher
    extends BaseObject {

    // ------------------------------------------------------------------------
    //                      f i e l d s
    // ------------------------------------------------------------------------

    // ------------------------------------------------------------------------
    //                      c o n s t r u c t o r
    // ------------------------------------------------------------------------

    private constructor() {
        super();
        this.init();
    }

    // ------------------------------------------------------------------------
    //                      p u b l i c
    // ------------------------------------------------------------------------

    public start(): void {
        console.debug("launcher.start()");
        // creates Vanilla.Websocket builder reference
        ROOT[UID] = {
            version: VERSION,
            create: this.newClient,

        };
    }

    // ------------------------------------------------------------------------
    //                      p r i v a t e
    // ------------------------------------------------------------------------

    private init(): void {
        // init application scope
        BaseObject.PREFIX = UID + "_"; // application uid become component prefix.
        // set console prefix
        console.uid = UID;
        if (DEBUG) {
            console.level = LogLevel.debug;
        } else {
            console.level = LogLevel.info;
        }
    }

    private newClient(params: any): SocketService {
        if (lang.isString(params)) {
            params = {
                host: params
            };
        } else {
            params = params || {};
        }
        return new SocketService(params);
    }

    // ------------------------------------------------------------------------
    //                      S I N G L E T O N
    // ------------------------------------------------------------------------

    private static __instance: launcher;

    static instance(): launcher {
        if (null == launcher.__instance) {
            launcher.__instance = new launcher();
        }
        return launcher.__instance;
    }

}

// ------------------------------------------------------------------------
//                      S T A R T   A P P L I C A T I O N
// ------------------------------------------------------------------------


launcher.instance().start();

// ------------------------------------------------------------------------
//                      e x p o r t
// ------------------------------------------------------------------------

export {ROOT}
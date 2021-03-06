import EventEmitter from "../events/EventEmitter";
import {EVENT_CLOSE, EVENT_ERROR, EVENT_MESSAGE, EVENT_OPEN, WebSocketChannel} from "./WebSocketChannel";
import lang from "../commons/lang";
import {Listener} from "../events/Events";

class SocketService
    extends EventEmitter {

    private readonly _params: any;

    private _socket: WebSocketChannel; // lazy initialized
    private _is_connected: boolean;

    // ------------------------------------------------------------------------
    //  c o n s t r u c t o r
    // ------------------------------------------------------------------------

    constructor(params: any) {
        super();
        if (lang.isString(params)) {
            this._params = {
                host: params
            };
        } else {
            this._params = params || {};
        }
        this._is_connected = false;
    }

    // ------------------------------------------------------------------------
    //  p u b l i c
    // ------------------------------------------------------------------------

    public get host(): string {
        return !!this._params ? this._params["host"] || "" : "";
    }

    public get isConnected(): boolean {
        return this._is_connected;
    }

    public send(message: any, callback?: any): void {
        this.socket.open().ready((is_ready: boolean, error: any) => {
            if (is_ready) {
                this.socket.send(message, callback);
            } else {
                lang.funcInvoke(callback, {"error": error})
            }
        });
    }

    public close() {
        if (!!this._socket) {
            if (this.socket.initialized) {
                this.socket.close();
            }
        }
    }

    public reset() {
        if (!!this._socket) {
            if (this.socket.initialized) {
                this.socket.reset();
            }
        }
    }

    // @ts-ignore
    public on(eventName: string, listener: Listener): void {
        super.on(this, eventName, listener);
    }

    // @ts-ignore
    public off(eventName?: string): void {
        super.off(this, eventName);
    }

    // ------------------------------------------------------------------------
    //  p r i v a t e
    // ------------------------------------------------------------------------

    private get socket(): WebSocketChannel {
        if (!this._socket) {
            // creates new socket
            this._socket = new WebSocketChannel(this._params);

            this._socket.on(this, EVENT_OPEN, this.onOpen);
            this._socket.on(this, EVENT_CLOSE, this.onClose);
            this._socket.on(this, EVENT_MESSAGE, this.onMessage);
            this._socket.on(this, EVENT_ERROR, this.onError);

        }
        return this._socket;
    }

    private onOpen(): void {
        this._is_connected = true;
        this.emit(EVENT_OPEN);
    }

    private onClose(): void {
        this._is_connected = false;
        this.emit(EVENT_CLOSE);
    }

    private onMessage(data: any): void {
        this.emit(EVENT_MESSAGE, data);
    }

    private onError(error: any): void {
        this._is_connected = false;
        this.emit(EVENT_ERROR, error);
    }
}

// ------------------------------------------------------------------------
//  E X P O R T
// ------------------------------------------------------------------------

export {
    SocketService,
    EVENT_CLOSE, EVENT_MESSAGE, EVENT_OPEN, EVENT_ERROR
}
/**
 * Extends standard console
 */
import random from "./random";
import {Dictionary} from "../collections/Dictionary";

enum LogLevel {
    error = 0,
    warn = 1,
    info = 2,
    debug = 3
}

class console_ext {

    // ------------------------------------------------------------------------
    //                      c o n s t
    // ------------------------------------------------------------------------

    // ------------------------------------------------------------------------
    //                      f i e l d s
    // ------------------------------------------------------------------------

    private readonly _class_levels: Dictionary<LogLevel>;

    private _uid: string;
    private _level: LogLevel;

    // ------------------------------------------------------------------------
    //                      c o n s t r u c t o r
    // ------------------------------------------------------------------------

    private constructor() {
        this._uid = random.guid();
        this._level = LogLevel.info;

        this._class_levels = new Dictionary();
    }

    // ------------------------------------------------------------------------
    //                      p r o p e r t i e s
    // ------------------------------------------------------------------------

    public get uid() {
        return this._uid;
    }

    public set uid(value: string) {
        this._uid = value;
    }

    public get level(): LogLevel {
        return this._level;
    }

    public set level(value: LogLevel) {
        this._level = value;
    }

    // ------------------------------------------------------------------------
    //                      p u b l i c
    // ------------------------------------------------------------------------

    public setClassLevel(class_name: string, level: LogLevel) {
        this._class_levels.put(class_name, level);
    }

    public error(scope: string, error: Error | string, ...args: any[]): void {
        console.error("[" + this.uid + "] " + scope, error, ...args);
    };

    public warn(scope: string, ...args: any[]): void {
        if (this.getLevel(scope) < LogLevel.warn) {
            return;
        }
        console.warn("[" + this.uid + "] " + scope, ...args);
    };

    public info(scope: string, ...args: any[]): void {
        if (this.getLevel(scope) < LogLevel.info) {
            return;
        }
        console.info("[" + this.uid + "] " + scope, ...args);
    };

    public debug(scope: string, ...args: any[]): void {
        if (this.getLevel(scope) < LogLevel.debug) {
            return;
        }
        console.log("[" + this.uid + "] " + scope, ...args);
    };

    public log(scope: string, ...args: any[]): void {
        if (this.getLevel(scope) < LogLevel.info) {
            return;
        }
        console.log("[" + this.uid + "] " + scope, ...args);
    }

    // ------------------------------------------------------------------------
    //                      p r i v a t e
    // ------------------------------------------------------------------------

    private getLevel(scope?: string): LogLevel {
        if (!!scope) {
            const class_name: string = scope.split(".")[0];
            if (this._class_levels.containsKey(class_name)) {
                return this._class_levels.get(class_name);
            }
        }
        return this.level;
    }

    // ------------------------------------------------------------------------
    //                      S I N G L E T O N
    // ------------------------------------------------------------------------

    private static __instance: console_ext;

    public static instance(): console_ext {
        if (null == console_ext.__instance) {
            console_ext.__instance = new console_ext();
        }
        return console_ext.__instance;
    }

}

// ------------------------------------------------------------------------
//                      e x p o r t
// ------------------------------------------------------------------------

export default console_ext.instance();
export {LogLevel};


/**
 * Utility class
 */

class langClass {

    // ------------------------------------------------------------------------
    //                      f i e l d s
    // ------------------------------------------------------------------------


    // ------------------------------------------------------------------------
    //                      c o n s t r u c t o r
    // ------------------------------------------------------------------------

    constructor() {

    }

    // ------------------------------------------------------------------------
    //                      p u b l i c
    // ------------------------------------------------------------------------

    public parse(value: any): any {
        try {
            if (this.isString(value)) {
                return JSON.parse(value);
            }
        } catch (err) {
        }
        return value;
    }

    // ------------------------------------------------------------------------
    //                      t o
    // ------------------------------------------------------------------------

    public toString(value: any): string {
        switch (typeof value) {
            case 'string':
            case 'number':
            case 'boolean':
                return value + '';
            case 'object':
                try {
                    // null is an object but is falsy. Swallow it.
                    return value === null ? '' : JSON.stringify(value);
                } catch (jsonError) {
                    return '{...}';
                }
            default:
                // Anything else will be replaced with an empty string
                // For example: undefined, Symbol, etc.
                return '';
        }
    }

    public toArray<T>(value: any | any[]) {
        return !!value
            ? this.isArray(value) ? value as Array<T> : [value as T]
            : [];
    }

    public toBoolean(value: any, def_val: boolean): boolean {
        return !!value
            ? value !== 'false' && value !== '0'
            : def_val;
    }

    public toFloat(value: any, def_value: number = 0.0, min?: number, max?: number): number {
        try {
            let result = parseFloat(value.replace(/,/g, '.'));
            result = this.isNaN(result) ? def_value : result;
            if (!this.isNaN(max) && result > (max || 0)) result = max || 0;
            if (!this.isNaN(min) && result < (min || 0)) result = min || 0;
            return result;
        } catch (err) {
            return def_value;
        }
    }

    public toInt(value: any, def_value: number = 0, min?: number, max?: number): number {
        try {
            let result = parseInt(value);
            result = this.isNaN(result) ? def_value : result;
            if (!this.isNaN(max) && result > (max || 0)) result = max || 0;
            if (!this.isNaN(min) && result < (min || 0)) result = min || 0;
            return result;
        } catch (err) {
            return def_value;
        }
    }

    // ------------------------------------------------------------------------
    //                      i s
    // ------------------------------------------------------------------------

    public isFunction(value: any): boolean {
        return typeof value == 'function';
    }

    public isObject(value: any): boolean {
        return value === Object(value);
    }

    public isArray(value: any): boolean {
        return !!Array.isArray
            ? Array.isArray(value)
            : value && typeof value == 'object' && typeof value.length == 'number' && toString.call(value) == '[object Array]' || false;
    }

    public isArguments(value: any): boolean {
        return value && typeof value == 'object' && typeof value.length == 'number' &&
            toString.call(value) == '[object Arguments]' || false;
    }

    public isBoolean(value: any): boolean {
        return value === true || value === false ||
            value && typeof value == 'object' && toString.call(value) == '[object Boolean]' || false;
    }

    public isString(value: any): boolean {
        return typeof value == 'string' ||
            value && typeof value == 'object' && toString.call(value) == '[object String]' || false;
    }

    public isNumber(value: any): boolean {
        return typeof value == 'number' ||
            value && typeof value == 'object' && toString.call(value) == '[object Number]' || false;
    }

    public isNaN(value: any): boolean {
        return isNaN(value);
    }

    static isDate(value: any): boolean {
        return value && typeof value == 'object' && toString.call(value) == '[object Date]' || false;
    }


    // ------------------------------------------------------------------------
    //                      u t i l s
    // ------------------------------------------------------------------------

    /**
     * Invoke a function. Shortcut for "func.call(this, ...args)"
     */
    public funcInvoke(func?: Function, ...args: any[]): any {
        const self = this;
        if (!!func && self.isFunction(func)) {
            if (args.length === 0) {
                return func.call(self);
            } else {
                return func.call(self, ...args);
            }
        }
        return null;
    }

    /**
     * Delays a function for the given number of milliseconds, and then calls
     * it with the arguments supplied.
     * NOTE: user "clearTimeout" with funcDelay response to
     */
    public funcDelay(func: Function, wait: number, ...args: any[]): any {
        return setTimeout(function () {
            return func.call(null, ...args);
        }, wait);
    }

    // ------------------------------------------------------------------------
    //                      p r i v a t e
    // ------------------------------------------------------------------------


    // ------------------------------------------------------------------------
    //                      S I N G L E T O N
    // ------------------------------------------------------------------------

    private static __instance: langClass;

    public static instance(): langClass {
        if (null == langClass.__instance) {
            langClass.__instance = new langClass();
        }
        return langClass.__instance;
    }


}

// ------------------------------------------------------------------------
//                      e x p o r t
// ------------------------------------------------------------------------

const lang: langClass = langClass.instance();
export default lang;



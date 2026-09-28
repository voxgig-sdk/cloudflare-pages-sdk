"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.CloudflarePagesError = void 0;
class CloudflarePagesError extends Error {
    isCloudflarePagesError = true;
    sdk = 'CloudflarePages';
    code;
    ctx;
    status = -1;
    // `err.notFound` rather than a magic number at every call site.
    get notFound() { return 404 === this.status; }
    constructor(code, msg, ctx) {
        super(msg);
        this.code = code;
        this.ctx = ctx;
    }
}
exports.CloudflarePagesError = CloudflarePagesError;
//# sourceMappingURL=CloudflarePagesError.js.map
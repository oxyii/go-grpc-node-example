import {credentials} from "@grpc/grpc-js"
import {TestClient} from "./proto/test"
import {StringValue} from "./proto/google/protobuf/wrappers"

const testClient = new TestClient("localhost:50051", credentials.createInsecure());

async function unary(): Promise<StringValue> {
    const req = new StringValue({value: "World"});
    return testClient.Unary(req);
}

async function clientStream(): Promise<StringValue> {
    return new Promise((resolve, reject) => {
        const stream = testClient.ClientStream((err, rsp) => {
            if (err) {
                reject(err);
            } else {
                resolve(rsp);
            }
        });
        for (let i = 0; i < 2; i++) {
            const req = new StringValue({value: `World${i}`});
            stream.write(req);
        }
        stream.end();
    });
}

async function serverStream(): Promise<StringValue[]> {
    return new Promise((resolve, reject) => {
        const ret = [];
        const req = new StringValue({value: "World"});
        const stream = testClient.ServerStream(req);
        stream.on("error", reject);
        stream.on("end", () => {
            resolve(ret);
        });
        stream.on("data", (rsp) => {
            ret.push(rsp);
        });
    });
}

async function bidiStream(): Promise<StringValue[]> {
    return new Promise((resolve, reject) => {
        const ret = [];
        const stream = testClient.BidiStream();
        stream.on("error", reject);
        stream.on("end", () => {
            resolve(ret);
        });
        stream.on("data", (rsp) => {
            ret.push(rsp);
        });
        for (let i = 0; i < 2; i++) {
            const req = new StringValue({value: `World${i}`});
            stream.write(req);
        }
        stream.end();
    });
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

async function all() {
    const rspUnary = await unary();
    console.log("unary " + rspUnary.value);

    const rspClientStream = await clientStream();
    console.log("clientStream " + rspClientStream.value);

    const rspServerStream = await serverStream();
    console.log("serverStream " + rspServerStream.map(v => v.value).join(", "));

    const rspBidiStream = await bidiStream();
    console.log("bidiStream " + rspBidiStream.map(v => v.value).join(", "));
}

all().catch(console.error);
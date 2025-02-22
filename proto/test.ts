/**
 * Generated by the protoc-gen-ts.  DO NOT EDIT!
 * compiler version: 3.21.12
 * source: test.proto
 * git: https://github.com/thesayyn/protoc-gen-ts */
import * as dependency_1 from "./google/protobuf/wrappers";
import * as grpc_1 from "@grpc/grpc-js";
interface GrpcUnaryServiceInterface<P, R> {
    (message: P, metadata: grpc_1.Metadata, options: grpc_1.CallOptions, callback: grpc_1.requestCallback<R>): grpc_1.ClientUnaryCall;
    (message: P, metadata: grpc_1.Metadata, callback: grpc_1.requestCallback<R>): grpc_1.ClientUnaryCall;
    (message: P, options: grpc_1.CallOptions, callback: grpc_1.requestCallback<R>): grpc_1.ClientUnaryCall;
    (message: P, callback: grpc_1.requestCallback<R>): grpc_1.ClientUnaryCall;
}
interface GrpcStreamServiceInterface<P, R> {
    (message: P, metadata: grpc_1.Metadata, options?: grpc_1.CallOptions): grpc_1.ClientReadableStream<R>;
    (message: P, options?: grpc_1.CallOptions): grpc_1.ClientReadableStream<R>;
}
interface GrpWritableServiceInterface<P, R> {
    (metadata: grpc_1.Metadata, options: grpc_1.CallOptions, callback: grpc_1.requestCallback<R>): grpc_1.ClientWritableStream<P>;
    (metadata: grpc_1.Metadata, callback: grpc_1.requestCallback<R>): grpc_1.ClientWritableStream<P>;
    (options: grpc_1.CallOptions, callback: grpc_1.requestCallback<R>): grpc_1.ClientWritableStream<P>;
    (callback: grpc_1.requestCallback<R>): grpc_1.ClientWritableStream<P>;
}
interface GrpcChunkServiceInterface<P, R> {
    (metadata: grpc_1.Metadata, options?: grpc_1.CallOptions): grpc_1.ClientDuplexStream<P, R>;
    (options?: grpc_1.CallOptions): grpc_1.ClientDuplexStream<P, R>;
}
interface GrpcPromiseServiceInterface<P, R> {
    (message: P, metadata: grpc_1.Metadata, options?: grpc_1.CallOptions): Promise<R>;
    (message: P, options?: grpc_1.CallOptions): Promise<R>;
}
export abstract class UnimplementedTestService {
    static definition = {
        Unary: {
            path: "/proto.Test/Unary",
            requestStream: false,
            responseStream: false,
            requestSerialize: (message: dependency_1.StringValue) => Buffer.from(message.serialize()),
            requestDeserialize: (bytes: Buffer) => dependency_1.StringValue.deserialize(new Uint8Array(bytes)),
            responseSerialize: (message: dependency_1.StringValue) => Buffer.from(message.serialize()),
            responseDeserialize: (bytes: Buffer) => dependency_1.StringValue.deserialize(new Uint8Array(bytes))
        },
        ClientStream: {
            path: "/proto.Test/ClientStream",
            requestStream: true,
            responseStream: false,
            requestSerialize: (message: dependency_1.StringValue) => Buffer.from(message.serialize()),
            requestDeserialize: (bytes: Buffer) => dependency_1.StringValue.deserialize(new Uint8Array(bytes)),
            responseSerialize: (message: dependency_1.StringValue) => Buffer.from(message.serialize()),
            responseDeserialize: (bytes: Buffer) => dependency_1.StringValue.deserialize(new Uint8Array(bytes))
        },
        ServerStream: {
            path: "/proto.Test/ServerStream",
            requestStream: false,
            responseStream: true,
            requestSerialize: (message: dependency_1.StringValue) => Buffer.from(message.serialize()),
            requestDeserialize: (bytes: Buffer) => dependency_1.StringValue.deserialize(new Uint8Array(bytes)),
            responseSerialize: (message: dependency_1.StringValue) => Buffer.from(message.serialize()),
            responseDeserialize: (bytes: Buffer) => dependency_1.StringValue.deserialize(new Uint8Array(bytes))
        },
        BidiStream: {
            path: "/proto.Test/BidiStream",
            requestStream: true,
            responseStream: true,
            requestSerialize: (message: dependency_1.StringValue) => Buffer.from(message.serialize()),
            requestDeserialize: (bytes: Buffer) => dependency_1.StringValue.deserialize(new Uint8Array(bytes)),
            responseSerialize: (message: dependency_1.StringValue) => Buffer.from(message.serialize()),
            responseDeserialize: (bytes: Buffer) => dependency_1.StringValue.deserialize(new Uint8Array(bytes))
        }
    };
    [method: string]: grpc_1.UntypedHandleCall;
    abstract Unary(call: grpc_1.ServerUnaryCall<dependency_1.StringValue, dependency_1.StringValue>, callback: grpc_1.sendUnaryData<dependency_1.StringValue>): void;
    abstract ClientStream(call: grpc_1.ServerReadableStream<dependency_1.StringValue, dependency_1.StringValue>, callback: grpc_1.sendUnaryData<dependency_1.StringValue>): void;
    abstract ServerStream(call: grpc_1.ServerWritableStream<dependency_1.StringValue, dependency_1.StringValue>): void;
    abstract BidiStream(call: grpc_1.ServerDuplexStream<dependency_1.StringValue, dependency_1.StringValue>): void;
}
export class TestClient extends grpc_1.makeGenericClientConstructor(UnimplementedTestService.definition, "Test", {}) {
    constructor(address: string, credentials: grpc_1.ChannelCredentials, options?: Partial<grpc_1.ChannelOptions>) {
        super(address, credentials, options);
    }
    Unary: GrpcPromiseServiceInterface<dependency_1.StringValue, dependency_1.StringValue> = (message: dependency_1.StringValue, metadata?: grpc_1.Metadata | grpc_1.CallOptions, options?: grpc_1.CallOptions): Promise<dependency_1.StringValue> => { if (!metadata) {
        metadata = new grpc_1.Metadata;
    } if (!options) {
        options = {};
    } return new Promise((resolve, reject) => super.Unary(message, metadata, options, (error: grpc_1.ServiceError, response: dependency_1.StringValue) => {
        if (error) {
            reject(error);
        }
        else {
            resolve(response);
        }
    })); };
    ClientStream: GrpWritableServiceInterface<dependency_1.StringValue, dependency_1.StringValue> = (metadata: grpc_1.Metadata | grpc_1.CallOptions | grpc_1.requestCallback<dependency_1.StringValue>, options?: grpc_1.CallOptions | grpc_1.requestCallback<dependency_1.StringValue>, callback?: grpc_1.requestCallback<dependency_1.StringValue>): grpc_1.ClientWritableStream<dependency_1.StringValue> => {
        return super.ClientStream(metadata, options, callback);
    };
    ServerStream: GrpcStreamServiceInterface<dependency_1.StringValue, dependency_1.StringValue> = (message: dependency_1.StringValue, metadata?: grpc_1.Metadata | grpc_1.CallOptions, options?: grpc_1.CallOptions): grpc_1.ClientReadableStream<dependency_1.StringValue> => {
        return super.ServerStream(message, metadata, options);
    };
    BidiStream: GrpcChunkServiceInterface<dependency_1.StringValue, dependency_1.StringValue> = (metadata?: grpc_1.Metadata | grpc_1.CallOptions, options?: grpc_1.CallOptions): grpc_1.ClientDuplexStream<dependency_1.StringValue, dependency_1.StringValue> => {
        return super.BidiStream(metadata, options);
    };
}

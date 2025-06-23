const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');

const PROTO_PATH = __dirname + '/protos/order.proto'
const packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {
        keepCase: true,
        longs: String,
        enums: String,
        defaults: true,
        oneofs: true
    }
);

const orderService = grpc.loadPackageDefinition(packageDefinition).order;
const orderClient = new orderService.Order('localhost:5003', grpc.credentials.createInsecure());

module.exports = orderClient;
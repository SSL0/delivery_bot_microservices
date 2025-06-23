const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');

class Client {
    constructor(protoPath) {
        const packageDefinition = protoLoader.loadSync(
            protoPath,
            {
                keepCase: true,
                longs: String,
                enums: String,
                defaults: true,
                oneofs: true
            }
        );
        
        this.protoDescriptor = grpc.loadPackageDefinition(packageDefinition);
    }
}

module.exports = Client;
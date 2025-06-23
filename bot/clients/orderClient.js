const grpc = require('@grpc/grpc-js');

const Client = require('./client')

class OrderClient extends Client {
    constructor(protoPath, serviceAddress){
        super(protoPath);
        this.stub = new this.protoDescriptor.order.Order(serviceAddress, grpc.credentials.createInsecure());
    }

    async createOrderByCart(cartId) {
        return new Promise((resolve, reject) => {
            this.stub.createOrderByCart({"cart_id": cartId}, (err, res) => {
                if (err) return reject(err);
                resolve(res);
            });
        });
    }
}

const PROTO_PATH = __dirname + '/protos/order.proto';
const ORDER_SERVICE_ADDRESS = 'order_service:5003';

const orderClient = new OrderClient(PROTO_PATH, ORDER_SERVICE_ADDRESS);

module.exports = orderClient
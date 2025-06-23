const grpc = require('@grpc/grpc-js');

const Client = require('./client')

class CartClient extends Client {
    constructor(protoPath, serviceAddress){
        super(protoPath);
        this.stub = new this.protoDescriptor.cart.Cart(serviceAddress, grpc.credentials.createInsecure());
    }

    async getCartIdByUserId(userId) {
        return new Promise((resolve, reject) => {
            this.stub.getCartIdByUserId({"user_id": userId}, (err, res) => {
                if (err) return reject(err);
                resolve(res.cart_id);
            });
        });
    }

    async getCart(cartId){
        return new Promise((resolve, reject) => {
            this.stub.getCart({"cart_id": cartId}, (err, res) => {
                if (err) return reject(err);
                resolve(res);
            });
        });
    
    }
}

const PROTO_PATH = __dirname + '/protos/cart.proto';
const CART_SERVICE_ADDRESS = 'localhost:5002';

const cartClient = new CartClient(PROTO_PATH, CART_SERVICE_ADDRESS);

module.exports = cartClient
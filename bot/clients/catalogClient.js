const grpc = require('@grpc/grpc-js');

const Client = require('./client')

class CatalogClient extends Client {
    constructor(protoPath, serviceAddress){
        super(protoPath);
        this.stub = new this.protoDescriptor.catalog.Catalog(serviceAddress, grpc.credentials.createInsecure());
    }

    async getProductByType(type){
        return new Promise((resolve, reject) => {
            this.stub.getProductsByType({"type": type}, (err, res) => {
                if (err) return reject(err);
                resolve(res.products);
            });
        });
    }

    async getProduct(id){
        return new Promise((resolve, reject) => {
            this.stub.getProduct({"id": id}, (err, res) => {
                if (err) return reject(err);
                resolve(res.product);
            });
        });
    }

    async getProductToppings(id){
        return new Promise((resolve, reject) => {
            this.stub.getProductToppings({"id": id}, (err, res) => {
                if (err) return reject(err);
                resolve(res.toppings);
            });
        });
    }

    async getTopping(id){
        return new Promise((resolve, reject) => {
            this.stub.getTopping({"id": id}, (err, res) => {
                if (err) return reject(err);
                resolve(res.topping);
            });
        });
    }
}

const PROTO_PATH = __dirname + '/protos/catalog.proto';
const CATALOG_SERVICE_ADDRESS = 'localhost:5001';

const catalogClient = new CatalogClient(PROTO_PATH, CATALOG_SERVICE_ADDRESS);

module.exports = catalogClient;
const { Markup, Scenes } = require('telegraf');

const cartScene = new Scenes.BaseScene('cart');

const cartClient = require('../../clients/cartClient')
const catalogClient = require('../../clients/catalogClient');
const { loadFileDescriptorSetFromObject } = require('@grpc/proto-loader');

cartScene.enter(async (ctx) => {
    const userId = 1001;
    const cartId = await cartClient.getCartIdByUserId(userId);
    const userCart = await cartClient.getCart(cartId);

    const displayCart = {}

    for(let i = 0; i < userCart.items.length; i++){
        const item = userCart.items[i];
        if(item.type == "product") {
            const product = await catalogClient.getProduct(item.item_id);

            displayCart[`${item.item_id}`] = {"name": product.name, "price": product.price, "quantity": item.quantity, "toppings": []}           
        } else if (item.type == "topping") {
            const topping = await catalogClient.getTopping(item.item_id);

            displayCart[`${topping.product_id}`].toppings.push({"name": topping.name, "price": topping.price, "quantity": topping.quantity})

        } else {
            console.error("unknown item type");
            return;
        }
    }

    const keyboard = Markup.inlineKeyboard([
            Markup.button.callback('Подтвердить заказ', 'order'),
            Markup.button.callback('Удалить из товар', 'delete'),
            Markup.button.callback('Очистить корзину', 'clear_cart'),
        ],
        { columns: 2 }
    );
    ctx.session.cart = displayCart;
    let output = '';
    let totalCost = 0.0;

    for (const [key, value] of Object.entries(displayCart)) {
        const price = parseFloat(value.price.substring(1));
        totalCost += price;
        output += `${value.name}(${value.price}) * ${value.quantity} = $${price * parseFloat(value.quantity)}\n`;
        value.toppings.forEach(topping => {
            output += `\t+ ${topping.name}(${topping.price})\n`;
        });
        output += '\n';
    }
    output += `Итого $${totalCost}`;
    ctx.reply(output, keyboard)
});

cartScene.action('order', (ctx) => ctx.scene.enter('order'));
cartScene.action('to_catalog', (ctx) => ctx.scene.enter('catalog'));
cartScene.action('clear_cart', async (ctx) => {
    await UserCart.findOneAndDelete({ userId: ctx.from.id });
    await ctx.editMessageText('Корзина очищена');
    await ctx.scene.enter('catalog');
});
  
module.exports = cartScene;
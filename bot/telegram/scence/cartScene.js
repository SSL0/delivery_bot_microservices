const { Markup, Scenes } = require('telegraf');

const cartScene = new Scenes.BaseScene('cart');

const cartClient = require('../../clients/cartClient')
const catalogClient = require('../../clients/catalogClient');
const orderClient = require('../../clients/orderClient');

cartScene.enter(async (ctx) => {
    const userId = ctx.from.id;

    const cartId = await cartClient.getCartIdByUserId(userId);
    ctx.session.cartId = cartId;

    const userCart = await cartClient.getCart(cartId);

    if(userCart.items.length == 0) {
        await ctx.reply('Корзина пуста, выберете товары в каталоге.');
        return;
    }

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

    for (const [_, value] of Object.entries(displayCart)) {
        const price = parseFloat(value.price.substring(1));
        totalCost += price;
        output += `${value.name}(${value.price}) * ${value.quantity} = $${price * parseFloat(value.quantity)}\n`;
        value.toppings.forEach(topping => {
            totalCost += parseFloat(topping.price.substring(1));
            output += `\t+ ${topping.name}(${topping.price})\n`;
        });
        output += '\n';
    }
    output += `Итого $${totalCost}`;
    await ctx.reply(output, keyboard)
});

cartScene.action('order', async (ctx) => {
    await orderClient.createOrderByCart(ctx.session.cartId);
    await ctx.editMessageText('Заказ успешно сделан, ожидайте доставки');
});

cartScene.action('clear_cart', async (ctx) => {
    await cartClient.removeCart(ctx.session.cartId);
    await ctx.editMessageText('Корзина очищена');
});
  
module.exports = cartScene;
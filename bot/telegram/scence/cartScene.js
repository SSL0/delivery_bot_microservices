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

    const displayCart = [];

    for (const cartItem of userCart.items) {
        if (cartItem.type === "product") {
            const product = await catalogClient.getProduct(cartItem.item_id);
    
            displayCart.push({
                cart_item_id: cartItem.id,
                product_id: cartItem.item_id,
                name: product.name,
                price: product.price,
                quantity: cartItem.quantity,
                toppings: []
            });
    
        } else if (cartItem.type === "topping") {
            const topping = await catalogClient.getTopping(cartItem.item_id);
    
            const parentProduct = displayCart.findLast(
                item => item.product_id === topping.product_id
            );
    
            if (!parentProduct) {
                console.error("Topping without parent product:", topping);
                continue;
            }

            parentProduct.toppings.push({
                cart_item_id: cartItem.id,
                name: topping.name,
                price: topping.price,
                quantity: cartItem.quantity
            });
            
        } else {
            console.error("Unknown item type:", cartItem.type);
            return;
        }
    }

    const keyboard = Markup.inlineKeyboard([
            Markup.button.callback('Подтвердить заказ', 'order'),
            Markup.button.callback('Удалить товар', 'delete'),
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

cartScene.action('delete', async (ctx) => {
    const buttons = [];
    for (const [_, value] of Object.entries(ctx.session.cart)) {
        buttons.push(Markup.button.callback(value.name, `delete_${value.cart_item_id}`));
    }

    const keyboard = Markup.inlineKeyboard(buttons, {columns: 1});
    await ctx.reply('Какой товар вы хотите удалить из корзины?', keyboard);
});


cartScene.action(/delete_(.+)/, async (ctx) => {
    const cartItemId = ctx.match[1];
    const cartItem = ctx.session.cart.find(
        item => item.cart_item_id == cartItemId
    );

    for(const topping of cartItem.toppings) {
        await cartClient.removeItem(topping.cart_item_id);
    }

    await cartClient.removeItem(cartItemId);

    await ctx.editMessageText('Товар успешно удален из корзины');
    await ctx.scene.enter('cart');
});

cartScene.action('order', async (ctx) => {
    ctx.session.cart = {};
    await orderClient.createOrderByCart(ctx.session.cartId);
    await ctx.editMessageText('Заказ успешно сделан, ожидайте доставки');
});

cartScene.action('clear_cart', async (ctx) => {
    ctx.session.cart = {};
    await cartClient.removeCart(ctx.session.cartId);
    await ctx.editMessageText('Корзина очищена');
});
  
module.exports = cartScene;
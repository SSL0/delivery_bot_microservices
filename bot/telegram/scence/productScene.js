const { Markup, Scenes } = require('telegraf');
const catalogClient = require('../../clients/catalogClient');
const cartClient = require('../../clients/cartClient');

const productScene = new Scenes.BaseScene('product');
productScene.enter(async (ctx) => {
    const productId = ctx.session.productId;

    const product = await catalogClient.getProduct(productId);
   
    if (!product) {
      await ctx.reply('Товар не найден');
      return ctx.scene.leave();
    }
    
    const keyboard = Markup.inlineKeyboard([
      [Markup.button.callback('➕ Добавить в корзину', 'add_to_cart')],
      [Markup.button.callback('🔙 Назад', 'back_to_catalog')]
    ]);
    
    let message = `*${product.name}*\n\n`;
    message += `${product.description}\n\n`;
    message += `*Цена:* ${product.price}`;
    
    if (product.image) {
      await ctx.replyWithPhoto(product.image, {
        caption: message,
        parse_mode: 'Markdown',
        ...keyboard
      });
    } else {
      await ctx.reply(message, { parse_mode: 'Markdown', ...keyboard });
    }
});

productScene.action('add_to_cart', async (ctx) => {
    const productId = ctx.session.productId;
    const product = await catalogClient.getProduct(productId);

    const toppings = await catalogClient.getProductToppings(productId);

    if (toppings.length > 0) {
      ctx.session.cartItem = { productId, toppingsIds: [] };
      
      const keyboard = Markup.inlineKeyboard([
        ...toppings.map(topping => [
          Markup.button.callback(
            `${topping.name} +${topping.price}`, 
            `topping_${topping.id}`
          )
        ]),
        [Markup.button.callback('✅ Завершить выбор', 'finish_toppings')]
      ], { columns: 2 });
      
      await ctx.editMessageText(`Выберите топпинги для "${product.name}":`, keyboard);
    } else {
      await addToCart(ctx, productId, []);
      await ctx.editMessageText(`Товар "${product.name}" добавлен в корзину!`);
      await ctx.scene.enter('catalog');
    }
});
  
productScene.action(/topping_(.+)/, async (ctx) => {
    const toppingId = ctx.match[1];
    const topping = await catalogClient.getTopping(toppingId);
    
    if (!ctx.session.cartItem.toppingsIds.includes(toppingId)) {
      ctx.session.cartItem.toppingsIds.push(toppingId);
      await ctx.answerCbQuery(`Добавлен: ${topping.name}`);
    } else {
      ctx.session.cartItem.toppingsIds = ctx.session.cartItem.toppingsIds.filter(id => id !== toppingId);
      await ctx.answerCbQuery(`Удалён: ${topping.name}`);
    }
});
  
productScene.action('finish_toppings', async (ctx) => {
    const { productId, toppingsIds } = ctx.session.cartItem;

    const product = await catalogClient.getProduct(productId);

    const addedCartItemId = await cartClient.addItem(productId, 'product', 1, ctx.from.id);
    if(addedCartItemId === undefined) {
        console.error(`failed to add product with id ${productId} to cart`);
        return;
    }
    toppingsIds.forEach(async (toppingId) => {  
        const addedCartItemId = await cartClient.addItem(toppingId, 'topping', 1, ctx.from.id);
        if(addedCartItemId === undefined) {
            console.error(`failed to add topping with id ${toppingId} to cart`);
            return;
        }
    });

    await ctx.editMessageText(`Товар "${product.name}" с выбранными топпингами добавлен в корзину!`);
    await ctx.scene.enter('catalog');
});
  
productScene.action('back_to_catalog', (ctx) => ctx.scene.enter('catalog'));
  
module.exports = productScene;
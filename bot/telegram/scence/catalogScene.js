const { Markup, Scenes } = require('telegraf');
const catalogScene = new Scenes.BaseScene('catalog');

const catalogClient = require('../../clients/catalogClient')

catalogScene.enter(async (ctx) => {
    const productTypes = [['pizza', 'Пицца'], ['burger', 'Бургер'], ['drink', 'Напиток'], ['side', 'Другое']];
    const keyboard = Markup.inlineKeyboard(
    	productTypes.map(type => [Markup.button.callback(type[1], `type_${type[0]}`)]),
    );
	
    await ctx.reply('Выберите категорию товара:', keyboard);
});

catalogScene.action(/type_(.+)/, async (ctx) => {
    const type = ctx.match[1];

    const products = await catalogClient.getProductByType(type);
    const keyboard = Markup.inlineKeyboard(
        products.map(product => [Markup.button.callback(
			`${product.name} - ${product.price}`, 
			`product_${product.id}`
        )]),
        { columns: 1 }
	);
    await ctx.editMessageText(`Товары в категории "${type}":`, keyboard);
});

catalogScene.action(/product_(.+)/, async (ctx) => {
    const productId = ctx.match[1];
    ctx.session.productId = productId;
    await ctx.scene.enter('product');
});

module.exports = catalogScene;
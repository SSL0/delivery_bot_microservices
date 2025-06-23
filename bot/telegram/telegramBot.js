const { Telegraf, Markup, Scenes, session } = require('telegraf');

const scenes = require('./scenes')


class TelegramBot {
    constructor(botToken){
        this._telegraf = new Telegraf(botToken);
        this._scences = {}
        this.registerScences();
    }

    registerScences() {
        this._telegraf.start(async (ctx) => {
            await ctx.reply('Добро пожаловать в сервис доставки!', Markup.keyboard([
              ['Каталог'],
              ['Корзина'],
              ['Помощь']
            ]).resize());
        });

        const stage = new Scenes.Stage(scenes);

        this._telegraf.use(session());
        this._telegraf.use(stage.middleware());

        this._telegraf.hears('Каталог', (ctx) => ctx.scene.enter('catalog'));
        this._telegraf.hears('Корзина', (ctx) => ctx.scene.enter('cart'));
        this._telegraf.hears('Помощь', (ctx) => ctx.reply(
            'Этот бот позволяет просматривать каталог товаров, добавлять их в корзину с топпингами и оформлять заказы.'
        ));
    }

    Start() {
        this._telegraf.launch();
    }
}

module.exports = {
    TelegramBot,
}
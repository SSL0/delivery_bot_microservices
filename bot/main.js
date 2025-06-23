const telgeramBot = require('./telegram/telegramBot')

const bot = new telgeramBot.TelegramBot(process.env.BOT_TOKEN);
bot.Start();
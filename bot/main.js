const telgeramBot = require('./telegram/telegramBot')

const clients = require('./clients/catalogClient')

const bot = new telgeramBot.TelegramBot(process.env.BOT_TOKEN);
bot.Start();
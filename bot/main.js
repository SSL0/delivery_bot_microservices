const telgeramBot = require('./telegram/telegramBot')

const clients = require('./clients/catalogClient')

const bot = new telgeramBot.TelegramBot('7925301168:AAG_PYl1sZw6AGSDQ98jpR8MfIr9wJHYhxw');
bot.Start();
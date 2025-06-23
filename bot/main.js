const telgeramBot = require('./telegram/telegramBot')

const clients = require('./clients/catalogClient')

const bot = new telgeramBot.TelegramBot('7925301168:AAFWhLdiRGYcr-OOgVoALkGnAtVhgSZvDJs');
bot.Start();
<p align="center">
      <img src="https://cashbee.me/img/shops/fl.svg" width="700">
</p>

<p align="center">
   <img src="https://img.shields.io/badge/Golang-version%201.18-blue" alt="Golang">   
   <img src="https://img.shields.io/badge/GoFLParcerXML-version%201.0-blue" alt="Project-Version">
</p>

## About

This project was created to optimize the search for orders on FL.RU. The reason for creating it is simple: the site does not show notifications of new orders for all categories, but only for selected ones. As for blocking, there will be no blocking, because when parsing XML file is used proxy SOCKS5 and your account is in complete safety. 

New requests come to you on Telegram.

## Documentation

Required language version **Golang 1.18**

## Build

For this application to work, you need to edit **message** in the **main.go** file located in **cmd/bot** in the root directory of the project and insert the chat ID of your Telegram account there.

Required language version **Golang 1.18**

After all the above steps you can run the application with the flag **"-tgBotToken "** and specify your bot's token, the token can be obtained by @BotFather.

The full startup command **"go run main.go -tgBotToken 5437359584:AAGHGhwi8nPaTKQHK7XNAjHHHH-rGhjgfdhyT "**

## Developers

- [OneByteForLife](https://github.com/OneByteForLife)

## License

- This software is protected under the MIT license!

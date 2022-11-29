# QRbot
This is a repository of a study project created as part of learning the golang programming language.
Telegram bot that creates a QR according to the text sent to it. It also has special settings for customization.
The QR generator was taken from [yeqown/go-qrcode](https://github.com/yeqown/go-qrcode) but further a little bit modified [1azar/go-qrcode](https://github.com/1azar/go-qrcode/tree/WithHalfTone-File-Image).

![QR Chan gif](https://github.com/1azar/QRbot/blob/v.1/assets/qrchan.gif)

# Implementation Details
![Project Graph](https://github.com/1azar/QRbot/blob/v.1/assets/godepgraphMajorNodes.png)

Unlike [v.0](https://github.com/1azar/QRbot/tree/v.0), the code in this [v.1](https://github.com/1azar/QRbot) branch has been designed to be as close as possible to the principles of [Uncle Bob's clean architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html).
I contains 4 specific packages (exept `main`):

## `domain pkg`
Contains the most stable entities dictated by business rules. This pkg has no external dependances. These are the QR settings structure for the user, the interface for the QR repository, the interface for the QR generator, and the QR entity. Also here declarated some constants.

## `usecases pkg`

  



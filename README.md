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
Contains the most stable entities dictated by business rules. This pkg has no external dependances. These are:
- The `QR settings structure` for the user. Also has constructor for default QR settings for new users.
- The `interface of QR repository`.
- The `interface of QR generator`.
- The `QR entity`. 
- Also here declarated some `constants`.

## `usecases pkg`
Contains application specific business rules. It encapsulates and implements all of the use cases of the app. These use cases orchestrate the flow of data to and from the entities. The following is implemented here:
- `Logger interface`. 
- `QR interactor struct`. Countains `QR Settings repository` within (domain entity), `QR Settings Buffer` - special struct which is needed to reduce data base interaction during setting QR Options by user. `Logger` to log important events of the app. Also `Find` and `Store` methods.
  



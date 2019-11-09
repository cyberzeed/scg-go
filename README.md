# scg-go
The assignment project from SCG recruiter. There are three module in this project.

## Series

This module is API services which be created to solve mathematical problems related to a set of numbers consisting of sequential numbers.

```
X, 5, 9, 15, 23, Y, Z
```

### API

__Find all number in series__

This API will present set of number in series follow specified size.

```
GET https://scg-go-odjtm7kfta-uc.a.run.app/series?size=<size>
```

Example
```
https://scg-go-odjtm7kfta-uc.a.run.app/series?size=7
```

__Find number in series by specified position__

When you assign a index of numbers to this API. It will calculate the value of position in series.

```
GET https://scg-go-odjtm7kfta-uc.a.run.app/series/<index>
```

Example
```
https://scg-go-odjtm7kfta-uc.a.run.app/series/0
https://scg-go-odjtm7kfta-uc.a.run.app/series/1
https://scg-go-odjtm7kfta-uc.a.run.app/series/2
```


## Restaurant

This module has a API which find the restaurants in specified area and present the results in JSON format.

### API

```
GET https://scg-go-odjtm7kfta-uc.a.run.app/restaurant/<area>
```

Example
```
https://scg-go-odjtm7kfta-uc.a.run.app/restaurant/bangsue
```

## Line

This module is API services for Line chatbot.

### API

__Broadcast message__

This API will send messages to everyone which follow Line account.
```
POST https://scg-go-odjtm7kfta-uc.a.run.app/message/broadcast
```

Example of request body
```json
{
    "messages": ["Hello world", "I am SCG-Go Bot."]
}
```

__Webhook__

This webhook will wait to receive requests and process events which be send from Line server. By default when user send message or sticker in chat room the chatbot will send back by same message or sticker. But if user send the text message that match with command the chatbot will execute that comand and send the result to user.

### Command

There are two command in webhook.

* series: \<index>
* restaurant: \<area>

The first command will call series API to find value of index.
The second command will call restaurant API and present name and address of restaurant in area.

### SCG-Go Line Account

Scan QR code for add SCG-Go account:

![Line Bot](https://raw.githubusercontent.com/cyberzeed/scg-go/master/qr-code.png)
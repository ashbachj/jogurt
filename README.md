# jogurt
This repository is used to manage a bot to alert users of the [Retro Game Boards](https://www.retrogameboards.com) Discord about updates to various websites. 

# How to use:
[Install Golang on your machine](https://golang.org/doc/install#install)

[Create a new Discord App](https://discordapp.com/developers/applications/me)

#### Add your bot to the server:
<https://discordapp.com/oauth2/authorize?client_id=${Client-ID}&scope=bot>

#### Install dependencies
Run the following command before running the application to install the discordgo library
```bash
go get https://github.com/bwmarrin/discordgo
```

#### Run the application
```bash
go run src/jogurt.go -t "${TOKEN}"
```

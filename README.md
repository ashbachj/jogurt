# jogurt
This repository is used to manage a bot to alert users of the [a Retro Game Boards](https://www.retrogameboards.com) Discord about updates to various websites. 

# How to use:
[a Install go on your machine](https://golang.org/doc/install#install)

[a Create a new Discord App](https://discordapp.com/developers/applications/me)

#### Add your bot to the server:
<https://discordapp.com/oauth2/authorize?client_id=${Client-ID}&scope=bot>

#### Run the application
```bash
go run src/jogurt.go -t "${TOKEN}"
```

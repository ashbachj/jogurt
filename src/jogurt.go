package main

import (
       "flag"
       "fmt"
       "bytes"
       "strings"
       "os"
       "os/signal"
       "syscall"
       "net/http"
       "io/ioutil"
       "encoding/json"

       "github.com/bwmarrin/discordgo"
)

var (
    Token	string
)

func init() {
     flag.StringVar(&Token, "t", "", "Owner Account Token") 
     flag.Parse()
     
     if Token == "" {
     	flag.Usage()
	os.Exit(1)
     }
}

func main() {
     discord, err := discordgo.New("Bot " + Token)

     if err != nil {
     	fmt.Println("Error establishing session!", err)
	return
     }
     fmt.Printf("Discord session made with token: %s\n", discord.Token)

     // Register a function to handle message events
     discord.AddHandler(messageCreate)

     // Open a websocket connection to Discord
     err = discord.Open()
     if err != nil {
     	fmt.Println("Something went wrong opening a connection!\n", err)
     }

     fmt.Println("Bot is now running. Press CTRL-C to exit.")
     sc := make(chan os.Signal, 1)
     signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
     <-sc

     // Finally
     discord.Close()
}

func getSuperNtJailbreak() string {
     var buffer bytes.Buffer
     var previous string
     
     url := "https://api.github.com/repos/SmokeMonsterPacks/Super-NT-Jailbreak/releases/latest"
     resp, err := http.Get(url)
     if err != nil {
     	fmt.Println("Error getting latest release", err)
	return "error"
     }
     defer resp.Body.Close()
     body, err := ioutil.ReadAll(resp.Body)

     if err != nil {
     	fmt.Println("Error reading response body", err)
	return "error"
     }

     result := make(map[string]interface{})
     json.Unmarshal(body, &result)

     if _, err := os.Stat("superNtJb.out"); err == nil {
     	in, err := ioutil.ReadFile("superNtJb.out")
	if err != nil {
	   fmt.Println("Error reading superNtJb.out", err)
	} else {
	  previous = string(in)
	}
     }

     latest := result["tag_name"].(string)
     if (strings.Compare(previous, latest) != 0) {
          err = ioutil.WriteFile("superNtJb.out", []byte(latest), 0644)
	  buffer.WriteString("Latest SuperNT Jailbreak: ")
	  buffer.WriteString(latest)
	  buffer.WriteString("\n")
	  buffer.WriteString(result["url"].(string))
     }

     return buffer.String()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
     // Ignore yourself
     if m.Author.ID == s.State.User.ID {
     	return
     }

     superNT := getSuperNtJailbreak()

     // If message is test, tell the truth
     if m.Content == "test" {
          s.ChannelMessageSend(m.ChannelID, "Sega is better than Nintendo")
	  if (strings.Compare(superNT, "") != 0) {
	    s.ChannelMessageSend(m.ChannelID, superNT)
	  }
     }
}

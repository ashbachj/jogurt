package main

import (
       "flag"
       "fmt"
       "bytes"
       "strings"
       "os"
       "os/signal"
       "syscall"
       "time"
       "net/http"
       "io/ioutil"
       "encoding/json"
       "math/rand"

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
     rand.Seed(time.Now().UTC().UnixNano())

     if err != nil {
     	fmt.Println("Error establishing session!", err)
	return
     }
     fmt.Printf("Discord session made with token: %s\n", discord.Token)

     // Register a function to handle message events
     discord.AddHandler(messageCreate)

     // Register a function to handle ready events
     discord.AddHandler(ready)

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

func getURL(url string) []uint8 {
     resp, err := http.Get(url)
     var body []uint8
     if err != nil {
     	fmt.Println("Error getting latest release", err)
	body = []uint8{0}
	return body
     }
     
     defer resp.Body.Close()
     body, err = ioutil.ReadAll(resp.Body)

     if err != nil {
     	fmt.Println("Error reading response body", err)
	body = []uint8{0}
	return body
     }

     return body
}

func readOutFile(filename string) string {
    var previous string
    if _, err := os.Stat(filename); err == nil {
      in, err := ioutil.ReadFile(filename)
      if err != nil {
        fmt.Println("Error reading superNtJb.out", err)
      } else {
        previous = string(in)
      }
    }

    return previous
}

func getRecentSuperNT() string {
     var previous string
     var buffer bytes.Buffer

     url := "https://github.com/SmokeMonsterPacks/Super-NT-Jailbreak/releases"
     previous = readOutFile("superNtJb.out")

     buffer.WriteString("Latest version is ")
     buffer.WriteString(previous)
     buffer.WriteString("\n\n")
     buffer.WriteString(url)

     return buffer.String()
}

func getSuperNtJailbreak() string {
    var buffer bytes.Buffer
    var previous string
    if (rand.Float64() < 0.8) {
      fmt.Println("Don't spam github")
      return ""
    }
     
    url := "https://api.github.com/repos/SmokeMonsterPacks/Super-NT-Jailbreak/releases/latest"

    body := getURL(url)
    if (len(body) <= 1) {
      fmt.Println("Error getting url")
      fmt.Println(url)
      return ""
    }
    result := make(map[string]interface{})
    json.Unmarshal(body, &result)

    previous = readOutFile("superNtJb.out")

    tag := result["tag_name"]
    if tag == nil {
       fmt.Println("Error reading results")
       fmt.Println(result)
       return buffer.String()
    }

    latest := tag.(string)

    if (strings.Compare(previous, latest) != 0) {
         err := ioutil.WriteFile("superNtJb.out", []byte(latest), 0644)
	 if err != nil {
	   fmt.Println("Error writing to file superNtJb.out")
	 }
         buffer.WriteString("@here Latest SuperNT Jailbreak: ")
         buffer.WriteString(latest)
         buffer.WriteString("\n\n")
         buffer.WriteString(result["url"].(string))
    }

    return buffer.String()
}

func getGDEmu(oem string) string {
    var buffer string

    var urlArray []string
    var orderOpen []string
    
    urlArray = append(urlArray, "https://gdemu.wordpress.com/ordering/ordering-") //gdemu/"
    urlArray = append(urlArray, oem) 
    url := strings.Join(urlArray, "")
    body := getURL(url)

    latest := string(body)

    if (strings.Contains(latest, "paypalemail' class='grunion-field-label email'>PayPal email<span>(required)</span>") == true) {
      orderOpen = append(orderOpen, "@here ")
      orderOpen = append(orderOpen, oem)
      orderOpen = append(orderOpen, " order page is open. Act fast!!\n\n")
      orderOpen = append(orderOpen, url)
      
      buffer = strings.Join(orderOpen, "")
      fmt.Println(latest)
    } else {
      fmt.Println(oem)
      fmt.Println("Preorders closed")
    } 

    return buffer
}

func waitUntilAligned(align time.Duration) {
  d := time.Until(time.Now().Add(align).Truncate(align))
  time.Sleep(d)
}

func ready(s *discordgo.Session, r *discordgo.Ready) {
  for {
    waitUntilAligned(3 * time.Minute)
    fmt.Println(time.Now())
    superNT := getSuperNtJailbreak()
    gdEmu := getGDEmu("gdemu")
    rhea := getGDEmu("rhea")
    phoebe := getGDEmu("phoebe")
    docbrown := getGDEmu("docbrown")

    s.ChannelMessageSend("446490232116871201", superNT)
    s.ChannelMessageSend("446490232116871201", gdEmu)
    if gdEmu != "" {
      time.Sleep(1 * time.Hour)
    }
    s.ChannelMessageSend("446490232116871201", rhea)
    if rhea != "" {
      time.Sleep(1 * time.Hour)
    }
    s.ChannelMessageSend("446490232116871201", phoebe)
    if phoebe != "" {
      time.Sleep(1 * time.Hour)
    }
    s.ChannelMessageSend("446490232116871201", docbrown)
    if docbrown != "" {
      time.Sleep(3 * time.Hour)
    }
  }
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
     // Ignore yourself
     if m.Author.ID == s.State.User.ID {
     	return
     }

     switch m.Content {
     // If message is !truth, tell the truth
       case "!truth": 
          s.ChannelMessageSend(m.ChannelID, "Sega is better than Nintendo")
       case "!superNT": 
          superNT := getRecentSuperNT()
          s.ChannelMessageSend(m.ChannelID, superNT)	
       default: 
          fmt.Println(m.Content)
     }
}

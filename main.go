package main

import "encoding/json"
import "fmt"
import "time"

import "github.com/r3labs/sse"
import "github.com/go-vgo/robotgo"

type Event struct {
	Type string              `json:"type"`
	Data []map[string]string `json:"data"`
}

func main() {
	fmt.Println("bpz")
	fpid, err := robotgo.FindIds("nes")
	if err != nil {
		panic(err)
	}
	if len(fpid) == 0 {
		fmt.Printf("No pid found")
		return
	}
	fmt.Printf("pid: %v", fpid)
	events := make(chan *sse.Event)
	client := sse.NewClient("https://bitgraph.network/s/ewogICJ2IjogMywKICAicSI6IHsKICAgICJmaW5kIjogeyJvdXQuYjAiOiB7ICJvcCI6IDEwNiB9LCAib3V0LmUuYSI6ICJxenc1ZWR6N3F1MDl1cHFlaHU0dHZ4ODJraGM5azAwejJ1MHRxOXU4Y2QifQogIH0sCiAgInIiOiB7CiAgICAiZiI6ICJbLltdIHwge21zZzogLm91dFswXS5zMX1dIgogIH0KfQ==")
	client.SubscribeChan("messages", events)
	client.SubscribeRaw(func(msg *sse.Event) {
		// Got some data!
		fmt.Printf("%#v\n", string(msg.Data))
		data := Event{}
		if len(msg.Data) == 0 {
			return
		}
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			panic(err)
		}
		if len(data.Data) == 0 {
			return
		}
		robotgo.ActiveName("nes")
		robotgo.ActivePID(fpid[0])
		fmt.Printf("data: %#v", data.Data[0]["msg"])
		robotgo.KeyToggle(data.Data[0]["msg"], "down")
		time.Sleep(250 * time.Millisecond)
		robotgo.KeyToggle(data.Data[0]["msg"], "up")
		time.Sleep(250 * time.Millisecond)
	})

	robotgo.ActiveName("nes")
	tit := robotgo.ActivePID(fpid[0])
	fmt.Printf("title: %v", tit)

	time.Sleep(5 * time.Second)
	robotgo.KeyToggle("right", "down")
	time.Sleep(500 * time.Millisecond)
	robotgo.KeyToggle("right", "up")
	time.Sleep(5 * time.Second)
	robotgo.KeyToggle("down", "down")
	time.Sleep(500 * time.Millisecond)
	robotgo.KeyToggle("down", "up")

}

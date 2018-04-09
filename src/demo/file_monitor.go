package demo

import (
	"github.com/fsnotify"
	"fmt"
	"log"
)

func NewMonitor() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("--err",err)
	}
	defer watcher.Close()

	fmt.Println("-->watcher", watcher)
	err = watcher.Add("config.toml")
	if err != nil {
		log.Fatal("---->watcher.Add:",err)
	}

	go func() {
		fmt.Println("--->1")
		for {
			fmt.Println("--->2")
			select {
			case event := <-watcher.Events:
				//log.Println("event:", event)
				fmt.Println("--event",event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
					//gConfig.ConfPath = "config_svr/config/config.toml" //"../conf/config.toml"
					//gConfig.Parse(&Iconfigs)
					//fmt.Println("-->Iconfigs",Iconfigs)
				}
				//nginx()
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	select {}
}
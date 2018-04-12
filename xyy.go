package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	"io"
	"github.com/hajimehoshi/oto"
	"github.com/hajimehoshi/go-mp3"
	"strconv"
	"time"
	"math"
	"regexp"
	"flag"
)

func play(song string)  {
	f, err := os.Open(song)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	d, err := mp3.NewDecoder(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer d.Close()

	p, err := oto.NewPlayer(d.SampleRate(), 2, 2, 8192)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer p.Close()
	if _, err := io.Copy(p, d); err != nil {
		fmt.Println(err)
		return
	}
}

func showLrc(lrc string)  {
	allwords, err := ioutil.ReadFile(lrc)
	if err != nil {
		fmt.Print(err)
	}
	str := string(allwords)
	lines := strings.Split(str, "\n")
	r, _ := regexp.Compile(`^\[\d{2}:\d{2}\.\d{2}\]`)
	var newTime, lastTime float64
	for _, line := range lines {
		if !r.MatchString(line){
			newTime += 0.3
		}else{
			minute, _ := strconv.ParseFloat(string([]rune(line)[2:3]), 0)
			second, _ := strconv.ParseFloat(string([]rune(line)[4:9]), 0)
			newTime = minute*60 + second
		}
		wait := time.Duration((newTime - lastTime) * math.Pow10(9))
		time.Sleep(wait)
		fmt.Println(line)
		lastTime = newTime
	}
}

func main() {
	song := flag.String("song", "qingchun.mp3", "播放的歌曲名称")
	lrc := flag.String("lrc", "qingchun.lrc", "lrc歌词文件")
	flag.Parse()
	go play(*song)
	showLrc(*lrc)
}
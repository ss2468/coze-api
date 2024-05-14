package coze

import (
	"context"
	"strings"
	"sync"
	"testing"
	"time"
)

const (
	/* fixme coze token和proxy */
	cookie     = ""
	msToken    = ""
	socksProxy = ""
)

func TestCookie(t *testing.T) {
	options := NewDefaultOptions("7353047124357365778", "1712645567468", 2, socksProxy)
	chat := New(cookie, msToken, options)
	t.Log(chat.makeCookie())
}

func TestChats(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(20)
	for i := 0; i < 20; i++ {
		go func() {
			TestChat(t)
			wg.Done()
		}()
	}
	wg.Wait()
	t.Log("finish")
}

func TestChat(t *testing.T) {
	options := NewDefaultOptions("7353047124357365778", "1712645567468", 2, socksProxy)
	chat := New(cookie, msToken, options)
	messages := []Message{
		{
			Role:    "system",
			Content: "你所有的回答要用英文返回，不得出现其他语言",
		},
		{
			Role:    "user",
			Content: "1900年有多少天？",
		},
	}

	timeout, withTimeout := context.WithTimeout(context.Background(), 120*time.Second)
	defer withTimeout()

	ch, err := chat.Reply(timeout, MergeMessages(messages))
	if err != nil {
		t.Fatal(err)
	}

	echo(ch, t)
}

func TestImages(t *testing.T) {
	options := NewDefaultOptions("7353047124357365778", "1712645567468", 2, socksProxy)
	chat := New(cookie, msToken, options)
	timeout, withTimeout := context.WithTimeout(context.Background(), 120*time.Second)
	defer withTimeout()

	image, err := chat.Images(timeout, "画一个二次元猫娘，1girl")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(image)
}

func echo(ch chan string, t *testing.T) {
	for {
		message, ok := <-ch
		if !ok {
			return
		}

		if strings.HasPrefix(message, "error:") {
			t.Fatal(message)
		}

		t.Log(message)
	}
}

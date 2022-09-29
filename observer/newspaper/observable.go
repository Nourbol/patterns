package newspaper

import (
	"fmt"
	"time"
)

type Publisher interface {
	NotifyAll(string)
	AddSubscriber(*Subscriber)
	RemoveSubscriber(Subscriber) error
}

type Newspaper struct {
	Title       string
	issues      [][]Page
	subscribers []*Subscriber
}

func NewNewspaper(title string) *Newspaper {
	return &Newspaper{
		Title: title,
	}
}

func (n *Newspaper) ReleaseIssue(pages []Page) {
	pagesNumber := len(pages)
	publishedDate := time.Now()
	for i := 0; i < pagesNumber; i++ {
		articleNumbers := len(pages[i].Articles)
		for j := 0; j < articleNumbers; j++ {
			pages[i].Articles[j].PublishedDate = publishedDate
		}
	}
	n.issues = append(n.issues, pages)
	n.NotifyAll("New issue was released!")
}

func (n *Newspaper) GetIssueByNumber(issueNo uint) []Page {
	return n.issues[issueNo]
}

func (n *Newspaper) AddSubscriber(subscriber *Subscriber) {
	n.subscribers = append(n.subscribers, subscriber)
}

func (n *Newspaper) RemoveSubscriber(subscriber Subscriber) error {
	index := -1
	for i := 0; i < len(n.subscribers); i++ {
		if *n.subscribers[i] == subscriber {
			index = i
		}
	}
	if index == -1 {
		return fmt.Errorf("subscriber not found")
	}
	n.subscribers = append(n.subscribers[:index], n.subscribers[index+1:]...)
	return nil
}

func (n *Newspaper) NotifyAll(message string) {
	for i := 0; i < len(n.subscribers); i++ {
		subscriber := *n.subscribers[i]
		subscriber.HandleEvent(message)
	}
}

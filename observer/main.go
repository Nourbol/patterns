package main

import (
	"fmt"
	n "patterns/observer/newspaper"
)

func main() {
	// Creating newspaper
	articles := []n.Article{
		*n.NewArticle("Alice", "Golang?", "..."),
		*n.NewArticle("Max", "Java!", "..."),
	}
	pages := []n.Page{
		*n.NewPage(articles...),
	}
	newspaper := n.NewNewspaper("Evrika")

	// Creating subscribers
	subscriber1 := n.Subscriber(n.NewPerson("Vlad"))
	subscriber2 := n.Subscriber(n.NewPerson("Bauka"))
	newspaper.AddSubscriber(&subscriber1)
	newspaper.AddSubscriber(&subscriber2)

	// Releasing issue
	newspaper.ReleaseIssue(pages)

	// Removing subscriber
	if err := newspaper.RemoveSubscriber(subscriber1); err != nil {
		panic(err)
	}
	fmt.Println("*************************************")

	// Releasing again
	newspaper.ReleaseIssue(pages)
}

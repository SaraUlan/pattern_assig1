package main

import "fmt"

type ListStrategy interface {
	AddItem(item string)
	GetItems() []string
}

type FavouriteList struct {
	items []string
}

func (p *FavouriteList) AddItem(item string) {
	p.items = append(p.items, item)
	fmt.Printf("This advertising '%p' added to favourite list. \n ", item)
}

func (p *FavouriteList) GetItems() []string {
	return p.items
}

type BuyList struct {
	items []string
}

func (p *BuyList) AddItem(item string) {
	p.items = append(p.items, item)
	fmt.Printf("This advertising '%p' added to buy list. \n ", item)

}

func (p *BuyList) GetItems() []string {
	return p.items
}

type ShareList struct {
	items []string
}

func (p *ShareList) AddItem(item string) {
	p.items = append(p.items, item)
	fmt.Printf("This advertising '%p' added to share list. \n ", item)
}

func (p *ShareList) GetItems() []string {
	return p.items
}

type SaveList struct {
	items []string
}

func (p *SaveList) AddItem(item string) {
	p.items = append(p.items, item)
	fmt.Printf("This advertising '%p' added to save list. \n ", item)
}

func (p *SaveList) GetItems() []string {
	return p.items
}

type ListContext struct {
	strategy ListStrategy
}

func (c *ListContext) SetStrategy(strategy ListStrategy) {
	c.strategy = strategy
}

func (c *ListContext) AddItemToList(item string) {
	c.strategy.AddItem(item)
}

func main() {
	context := ListContext{}

	FavouriteStrategy := &FavouriteList{}
	context.SetStrategy(FavouriteStrategy)
	context.AddItemToList("Laptop")

	buyStrategy := &BuyList{}
	context.SetStrategy(buyStrategy)
	context.AddItemToList("Pen")

	shareStrategy := &ShareList{}
	context.SetStrategy(shareStrategy)
	context.AddItemToList("lipstick")

	saveStrategy := &SaveList{}
	context.SetStrategy(saveStrategy)
	context.AddItemToList("Ipad")
}

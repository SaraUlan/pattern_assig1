package main

import "fmt"

type ProductCustomer interface {
	Update(message string)
}

type ProductManager struct {
	customer ProductCustomer
}

func NewProductManager(customer ProductCustomer) *ProductManager {
	return &ProductManager{
		customer: customer,
	}
}

func (em *ProductManager) Notify(message string) {
	em.customer.Update(message)
}

type Postamat struct {
	events *ProductManager
}

func NewPostamat(customer ProductCustomer) *Postamat {
	return &Postamat{
		events: NewProductManager(customer),
	}
}

func (p *Postamat) PackageReady() {
	message := "The product has arrived at the postamat machine and you can get it"
	p.events.Notify(message)
}

// main customer
type MainCustomer struct {
	Name string
}

func (c *MainCustomer) Update(message string) {
	fmt.Printf("%s received a notification: %s\n", c.Name, message)
}

func main() {
	customer := &MainCustomer{Name: "Customer1"}
	postamat := NewPostamat(customer)

	// Simulate the arrival of product at the postamat
	postamat.PackageReady()
}

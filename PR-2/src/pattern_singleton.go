package main

import (
	"fmt"
	"sync"
)

// PaymentConfig - конфигурация платёжного модуля.
// В системе должен существовать только один экземпляр конфигурации.
type PaymentConfig struct {
	provider   string
	apiURL     string
	timeoutSec int
}

var (
	instance *PaymentConfig
	once     sync.Once
)

// GetInstance возвращает единственный экземпляр конфигурации.
func GetInstance() *PaymentConfig {
	once.Do(func() {
		instance = &PaymentConfig{
			provider:   "SBP",
			apiURL:     "https://pay.example.local",
			timeoutSec: 30,
		}
		fmt.Println("Инициализация PaymentConfig выполнена один раз")
	})
	return instance
}

func (p *PaymentConfig) SetProvider(provider string) {
	p.provider = provider
}

func (p *PaymentConfig) GetProvider() string {
	return p.provider
}

func (p *PaymentConfig) GetAPIURL() string {
	return p.apiURL
}

// OrderProcessor использует глобальную конфигурацию платёжного модуля.
type OrderProcessor struct {
	config *PaymentConfig
}

func NewOrderProcessor() *OrderProcessor {
	return &OrderProcessor{config: GetInstance()}
}

func (o *OrderProcessor) PrintConfig() {
	fmt.Printf("Текущий провайдер: %s, API: %s\n", o.config.GetProvider(), o.config.GetAPIURL())
}

func main() {
	first := NewOrderProcessor()
	first.PrintConfig()

	second := NewOrderProcessor()
	second.config.SetProvider("CARD")
	second.PrintConfig()

	fmt.Printf("Адрес конфигурации first:  %p\n", first.config)
	fmt.Printf("Адрес конфигурации second: %p\n", second.config)
}

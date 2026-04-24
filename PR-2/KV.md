## 9. Контрольные вопросы для защиты работы

### 1. На какие три группы делятся шаблоны GoF?

Шаблоны GoF делятся на три группы:

* **Порождающие (Creational)** — управляют созданием объектов
  Примеры:

  * Factory Method
  * Abstract Factory
  * Singleton

* **Структурные (Structural)** — описывают композицию классов и объектов
  Примеры:

  * Adapter
  * Decorator
  * Facade

* **Поведенческие (Behavioral)** — отвечают за взаимодействие объектов
  Примеры:

  * Strategy
  * Observer
  * State

---

### 2. Что такое паттерн Strategy? Чем он отличается от State? Как реализовать Strategy в Go?

**Strategy** — это паттерн, позволяющий выбирать алгоритм во время выполнения.

**Отличие от State:**

* Strategy — алгоритм выбирается извне
* State — поведение меняется в зависимости от внутреннего состояния объекта

**Реализация в Go через интерфейсы:**

```go
type Strategy interface {
    Execute(a, b int) int
}

type Add struct{}
func (Add) Execute(a, b int) int { return a + b }

type Multiply struct{}
func (Multiply) Execute(a, b int) int { return a * b }

type Context struct {
    strategy Strategy
}

func (c *Context) SetStrategy(s Strategy) {
    c.strategy = s
}

func (c *Context) Execute(a, b int) int {
    return c.strategy.Execute(a, b)
}
```

---

### 3. Роли Subject и Observer в Observer

* **Subject (наблюдаемый)** — хранит список наблюдателей и уведомляет их
* **Observer (наблюдатель)** — получает обновления

**Пример:**
Подписка на изменения данных (например, обновление UI при изменении модели)

**Реализация через каналы в Go:**

```go
type Subject struct {
    observers []chan string
}

func (s *Subject) Subscribe() chan string {
    ch := make(chan string)
    s.observers = append(s.observers, ch)
    return ch
}

func (s *Subject) Notify(msg string) {
    for _, ch := range s.observers {
        ch <- msg
    }
}
```

---

### 4. Отличие Factory Method от Abstract Factory

* **Factory Method** — создаёт один продукт
* **Abstract Factory** — создаёт семейство связанных объектов

**Влияние отсутствия наследования в Go:**

* используется **композиция и интерфейсы**, а не наследование
* фабрики реализуются через функции или интерфейсы

---

### 5. Принцип Open/Closed Principle

Объекты должны быть:

* **открыты для расширения**
* **закрыты для изменения**

**Паттерны:**

* Strategy
* Decorator
* Observer

**В Go:**
реализуется через интерфейсы:

```go
type Shape interface {
    Draw()
}
```

Добавление новой реализации не требует изменения существующего кода.

---

### 6. Композиция vs Embedding в Decorator

* **Композиция** — объект содержит другой объект
* **Embedding** — встраивание структуры (анонимное поле)

**Пример:**

```go
type Notifier interface {
    Send(msg string)
}

type Email struct{}
func (Email) Send(msg string) {}

type SMSDecorator struct {
    wrapped Notifier
}

func (d SMSDecorator) Send(msg string) {
    d.wrapped.Send(msg)
    // дополнительное поведение
}
```

---

### 7. Почему Singleton тяжело тестировать?

* глобальное состояние
* сложно подменить зависимости

**Потокобезопасный Singleton:**

```go
var instance *Singleton
var once sync.Once

func GetInstance() *Singleton {
    once.Do(func() {
        instance = &Singleton{}
    })
    return instance
}
```

**Альтернатива:**

* Dependency Injection

---

### 8. Facade vs Adapter

* **Facade** — упрощает интерфейс системы
* **Adapter** — преобразует интерфейс

**В Go:**

* Facade — функция/структура-обёртка
* Adapter — структура с нужным интерфейсом

---

### 9. Реализация выбранного паттерна

(пример: Strategy)

Использованы:

* интерфейсы — для абстракции
* композиция — для внедрения стратегии
* отсутствие наследования компенсируется интерфейсами

---

### 10. Какая группа паттернов используется чаще?

Чаще используются **поведенческие**, так как:

* Go ориентирован на простые взаимодействия
* активно используются интерфейсы и каналы

**Замены:**

* Observer → каналы
* Iterator → range
* Singleton → DI

---

### 11. Неявная реализация интерфейсов в Go

В Go:

* интерфейс реализуется автоматически (implicit)

В других языках:

* требуется явное указание (`implements`)

**Влияние:**

* меньше связности
* легче применять паттерны
* проще тестирование

---

### 12. Dependency Injection в Go

**Пример:**

```go
type DB interface {
    Save(data string)
}

type Service struct {
    db DB
}

func NewService(db DB) *Service {
    return &Service{db: db}
}
```

**Использование:**

```go
db := &MySQL{}
service := NewService(db)
```

**Преимущества:**

* тестируемость
* гибкость
* отсутствие глобального состояния

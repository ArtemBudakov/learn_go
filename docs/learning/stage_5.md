# info about errors

В Go "стандартная ошибка" — это реализация интерфейса `error`, встроенного в язык.

Любой тип, у которого есть метод `Error() string`, автоматически реализует интерфейс `error` 
и может использоваться как ошибка

## Стандартная ошибка
```go
import "errors"

func doSomething() error {
    return errors.New("something went wrong")
}
```
- errors.New(...) — это стандартный способ создать ошибку.
- Возвращаемое значение — объект типа error.


## Как обрабатывать стандартные ошибки
```go
err := doSomething()
if err != nil {
    fmt.Println("Error:", err)
}
```


## Полезные функции:
- проверка, является ли ошибка определённой.
```go
errors.Is(err, target)
```

- извлечение конкретного типа из error.
```go
errors.As(err, &target)
```

- оборачивание ошибки (%w).
```go
fmt.Errorf("something: %w", err)

```


# Почему recover работает только в defer?
Потому что `recover` только возвращает `panic`, если он уже был поднят. 
А `panic` инициирует «выход вверх по стеку» `stack unwinding`. 
Поэтому `recover()` нужно вызвать на пути этого выхода — то есть в `defer`.

- Если ты не вызываешь `recover()`, `panic` завершит выполнение программы.
- `recover()` не остановит `panic`, если он не находится внутри `defer`.

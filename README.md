# Задача – научиться обрабатывать ошибки в функциях

## Постановка задачи

Сегодня ты учишься базовой модели обработки ошибок в Go: функции возвращают **результат и ошибку** вторым значением. Твоя цель — научиться:

- писать функции с сигнатурами вида `func f(...) (T, error)`,
- вызывать такие функции и корректно проверять `err`,
- использовать короткую запись `if ... ; err != nil { ... }`.

### Материалы для самостоятельного изучения (прочитай вначале)

1. Return and handle an error. ([Go.dev][1])
2. Tour of Go — раздел **Errors**. ([Go.dev][2])
3. Go by Example — **Errors**. ([Go by Example][3])
4. Метанит (RU) — обзор обработки ошибок. ([Metanit][4])
5. Better Stack Community — краткий гайд с практическими советами. ([Better Stack][5])

## Шпаргалка

- Возврат без ошибки: `return value, nil`
- Возврат с ошибкой:

  ```go
  import "errors"
  return zeroValue, errors.New("описание проблемы")
  ```

- Базовая обработка:

  ```go
  v, err := f()
  if err != nil {
      // обработка/логирование
      return
  }
  // здесь можно использовать v
  ```

- Короткая запись в `if`:

  ```go
  if v, err := f(); err != nil {
      // обработка
  } else {
      // используем v
  }
  ```

## Подзадача 1 — Безопасное деление

Реализуй:

```go
// SafeDiv делит a на b и возвращает результат.
// Если b == 0 — верни ошибку.
func SafeDiv(a, b float64) (float64, error) {
    // твой код
}
```

В `main()` протестируй: `SafeDiv(10, 2)`, `SafeDiv(5, 0)`.
Используй **короткую запись**:

```go
if res, err := SafeDiv(10, 2); err != nil {
    fmt.Println("ошибка деления:", err)
} else {
    fmt.Println("результат:", res)
}
```

## Подзадача 2 — Первый элемент среза

Напиши:

```go
// FirstOrErr возвращает первый элемент среза строк.
// Если срез пустой — верни ошибку.
func FirstOrErr(items []string) (string, error) {
    // твой код
}
```

Проверь на `[]string{"Интукод", "Go"}` и на `[]string{}`.
Важно: не обращайся к `items[0]`, пока не проверишь `len(items)` — иначе будет паника.

## Подзадача 3 — Валидатор пользователя

Есть структура (можешь скопировать):

```go
type Student struct {
    Name string
    Age  int
}
```

Сделай функцию:

```go
// NewStudent создаёт студента. Если имя пустое или возраст < 0 — верни ошибку.
func NewStudent(name string, age int) (Student, error) {
    // твой код
}
```

В `main()` проверь кейсы: корректные данные, пустое имя, отрицательный возраст.
Потренируй оба варианта обработки: обычный `val, err := ...` и короткий `if ... ; err != nil`.

## Подзадача 4 — Ошибка из чужой функции

Напиши функцию-обёртку, которая внутри вызывает `strconv.ParseFloat`:

```go
// ParsePrice возвращает цену из строки в формате "123.45".
// Если формат неверный — верни ошибку из ParseFloat.
func ParsePrice(s string) (float64, error) {
    // твой код
}
```

## Подзадача 5

Ранее ты писал методы для разных структур в рамках задач

- https://github.com/intocode/go-methods-counter
- https://github.com/intocode/go-methods-todo-easy
- https://github.com/intocode/go-methods-phonebook-easy
- https://github.com/intocode/go-methods-creative

Сейчас тебе нужно вернуться к этим таскам и доработать методы так, чтобы они правильно возвращали ошибку.

Если функция не принимает входных параметров, то можно не писать проверку. Например функция `Increment()` для счетчика просто увеличивала значение счетчика на единицу и для неё сложно придумать сценарий, когда могла бы возникнуть ошибка.

Также не забудь затем обработать возврат ошибок в `main()`.

[1]: https://go.dev/doc/tutorial/handle-errors 'Return and handle an error'
[2]: https://go.dev/tour/methods/19 'Errors'
[3]: https://gobyexample.com/errors 'Go by Example: Errors'
[4]: https://metanit.com/go/tutorial/11.1.php 'Go | Изоморфная обработка ошибок'
[5]: https://betterstack.com/community/guides/scaling-go/golang-errors/ 'The Fundamentals of Error Handling in Go'

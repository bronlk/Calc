# Проект не доработан, просьба проверять завтра или после-завтра.

Скомпилируйте проект на linux/macOS - зайдите в orch и напишите go build, Linux/MacOs: ./orch
                                      зайдите в calc и напишите go build, Linux/MacOs: ./calс
                                      

Что-бы добавить пример нужно:
1) Зарегистрировать юзера ---- curl -X POST http://localhost:8080/register -H "Content-Type: application/json" -d '{"login": "te123st123user", "password": "testpassword"}' // логин должен быть уникальным
  
2) Войти в юзера ---- curl -X POST http://localhost:8080/login -H "Content-Type: application/json" -d '{"login": "te123st123user", "password": "testpassword"}' // без логина вы не сможете добавить пример

3) Добавить пример ---- curl -X POST http://localhost:8090/add_expression -d "1+3" //вы можете добавить любой пример,НО КАЛЬКУЛЯТОР УМЕЕТ РЕШАТЬ ТОЛЬКО ПРОСТЫЕ ПРИМЕРЫ т.е. примеры только из 2ух чисел

После этого в терминал будет выведен ответ.

Так же вы можете сделать следующие curl запросы:

1) curl -X POST http://localhost:8090/list_expressions //выводит все примеры





Ограничения:
Калькулятор не умеет вычислять сложные выражения больше - 2х чисел
![image](https://github.com/bronlk/Calc/assets/71665828/ef15bf19-41b3-4bf1-9b69-2fad82b66c8c)

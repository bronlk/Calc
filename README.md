# Осталось доделать только readme весь код рабочий ( сегодня все будет ) 
Весь калькулятор был написан и проверен на linux, у меня нету возможности проверить windows / macOs. Если у вас windows или macOs, и у вас не запускается проект то оцените его по скриншотам.

Скомпилируйте проект на linux - зайдите в orch и напишите go build, Linux: ./orch
                                зайдите в calc и напишите go build, Linux: ./calс
                                      

Что-бы добавить пример нужно:
1) Зарегистрировать юзера ---- curl -X POST http://localhost:8080/register -H "Content-Type: application/json" -d '{"login": "te123st123user", "password": "testpassword"}' // логин должен быть уникальным
  
2) Войти в юзера ---- curl -X POST http://localhost:8080/login -H "Content-Type: application/json" -d '{"login": "te123st123user", "password": "testpassword"}' // после логина в терминал напишется jwt token он нужен что бы добавить пример.

3) Добавить пример ----curl -X POST http://localhost:8080/save_expression -H "Content-Type: application/json" -d '{"expression": "2+4", "token": "сюда вы должны вставить ваш токен"}'
 //вы можете добавить любой пример,НО КАЛЬКУЛЯТОР УМЕЕТ РЕШАТЬ ТОЛЬКО ПРОСТЫЕ ПРИМЕРЫ т.е. примеры только из 2ух чисел

После этого в терминал будет выведен ответ.



Так же вы можете вывести все примеры:

1) curl -X POST http://localhost:8080/list_expressions //выводит все примеры





Ограничения:
Калькулятор не умеет вычислять сложные выражения больше - 2х чисел
![image](https://github.com/bronlk/Calc/assets/71665828/ef15bf19-41b3-4bf1-9b69-2fad82b66c8c)

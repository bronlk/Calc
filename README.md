Проект не доработан, просьба проверять завтра или после-завтра.

# sprint2final_proj

Скомпилируйте проект - зайдите в orch и напишите go build, Linux/MacOs: ./orch, Widnows: orch.exe
                       зайдите в calc и напишите go build, Linux/MacOs: ./calc calc1, Widnows: calc.exe calc1 
                       calc1 - идентификатор калькулатора
                       



Запустите bash файл написав -bash status или ./status.sh который будет показывать состояние агентов и статус вычисления выражения- bash status



Возможные функции:
curl -X POST http://localhost:8090/add_expression -d "1+3" //добавить пример

curl -X POST http://localhost:8090/list_agents //выводит все работующие агенты

curl -X POST http://localhost:8090/list_expressions //выводит все примеры

curl -X POST http://localhost:8080/register -H "Content-Type: application/json" -d '{"login": "te123st123user", "password": "testpassword"}'

curl -X POST http://localhost:8080/login -H "Content-Type: application/json" -d '{"login": "te123st123user", "password": "testpassword"}'

curl -X POST http://localhost:8080/logout -H "Content-Type: application/json" -d '{"token": "ваш jwt токен"}'


Ограничения:
Калькулятор не умеет вычислять сложные выражения больше - 2х чисел
![image](https://github.com/bronlk/Calc/assets/71665828/ef15bf19-41b3-4bf1-9b69-2fad82b66c8c)


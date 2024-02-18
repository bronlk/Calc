# sprint2final_proj

Скомпилируйте проект - зайдите в orch и напишите go build, Linux/MacOs: ./orch, Widnows: orch.exe
                       зайдите в calc и напишите go build, Linux/MacOs: ./calc calc1, Widnows: calc.exe calc1 
                       calc1 - идентификатор калькулатора
                       



Запустите bash файл ./status.sh который будет показывать состояние агентов и статус вычисления выражения- bash status



Возможные функции:
curl -X POST http://localhost:8090/add_expression -d "1+3" //добавить пример

curl -X POST http://localhost:8090/list_agents //выводит все работующие агенты

curl -X POST http://localhost:8090/list_expressions //выводит все примеры

Ограничения:
Калькулятор не умеет вычислять сложные выражения больше - 2х чисел

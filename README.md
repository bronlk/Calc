# sprint2final_proj

Скомпилируйте проект - зайдите в orch и напишите go build, Linux/MacOs: ./orch, Widnows: orch.exe
                       зайдите в calc и напишите go build, Linux/MacOs: ./calc, Widnows: calc.exe

Запустите bash файл status который будет показывать состояние агентов - bash status

 Попробуйте следующие функции


Возможные функции:
curl -X POST http://localhost:8090/add_expression -d "4+5" //добавить пример

curl http://localhost:8090/get_expression  //передать пример для калькулятора

curl http://localhost:8090/set_expression_result  //вывести ответ

curl http://localhost:8090/list_agents //выводит все работующие агенты

curl http://localhost:8090/list_expressions //выводит все примеры

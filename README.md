# sprint2final_proj

Скампилируйте проект - зайдите в orch и напишите go build, Linux/MacOs: ./orch, Widnows: orch.exe
                       зайдите в calc и напишите go build, Linux/MacOs: ./calc, Widnows: calc.exe

 Попробуйте следующие функции


Возможные функции:
curl http://localhost:8090/add_expression  //добавить пример

curl http://localhost:8090/get_expression  //передать пример для калькулятора

curl http://localhost:8090/set_expression_result  //вывести ответ

curl http://localhost:8090/list  //вывовит все примеры

curl http://localhost:8090/list_calc  //статусы калькуляторов

curl http://localhost:8090/history  //история примеров

curl http://localhost:8090/list_agents //выводит все работующие агенты

curl http://localhost:8090/list_expressions //выводит все примеры

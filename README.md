# sprint2final_proj


final spring2_proj 
1) Скачиваем проект и раз архивируем 
2) Заходим в orch и делаем go build и запускаем Linux/MacOs ./orch, Windows orch.exe
3) Заходим в calc и делаем go build и запускаем Linux/MacOs ./calc, Windows calc.exe



Возможные функции:
curl http://localhost:8090/add_expression  //добавить пример

curl http://localhost:8090/get_expression  //передать пример для калькулятора

curl http://localhost:8090/set_expression_result  //вывести ответ

curl http://localhost:8090/list  //вывовит все примеры

curl http://localhost:8090/list_calc  //статусы калькуляторов

curl http://localhost:8090/history  //история примеров

curl http://localhost:8090/list_agents //выводит все работующие агенты

curl http://localhost:8090/list_expressions //выводит все примеры

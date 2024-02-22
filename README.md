Мой дискорд - bron.lk пишите если будут проблемы

# sprint2final_proj

Скомпилируйте проект - зайдите в orch и напишите go build, Linux/MacOs: ./orch, Widnows: orch.exe
                       зайдите в calc и напишите go build, Linux/MacOs: ./calc calc1, Widnows: calc.exe calc1 
                       calc1 - идентификатор калькулатора
                       



Запустите bash файл написав -bash status или ./status.sh который будет показывать состояние агентов и статус вычисления выражения- bash status



Возможные функции:
curl -X POST http://localhost:8090/add_expression -d "1+3" //добавить пример

curl -X POST http://localhost:8090/list_agents //выводит все работующие агенты

curl -X POST http://localhost:8090/list_expressions //выводит все примеры

Ограничения:
Калькулятор не умеет вычислять сложные выражения больше - 2х чисел
![Снимок экрана от 2024-02-22-14-38-11](https://github.com/IvanK09/sprint2final_proj/assets/71665828/14531f97-a0ae-44ee-bfc0-c31553f92e69)

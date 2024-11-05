Здравствуйте, в этом репозитории можно посмотреть решение домашнего задания по конфигурационному управлению.

Задание №1
Разработать эмулятор для языка оболочки ОС. Необходимо сделать работу
эмулятора как можно более похожей на сеанс shell в UNIX-подобной ОС.
Эмулятор должен запускаться из реальной командной строки, а файл с
виртуальной файловой системой не нужно распаковывать у пользователя.
Эмулятор принимает образ виртуальной файловой системы в виде файла формата
tar. Эмулятор должен работать в режиме CLI.
Конфигурационный файл имеет формат csv и содержит:
• Имя пользователя для показа в приглашении к вводу.
• Имя компьютера для показа в приглашении к вводу.
• Путь к архиву виртуальной файловой системы.
• Путь к лог-файлу.
Лог-файл имеет формат csv и содержит все действия во время последнего
сеанса работы с эмулятором. Для каждого действия указаны дата и время. Для
каждого действия указан пользователь.
Необходимо поддержать в эмуляторе команды ls, cd и exit, а также
следующие команды:
1. chown.
2. find.
   

Все функции эмулятора должны быть покрыты тестами, а для каждой из
поддерживаемых команд необходимо написать 3 теста.
Задание №2

Разработать инструмент командной строки для визуализации графа
зависимостей, включая транзитивные зависимости. Сторонние средства для
получения зависимостей использовать нельзя.
Зависимости определяются по имени пакета ОС Alpine Linux (apk). Для
описания графа зависимостей используется представление Graphviz.
Визуализатор должен выводить результат в виде сообщения об успешном
выполнении и сохранять граф в файле формата png.
Конфигурационный файл имеет формат ini и содержит:
• Путь к программе для визуализации графов.
• Имя анализируемого пакета.
• Путь к файлу с изображением графа зависимостей.
• Максимальная глубина анализа зависимостей.
• URL-адрес репозитория.
Все функции визуализатора зависимостей должны быть покрыты тестами.


Задание №3
Разработать инструмент командной строки для учебного конфигурационного
языка, синтаксис которого приведен далее. Этот инструмент преобразует текст из
входного формата в выходной. Синтаксические ошибки выявляются с выдачей
сообщений.

Входной текст на учебном конфигурационном языке принимается из
файла, путь к которому задан ключом командной строки. Выходной текст на
языке toml попадает в файл, путь к которому задан ключом командной строки.

Однострочные комментарии:
// Это однострочный комментарий

Многострочные комментарии:

|#
Это многострочный
комментарий
#|

Словари:

[
 имя => значение,
 имя => значение,
 имя => значение,
 ...
]

Имена:

[_A-Z][_a-zA-Z0-9]*

Значения:

• Числа.
• Строки.
• Словари.

Строки:

@"Это строка"

Объявление константы на этапе трансляции:

var имя := значение

Вычисление константы на этапе трансляции:

$(имя)

Результатом вычисления константного выражения является значение.
Все конструкции учебного конфигурационного языка (с учетом их
возможной вложенности) должны быть покрыты тестами. Необходимо показать 3
примера описания конфигураций из разных предметных областей.

Задание №4

Разработать ассемблер и интерпретатор для учебной виртуальной машины
(УВМ). Система команд УВМ представлена далее.
Для ассемблера необходимо разработать читаемое представление команд
УВМ. Ассемблер принимает на вход файл с текстом исходной программы, путь к
которой задается из командной строки. Результатом работы ассемблера является
бинарный файл в виде последовательности байт, путь к которому задается из
командной строки. Дополнительный ключ командной строки задает путь к файлулогу, в котором хранятся ассемблированные инструкции в духе списков
“ключ=значение”, как в приведенных далее тестах.

Интерпретатор принимает на вход бинарный файл, выполняет команды УВМ
и сохраняет в файле-результате значения из диапазона памяти УВМ. Диапазон
также указывается из командной строки.

Форматом для файла-лога и файла-результата является json.
Необходимо реализовать приведенные тесты для всех команд, а также
написать и отладить тестовую программу.

![image](https://github.com/user-attachments/assets/d324adfb-82e2-409e-a4ab-99c733879c11)


Размер команды: 5 байт. Операнд: поле B. Результат: регистр по адресу,
которым является поле C.

Тест (A=4, B=368, C=63):

0x84, 0x0B, 0x00, 0xC0, 0x0F


![image](https://github.com/user-attachments/assets/6c839bae-ae91-493f-8d90-113ffc0ccc2d)

Размер команды:
5 байт. Операнд: значение в памяти по адресу, которым
является регистр по адресу, которым является поле C. 
Результат:
регистр по адресу, которым является поле B.

Тест (A=3, B=18, C=10):

0x93, 0x14, 0x00, 0x00, 0x00

Запись значения в память

![image](https://github.com/user-attachments/assets/4fc61401-5cb9-4502-8055-5af5b90ae088)


Размер команды: 5 байт. Операнд: регистр по адресу, которым является поле
C. Результат: значение в памяти по адресу, которым является поле B.

Тест (A=5, B=235, C=14):

0x5D, 0x07, 0x70, 0x00, 0x00

![image](https://github.com/user-attachments/assets/0dff7b6f-4a77-4c8c-88b4-3a9b81795bb6)


Размер команды: 5 байт. Первый операнд: регистр по адресу, которым
является поле B. Второй операнд: значение в памяти по адресу, которым является
поле C. Результат: регистр по адресу, которым является поле B.

Тест (A=1, B=35, C=497):

0x19, 0xE3, 0x03, 0x00, 0x00

Тестовая программа
Выполнить поэлементно операцию "<=" над двумя векторами длины 5.
Результат записать в первый вектор.
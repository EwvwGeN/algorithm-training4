|C. Быстрый алгоритм Дейкстры||
|-|-|
Ограничение времени|	5 секунд
Ограничение памяти|	512Mb
Ввод|	стандартный ввод или input.txt
Вывод|	стандартный вывод или output.txt
<hr>

Вам дано описание дорожной сети страны. Ваша задача – найти длину кратчайшего пути между городами $А$ и $B$.

## Формат ввода
Сеть дорог задана во входном файле следующим образом: первая строка содержит числа $N и K$ $(1 ≤ N ≤ 100000, 0 ≤ K ≤ 300000)$, где $K$ – количество дорог. Каждая из следующих K строк содержит описание дороги с двусторонним движением – три целых числа $a_i, b_i$ и $l_i (1 ≤ a_i, b_i ≤ N, 1 ≤ l_i ≤ 10^6)$. Это означает, что имеется дорога длины $li$, которая ведет из города $ai$ в город $bi$. В последней строке находятся два числа $А$ и $В$ – номера городов, между которыми надо посчитать кратчайшее расстояние $(1 ≤ A, B ≤ N)$

## Формат вывода
Выведите одно число – расстояние между нужными городами. Если по дорогам от города А до города В доехать невозможно, выведите $–1$.

|Пример||
|-|-|
Ввод|	Вывод
6 4|115
1 2 7|
2 4 8|
4 5 1|
4 3 100|
3 1|


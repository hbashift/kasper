def count_unique_states(a, b, c, d):
    # Начальное состояние
    initial_state = (a, b, c, d)

    # Множество для хранения уникальных состояний
    states = set()
    states.add(initial_state)

    # Очередь для обработки состояний
    queue = [initial_state]

    # Обход в ширину для генерации всех состояний
    while queue:
        current_state = queue.pop(0)
        a, b, c, d = current_state

        # Все возможные переходы
        new_states = []
        if a >= 2:
            new_states.append((a - 2, b, c + 1, d + 1))
        if b >= 2:
            new_states.append((a, b - 2, c + 1, d + 1))
        if c >= 2:
            new_states.append((a + 1, b + 1, c - 2, d))
        if d >= 2:
            new_states.append((a + 1, b + 1, c, d - 2))

        # Добавление новых состояний
        for new_state in new_states:
            if new_state not in states:
                states.add(new_state)
                queue.append(new_state)

    # Возвращаем количество уникальных состояний
    return len(states)

# Начальное состояние
a, b, c, d = 2, 0, 2, 3

# Получаем количество уникальных состояний
unique_states_count = count_unique_states(a, b, c, d)
print(unique_states_count)

import math

def torus_volume(R, r):
    """
    Вычисляет объём тора.

    :param R: Большой радиус (расстояние от центра тора до центра трубки)
    :param r: Малый радиус (радиус трубки)
    :return: Объём тора
    """
    volume = 2 * math.pi**2 * R * r**2
    return volume

# Пример использования
R = 9  # Большой радиус
r = 3  # Малый радиус

volume = torus_volume(R, r)
print(f"Объём тора с большим радиусом {R} и малым радиусом {r}: {volume:.2f}")



if __name__ == "__main__":
    print((torus_volume(9, 3) / torus_volume(7.5, 1.5))/math.pi)
    print(1.4/math.pi)
    # 1.5 и 7.5
 # 1598.88
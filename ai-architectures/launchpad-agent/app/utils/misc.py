def float_clamp(value: float | str, min_value: float = 0.0, max_value: float = 1.0, default_value: float = 0.0) -> float:

    try:
        value = float(value)
    except Exception:
        value = default_value

    return max(min(value, max_value), min_value)

def dsu(n: int, relations: list[tuple[int, int]]) -> list[int]:
    parent = list(range(n))
    rank = [0] * n

    def find(x: int) -> int:
        if parent[x] != x:
            parent[x] = find(parent[x])
        return parent[x]

    def unite(x: int, y: int) -> None:
        x_root = find(x)
        y_root = find(y)

        if x_root == y_root:
            return

        if rank[x] < rank[y]:
            parent[x_root] = y_root
        
        elif rank[x] > rank[y]:
            parent[y_root] = x_root
        
        else:
            parent[y_root] = x_root
            rank[x_root] += 1

    for x, y in relations:
        unite(x, y)

    return parent
    
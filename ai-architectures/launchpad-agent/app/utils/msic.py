def float_clamp(value: float | str, min_value: float = 0.0, max_value: float = 1.0, default_value: float = 0.0) -> float:

    try:
        value = float(value)
    except Exception:
        value = default_value

    return max(min(value, max_value), min_value)
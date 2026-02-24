REQUESTS
start_at / end_at semantics:

- NULL / NULL   -> non-time-based request
- Both set      -> reserved / absent during period
- Start only    -> starts at a point, open-ended
- end only      -> deadline-based request
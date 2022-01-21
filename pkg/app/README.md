The domain level implementation goes here.
For easier business and application logic separation these two are separated into `api` and `core`.
The `api` implements the application logic that should implement interfaces declared in `ports/left.go`.
The `core` should focus on domain level details.

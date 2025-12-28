# Signup

```json
curl -X POST http://localhost:8080/api/signup \
  -H "Content-Type: application/json" \
  -d '{
    "username": "newuser",
    "email": "newuser@example.com",
    "password": "securepassword123"
  }'
  ```

  # Login

  ```json
  curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "newuser@example.com",
    "password": "securepassword123"
  }'
  ```

  # Create Movie

  ```json
  curl -X POST 'http://localhost:8080/api/admin/movies' \
-H 'Content-Type: application/json' \
-H 'Authorization: YOUR_JWT_TOKEN_HERE' \
-d '{
  "title": "Inception",
  "director": "Christopher Nolan",
  "releaseDate": "2010-07-16",
  "duration": 148,
  "description": "A thief who enters the dreams of others to steal secrets from their subconscious.",
  "genre": "Sci-Fi",
  "posterURL": "https://example.com/inception-poster.jpg"
}'
```

# Reservation(Admin)

```json
curl -X GET 'http://localhost:8080/api/admin/reservations' \
  -H 'Authorization: YOUR_JWT_TOKEN_HERE'
```

# Reservation(User)

```json
curl -X GET 'http://localhost:8080/api/user/reservations' \
-H 'Authorization: YOUR_JWT_TOKEN_HERE'
```

# Add Showtime

```json
curl -X POST 'http://localhost:8080/api/admin/showtimes' \
-H 'Content-Type: application/json' \
-H 'Authorization: YOUR_JWT_TOKEN_HERE' \
-d '{
  "movieId": 1,
  "startTime": "2024-03-15T19:30:00Z",
  "endTime": "2024-03-15T21:30:00Z",
  "availableSeats": 100,
  "price": 12.50
}'
```

# Add Reservation(User)

```json
curl -X POST 'http://localhost:8080/api/reservations' \
-H 'Content-Type: application/json' \
-H 'Authorization: Bearer YOUR_JWT_TOKEN_HERE' \
-d '{
  "showtimeId": "d48c8908-7935-4d12-8736-a75410a50d40",
  "seatNumbers": "A1"
}'
```

# Cancel Reservation(User)

```json
curl -X DELETE \
  'http://localhost:8080/api/reservations/ID' \
  -H 'Authorization: YOUR_JWT_TOKEN_HERE'
  ```


# Get Seats Available


```json
curl -X GET \
  'http://localhost:8080/api/showtimes/SHOWTIME_ID' \
  -H 'Authorization: YOUR_JWT_TOKEN_HERE'
  ```



# User Show times(user)

```json
curl -X GET 'http://localhost:8080/api/movies/MOVIE_ID/showtimes' \
-H 'Authorization: YOUR_JWT_TOKEN_HERE'
```

# Get Reservations(Admin)

```json
curl -X GET 'http://localhost:8080/api/admin/reservations' \
-H 'Authorization: YOUR_JWT_TOKEN_HERE'
```

# Promote User to Admin

```json
curl -X POST 'http://localhost:8080/api/admin/users/USER_ID/promote' \
  -H 'Authorization: YOUR_JWT_TOKEN_HERE' \
  -H 'Content-Type: application/json'
```
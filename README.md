# SaffronStays API

The SaffronStays API provides room occupancy and rate information for Airbnb rooms over specified timeframes.

## API Endpoints

### Get Room Information

- **Endpoint**: `GET /api/room/{room_id}`
- **Description**: Retrieves the average occupancy percentage over the last 5 months and the average, highest, and lowest night rates for the next 30 days for a specified room.
- **Path Parameters**:
  - `room_id` (int): The ID of the room for which to fetch the information.
- **Responses**:
  - **200 OK**: Returns a JSON object with the following fields:
    - `occupancy_percentage` (float64): The average occupancy percentage for the room over the last 5 months.
    - `avg_rate` (float64): The average night rate for the room over the next 30 days.
    - `high_rate` (float64): The highest night rate for the room over the next 30 days.
    - `low_rate` (float64): The lowest night rate for the room over the next 30 days.
  - **404 Not Found**: If the room ID does not exist or no data is available.
  - **500 Internal Server Error**: For unexpected server errors.

## Technologies Used

The following technologies were used to create the service:

- `Supabase`: Used to host the service's database
- `Vercel`: Used to host the service

## How to Start the Service

To start the service, ensure the environment variables are set up correctly:

1. Clone the repository.
2. Copy the `.env.sample` file to `.env` and set up your `.env` file with the required environment variables, including your own database URL in the `DB_URL` variable.

  ```bash
    cp .env.sample .env
  ```

3. Run the following command:

  ```bash
  go run .
  ```

The service will be available at http://localhost:8000.

## Project File Tree
```bash
saffronstays-api
├── main.go
├── handlers
│   └── room_handler.go
├── models
│   ├── occupancy.go
│   └── rate.go
├── config
│   ├── pgxpool_config.go
│   └── db.go
├── .env.sample
├── go.mod
├── go.sum
├── README.md
└── .gitignore
```

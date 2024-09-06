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

## Environment Variables

The following environment variables are required to run the service:

- `DB_HOST`: Database host URL
- `DB_USER`: Database username
- `DB_PASSWORD`: Database password
- `DB_NAME`: Name of the database
- `DB_PORT`: Port number for the database (default: 5432)
- `PORT`: Port on which the API will run (default: 8000)

## How to Start the Service

To start the service, ensure the environment variables are set up correctly, and use Docker:

1. Clone the repository.
2. Set up your `.env` file with the required environment variables.
3. Run the following commands:

   ```bash
   docker-compose up --build
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
│   └── db.go
├── .env.sample
├── go.mod
├── go.sum
├── Dockerfile
├── README.md
├── .gitignore
└── docker-compose.yml
```

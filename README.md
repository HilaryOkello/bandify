# Groupie Tracker

## Overview

Groupie Tracker is a web application designed to help users track their favorite musical groups and their activities. The application interacts with a provided API that contains information about artists, concert locations, dates, and their relations.

## Features
- **User-Friendly Interface**: Intuitive design for easy navigation.
- **Data Visualization**: Display information using various formats like cards, and graphics.
- **Event Creation**: Trigger actions that communicate with the server for dynamic updates.
- **Error Handling**: Ensure the application handles errors gracefully and provides meaningful feedback to users.
- **Responsive Design**: The website is designed to be responsive and accessible on various devices.

## Technologies Used
- Backend: Go (Golang) for server-side logic and API interaction.
- Frontend: HTML, CSS for building a dynamic and responsive user interface.
- Unit Testing: Go's testing framework for ensuring code correctness.


## Project Structure
The project is organized into several packages:
1. `server`: Contains files that handle:
    -   Storing the data structures representing the API data such as artist, location, date, and relation.
    -   API requests and data fetching.
    -   Managing the HTTP routes and handlers for serving the web pages and processing requests.
2. `templates`: Contains the HTML templates for the homepage, artist information page and error page.
3. `static`: Contains CSS and other static assets for the website.

## Set up and Usage
Ensure you have Go downloaded and installed in your machine. To use the program:
1. Clone the repository:
```bash
   git clone https://github.com/Joan2509/groupie-tracker.git
```
2. Navigate to the project directory and run the application by respectively typing the following commands:
```bash
cd groupie-tracker
go run main.go
```
3. Open your browser and navigate to `http://localhost:3000/`

## Error Handling
- `400: Bad Request`: Returned when the server cannot process the request due to a client error say empty searches.
- `404: Not Found`: Returned when the server can not find the requested resource/page.
- `405: Method Allowed`: Returned when the HTTP method used is not supported for the specified resource.
- `500: Internal Server Error`: Generic error message returned when an issue is encountered on the server side. 

## Contributing
Contributions are welcome! Please open an issue or submit a pull request.

## License
This project is licensed under the MIT License. See the LICENSE file for details.

## Authors
[Hillary Okello](https://learn.zone01kisumu.ke/git/hilaokello/)

[Joan Wambugu](https://learn.zone01kisumu.ke/git/jwambugu/)
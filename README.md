# Groupie Tracker

Groupie Tracker is a web-based application that utilizes a given RESTful API to display information about artists, their concert locations, and dates in a user-friendly format. This project emphasizes client-server interactions, data manipulation, and event-driven programming.

## Project Overview

The provided API contains four main parts:

1. **Artists**: Information about bands and artists, including:
   - Name(s)
   - Image
   - Year of activity commencement
   - First album release date
   - Band members

2. **Locations**: Last and/or upcoming concert locations.

3. **Dates**: Last and/or upcoming concert dates.

4. **Relation**: Links between artists, locations, and dates.

Using this data, the application enables users to explore artist information through various data visualizations, such as cards, tables, lists, and graphics. The project also includes an interactive event feature that triggers a client-server interaction.

## Features

- **Artist Profiles**: Display detailed information about bands/artists, including their members, images, and activity timelines.
- **Concert Locations**: Visualize the locations of past and upcoming concerts.
- **Concert Dates**: Explore the schedules of concerts through intuitive layouts.
- **Interactive Events**: Trigger actions that involve client-server communication using the request-response model.

## Technologies

- **Backend**: Written entirely in Go.
- **Frontend**: HTML, CSS, and JavaScript (if applicable).
- **Data**: Managed using JSON provided by the API.

## Requirements

- **Server Stability**: The server and website must never crash.
- **Error Handling**: All pages should handle errors gracefully.
- **Code Quality**: Follow best practices in code structure and readability.
- **Unit Testing**: Include unit tests to verify functionality.

## Usage

### Running the Project

1. Clone the repository:
   ```bash
   git clone https://learn.zone01oujda.ma/git/aberhili/groupie-tracker.git
   cd groupie-tracker

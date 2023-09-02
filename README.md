# CID-10-API

## Overview

The CID-10-API is developed using the Go programming language and is intended for educational purposes. It provides an interface to add and search for CID codes (International Statistical Classification of Diseases and Related Health Problems). All interactions with the API are authenticated using JWT (JSON Web Tokens).

## Features

- User authentication via JWT
- Add CID codes to a MongoDB database
- Search for existing CID codes in the database

## Getting Started

### Prerequisites

Ensure you have Go installed and MongoDB set up.

### Installation & Configuration

1. Clone the repository:
git clone <https://github.com/rafajulio/cid-10-api.git>


2. Setup your `.env` file with the following parameters:
MONGOURI=<Your_MongoDB_Connection_URL>
JWTSECRET=<Your_JWT_Secret>

Replace `<Your_JWT_Secret>` with your JWT secret key and `<Your_MongoDB_Connection_URL>` with your MongoDB connection string.

3. In the `extra` folder, there is a provided `.json` file containing all the CID-10 codes sourced from the official [cms.gov](https://www.cms.gov/medicare/icd-10/2024-icd-10-cm) website. You can use this file to populate your database or for other related purposes.

### Running the Application

To run the API, simply run `go run main.go` on your terminal


## Contributing

Feel free to fork, raise issues, or submit Pull Requests. Feedback and contributions are always welcome.


## Acknowledgements

This project was developed for personal learning. Special thanks to the Go community and the resources available that made this project possible.




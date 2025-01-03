# De To' API

API for the De to' platform. Designed to facilitate interaction between businesses and users.

## Features

- **Business Registration**: Allows users to register their businesses and product catalogs.
    
- **Nearby Product Search**: Users can search for products based on their geographic location.
    
- **Business Rating**: Users can rate the businesses they visit.
    
- **Inventory Updates**: To maintain platform accuracy, business owners are required to update their catalogs periodically.
    
- **Reminder Notifications**: Business owners receive reminders to update their inventory.
    
- **Incentives and Penalties**: Businesses will receive incentives or penalties based on their behavior.
    
- (more features to be defined)
    

## Installation

To run this API locally, follow these steps:

### Prerequisites

- **Go** installed. If you don't have it, you can download and install it from here.
    

### Installation Steps

1. **Clone the repository**:

```bash
git clone https://github.com/Dario61K/De-To-API.git
cd de-to-api
```

2. **Install dependencies**: The API uses Go modules to manage dependencies. Run the following command to install them:

```bash
go mod tidy
```

3. **Set up environment variables in your system (currently not using .env file)**
    
    (Still defining all necessary environment variables)
    
4. **Run the API**: Start the server by running the following command:

```bash
go run cmd/main.go
```

- **Test the API**: Once the server is running, the API will be available at `http://localhost:8080` by default, unless another port is specified in the environment variables.# De-To-API

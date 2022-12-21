# Tech Test

Tech test related to CRUD user 

## Project Structure

- **Delivery**: A presenter layer to decide how the data will be presented. In this project I use REST API
- **Usecase**: Where business logic happens
- **Repository**: A non business logic layer to communicate with datastore like Database, Cache, and etc.
- **Model**: To define entities and its method that used in the project
- **Config**: Folder to store configuration
- **Utils**: An utility method that used in the project and not a part of business logic
- **Database**: Additional folder to store database migration and seed
# Bookstore API

This is a simple RESTful API for managing a bookstore, built using Go.

## Installation

1. Make sure you have Go installed. If not, you can download and install it from [the official Go website](https://golang.org/dl/).
2. Clone this repository:

    ```
    https://github.com/MRA11145/booksManagement.git
    ```

3. Navigate to the project directory:

    ```
    cd booksManagement
    ```

4. Run the project:

    ```
    go run main.go
    ```

## Usage

Once the server is running, you can access the following APIs:

1. **Fetch all books**:  
    ```
    GET http://localhost:8000/api/books
    ```

2. **Fetch information about a single book by ID**:  
    ```
    GET http://localhost:8000/api/books/{id}
    ```

3. **Add a single book**:  
    ```
    POST http://localhost:8000/api/books
    ```

    Request Body:
    ```json
    {
        "ID": "5",
        "Isbn": "5660",
        "Title": "New Book",
        "Author": {
            "FirstName": "Jane",
            "LastName": "Doe"
        }
    }
    ```

4. **Delete a book by ID**:  
    ```
    DELETE http://localhost:8000/api/books/{id}
    ```

5. **Update the information of a book by ID**:  
    ```
    PUT http://localhost:8000/api/books/{id}
    ```

    Request Body:
    ```json
    {
        "Isbn": "5660",
        "Title": "Updated Book",
        "Author": {
            "FirstName": "John",
            "LastName": "Doe"
        }
    }
    ```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

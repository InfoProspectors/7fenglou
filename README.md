# 7fenglou 

This project is designed to extract data from the 7fenglou API and save it to an Excel file.

## Description

The `7fenglou` project fetches data from the 7fenglou API and saves it to an Excel file. It retrieves information such as thread IDs, post IDs, user IDs, and category IDs from the API response.

The project uses the Go programming language and relies on the following external libraries:

- `github.com/xuri/excelize/v2` for creating and manipulating Excel files.

## Getting Started

### Prerequisites

- Go 1.16 or higher

### Installation

1. Clone the repository:

   ```shell
   git clone https://github.com/InfoProspectors/7fenglou.git
   ```
2. Navigate to the project directory:

   ```shell
   cd 7fenglou
   ```
   
3. Install dependencies:

   ```shell
    go mod download
   ```

## Usage

1. Open the main.go file in a text editor.

2. Locate the BaseURL constant and update it with the base URL of the 7fenglou API.

3. Run the project:

    ```shell
    go run main.go
    ```

4. The data will be fetched from the API, saved to a JSON file (data.json), and then extracted and saved to an Excel file (7fenglou.xlsx).

## Configuration

The following parameters can be modified in the `main.go` file:

- `BaseURL`: The base URL of the 7fenglou API.

- `PageLimit`: The maximum number of data entries to retrieve per page.

## Contributing

Contributions are welcome! If you find any issues or want to enhance the project, please feel free to submit a pull request.

## License

This project is licensed under the  Apache License [License](./LICENSE).


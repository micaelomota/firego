# FireGo

FireGo is a command-line interface (CLI) tool designed to simplify the management of Firebase Firestore data using Go. This project was created to address the challenges and monotony of modifying data on the Firestore UI. The primary objective of FireGo is to provide a comprehensive set of tools that enable users to effortlessly manage their Firestore database directly from the terminal.

## Features

- Streamlined data manipulation: FireGo offers a range of commands and functionalities that make it easier to add, update, delete, and query data in Firestore.
- Interactive interface: The CLI provides an interactive interface that allows users to navigate through their Firestore collections and documents effortlessly.
- Batch operations: FireGo supports batch operations, enabling users to perform multiple data modifications in a single command.
- Customizable queries: Users can leverage FireGo's query capabilities to retrieve specific data from their Firestore database based on various criteria.
- Seamless integration: FireGo seamlessly integrates with existing Go projects, allowing developers to incorporate Firestore management into their workflows effortlessly.

## Installation

To install FireGo, follow these steps:

1. Ensure you have Go installed on your machine.
2. Open your terminal and run the following command:

```shell
go get github.com/micaelomota/firego
```

3. Once the installation is complete, you can start using FireGo by running the `firego` command in your terminal.

## Usage

To use FireGo, execute the following command:

```shell
firego [command]
```

Replace `[command]` with one of the available commands provided by FireGo. For example, to add data to your Firestore database, you can use the `add` command:

```shell
firego add --collection users --data '{"name": "John Doe", "age": 30}'
```

For a complete list of available commands and their usage, refer to the FireGo documentation.

## Contributing

We welcome contributions from the community to enhance FireGo's functionality and usability. If you would like to contribute, please follow our [contribution guidelines](https://github.com/micaelomota/firego/blob/main/CONTRIBUTING.md).

## License

FireGo is licensed under the [MIT License](https://github.com/micaelomota/firego/blob/main/LICENSE).

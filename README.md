# RabbitMQ Broadcast and Data Integrity

This repository contains examples and utilities for implementing RabbitMQ broadcast and testing data integrity over RabbitMQ protocol.

## Overview

RabbitMQ is a lightweight messaging protocol widely used in IoT (Internet of Things) and other applications requiring efficient message delivery. This repository provides examples and utilities to demonstrate RabbitMQ broadcast functionality and perform data integrity testing.

## Features

- **RabbitMQ Broadcast**: Examples showcasing how to broadcast messages to multiple subscribers using RabbitMQ.
  
- **Data Integrity Testing**: Utilities and scripts to test data integrity over RabbitMQ, ensuring messages are delivered intact and in the correct order.

## Setup

1. **Clone the Repository**: Clone this repository to your local machine using Git:

    ```bash
    git clone https://github.com/hebobibun/rabbitmq-integrity.git
    ```

## Testing Data Integrity

1. **Run Publisher**: Start the publisher service to publish messages to Beanstalkd.

    ```bash
    cd pub
    go run main.go
    ```

2. **Run Subscriber 1**: Start the subscriber service to receive and process messages from Beanstalkd.

    ```bash
    cd sub1
    go run main.go
    ```

3. **Verify Integrity**: Monitor the subscriber output to ensure that messages are received intact and in the correct order.

## Contributing

Contributions are welcome! If you find bugs, have ideas for improvements, or wish to add new features, feel free to open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).

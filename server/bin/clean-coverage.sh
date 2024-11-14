#!/bin/bash

# Define the list of packages to remove from the coverage report
# WARNING: this removes all sub directories from the coverage report  as well!
packages_to_remove=(
    # CLI doesn't require unit testing
    "github.com/capsa-gg/capsa/server/internal/cmd"

    # Database is code generated and doesn't need testing
    "github.com/capsa-gg/capsa/server/internal/data/database"

    # Some data/infrastructure are just abstraction layers for existing packages, they don't need unit tests
    "github.com/capsa-gg/capsa/server/internal/data/emails"
    "github.com/capsa-gg/capsa/server/internal/data/logchunks"
    "github.com/capsa-gg/capsa/server/internal/infrastructure/blobstorage"
    "github.com/capsa-gg/capsa/server/internal/infrastructure/logger"
    "github.com/capsa-gg/capsa/server/internal/infrastructure/mailer"
    "github.com/capsa-gg/capsa/server/internal/infrastructure/migrator"

    # Interactor doesn't need unit tests, it's just a data model with some tiny helpers
    "github.com/capsa-gg/capsa/server/internal/interactor"

    # HTTP logic does not require unit testing
    "github.com/capsa-gg/capsa/server/internal/server/handlers"
    "github.com/capsa-gg/capsa/server/internal/server/middleware"

    # Swagger is auto generated and doesn't need unit tests
    "github.com/capsa-gg/capsa/server/swagger"
)

# Function to check if a line should be removed based on the packages to remove
should_remove() {
    for pkg in "${packages_to_remove[@]}"; do
        if [[ $1 == $pkg* ]]; then
            return 0
        fi
    done
    return 1
}

# Check if an input file was provided as an argument
if [ $# -eq 0 ]; then
    echo "Error: No input file provided."
    echo "Usage: $0 <input_file> (often coverage_raw.txt)"
    echo "This will output the filtered coverage report to coverage.txt"
    exit 1
fi

# Read the coverage report file line by line
while IFS= read -r line; do
    if ! should_remove "$line"; then
        echo "$line"
    fi

# Read file and output
done < "$1" > coverage.txt

echo "Filtered coverage report saved to coverage.txt"

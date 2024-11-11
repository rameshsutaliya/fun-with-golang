#!/bin/bash

# Set the minimum coverage threshold (90% in this case)
MIN_COVERAGE=80

# Define the test coverage output file
COVERAGE_FILE="coverage.out"

# Run Go tests with coverage profile
echo "Running Go tests with coverage..."
go test -coverprofile=$COVERAGE_FILE ./ch1/... ./neetcode/...

# Check if tests failed
if [ $? -ne 0 ]; then
  echo "Tests failed! Exiting..."
  exit 1
fi

# Calculate the coverage percentage from the coverage report
COVERAGE=$(go tool cover -func=$COVERAGE_FILE | grep total: | awk '{print $3}' | sed 's/%//')

# Display the coverage
echo "Total test coverage: $COVERAGE%"

# Check if coverage meets the minimum threshold
if [ $(echo "$COVERAGE < $MIN_COVERAGE" | bc) -eq 1 ]; then
  echo "ERROR: Code coverage is below $MIN_COVERAGE%. Current coverage: $COVERAGE%"
  exit 1
else
  echo "Code coverage is sufficient. Current coverage: $COVERAGE%"
fi

# Optional: Output a detailed test coverage report in HTML format for review
echo "Generating detailed test coverage report..."
go tool cover -html=$COVERAGE_FILE -o coverage.html

# Provide the location of the coverage report
echo "Detailed test coverage report is saved as 'coverage.html'."

exit 0

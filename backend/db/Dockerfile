# Use the official Alpine Linux as the base image
FROM alpine:latest

# Set environment variables for PostgreSQL
ENV POSTGRES_USER youruser
ENV POSTGRES_PASSWORD yourpass
ENV POSTGRES_DB yourdb

# Install PostgreSQL and required libraries
RUN apk add --no-cache postgresql postgresql-client

# Initialize a PostgreSQL database cluster
RUN initdb -D /var/lib/postgresql/data

# Start the PostgreSQL server
CMD ["postgres", "-D", "/var/lib/postgresql/data"]

# Expose the default PostgreSQL port
EXPOSE 5432 

-- Database initialization script for Book Store Application

-- Use the bookstore database
USE bookstore;

-- Create books table
CREATE TABLE IF NOT EXISTS books (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    year INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_title (title),
    INDEX idx_author (author),
    INDEX idx_year (year)
);


-- Insert some sample data
INSERT INTO books (title, author, year) VALUES
    ('The Go Programming Language', 'Alan A. A. Donovan and Brian W. Kernighan', 2015),
    ('Clean Code: A Handbook of Agile Software Craftsmanship', 'Robert C. Martin', 2008),
    ('Designing Data-Intensive Applications', 'Martin Kleppmann', 2017),
    ('Clean Architecture', 'Robert C. Martin', 2017),
    ('Kubernetes Up & Running', 'Kelsey Hightower, Brendan Burns, Joe Beda', 2022),
    ('Docker Deep Dive', 'Nigel Poulton', 2020);

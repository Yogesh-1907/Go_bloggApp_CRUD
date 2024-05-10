DROP SCHEMA IF EXISTS blogging_db CASCADE;

CREATE SCHEMA blogging_db;

SET schema 'blogging_db';

CREATE TABLE posts (
  id serial PRIMARY KEY,
  title varchar(255) NOT NULL DEFAULT NULL,
  content text NOT NULL DEFAULT NULL,
  author varchar(255) NOT NULL DEFAULT NULL
);

INSERT INTO posts (title, content, author) VALUES
('How to Create a Blog Application with Golang and GORM', 'In this post, we will show you how to create a blog application with Golang and GORM. We will start by creating a new Golang project and then we will install the GORM package. We will then create a database and a table for our blog posts. Finally, we will create a service layer and a repository layer for our blog application.', 'Emily Garcia'),
('Using GORM to Persist Data to a Database', 'In this post, we will show you how to use GORM to persist data to a database. We will start by creating a new GORM model and then we will save the model to the database. We will then fetch the model from the database and print it to the console.', 'Emily Garcia'),
('Paginate Results in GORM', 'In this post, we will show you how to use GORM to paginate results. We will start by creating a new GORM query and then we will paginate the results of the query. We will then print the paginated results to the console.', 'Emily Garcia'),
('Creating Relationships between Models in GROM', 'In this post, we will show you how to use GORM to create relationships between models. We will start by creating two GORM models and then we will create a relationship between the two models. We will then save the models to the database and fetch the models from the database.', 'Michael Davis'),
('Querying Data with Conditions', 'In this post, we will show you how to use GORM to query data with conditions. We will start by creating a new GORM query and then we will add conditions to the query. We will then execute the query and print the results to the console.', 'Susan Carter'),
('Updating Data using GORM', 'In this post, we will show you how to use GORM to update data. We will start by creating a new GORM model and then we will update the models data. We will then save the model to the database.', 'Peter Jones'),
('Deleting Data using GORM', 'In this post, we will show you how to use GORM to delete data. We will start by creating a new GORM model and then we will delete the model from the database.', 'Peter Jones'),
('How to Use GORM to Create a Custom Database Migration', 'In this post, we will show you how to use GORM to create a custom database migration. We will start by creating a new GORM migration file and then we will add code to the migration file. We will then run the migration file to create the database schema.', 'Peter Jones'),
('How to Use GORM to Seed Data into a Database', 'In this post, we will show you how to use GORM to seed data into a database. We will start by creating a new GORM seed file and then we will add data to the seed file. We will then run the seed file to seed the database with data.', 'Jane Doe'),
('How to Optimize Database Performance', 'In this post, we will show you how to use GORM to optimize database performance. We will start by identifying the bottlenecks in our application and then we will implement optimizations to improve performance.', 'John Smith');

SELECT * FROM posts;
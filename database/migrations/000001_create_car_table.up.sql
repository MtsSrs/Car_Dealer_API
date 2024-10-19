CREATE TABLE IF NOT EXISTS cars(
   car_id serial PRIMARY KEY,
   car_manufacturer VARCHAR (255) NOT NULL,
   car_name VARCHAR (255) NOT NULL,
   car_model VARCHAR (300) NOT NULL,
   car_year_model VARCHAR(5) NOT NULL,
   car_price FLOAT NOT NULL
);




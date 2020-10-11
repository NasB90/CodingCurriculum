USE fernweh;

CREATE TABLE tours (
    id INT PRIMARY KEY NOT NULL,
    category VARCHAR(255) NOT NULL,
    product_name VARCHAR(255) NOT NULL,
    details VARCHAR(255) NOT NULL,
    duration VARCHAR(255) NOT NULL, 
    cost FLOAT(7,2) NOT NULL,
    img VARCHAR(255) Not NULL
);

INSERT INTO tours (
    id,category,product_name,details,duration,cost,img
) 

VALUES
(1, 'tour', 'Desert trip', 'description','8', 20, 'images/jeep-girl.jpg'),
(2, 'tour', 'Day Trip to Oman', 'description','12', 100, 'images/oman.jpg'),
(3, 'tour', 'Sand Dune Bashing...','description', '8', 20, 'images/bbq.jpg'),
(4, 'tour', 'Dubai Eye', 'description','8', 20, 'images/dubai-eye.jpg'),
(5, 'tour', 'Sightseeing', 'description','8', 20, 'images/burj-khaifa.jpeg');
-- (21, 'transfer', 'AUH - Abu Dhabi', 'description','20 min', 40, 'images/jeep-girl.jpg');
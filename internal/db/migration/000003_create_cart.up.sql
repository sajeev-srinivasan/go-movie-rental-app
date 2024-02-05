create table cart (
    id integer primary key ,
    userId integer,
    noOfItems integer
);


create table cartItem (
    id integer PRIMARY KEY ,
    cartId integer references cart (id),
    movieId varchar(50) references movies (id),
    movieQuantity integer
)
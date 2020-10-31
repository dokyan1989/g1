create table `products` (
    `id` bigint not null auto_increment,
    `name` varchar(255) not null,
    `price` decimal(20,4),
    primary key (id)
);

insert into products (`name`, `price`)
values ('macbook', 100000);

insert into products (`name`, `price`)
values ('iphone 12', 50000);

insert into products (`name`, `price`)
values ('ipad pro', 70000);
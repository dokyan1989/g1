create table `employees` (
  `emp_no` bigint not null auto_increment,
  `birth_date` date not null,
  `first_name` varchar(14) not null,
  `last_name` varchar(16) not null,
  `gender` enum('M', 'F') not null,
  `hire_date` date not null,
  primary key (emp_no)
);

create table `departments` (
  `dept_no` char(4) not null,
  `dept_name` varchar(40) not null,
  primary key (`dept_no`),
  unique key (`dept_name`)
);

create table `dept_emp` (
  `emp_no` bigint not null,
  `dept_no` char(4) not null,
  `from_date` date not null,
  `to_date` date not null,
  key (`emp_no`),
  key (`dept_no`),
  foreign key (`emp_no`) references `employees` (`emp_no`) on delete cascade,
  foreign key (`dept_no`) references `departments` (`dept_no`) on delete cascade,
  primary key (`emp_no`, `dept_no`)
);

create table `dept_manager` (
  `emp_no` bigint not null,
  `dept_no` char(4) not null,
  `from_date` date not null,
  `to_date` date not null,
  key (`emp_no`),
  key (`dept_no`),
  foreign key (`emp_no`) references `employees` (`emp_no`) on delete cascade,
  foreign key (`dept_no`) references `departments` (`dept_no`) on delete cascade,
  primary key (`emp_no`, `dept_no`)
);

create table `titles` (
  `emp_no` bigint not null,
  `title` varchar(50) not null,
  `from_date` date not null,
  `to_date` date,
  key (`emp_no`),
  foreign key (`emp_no`) references `employees` (`emp_no`) on delete cascade,
  primary key (`emp_no`, `title`, `from_date`)
);

create table `salaries` (
  `emp_no` bigint not null,
  `salary` decimal(20, 4) not null,
  `from_date` date not null,
  `to_date` date not null,
  key (`emp_no`),
  foreign key (`emp_no`) references `employees` (`emp_no`) on delete cascade,
  primary key (`emp_no`, `from_date`)
);
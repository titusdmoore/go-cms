CREATE TABLE Testing (
    ID int NOT NULL AUTO_INCREMENT,
    FirstName varchar(255),
    LastName varchar(255),
    Username varchar(255) NOT NULL,

    PRIMARY KEY (ID),
    UNIQUE (Username)
);

INSERT INTO Testing (FirstName, LastName, Username)
    VALUE ("Titus", "Moore", "tmoore");

INSERT INTO Testing (FirstName, LastName, Username)
    VALUE ("John", "Doe", "jdoe");

INSERT INTO Testing (FirstName, LastName, Username)
    VALUE ("Cara", "Doe", "cdoe");

INSERT INTO Testing (FirstName, LastName, Username)
    VALUE ("Frank", "Williams", "fwilliams");

INSERT INTO Testing (FirstName, LastName, Username)
    VALUE ("Lillith", "Saxon", "lsaxon");

INSERT INTO Testing (FirstName, LastName, Username)
    VALUE ("Billy", "Beckman", "bbeckman");

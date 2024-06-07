-- Configuration Tables
CREATE TABLE Options (
    ID int NOT NULL AUTO_INCREMENT,
    optionName varchar(255) NOT NULL,
    optionValue varchar(1024) NOT NULL,

    PRIMARY KEY (ID),
    UNIQUE (optionName)
);

-- Users and Auth Tables
CREATE TABLE Accounts (
    ID int NOT NULL AUTO_INCREMENT,
    firstName varchar(255) NOT NULL,
    lastName varchar(255) NOT NULL,
    username varchar(75) NOT NULL,
    email varchar(255) NOT NULL,


    PRIMARY KEY (ID),
    UNIQUE (username)
);

CREATE TABLE Account_Attributes (
    ID int NOT NULL AUTO_INCREMENT,
    accountId int NOT NULL,
    attributeName varchar(144),
    attributeValue varchar(1024),

    PRIMARY KEY (ID),
    FOREIGN KEY (accountId) REFERENCES Accounts(ID)
);

CREATE TABLE Roles (
    ID int NOT NULL AUTO_INCREMENT,
    name varchar(50) NOT NULL,
    permissions JSON,

    PRIMARY KEY (ID),
    UNIQUE (NAME)
);

CREATE TABLE Accounts_Roles (
    ID int NOT NULL AUTO_INCREMENT,
    accountId int NOT NULL,
    roleId int NOT NULL,

    PRIMARY KEY (ID),
    FOREIGN KEY (accountId) REFERENCES Accounts(ID),
    FOREIGN KEY (roleId) REFERENCES Roles(ID)
);

-- Generic Data Types
CREATE TABLE Pages (
    ID int NOT NULL AUTO_INCREMENT,
    title varchar(144) NOT NULL,
    slug varchar(144) NOT NULL,
    content longtext,
    status ENUM('Draft', 'Pending_Review', 'Scheduled', 'Published') NOT NULL DEFAULT 'Draft',
    publishedBy int NOT NULL,

    publishedDate datetime NOT NULL,
    lastUpdated datetime,

    PRIMARY KEY (ID),
    FOREIGN KEY (publishedBy) REFERENCES Accounts(ID),
    UNIQUE (slug)
);

CREATE TABLE Page_Attributes (
    ID int NOT NULL AUTO_INCREMENT,
    pageId int NOT NULL,
    attributeName varchar(144),
    attributeValue varchar(1024),

    PRIMARY KEY (ID),
    FOREIGN KEY (pageId) REFERENCES Pages(ID)
);

-- Tell the DB that we have completed the schema creation
INSERT INTO Options (optionName, optionValue) VALUES ('initialized', 'true');

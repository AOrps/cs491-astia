CREATE DATABASE Astia;

USE Astia;

CREATE TABLE Targets (
    CampaignRunner VARCHAR (64),    /*Stores username of file uploader. Associates each target to the uploader.*/
    CampaignID VARCHAR (64),        /*Stores the generated campaign ID. Associates each target to a specific campaign.*/
    TargetName VARCHAR (255),       /*Full name of the target. From uploaded csv/xlsx.*/
    PhoneNumber VARCHAR(32),        /*Phone number of the target. From uploaded csv/xlsx.*/
    Email VARCHAR(255),             /*Email of the target. From uploaded csv/xlsx.*/
    Position VARCHAR (128),         /*The target's position in the company. From uploaded csv/xlsx.*/
    ReportsTo VARCHAR (255),        /*Who the target reports to/the target's boss. From uploaded csv/xlsx.*/
    Status VARCHAR (32),            /*Status of the target's smish. IE: Not Received, Received, Clicked, Entered Credentials.*/
    CONSTRAINT PRIMARY KEY (CampaignID, PhoneNumber)   /*CampaignID and PhoneNumber are the primary key.*/
);

ALTER TABLE Targets ALTER Status SET DEFAULT "Not Received"; /*Set the default value of Status to be "Not Received"*/

/*Test Data*/
INSERT INTO Targets VALUES
    ("NintendoITDept", "a8d9a8fe7f805966ba16b44a117b2fdb9e766c8e35f89ff3ea1aa11f001fdf4f", "Shigeru Miyamoto", "(875) 662-9600", "smiyamoto@nintendo.com", "CEO", "BoD", DEFAULT),
    ("NintendoITDept", "a8d9a8fe7f805966ba16b44a117b2fdb9e766c8e35f89ff3ea1aa11f001fdf4f", "Masahiro Sakurai", "(914) 381-7054", "msakurai@nintendo.com", "Game Director", "Shigeru Miyamoto", DEFAULT),
    ("NintendoITDept", "a8d9a8fe7f805966ba16b44a117b2fdb9e766c8e35f89ff3ea1aa11f001fdf4f", "Tadashi Sugiyama", "(423) 715-3930", "tsugiyama@nintendo.com", "Head Graphic Designer", "Masahiro Sakurai", DEFAULT),
    ("NintendoITDept", "073346bc86b698f12a6d3e31de7a91380b5462dec21ddc07dc8c389a9b81c682", "Takashi Tezuka", "(108) 217-8256", "ttezuka@nintendo.com", "Game Director", "Shigeru Miyamoto", DEFAULT),
    ("NintendoITDept", "073346bc86b698f12a6d3e31de7a91380b5462dec21ddc07dc8c389a9b81c682", "Satoru Takizawa", "(252) 463-1139", "stakizawa@nintendo.com", "Head Graphic Designer", "Takashi Tezuka", DEFAULT),
    ("NintendoITDept", "073346bc86b698f12a6d3e31de7a91380b5462dec21ddc07dc8c389a9b81c682", "Eiji Aonuma", "(860) 317-4957", "eaonuma@nintendo.com", "Head Programmer", "Takashi Tezuka", DEFAULT)
;
<?php

/*
  This PHP Code uploads a csv file to a folder called 'uploads'.
  It then injests the file and stores it contents into a database.
  Lastly, the uploaded file is deleted.

  Current limitations:
  -At the present the csv file has to have its data ordered by the following columns:
    -'Name, Phone Number, Email, Position, Reports To'
  -The file can not have a top row for headers, only data.

  To-Do:
  -Still needs a method to store the username into the data. This will be needed to for graphs that link data between multiple campaigns.
*/


if(isset($_POST['submit']))	{ 

  $file = $_FILES['file'];
  $fileName = $_FILES['file']['name'];
  $fileTmpName = $_FILES['file']['tmp_name'];
  $fileSize = $_FILES['file']['size'];
  $fileError = $_FILES['file']['error'];
  $fileType = $_FILES['file']['type'];

  $tempFileExt = explode('.', $fileName);
  $fileExt = strtolower(end($tempFileExt));

  $allowed = array('csv');

  $campaignID = hash('sha256', time()); //can add username to pre-hash to avoid conflicts

  if (in_array($fileExt, $allowed)){
      if ($fileError === 0){
        if ($fileSize < 100000){
          $newFileName = $campaignID.".".$fileExt;
          $fileDestination = "../uploads/".$newFileName;
          move_uploaded_file($fileTmpName, $fileDestination);

          echo "Upload successful.";

          InsertData($fileDestination, "user", $campaignID); //needs a way to be passed the current user

          unlink($fileDestination);

        }else{
          echo "File size cannot exceed 100MB.";
        }
      }else{
        echo "Error: File could not be uploaded.";
      }
  }else{
    echo "File type not allowed.";
  }

}

  function connDB(){
    $dbservername = "localhost";
    $dbusername = "root";
    $dbpassword = "";
    $database = "Astia";
    
    $connection = new mysqli($dbservername, $dbusername, $dbpassword, $database);

      if($connection->connect_error) {
      die("Connection failed: " . $connection->connect_error);
    }

      return $connection;
  }
  
  function InsertData($filePath, $user, $campaignID){

    $connection = connDB();

    $file = fopen($filePath, "r");

    while (($row = fgetcsv($file)) !== FALSE) {
      $query = $connection->prepare('INSERT INTO targets VALUES (?, ?, ?, ?, ?, ?, ?, DEFAULT)');
      $query->bind_param("sssssss", $user, $campaignID, $row[0], $row[1], $row[2], $row[3], $row[4]);	
      $query->execute();
    }

  }

  function ExtractData($campaignID)

?>
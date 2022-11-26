# ------------------------------------------------- #
# Calling this function once will send sms to each  |
# and every number extracted from csv file.         |
# ------------------------------------------------- #

#importing dependencies
import pandas as pd
from twilio.rest import Client

# pass the excel file name
def send_message(str, body):       
    # Reading the csv file 
    data = pd.read_csv(str)  
    # Extracting specific column to a list
    number_column = data['PhoneNumber'].tolist() 
    account_sid = ''   # twilio account sid
    auth_token = ''    # twilio auth token
    client = Client(account_sid, auth_token)
    for i in range(len(number_column)):
        message = client.messages.create(
            body= body,
            from_='',   # number provided from twilio
            to=number_column[i]
        )


template_dict = {
    "UPS Tracking": "Your package, getting tracked by ParcelApp LLC can be viewed here: (URL)",
    "OTP": "Your OTP is 930, also go to (URL) here",
    "Join Religious Org": "Have you thought out god recently? Click here: (URL) to learn more",
    "Research Polls": "Hi, this is Research-Polls! Interested in voicing your opinions on New York? Your confidentiality is guaranteed! Click on the link to get started: (URL). Not interested in voicing your opinions? Reply STOP. Want to receive more polls? Reply Yes",
    "Sus Bank": "A suspicious ip address logind in from, deny them to revoke access: (URL)"
}

template = input("Select a template. Available templates are listed below\nUPS Tracking\nOTP\nJoin Religious Org\nResearch Polls\nSus Bank\nType here to test: ")

send_message(csvfile.csv, template_dict[template]) # calling function by passing the csv file name

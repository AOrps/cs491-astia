import sys
import os
from twilio.rest import Client
from dotenv import load_dotenv
import pandas as pd

load_dotenv(".env")


# Find your Account SID and Auth Token at twilio.com/console
account_sid = os.getenv('TWILIO_ACCOUNT_SID')
auth_token = os.getenv('TWILIO_AUTH_TOKEN')
client = Client(account_sid, auth_token)


# Function for sending sms
def send_msg():
    number_column = data['PhoneNumber'].tolist()    # Extracting Phone Number from PhoneNumber column
    name_column = data['Name']                      # Extracting Name from Name column
    for i in range(len(number_column)):             # Looping through each numbers for sending sms to all
        message = client.messages \
            .create(
            body=template,
            from_='',                               # Find this number in your twilio Account
            to=number_column[i]
        )
        print("Sent to " + name_column[i] + " Body: " + message.body)


# Remote sites can be replaced by malicious sites 
remote_site = ["astia-smish.ml", "astia-smish.ml", "astia-smish.ml", "astia-smish.ml", "astia-smish.ml"]
template_dict = {
    "UPS Tracking": f"Your package, getting tracked by ParcelApp LLC can be viewed here: {remote_site[0]}",
    "OTP": f"Your OTP is 930, also go to {remote_site[1]} here",
    "Join Religious Org": f"Have you thought out god recently? Click here: {remote_site[2]} to learn more",
    "Research Polls": f"Hi, this is Research-Polls! Interested in voicing your opinions on New York? Your confidentiality is guaranteed! Click on the link to get started: {remote_site[3]}. Not interested in voicing your opinions? Reply STOP. Want to receive more polls? Reply Yes",
    "Sus Bank": f"A suspicious ip address logind in from, deny them to revoke access: {remote_site[4]}"
}

file_name = sys.argv[1]
template = template_dict[sys.argv[2]]


# Determining whether the file is excel or csv and calling the function for smishing
if file_name.endswith(".csv"):                            
    data = pd.read_csv(file_name)
    send_msg()
if file_name.endswith(".xlsx"):
    data = pd.read_excel(file_name)
    send_msg()

import sys
import os
from twilio.rest import Client
from dotenv import load_dotenv
load_dotenv(".env")
# Find your Account SID and Auth Token at twilio.com/console

account_sid = os.getenv('TWILIO_ACCOUNT_SID')
auth_token = os.getenv('TWILIO_AUTH_TOKEN')
client = Client(account_sid, auth_token)

message = client.messages \
                .create(
                     body=sys.argv[2],
                     from_='', # Find in your twilio Account
                     to=sys.argv[1]
                 )

print("Sent to " + sys.argv[1] +" Body: " + message.body)
# meetings_api
main.go is the only file which handles the requests and returns the response.
Database name: appointy_task , Collection Name: test

## sample requests:-

1. Create new meeting :
http://127.0.0.1:10000/meetings 
  - (post request) 
  - (post data should contain json in this format)
  ![alt text](https://github.com/1-mukesh-1/meetings_api/blob/main/img/meetings.PNG)
  
2. Get meeting details from meeting id :
http://127.0.0.1:10000/meeting/{id} ex: http://127.0.0.1:10000/meeting/1
  - (Get request)
  - (Get data doesn't need and data, data should be in the url)

3. Get meetings which are between {start time} and {end time}
http://127.0.0.1:10000/meetings?start=<start_time>&end=<end_time>
ex: http://127.0.0.1:10000/meetings?start=2018-01-20%2004:35&end=2018-01-20%2006:35
  - this url displays all the meetings between 4:35 to 6:35 on date 2018-01-20
  - (Get request)

4. Get meetings in which participant.email is present
http://127.0.0.1:10000/meetings?participant=<email>
http://127.0.0.1:10000/meetings?participant=chmukesh1612@gmail.com
  - this url displays all the meetings in which <email> is present
  - (Get request)

# Sample Screenshots:

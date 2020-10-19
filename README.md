# meetings_api

## sample requests:-

http://127.0.0.1:10000/meetings 
  - insert document 
  - (post request) 
  - (post data should contain json in this format)
  type Meeting struct {
    Id           string
    Title        string
    Participants []Part{
      Part{
        Name  string
        Email string
        Rsvp  string
      },
      Part{
        Name  string
        Email string
        Rsvp  string
      },
      Part{
        Name  string
        Email string
        Rsvp  string
      },........ any number of participants
    }
    Start_Time   time.Time
    End_Time     time.Time
    Timestamp    time.Time
  }
  
  

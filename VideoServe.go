package randvid

import (
    //"fmt"
    "net/http"
    "math/rand"
    "time"
    "html/template"

    "appengine"
    "appengine/datastore"
    "appengine/user"
)

type InsertLink struct {
  Video string
  Rand int
}

type Greeting struct {
        Author  string
        Content string
        Date    time.Time
}

var guestbookTemplate = template.Must(template.New("book").Parse(`
<html>
  <head>
    <title>Go Guestbook</title>
  </head>
  <body>
    {{range .}}
      {{with .Author}}
        <p><b>{{.}}</b> wrote:</p>
      {{else}}
        <p>An anonymous person wrote:</p>
      {{end}}
      <pre>{{.Content}}</pre>
    {{end}}
    <form action="/sign" method="post">
      <div><textarea name="content" rows="3" cols="60"></textarea></div>
      <div><input type="submit" value="Sign Guestbook"></div>
    </form>
  </body>
</html>
`))

const body = `
<html>
  <head>
    <title>Colin's Random-Video</title>
    <meta http-equiv="refresh" content="2; URL='{{.Video}}'" />
  </head>
  <body bgcolor="#ffffff">
    <center>
      Please wait to be redirected. If you are not redirected please click <a href="{{.Video}}"> here</a></br>
      New Videos are added every day. Check back often!</br>
      Your random number is {{.Rand}}
    </center>
  </body>
</html>
`

const bodySFW = `
<html>
  <head>
    <title>Colin's Random-Video</title>
    <meta http-equiv="refresh" content="2; URL='{{.Video}}'" />
  </head>
  <body bgcolor="#ffffff">
    <center>
      Please wait to be redirected. If you are not redirected please click <a href="{{.Video}}"> here</a></br>
      New Videos are added every day. Check back often!</br>
      Your random number is {{.Rand}} </br>
      <b>YOU ARE USING THE SFW VERSION</b>
    </center>
  </body>
</html>
`

func RandLinkSFW() (string, int){
  VideoListSFW := []string{
    "https://www.youtube.com/watch?v=tE1HDSipRxU",                          // 0 Steve Irwin
    "http://img.4plebs.org/boards/f/image/1434/92/1434921182936.swf",       // 1  Just Do It
    "http://img.4plebs.org/boards/f/image/1442/53/1442530025364.swf",       // 2  SpongeBob
    //"http://img.4plebs.org/boards/f/image/1396/24/1396243585761.swf",       // 3  Big Shake
    //"http://img.4plebs.org/boards/f/image/1421/12/1421126458648.swf",       // 4  2012
    "https://www.youtube.com/watch?v=G2e_M06YDyY",                          // 5  Sieze the Day
    "https://www.youtube.com/watch?v=DZGINaRUEkU",                          // 6  Symphony of Science
    "https://www.youtube.com/watch?v=otnyM9RJG4o",                          // 7  Power of Music
    //"https://www.youtube.com/watch?v=2MN1vXO5JeI",                          // 8  Singalong
    "https://www.youtube.com/watch?v=u_jRgv-UqBU",                          // 9  Billy, Walmart Greeter
    "https://www.youtube.com/watch?v=DqC7H7_Noi8",                          // 10  Neil Armstrong Tribute
    "https://www.youtube.com/watch?v=2rEuie5lpGA",                          // 11 SURF THE NET
    "http://i.4cdn.org/wsg/1444358103893.webm",                             // 12 Thug Cat
    //"http://i.4cdn.org/wsg/1443246935183.webm",                             // 13 Plane Crash
    "http://i.4cdn.org/wsg/1443251102863.webm",                             // 14 Old Spice
    "http://i.4cdn.org/wsg/1443254714327.webm",                             // 15 McDonald's Remix
    "http://i.4cdn.org/wsg/1443913579648.webm",                             // 16 Sledge Hammer
    //"http://i.4cdn.org/wsg/1444112721405.webm",                             // 17 Frozen Jesse Pinkman
    "http://i.4cdn.org/wsg/1443246656431.webm",                             // 18 Talking Carl
    "http://i.4cdn.org/wsg/1443199345935.webm",                             // 19 Mario Head Bang
    //"http://i.4cdn.org/wsg/1443246507978.webm",                             // 20 Gin+Juice
    "http://i.4cdn.org/wsg/1444262136311.webm",                             // 21 Tomato -> Fan
    //"http://i.4cdn.org/wsg/1444262646541.webm",                             // 22 Old Man + Rollerskates
    "http://i.4cdn.org/wsg/1444268978629.webm",                             // 23 Rugby Kid
    "http://i.4cdn.org/wsg/1444276631906.webm",                             // 24 Beautiful Science
    //"http://i.4cdn.org/wsg/1444301298445.webm",                             // 25 Hitler Leek
    "http://i.4cdn.org/wsg/1444314579514.webm",                             // 26 Bobcat Loading
    "http://i.4cdn.org/wsg/1444339262566.webm",                             // 27 Throwing Knives
    "http://i.4cdn.org/wsg/1444339554930.gif",                              // 28 Perfect Circle
    "http://i.4cdn.org/wsg/1444372368654.webm",                             // 29 Yodeling
    //"http://i.4cdn.org/wsg/1444415934593.webm",                             // 30 Leek Gun
    "http://i.4cdn.org/wsg/1444448380633.webm",                             // 31 Cow Bell
    "http://i.4cdn.org/wsg/1444467926220.webm",                             // 32 Arnold Palmer
    //"https://www.youtube.com/watch?v=2svVkkNuSq0",                          // 33 Stop to my Beat
    "https://www.youtube.com/watch?v=DX_eeOZVS2o",                          // 34 Microwave Dance
    "http://i.4cdn.org/wsg/1443899505305.webm",                             // 35 I Don't Need a Jacket
    //"http://i.4cdn.org/wsg/1444055325919.webm",                             // 36 GTFO
    "http://i.4cdn.org/wsg/1441602608583.webm",                             // 37 Steven Universe
    "http://i.4cdn.org/wsg/1444244530145.webm",                             // 38 BAD BOYZ
    "http://i.4cdn.org/wsg/1444254081149.webm",                             // 39 BEER ME
    "http://i.4cdn.org/wsg/1444371392998.webm",                             // 40 Pro Dad
    "http://i.4cdn.org/wsg/1444440546084.webm",                             // 41 Drum Dog
    //"http://i.4cdn.org/wsg/1440917752642.webm",                             // 42 Get Down Cat
    "http://i.4cdn.org/wsg/1440917996496.webm",                             // 43 Chicken Pokemon
    "http://i.4cdn.org/wsg/1442157378169.webm",                             // 44 Dubstep Dog
    "http://i.4cdn.org/wsg/1442157620314.webm",                             // 45 Shovel + Head
    "http://i.4cdn.org/wsg/1442218215289.webm",                             // 46 Sly Kid
    "http://i.4cdn.org/wsg/1442384146512.webm",                             // 47 Arriba
    //"http://i.4cdn.org/wsg/1443312573923.webm",                             // 48 Hot Boy Dog
    "http://i.4cdn.org/wsg/1440991924939.webm",                             // 49 More Doge
    "http://i.4cdn.org/wsg/1443393958007.webm",                             // 50 Even More Doge
    "http://i.imgur.com/J7VGU2g.gifv",                                      // 51 Kittens + Puppies
    "http://i.imgur.com/ZuMSuvM.gifv",                                      // 52 Doge on Ledge
  }
  r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
  vid := r.Intn(len(VideoListSFW))
  return VideoListSFW[vid], vid
}

func RandLinkNSFW() (string, int){
  VideoListNSFW := []string{
    "https://www.youtube.com/watch?v=tE1HDSipRxU",                          // 0 Steve Irwin
    "http://img.4plebs.org/boards/f/image/1434/92/1434921182936.swf",       // 1  Just Do It
    "http://img.4plebs.org/boards/f/image/1442/53/1442530025364.swf",       // 2  SpongeBob
    "http://img.4plebs.org/boards/f/image/1396/24/1396243585761.swf",       // 3  Big Shake
    "http://img.4plebs.org/boards/f/image/1421/12/1421126458648.swf",       // 4  2012
    "https://www.youtube.com/watch?v=G2e_M06YDyY",                          // 5  Sieze the Day
    "https://www.youtube.com/watch?v=DZGINaRUEkU",                          // 6  Symphony of Science
    "https://www.youtube.com/watch?v=otnyM9RJG4o",                          // 7  Power of Music
    "https://www.youtube.com/watch?v=2MN1vXO5JeI",                          // 8  Singalong
    "https://www.youtube.com/watch?v=u_jRgv-UqBU",                          // 9  Billy, Walmart Greeter
    "https://www.youtube.com/watch?v=DqC7H7_Noi8",                          // 10  Neil Armstrong Tribute
    "https://www.youtube.com/watch?v=2rEuie5lpGA",                          // 11 SURF THE NET
    "http://i.4cdn.org/wsg/1444358103893.webm",                             // 12 Thug Cat
    "http://i.4cdn.org/wsg/1443246935183.webm",                             // 13 Plane Crash
    "http://i.4cdn.org/wsg/1443251102863.webm",                             // 14 Old Spice
    "http://i.4cdn.org/wsg/1443254714327.webm",                             // 15 McDonald's Remix
    "http://i.4cdn.org/wsg/1443913579648.webm",                             // 16 Sledge Hammer
    "http://i.4cdn.org/wsg/1444112721405.webm",                             // 17 Frozen Jesse Pinkman
    "http://i.4cdn.org/wsg/1443246656431.webm",                             // 18 Talking Carl
    "http://i.4cdn.org/wsg/1443199345935.webm",                             // 19 Mario Head Bang
    "http://i.4cdn.org/wsg/1443246507978.webm",                             // 20 Gin+Juice
    "http://i.4cdn.org/wsg/1444262136311.webm",                             // 21 Tomato -> Fan
    "http://i.4cdn.org/wsg/1444262646541.webm",                             // 22 Old Man + Rollerskates
    "http://i.4cdn.org/wsg/1444268978629.webm",                             // 23 Rugby Kid
    "http://i.4cdn.org/wsg/1444276631906.webm",                             // 24 Beautiful Science
    "http://i.4cdn.org/wsg/1444301298445.webm",                             // 25 Hitler Leek
    "http://i.4cdn.org/wsg/1444314579514.webm",                             // 26 Bobcat Loading
    "http://i.4cdn.org/wsg/1444339262566.webm",                             // 27 Throwing Knives
    "http://i.4cdn.org/wsg/1444339554930.gif",                              // 28 Perfect Circle
    "http://i.4cdn.org/wsg/1444372368654.webm",                             // 29 Yodeling
    "http://i.4cdn.org/wsg/1444415934593.webm",                             // 30 Leek Gun
    "http://i.4cdn.org/wsg/1444448380633.webm",                             // 31 Cow Bell
    "http://i.4cdn.org/wsg/1444467926220.webm",                             // 32 Arnold Palmer
    "https://www.youtube.com/watch?v=2svVkkNuSq0",                          // 33 Stop to my Beat
    "https://www.youtube.com/watch?v=DX_eeOZVS2o",                          // 34 Microwave Dance
    "http://i.4cdn.org/wsg/1443899505305.webm",                             // 35 I Don't Need a Jacket
    "http://i.4cdn.org/wsg/1444055325919.webm",                             // 36 GTFO
    "http://i.4cdn.org/wsg/1441602608583.webm",                             // 37 Steven Universe
    "http://i.4cdn.org/wsg/1444244530145.webm",                             // 38 BAD BOYZ
    "http://i.4cdn.org/wsg/1444254081149.webm",                             // 39 BEER ME
    "http://i.4cdn.org/wsg/1444371392998.webm",                             // 40 Pro Dad
    "http://i.4cdn.org/wsg/1444440546084.webm",                             // 41 Drum Dog
    "http://i.4cdn.org/wsg/1440917752642.webm",                             // 42 Get Down Cat
    "http://i.4cdn.org/wsg/1440917996496.webm",                             // 43 Chicken Pokemon
    "http://i.4cdn.org/wsg/1442157378169.webm",                             // 44 Dubstep Dog
    "http://i.4cdn.org/wsg/1442157620314.webm",                             // 45 Shovel + Head
    "http://i.4cdn.org/wsg/1442218215289.webm",                             // 46 Sly Kid
    "http://i.4cdn.org/wsg/1442384146512.webm",                             // 47 Arriba
    "http://i.4cdn.org/wsg/1443312573923.webm",                             // 48 Hot Boy Dog
    "http://i.4cdn.org/wsg/1440991924939.webm",                             // 49 More Doge
    "http://i.4cdn.org/wsg/1443393958007.webm",                             // 50 Even More Doge
    "http://i.imgur.com/J7VGU2g.gifv",                                      // 51 Kittens + Puppies
    "http://i.imgur.com/ZuMSuvM.gifv",                                      // 52 Doge on Ledge
    //"http://i.4cdn.org/wsg/1445033632154.webm",                             // 53 First Kiss/Life Insurance
    "http://i.4cdn.org/wsg/1445043480702.webm",                             // 54 Dancing Birdz
    "http://i.4cdn.org/wsg/1445058535377.webm",                             // 55 Wendy's Commercial
    "http://i.4cdn.org/wsg/1445068822578.webm",                             // 56 Water Bottle Kick
    "http://i.4cdn.org/wsg/1445110085377.webm",                             // 57 Pile of Balls
    "http://i.4cdn.org/wsg/1445110169319.webm",                             // 58 Mimicking Bird
    "http://i.4cdn.org/wsg/1445110618823.webm",                             // 59 Drum Keyboard
    "http://i.4cdn.org/wsg/1445110688973.webm",                             // 60 Rolaids
    "http://i.4cdn.org/wsg/1444956089069.webm",                             // 61 Terrible Email
    //"http://i.4cdn.org/wsg/1444966036876.webm",                           // 62 Trap + Horse
    "http://i.4cdn.org/wsg/1444985967218.webm",                             // 63 Racist SpongeBob
    "http://i.4cdn.org/wsg/1445031045338.webm",                             // 64 White Rapping
    "http://i.4cdn.org/wsg/1444877206633.webm",                             // 65 Amanda Berry Rap
    "http://i.4cdn.org/wsg/1444944916366.webm",                             // 66 Ukrainian Army Fail
    "http://i.4cdn.org/wsg/1444954291292.webm",                             // 67 Trump Dogg
    "http://i.4cdn.org/wsg/1444681180896.webm",                             // 68 Eminem Bill Cosby
    "http://i.4cdn.org/wsg/1444687230838.webm",                             // 69 Barbie March
    "http://i.4cdn.org/wsg/1444691894641.webm",                             // 70 Steal yo Girl
    "http://i.4cdn.org/wsg/1444697007201.webm",                             // 71 Jigsaw
    "http://i.4cdn.org/wsg/1444704274634.webm",                             // 72 British Jokes
    "http://i.4cdn.org/wsg/1444707961154.webm",                             // 73 Grinch Yoga
    "http://i.4cdn.org/wsg/1444719181433.webm",                             // 74 Chatty Patty
    "http://i.4cdn.org/wsg/1444728695622.webm",                             // 75 Lizard vs Cat
    "http://i.4cdn.org/wsg/1444729232374.webm",                             // 76 Raccoon
    "",                             // 77
    "",                             // 78
    "",                             // 79
    "",                             // 80
  }
  r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
  vid := r.Intn(len(VideoListNSFW))
  return VideoListNSFW[vid], vid
}

func IndexSFW(w http.ResponseWriter, r *http.Request) {
  w.Header().Add("Content Type", "text/html")
  tmpl, err := template.New("video").Parse(bodySFW)
  video, rand := RandLinkSFW()
  if err == nil {
    redirect := InsertLink{video, rand}
    tmpl.Execute(w, redirect)
  } else {
    panic(err)
  }
}

func IndexNSFW(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content Type", "text/html")
    tmpl, err := template.New("video").Parse(body)
    video, rand := RandLinkNSFW()
    if err == nil {
      redirect := InsertLink{video, rand}
      tmpl.Execute(w, redirect)
    } else {
      panic(err)
    }
}

// guestbookKey returns the key used for all guestbook entries.
func guestbookKey(c appengine.Context) *datastore.Key {
        // The string "default_guestbook" here could be varied to have multiple guestbooks.
        return datastore.NewKey(c, "Guestbook", "default_guestbook", 0, nil)
}

// [START func_root]
func root(w http.ResponseWriter, r *http.Request) {
        c := appengine.NewContext(r)
        // Ancestor queries, as shown here, are strongly consistent with the High
        // Replication Datastore. Queries that span entity groups are eventually
        // consistent. If we omitted the .Ancestor from this query there would be
        // a slight chance that Greeting that had just been written would not
        // show up in a query.
        // [START query]
        q := datastore.NewQuery("Greeting").Ancestor(guestbookKey(c)).Order("-Date").Limit(10)
        // [END query]
        // [START getall]
        greetings := make([]Greeting, 0, 10)
        if _, err := q.GetAll(c, &greetings); err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }
        // [END getall]
        if err := guestbookTemplate.Execute(w, greetings); err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
        }
}
// [END func_root]
// [START func_sign]
func sign(w http.ResponseWriter, r *http.Request) {
        c := appengine.NewContext(r)
        g := Greeting{
                Content: r.FormValue("content"),
                Date:    time.Now(),
        }
        if u := user.Current(c); u != nil {
                g.Author = u.String()
        }
        // We set the same parent key on every Greeting entity to ensure each Greeting
        // is in the same entity group. Queries across the single entity group
        // will be consistent. However, the write rate to a single entity group
        // should be limited to ~1/second.
        key := datastore.NewIncompleteKey(c, "Greeting", guestbookKey(c))
        _, err := datastore.Put(c, key, &g)
        if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }
        http.Redirect(w, r, "/suggest", http.StatusFound)
}
// [END func_sign]

func init() {
  http.HandleFunc("/sfw", IndexSFW)
  http.HandleFunc("/suggest", root)
  http.HandleFunc("/sign", sign)
  http.HandleFunc("/", IndexNSFW)
}

// goapp deploy -application rand-vid app.yaml

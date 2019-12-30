package main

/* 
There is some more naivety in the way we are dealing with files which could create a very nasty bug down the line.
When we RecordWin, we Seek back to the start of the file and then write the new data—but what if the new data was smaller than what was there before?
In our current case, this is impossible. We never edit or delete scores so the data can only get bigger. However, it would be irresponsible for us to leave the code like this; it's not unthinkable that a delete scenario could come up.
How will we test for this though? What we need to do is first refactor our code so we separate out the concern of the kind of data we write, from the writing. We can then test that separately to check it works how we hope.
We'll create a new type to encapsulate our "when we write we go from the beginning" functionality. I'm going to call it Tape. Create a new file with the following:

Notice that we're only implementing Write now, as it encapsulates the Seek part. This means our FileSystemStore can just have a reference to a Writer instead.
type FileSystemPlayerStore struct {
    database io.Writer
    league   League
}
Update the constructor to use Tape
func NewFileSystemPlayerStore(database io.ReadWriteSeeker) *FileSystemPlayerStore {
    database.Seek(0, 0)
    league, _ := NewLeague(database)

    return &FileSystemPlayerStore{
        database: &tape{database},
        league:   league,
    }
}
Finally, we can get the amazing payoff we wanted by removing the Seek call from RecordWin. Yes, it doesn't feel much, but at least it means if we do any other kind of writes we can rely on our Write to behave how we need it to. 
Plus it will now let us test the potentially problematic code separately and fix it.
*/

import (
	"os"
)

type tape struct {
    file *os.File
}

func (t *tape) Write(p []byte) (n int, err error) {
	t.file.Truncate(0)
    t.file.Seek(0, 0)
    return t.file.Write(p)
}

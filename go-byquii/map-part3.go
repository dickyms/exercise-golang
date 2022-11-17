// belum selesai


var ErrNotFound = errors.New("could not find the word you were looking for")

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error){
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(w)

func TestSearch(t *testing.T) {
	dic := Dictionary{"test": "this is just a test"}

	assertStrings:= func(t testing.TB, got, want error) {
		t.Helper()
		if got != want {
			t.Errorf("got error %q want %q", got, want)
		}
	}

	t.Run("know word", func(t *testing.T){
		_, got := dic.Search("unknown")
		assertStrings(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dic := Dictionary{}
		word := "test"
		definition := "this is just a test"

		err := dic.Add(word, definition)
		assertError(t, err, nil)
		assertDefinition(t, dic, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dic := Dictionary{word: definition}
		err := dic.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dic, word, definition)
	})
}
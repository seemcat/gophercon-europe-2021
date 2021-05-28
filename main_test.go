package main

import (
	"testing"
	"reflect"
)

func TestPrintTweets(t *testing.T) {
	var secondTweet = []string{"2", "This is our second default tweet!", "maricris@magic.link"}
	got := PrintTweets(secondTweet)
	want := secondTweet

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/vanng822/go-solr/solr"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"gopkg.in/zabawaba99/firego.v1"
)

func main() {
	opt := option.WithCredentialsFile("/Users/bhargav_skr/Downloads/profile-1164ab738a7e.json")
	_, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	f := firego.New("https://profile-7a716.firebaseio.com/profiles", nil)

	d, err := ioutil.ReadFile("/Users/bhargav_skr/Downloads/profile-1164ab738a7e.json")
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	conf, err := google.JWTConfigFromJSON(d, "https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/firebase.database")
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	_ = firego.New("https://profile-7a716.firebaseio.com", conf.Client(oauth2.NoContext))
	//fmt.Printf("%s", fb)
	var v map[string]Profile
	if err := f.Value(&v); err != nil {
		log.Fatal(err)
	}
	//fmt.Print(v)
	var profiles []solr.Document

	for k, v := range v {
		//	fmt.Println("k:", k, "v:", v)
		doc := make(solr.Document)
		marshallv, _ := json.Marshal(v)
		doc.Set(k, marshallv)
		//fmt.Println(doc)
		profiles = append(profiles, doc)
	}
	//fmt.Println(len(profiles))

	si, _ := solr.NewSolrInterface("http://localhost:8983/solr", "profile")
	//si, _ := solr.NewSolrInterface("http://localhost:8983/solr", "core0")
	// query := solr.NewQuery()
	// query.Q("*:*")
	// s := si.Search(query)
	// r, _ := s.Result(nil)
	// fmt.Println(r.Results.Docs)
	//docs := make([]solr.Document, 0, 5)
	//for i := 0; i < 5; i++ {
	//	docs = append(docs, solr.Document{"id": fmt.Sprintf("test_id_%d", i), "title": fmt.Sprintf("add sucess %d", i)})
	//}
	//fmt.Println(docs)
	//res, _ := si.Add(docs, 0, nil)

	//fmt.Println(profiles)
	res, _ := si.Add(profiles, 0, nil)
	res2, _ := si.Commit()
	// not sure what we can test here but at least run and see thing flows
	if res == nil {
		log.Fatalf("Add response should not be nil")
	}
	fmt.Println(res)
	if res.Success != true {
		log.Fatalf("res.Success should be true but got false")
	}

	if res2 == nil {
		log.Fatalf("Commit response should not be nil")
	}

}

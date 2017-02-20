package main

import (
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {

	//tp := github.BasicAuthTransport{
	//	Username: strings.TrimSpace("heidsoft"),
	//	Password: strings.TrimSpace("LiuGanBing72712601017"),
	//}

	//client := github.NewClient(tp.Client())


	//var allRepos []*github.Repository
	//for {
	//	repos, resp, err := client.Repositories.ListByOrg("github", opt)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	fmt.Println(repos)
	//	allRepos = append(allRepos, repos...)
	//	if resp.NextPage == 0 {
	//		break
	//	}
	//	opt.ListOptions.Page = resp.NextPage
	//}

	//repos, _, err := client.Repositories.List("fork", nil)
	//if err != nil{
	//	fmt.Println(err)
	//}



	//86f7199490ed1596c8be2ffd7872701fdb550b1a

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "86f7199490ed1596c8be2ffd7872701fdb550b1a"},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)


	opt := &github.RepositoryListOptions{
		Visibility:"public",
		Type:"private",
	}

	// list all repositories for the authenticated user
	repos, _, err := client.Repositories.List("heidsoft", opt)
	if err != nil{
		fmt.Println(err)
	}
	if repos != nil {
		for _,repo := range repos{
			fmt.Println(*repo.Name)
		}
	}
}